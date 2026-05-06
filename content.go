package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// ContentService handles content operations.
type ContentService struct {
	client *MeibelgoClient
}

// ListDatasourceContentOptions contains optional parameters for ListDatasourceContent.
type ListDatasourceContentOptions struct {
	// Filter content by path prefix
	Prefix interface{}
	// Token for pagination to get next page of results
	ContinuationToken interface{}
	// Maximum number of items to return (1-10000)
	Limit *int64
}

// ListDatasourceContent List datasource content
//
// List files and directories in a datasource with optional filtering and pagination
func (s *ContentService) ListDatasourceContent(ctx context.Context, datasourceId string, opts *ListDatasourceContentOptions) (*ListContentResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/content"
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

	var result ListContentResponse
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

// UploadDatasourceContent Upload Content
//
// Proxy upload with zero-copy streaming.
// 
// This endpoint maintains the multipart form data structure and streams
// it directly to the backend service without buffering files in memory.
// The multipart parsing happens on the backend service side.
func (s *ContentService) UploadDatasourceContent(ctx context.Context, datasourceId string) (*UploadContentResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/content"

	var result UploadContentResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
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

// GetDatasourceUploadStatus Get upload status
//
// Get the current status of a content upload operation
func (s *ContentService) GetDatasourceUploadStatus(ctx context.Context, datasourceId string, uploadId string) (*string, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/content/upload-status/" + fmt.Sprintf("%v", uploadId)

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

// StreamDatasourceUploadProgress Stream upload progress events (legacy)
//
// Subscribe to real-time upload progress updates via Server-Sent Events. Consider using /uploads/{upload_id}/progress instead.
func (s *ContentService) StreamDatasourceUploadProgress(ctx context.Context, datasourceId string, uploadId string) error {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/content/upload-progress/" + fmt.Sprintf("%v", uploadId)

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetDatasourceContentMetadata Get content metadata
//
// Get metadata information for a file or directory in the datasource
func (s *ContentService) GetDatasourceContentMetadata(ctx context.Context, datasourceId string, path string) (*GetContentResponse, error) {
	reqPath := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/content/" + fmt.Sprintf("%v", path)

	var result GetContentResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   reqPath,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteDatasourceContent Delete content
//
// Delete a file or directory from the datasource
func (s *ContentService) DeleteDatasourceContent(ctx context.Context, datasourceId string, path string) (*DeleteContentResponse, error) {
	reqPath := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/content/" + fmt.Sprintf("%v", path)

	var result DeleteContentResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   reqPath,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DownloadDatasourceContent Download content file
//
// Download a file from the datasource with streaming support for large files
func (s *ContentService) DownloadDatasourceContent(ctx context.Context, datasourceId string, path string) error {
	reqPath := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/content/" + fmt.Sprintf("%v", path) + "/download"

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   reqPath,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
