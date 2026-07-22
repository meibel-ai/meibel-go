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

// DocumentsListDeepTransformsOptions contains optional parameters for ListDeepTransforms.
type DocumentsListDeepTransformsOptions struct {
	// Number of jobs to skip
	Offset *int64
	// Maximum number of jobs to return
	Limit *int64
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

// ListDeepTransforms List deep-transform jobs
//
// List the calling customer's deep-transform jobs, newest first. Scoped to the customer (and project, when a project header is set). Paginated via `offset`/`limit`.
func (s *DocumentsService) ListDeepTransforms(ctx context.Context, opts *DocumentsListDeepTransformsOptions) *PageIterator[DeepTransformJob] {
	path := "/documents/deep-transform"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[DeepTransformJob], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Jobs []DeepTransformJob `json:"jobs"`
			NextCursor string `json:"next_cursor"`
		}

		err := s.client.http.Do(ctx, RequestOptions{
			Method: "GET",
			Path:   path,
			Query:  query,
		}, &resp)
		if err != nil {
			return nil, err
		}

		return &Page[DeepTransformJob]{
			Items:      resp.Jobs,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// DocumentsSubmitDeepTransformOptions contains parameters for SubmitDeepTransform.
type DocumentsSubmitDeepTransformOptions struct {
	// Document file to extract from
	File string
	// JSON Schema (as a JSON string) of the entities to extract
	Schema interface{}
	// Name of the root entity in the schema. Optional: resolved from the schema's title or inferred when omitted.
	RootName *string
	// Optional domain guidance for the extraction
	Guidance *string
	// Optional cap on the number of pages to process
	MaxPages *int64
}

// SubmitDeepTransform Submit a deep-transform extraction from a file upload (async)
//
// Upload a document and submit an extraction against a JSON schema, returning immediately with a job id. To reuse an already-parsed document instead of uploading, use POST /documents/deep-transform/from-document. Poll status via GET /documents/deep-transform/{job_id} and download artifacts once it succeeds. Submission is idempotent on the (document, schema) pair.
func (s *DocumentsService) SubmitDeepTransform(ctx context.Context, opts DocumentsSubmitDeepTransformOptions) (*SubmitDeepTransformResponse, error) {
	path := "/documents/deep-transform"
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
		formFields["schema"] = sv
	case nil:
		// skip
	default:
		b, _ := json.Marshal(sv)
		formFields["schema"] = string(b)
	}
	formFields["root_name"] = fmt.Sprintf("%v", opts.RootName)
	formFields["guidance"] = fmt.Sprintf("%v", opts.Guidance)
	formFields["max_pages"] = fmt.Sprintf("%v", opts.MaxPages)

	var result SubmitDeepTransformResponse
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

// SubmitDeepTransformFrom Submit a deep-transform extraction reusing a parsed document (async)
//
// Submit an extraction that reuses an already-parsed document (by `document_job_id` from POST /documents) instead of re-parsing an upload. Returns immediately with a job id. Poll status via GET /documents/deep-transform/{job_id} and download artifacts once it succeeds. Submission is idempotent on the (document, schema) pair.
func (s *DocumentsService) SubmitDeepTransformFrom(ctx context.Context, body SubmitDeepTransformFromDocument) (*SubmitDeepTransformResponse, error) {
	path := "/documents/deep-transform/from-document"

	var result SubmitDeepTransformResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDeepTransformStatus Get deep-transform job status
//
// Check status and, once succeeded, the list of downloadable artifacts.
func (s *DocumentsService) GetDeepTransformStatus(ctx context.Context, jobId string) (*DeepTransformJob, error) {
	path := "/documents/deep-transform/" + fmt.Sprintf("%v", jobId)

	var result DeepTransformJob
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DownloadDeepTransformArtifact Download a deep-transform artifact
//
// Download a named artifact (e.g. output.json) produced by a succeeded job. Ownership is verified against the customer header before any bytes are returned.
func (s *DocumentsService) DownloadDeepTransformArtifact(ctx context.Context, jobId string, name string) (*string, error) {
	path := "/documents/deep-transform/" + fmt.Sprintf("%v", jobId) + "/artifact/" + fmt.Sprintf("%v", name)

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Parse Parse a document (async)
//
// Upload a document for asynchronous parsing. Returns a job ID to track progress.
func (s *DocumentsService) Parse(ctx context.Context, file io.Reader, fileName string) (*ParseDocumentResponse, error) {
	path := "/documents"

	uploadFields := []UploadField{
		{FieldName: "file", Reader: file, FileName: fileName},
	}
	formFields := map[string]string{}

	var result ParseDocumentResponse
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

	uploadFields := []UploadField{
		{FieldName: "file", Reader: file, FileName: fileName},
	}
	formFields := map[string]string{}

	var result ProcessDocumentResponse
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

// GetStructuredResult Get structured parse result
//
// Download the fully structured parse result (the json format): pages, typed elements, tables, chart data, chart OCR text, and bounding boxes. The response schema (StructuredDocument) is defined by the parsing engine and hoisted into this spec by the OpenAPI generator.
func (s *DocumentsService) GetStructuredResult(ctx context.Context, jobId string) (*ParseStructuredDocument, error) {
	path := "/documents/" + fmt.Sprintf("%v", jobId) + "/structured"

	var result ParseStructuredDocument
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
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

// DocumentsTransformOptions contains parameters for Transform.
type DocumentsTransformOptions struct {
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
func (s *DocumentsService) Transform(ctx context.Context, opts DocumentsTransformOptions) (*TransformDocumentResponse, error) {
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

// DocumentsSubmitTransformOptions contains parameters for SubmitTransform.
type DocumentsSubmitTransformOptions struct {
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
func (s *DocumentsService) SubmitTransform(ctx context.Context, opts DocumentsSubmitTransformOptions) (*SubmitDocumentTransformResponse, error) {
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

// Move Move documents into a datasource (async)
//
// Move documents (identified by their parse job IDs, e.g. the job_id returned by parseDocument) into an existing datasource or a newly created one. Returns a workflow_id to poll for completion.
func (s *DocumentsService) Move(ctx context.Context, body MoveDocumentsRequest) (*MoveDocumentsResponse, error) {
	path := "/documents/move"

	var result MoveDocumentsResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
