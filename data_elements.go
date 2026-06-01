package v2

import (
	"context"
	"fmt"
	"net/url"
)

// DataElementsService handles DataElements operations.
type DataElementsService struct {
	client *MeibelClient
}

// DataElementsListOptions contains optional parameters for List.
type DataElementsListOptions struct {
	// Cursor for pagination
	Cursor interface{}
	// Maximum items to return
	Limit *int64
}

// DataElementsSearchOptions contains optional parameters for Search.
type DataElementsSearchOptions struct {
	// Cursor for pagination
	Cursor interface{}
	// Maximum items to return
	Limit *int64
}

// Get Get Data Element
func (s *DataElementsService) Get(ctx context.Context, dataElementId string, datasourceId string) (*DataElementResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId)

	var result DataElementResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update Update Data Element
func (s *DataElementsService) Update(ctx context.Context, dataElementId string, datasourceId string, body UpdateDataElementRequest) (*DataElementResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId)

	var result DataElementResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "PUT",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// List List Data Elements
func (s *DataElementsService) List(ctx context.Context, datasourceId string, opts *DataElementsListOptions) *PageIterator[DataElementResponse] {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements"
	query := url.Values{}
	if opts != nil && opts.Cursor != nil {
		query.Set("cursor", fmt.Sprintf("%v", opts.Cursor))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[DataElementResponse], error) {
		if cursor != "" {
			query.Set("cursor", cursor)
		}

		var resp struct {
			Items []DataElementResponse `json:"items"`
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

		return &Page[DataElementResponse]{
			Items:      resp.Items,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// Search Search Data Elements
func (s *DataElementsService) Search(ctx context.Context, datasourceId string, body DataElementSearchRequest, opts *DataElementsSearchOptions) (*DataElementListResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/search"
	query := url.Values{}
	if opts != nil && opts.Cursor != nil {
		query.Set("cursor", fmt.Sprintf("%v", opts.Cursor))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}

	var result DataElementListResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Query:  query,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
