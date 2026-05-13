package v2

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"reflect"
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
	// File path, URL, or GCS URI to transform
	File string
	// Schema name/ID or inline JSON Schema
	Schema interface{}
	// LLM model override
	Model interface{}
	// Extraction instructions override
	Prompt interface{}
	// Prompt template reference
	PromptId interface{}
	// Max wait time in seconds (sync only)
	TimeoutSeconds interface{}
}

// Transform Transform a document using AI extraction (sync)
//
// Submit a document for AI-powered structured extraction and block until complete. Internally orchestrates a system agent session, polls for completion, and returns the extracted data.
func (s *DocumentsService) Transform(ctx context.Context, opts TransformOptions) (*TransformDocumentResponse, error) {
	path := "/documents/transform"
	var err error

	schemaResolved, err := resolveSchema(opts.Schema)
	if err != nil {
		return nil, err
	}

	body := TransformDocumentRequest{
		File: opts.File,
		ArtifactSchema: schemaResolved,
		Model: opts.Model,
		Prompt: opts.Prompt,
		PromptId: opts.PromptId,
		TimeoutSeconds: opts.TimeoutSeconds,
	}

	var result TransformDocumentResponse
	err = s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SubmitTransformOptions contains parameters for SubmitTransform.
type SubmitTransformOptions struct {
	// File path, URL, or GCS URI to transform
	File string
	// Schema name/ID or inline JSON Schema
	Schema interface{}
	// LLM model override
	Model interface{}
	// Extraction instructions override
	Prompt interface{}
	// Prompt template reference
	PromptId interface{}
	// Max wait time in seconds (sync only)
	TimeoutSeconds interface{}
}

// SubmitTransform Submit a document transform (async)
//
// Submit a document for AI-powered extraction and return immediately. Poll for completion via client.sessions.get(execution_id).
func (s *DocumentsService) SubmitTransform(ctx context.Context, opts SubmitTransformOptions) (*SubmitDocumentTransformResponse, error) {
	path := "/documents/transform/submit"
	var err error

	schemaResolved, err := resolveSchema(opts.Schema)
	if err != nil {
		return nil, err
	}

	body := TransformDocumentRequest{
		File: opts.File,
		ArtifactSchema: schemaResolved,
		Model: opts.Model,
		Prompt: opts.Prompt,
		PromptId: opts.PromptId,
		TimeoutSeconds: opts.TimeoutSeconds,
	}

	var result SubmitDocumentTransformResponse
	err = s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
