package v2

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

// DocumentsService handles Documents operations.
type DocumentsService struct {
	client *MeibelClient
}

// ProcessDocumentOptions contains optional parameters for ProcessDocument.
type ProcessDocumentOptions struct {
	// Result format: markdown, annotated, docling, json
	Format *string
}

// GetDocumentResultOptions contains optional parameters for GetDocumentResult.
type GetDocumentResultOptions struct {
	// Result format: markdown, annotated, docling, json
	Format *string
}

// ParseDocument Parse a document (async)
//
// Upload a document for asynchronous parsing. Returns a job ID to track progress.
func (s *DocumentsService) ParseDocument(ctx context.Context, file io.Reader, fileName string) (*ParseDocumentResponse, error) {
	path := "/documents"

	var result ParseDocumentResponse
	files := []UploadField{
		{FieldName: "file", Reader: file, FileName: fileName},
	}
	formFields := map[string]string{}

	err := s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, files, formFields, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ProcessDocument Parse a document (sync)
//
// Upload a document and block until parsing is complete. Returns the full parsed result.
func (s *DocumentsService) ProcessDocument(ctx context.Context, file io.Reader, fileName string, opts *ProcessDocumentOptions) (*ProcessDocumentResponse, error) {
	path := "/documents/process"
	query := url.Values{}
	if opts != nil && opts.Format != nil {
		query.Set("format", fmt.Sprintf("%v", *opts.Format))
	}

	var result ProcessDocumentResponse
	files := []UploadField{
		{FieldName: "file", Reader: file, FileName: fileName},
	}
	formFields := map[string]string{}

	err := s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Query:  query,
	}, files, formFields, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDocumentStatus Get document parsing status
//
// Check the status of a document parsing job, including progress statistics.
func (s *DocumentsService) GetDocumentStatus(ctx context.Context, jobId string) (*DocumentStatus, error) {
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

// GetDocumentResult Get parsed document result
//
// Download the parsed result of a completed document parsing job.
func (s *DocumentsService) GetDocumentResult(ctx context.Context, jobId string, opts *GetDocumentResultOptions) (*string, error) {
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

// ListDocumentChildren List child documents
//
// For container files (ZIP, TAR, EML), list the child documents extracted from the container.
func (s *DocumentsService) ListDocumentChildren(ctx context.Context, jobId string) (*[]DocumentChild, error) {
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

// StreamDocumentTrace Stream document parsing trace
//
// Subscribe to real-time parsing progress via Server-Sent Events.
func (s *DocumentsService) StreamDocumentTrace(ctx context.Context, jobId string) (*EventStream[interface{}], error) {
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
