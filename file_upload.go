package meibelgo

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

// FileUploadService handles File Upload operations.
type FileUploadService struct {
	client *MeibelgoClient
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

// UploadContent Upload Content (async)
func (s *FileUploadService) UploadContent(ctx context.Context, file io.Reader, fileName string) (*UploadContentResponse, error) {
	path := "/datasources/uploads"

	var result UploadContentResponse
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

// UploadAndListContent Upload Content (sync)
func (s *FileUploadService) UploadAndListContent(ctx context.Context, file io.Reader, fileName string) (*FileUploadSyncResponse, error) {
	path := "/datasources/uploads/process"

	var result FileUploadSyncResponse
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
func (s *FileUploadService) StreamUploadProgress(ctx context.Context, uploadId string) (*EventStream[interface{}], error) {
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

// ListContent List Content
func (s *FileUploadService) ListContent(ctx context.Context, datasourceId string, opts *ListContentOptions) *PageIterator[ContentItem] {
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
