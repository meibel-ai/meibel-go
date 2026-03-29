package meibelgo

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

// ContentService handles Content operations.
type ContentService struct {
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

// ListContent List Content
func (s *ContentService) ListContent(ctx context.Context, datasourceId string, opts *ListContentOptions) *PageIterator[string] {
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

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[string], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Items []string `json:"items"`
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

		return &Page[string]{
			Items:      resp.Items,
			NextCursor: resp.NextCursor,
		}, nil
	})
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

// GetContentMetadata Get Content Metadata
func (s *ContentService) GetContentMetadata(ctx context.Context, datasourceId string, path string) (*string, error) {
	reqPath := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content/" + fmt.Sprintf("%v", path) + "/metadata"

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   reqPath,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DownloadContent Download Content
func (s *ContentService) DownloadContent(ctx context.Context, datasourceId string, path string) error {
	reqPath := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content/" + fmt.Sprintf("%v", path) + "/download"

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   reqPath,
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

// DeleteContent Delete Content
func (s *ContentService) DeleteContent(ctx context.Context, datasourceId string, path string) (*string, error) {
	reqPath := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/content/" + fmt.Sprintf("%v", path)

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   reqPath,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
