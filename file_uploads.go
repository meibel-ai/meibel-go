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

// UploadContent Upload Content (async)
func (s *FileUploadsService) UploadContent(ctx context.Context, files io.Reader, filesName string, datasourceId string, name string, description string, metadataConfig MetadataConfigRequest) (*UploadContentResponse, error) {
	path := "/datasources/uploads"

	formFields := map[string]string{
		"datasource_id": fmt.Sprintf("%v", datasourceId),
		"name": fmt.Sprintf("%v", name),
		"description": fmt.Sprintf("%v", description),
		"metadata_config": fmt.Sprintf("%v", metadataConfig),
	}

	var result UploadContentResponse
	uploadFields := []UploadField{
		{FieldName: "files", Reader: files, FileName: filesName},
	}

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
func (s *FileUploadsService) UploadAndListContent(ctx context.Context, files io.Reader, filesName string, datasourceId string, name string, description string, metadataConfig MetadataConfigRequest, triggerIngest bool) (*FileUploadSyncResponse, error) {
	path := "/datasources/uploads/process"

	formFields := map[string]string{
		"datasource_id": fmt.Sprintf("%v", datasourceId),
		"name": fmt.Sprintf("%v", name),
		"description": fmt.Sprintf("%v", description),
		"metadata_config": fmt.Sprintf("%v", metadataConfig),
		"trigger_ingest": fmt.Sprintf("%v", triggerIngest),
	}

	var result FileUploadSyncResponse
	uploadFields := []UploadField{
		{FieldName: "files", Reader: files, FileName: filesName},
	}

	err := s.client.http.DoUpload(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
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
