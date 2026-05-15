package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

// DocumentsService handles Documents operations.
type DocumentsService struct {
	client *MeibelClient
}

// DocumentsProcessOptions contains optional parameters for Process.
type DocumentsProcessOptions struct {
	// Result format: markdown, annotated, docling, json
	Format *string
}

// DocumentsGetResultOptions contains optional parameters for GetResult.
type DocumentsGetResultOptions struct {
	// Result format: markdown, annotated, docling, json
	Format *string
}

// Parse Parse a document (async)
//
// Upload a document for asynchronous parsing. Returns a job ID to track progress.
func (s *DocumentsService) Parse(ctx context.Context, file io.Reader, fileName string) (*ParseDocumentResponse, error) {
	path := "/documents"

	var result ParseDocumentResponse
	uploadFields := []UploadField{
		{FieldName: "file", Reader: file, FileName: fileName},
	}
	formFields := map[string]string{}

	err := s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, uploadFields, formFields, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Process Parse a document (sync)
//
// Upload a document and block until parsing is complete. Returns the full parsed result.
func (s *DocumentsService) Process(ctx context.Context, file io.Reader, fileName string, opts *DocumentsProcessOptions) (*ProcessDocumentResponse, error) {
	path := "/documents/process"
	query := url.Values{}
	if opts != nil && opts.Format != nil {
		query.Set("format", fmt.Sprintf("%v", *opts.Format))
	}

	var result ProcessDocumentResponse
	uploadFields := []UploadField{
		{FieldName: "file", Reader: file, FileName: fileName},
	}
	formFields := map[string]string{}

	err := s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Query:  query,
	}, uploadFields, formFields, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetStatus Get document parsing status
//
// Check the status of a document parsing job, including progress statistics.
func (s *DocumentsService) GetStatus(ctx context.Context, jobId string) (*DocumentStatus, error) {
	path := "/documents/" + fmt.Sprintf("%v", jobId)

	var result DocumentStatus
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetResult Get parsed document result
//
// Download the parsed result of a completed document parsing job.
func (s *DocumentsService) GetResult(ctx context.Context, jobId string, opts *DocumentsGetResultOptions) (*string, error) {
	path := "/documents/" + fmt.Sprintf("%v", jobId) + "/result"
	query := url.Values{}
	if opts != nil && opts.Format != nil {
		query.Set("format", fmt.Sprintf("%v", *opts.Format))
	}

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListChildren List child documents
//
// For container files (ZIP, TAR, EML), list the child documents extracted from the container.
func (s *DocumentsService) ListChildren(ctx context.Context, jobId string) (*[]DocumentChild, error) {
	path := "/documents/" + fmt.Sprintf("%v", jobId) + "/children"

	var result []DocumentChild
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// StreamTrace Stream document parsing trace
//
// Subscribe to real-time parsing progress via Server-Sent Events.
func (s *DocumentsService) StreamTrace(ctx context.Context, jobId string) (*EventStream[interface{}], error) {
	path := "/documents/" + fmt.Sprintf("%v", jobId) + "/trace"

	resp, err := s.client.http.DoStream(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	})
	if err != nil {
		return nil, err
	}

	return JSONEventStream[interface{}](resp), nil
}

// TransformOptions contains parameters for Transform.
type TransformOptions struct {
	// Document file to transform
	File string
	// JSON Schema dict (as JSON string) or schema name/ID
	Schema interface{}
	// LLM model override
	Model *string
	// Extraction instructions override
	Prompt *string
	// Prompt template reference
	PromptId *string
	// Max wait time in seconds (sync only)
	TimeoutSeconds *int64
}

// Transform Transform a document using AI extraction (sync)
//
// Upload a document for AI-powered structured extraction and block until complete. The file is uploaded to cloud storage and processed by a system agent.
func (s *DocumentsService) Transform(ctx context.Context, opts TransformOptions) (*TransformDocumentResponse, error) {
	path := "/documents/transform"
	var err error

	schemaResolved, err := resolveSchema(opts.Schema)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(opts.File)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	formFields := map[string]string{}
	switch sv := schemaResolved.(type) {
	case string:
		formFields["artifact_schema"] = sv
	case nil:
		// skip
	default:
		b, _ := json.Marshal(sv)
		formFields["artifact_schema"] = string(b)
	}
	formFields["model"] = fmt.Sprintf("%v", opts.Model)
	formFields["prompt"] = fmt.Sprintf("%v", opts.Prompt)
	formFields["prompt_id"] = fmt.Sprintf("%v", opts.PromptId)
	formFields["timeout_seconds"] = fmt.Sprintf("%v", opts.TimeoutSeconds)

	var result TransformDocumentResponse
	err = s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, []UploadField{
		{FieldName: "file", Reader: f, FileName: filepath.Base(opts.File)},
	}, formFields, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SubmitTransformOptions contains parameters for SubmitTransform.
type SubmitTransformOptions struct {
	// Document file to transform
	File string
	// JSON Schema dict (as JSON string) or schema name/ID
	Schema interface{}
	// LLM model override
	Model *string
	// Extraction instructions override
	Prompt *string
	// Prompt template reference
	PromptId *string
	// Max wait time in seconds (sync only)
	TimeoutSeconds *int64
}

// SubmitTransform Submit a document transform (async)
//
// Upload a document for AI-powered extraction and return immediately. Poll for completion via client.sessions.get(execution_id).
func (s *DocumentsService) SubmitTransform(ctx context.Context, opts SubmitTransformOptions) (*SubmitDocumentTransformResponse, error) {
	path := "/documents/transform/submit"
	var err error

	schemaResolved, err := resolveSchema(opts.Schema)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(opts.File)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	formFields := map[string]string{}
	switch sv := schemaResolved.(type) {
	case string:
		formFields["artifact_schema"] = sv
	case nil:
		// skip
	default:
		b, _ := json.Marshal(sv)
		formFields["artifact_schema"] = string(b)
	}
	formFields["model"] = fmt.Sprintf("%v", opts.Model)
	formFields["prompt"] = fmt.Sprintf("%v", opts.Prompt)
	formFields["prompt_id"] = fmt.Sprintf("%v", opts.PromptId)
	formFields["timeout_seconds"] = fmt.Sprintf("%v", opts.TimeoutSeconds)

	var result SubmitDocumentTransformResponse
	err = s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, []UploadField{
		{FieldName: "file", Reader: f, FileName: filepath.Base(opts.File)},
	}, formFields, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
