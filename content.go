package meibelgo

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

// ContentService handles Content operations.
type ContentService struct {
	client *MeibelClient
}

// ListContentOptions contains optional parameters for ListContent.
type ListContentOptions struct {
	// Filter content by path prefix
	Prefix interface{}
	// Token for pagination
	ContinuationToken interface{}
	// Maximum items to return
	Limit *int64
}

// StreamUploadProgress Stream upload progress events
//
// Subscribe to real-time upload progress updates via Server-Sent Events
func (s *ContentService) StreamUploadProgress(ctx context.Context, uploadId string) error {
	path := "/uploads/" + fmt.Sprintf("%v", uploadId) + "/progress"

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// TriggerIngest Trigger ingest
//
// Trigger ingestion for a datasource
func (s *ContentService) TriggerIngest(ctx context.Context, datasourceId string) (*string, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/trigger-ingest"

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

// ListContent List Content
func (s *ContentService) ListContent(ctx context.Context, datasourceId string, opts *ListContentOptions) (*string, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content"
	query := url.Values{}
	if opts != nil && opts.Prefix != nil {
		query.Set("prefix", fmt.Sprintf("%v", opts.Prefix))
	}
	if opts != nil && opts.ContinuationToken != nil {
		query.Set("continuation_token", fmt.Sprintf("%v", opts.ContinuationToken))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
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

// UploadContent Upload Content
func (s *ContentService) UploadContent(ctx context.Context, datasourceId string, file io.Reader, fileName string) (*string, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content"

	var result string
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

// StreamUploadProgress Stream Upload Progress
func (s *ContentService) StreamUploadProgress(ctx context.Context, uploadId string) error {
	path := "/v2/uploads/" + fmt.Sprintf("%v", uploadId) + "/progress"

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// TriggerIngest Trigger Ingest
func (s *ContentService) TriggerIngest(ctx context.Context, datasourceId string) (*string, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/trigger-ingest"

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
