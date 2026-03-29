package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// TagDescriptionsService handles Tag Descriptions operations.
type TagDescriptionsService struct {
	client *MeibelgoClient
}

// ListTagTables List Tag Tables
func (s *TagDescriptionsService) ListTagTables(ctx context.Context, datasourceId string) *PageIterator[string] {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tag-tables"
	query := url.Values{}

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

// ListTagColumns List Tag Columns
func (s *TagDescriptionsService) ListTagColumns(ctx context.Context, datasourceId string, tableName string) *PageIterator[string] {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tag-tables/" + fmt.Sprintf("%v", tableName) + "/columns"
	query := url.Values{}

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

// UpdateTagTableDescription Update Tag Table Description
func (s *TagDescriptionsService) UpdateTagTableDescription(ctx context.Context, datasourceId string, tableName string, body UpdateTagDescriptionRequest) (*TagTable, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tag-tables/" + fmt.Sprintf("%v", tableName)

	var result TagTable
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

// UpdateTagColumnDescription Update Tag Column Description
func (s *TagDescriptionsService) UpdateTagColumnDescription(ctx context.Context, datasourceId string, tableName string, columnName string, body UpdateTagDescriptionRequest) (*TagColumn, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tag-tables/" + fmt.Sprintf("%v", tableName) + "/columns/" + fmt.Sprintf("%v", columnName)

	var result TagColumn
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
