package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

// HTTPClient handles all HTTP communication with the API.
type HTTPClient struct {
	baseURL    string
	httpClient *http.Client
	headers    map[string]string
}

// HTTPClientConfig holds configuration for the HTTP client.
type HTTPClientConfig struct {
	BaseURL    string
	Timeout    time.Duration
	Headers    map[string]string
	HTTPClient *http.Client
}

// NewHTTPClient creates a new HTTP client with the given configuration.
func NewHTTPClient(config HTTPClientConfig) *HTTPClient {
	client := config.HTTPClient
	if client == nil {
		client = &http.Client{
			Timeout: config.Timeout,
		}
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["User-Agent"] = "meibel-ai-api-go/0.4.0"
	for k, v := range config.Headers {
		headers[k] = v
	}

	return &HTTPClient{
		baseURL:    config.BaseURL,
		httpClient: client,
		headers:    headers,
	}
}

// RequestOptions holds options for a single request.
type RequestOptions struct {
	Method     string
	Path       string
	Query      url.Values
	Headers    map[string]string
	Body       interface{}
	FormFields map[string]string // Form fields to send as application/x-www-form-urlencoded (mutually exclusive with Body)
	Stream     bool
}

// Do performs an HTTP request and decodes the response into result.
func (c *HTTPClient) Do(ctx context.Context, opts RequestOptions, result interface{}) error {
	req, err := c.newRequest(ctx, opts)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &NetworkError{Err: err}
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return err
	}

	if result == nil || resp.StatusCode == http.StatusNoContent {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// DoStream performs an HTTP request and returns the response body for streaming.
func (c *HTTPClient) DoStream(ctx context.Context, opts RequestOptions) (*http.Response, error) {
	opts.Stream = true
	req, err := c.newRequest(ctx, opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &NetworkError{Err: err}
	}

	if err := c.checkResponse(resp); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return resp, nil
}

// UploadField represents a file field in a multipart upload.
type UploadField struct {
	FieldName string
	Reader    io.Reader
	FileName  string
}

// DoUpload performs a multipart file upload and decodes the JSON response.
// Uses io.Pipe for zero-copy streaming (no buffering the entire file in memory).
func (c *HTTPClient) DoUpload(ctx context.Context, opts RequestOptions, files []UploadField, formFields map[string]string, result interface{}) error {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer writer.Close()
		for _, f := range files {
			part, err := writer.CreateFormFile(f.FieldName, f.FileName)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			if _, err := io.Copy(part, f.Reader); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		for k, v := range formFields {
			if err := writer.WriteField(k, v); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()

	u, err := url.Parse(c.baseURL + opts.Path)
	if err != nil {
		return err
	}
	if opts.Query != nil {
		u.RawQuery = opts.Query.Encode()
	}

	method := opts.Method
	if method == "" {
		method = http.MethodPost
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), pr)
	if err != nil {
		return err
	}

	// Set headers — Content-Type must be multipart, not application/json
	for k, v := range c.headers {
		if k != "Content-Type" {
			req.Header.Set(k, v)
		}
	}
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &NetworkError{Err: err}
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return err
	}

	if result == nil || resp.StatusCode == http.StatusNoContent {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *HTTPClient) newRequest(ctx context.Context, opts RequestOptions) (*http.Request, error) {
	u, err := url.Parse(c.baseURL + opts.Path)
	if err != nil {
		return nil, err
	}

	if opts.Query != nil {
		u.RawQuery = opts.Query.Encode()
	}

	var body io.Reader
	isFormEncoded := false
	if len(opts.FormFields) > 0 {
		// URL-encode form fields
		form := url.Values{}
		for k, v := range opts.FormFields {
			form.Set(k, v)
		}
		body = bytes.NewReader([]byte(form.Encode()))
		isFormEncoded = true
	} else if opts.Body != nil {
		data, err := json.Marshal(opts.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(data)
	}

	method := opts.Method
	if method == "" {
		method = http.MethodGet
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set default headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	// Set request-specific headers
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	// Override Content-Type for form-encoded requests
	if isFormEncoded {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Set accept header for streaming
	if opts.Stream {
		req.Header.Set("Accept", "text/event-stream")
	}

	return req, nil
}

func (c *HTTPClient) checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)

	var errBody map[string]interface{}
	if err := json.Unmarshal(body, &errBody); err != nil {
		errBody = map[string]interface{}{"message": string(body)}
	}

	message := "request failed"
	if msg, ok := errBody["message"].(string); ok {
		message = msg
	} else if msg, ok := errBody["error"].(string); ok {
		message = msg
	}

	var code string
	if c, ok := errBody["code"].(string); ok {
		code = c
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return &AuthenticationError{APIError: APIError{Status: resp.StatusCode, Message: message, Code: code, Body: errBody}}
	case http.StatusForbidden:
		return &AuthorizationError{APIError: APIError{Status: resp.StatusCode, Message: message, Code: code, Body: errBody}}
	case http.StatusNotFound:
		return &NotFoundError{APIError: APIError{Status: resp.StatusCode, Message: message, Code: code, Body: errBody}}
	case http.StatusUnprocessableEntity:
		return &ValidationError{APIError: APIError{Status: resp.StatusCode, Message: message, Code: code, Body: errBody}}
	case http.StatusTooManyRequests:
		retryAfter := 0
		if ra := resp.Header.Get("Retry-After"); ra != "" {
			fmt.Sscanf(ra, "%d", &retryAfter)
		}
		return &RateLimitError{
			APIError:   APIError{Status: resp.StatusCode, Message: message, Code: code, Body: errBody},
			RetryAfter: retryAfter,
		}
	default:
		if resp.StatusCode >= 500 {
			return &ServerError{APIError: APIError{Status: resp.StatusCode, Message: message, Code: code, Body: errBody}}
		}
		return &APIError{Status: resp.StatusCode, Message: message, Code: code, Body: errBody}
	}
}
