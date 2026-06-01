package v2

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

// FileUploadsService handles FileUploads operations.
type FileUploadsService struct {
	client *MeibelClient
}

// FileUploadsListContentOptions contains optional parameters for ListContent.
type FileUploadsListContentOptions struct {
	// Filter content by path prefix
	Prefix interface{}
	// Token for pagination
	ContinuationToken interface{}
	// Maximum items to return
	Limit *int64
}

// FileUploadsUploadAndListContentOptions contains optional parameters for UploadAndListContent.
type FileUploadsUploadAndListContentOptions struct {
	// Start ingestion after upload completes. Returns ingest_url to poll for status.
	TriggerIngest *bool
}

// ListContent List Content
func (s *FileUploadsService) ListContent(ctx context.Context, datasourceId string, opts *FileUploadsListContentOptions) *PageIterator[ContentItem] {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content"
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

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[ContentItem], error) {
		if cursor != "" {
			query.Set("continuation_token", cursor)
		}

		var resp struct {
			Items []ContentItem `json:"items"`
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

		return &Page[ContentItem]{
			Items:      resp.Items,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// UploadContent Upload Content (async)
func (s *FileUploadsService) UploadContent(ctx context.Context, datasourceId string, files io.Reader, filesName string) (*UploadContentResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content"

	uploadFields := []UploadField{
		{FieldName: "files", Reader: files, FileName: filesName},
	}
	formFields := map[string]string{}

	var result UploadContentResponse
	err := s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, uploadFields, formFields, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UploadAndListContent Upload Content (sync)
func (s *FileUploadsService) UploadAndListContent(ctx context.Context, datasourceId string, files io.Reader, filesName string, opts *FileUploadsUploadAndListContentOptions) (*FileUploadSyncResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content/process"
	query := url.Values{}
	if opts != nil && opts.TriggerIngest != nil {
		query.Set("trigger_ingest", fmt.Sprintf("%v", *opts.TriggerIngest))
	}

	uploadFields := []UploadField{
		{FieldName: "files", Reader: files, FileName: filesName},
	}
	formFields := map[string]string{}

	var result FileUploadSyncResponse
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

// StreamUploadProgress Stream Upload Progress
func (s *FileUploadsService) StreamUploadProgress(ctx context.Context, uploadId string) (*EventStream[interface{}], error) {
	path := "/datasources/uploads/" + fmt.Sprintf("%v", uploadId) + "/progress"

	resp, err := s.client.http.DoStream(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	})
	if err != nil {
		return nil, err
	}

	return JSONEventStream[interface{}](resp), nil
}
