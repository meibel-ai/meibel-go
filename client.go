package meibelgo

import (
	"net/http"
	"time"
)

// MeibelgoClient is the main client for the meibel API.
type MeibelgoClient struct {
	http *HTTPClient

	Blueprints *BlueprintsService
	Datasources *DatasourcesService
}

// ClientOption is a function that configures the client.
type ClientOption func(*clientOptions)

type clientOptions struct {
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
	headers    map[string]string
}

func defaultClientOptions() *clientOptions {
	return &clientOptions{
		baseURL: "https://api.meibel.ai/v1",
		timeout: 30 * time.Second,
		headers: make(map[string]string),
	}
}

// WithBaseURL sets the base URL for API requests.
func WithBaseURL(url string) ClientOption {
	return func(o *clientOptions) {
		o.baseURL = url
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *clientOptions) {
		o.timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(o *clientOptions) {
		o.httpClient = client
	}
}

// WithAPIKey sets the API key for authentication.
func WithAPIKey(key string) ClientOption {
	return func(o *clientOptions) {
		o.headers["Meibel-API-Key"] = key
	}
}

// WithBearerToken sets the bearer token for authentication.
func WithBearerToken(token string) ClientOption {
	return func(o *clientOptions) {
		o.headers["Authorization"] = "Bearer " + token
	}
}

// WithHeader sets a custom header for all requests.
func WithHeader(key, value string) ClientOption {
	return func(o *clientOptions) {
		o.headers[key] = value
	}
}

// NewClient creates a new MeibelgoClient with the given options.
func NewClient(opts ...ClientOption) *MeibelgoClient {
	cfg := defaultClientOptions()
	for _, opt := range opts {
		opt(cfg)
	}

	httpClient := NewHTTPClient(HTTPClientConfig{
		BaseURL:    cfg.baseURL,
		Timeout:    cfg.timeout,
		Headers:    cfg.headers,
		HTTPClient: cfg.httpClient,
	})

	c := &MeibelgoClient{
		http: httpClient,
	}

	c.Blueprints = &BlueprintsService{client: c}
	c.Datasources = &DatasourcesService{client: c}
	c.Blueprints.Executions = &ExecutionsService{client: c}
	c.Blueprints.Instances = &InstancesService{client: c}
	c.Datasources.Content = &ContentService{client: c}
	c.Datasources.Dataelements = &DataelementsService{client: c}
	c.Datasources.Rag = &RagService{client: c}
	c.Datasources.Tag = &TagService{client: c}

	return c
}
