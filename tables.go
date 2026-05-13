package v2

import (
	"context"
	"fmt"
	"net/url"
)

// TablesService handles Tables operations.
type TablesService struct {
	client *MeibelClient
}

// TablesListOptions contains optional parameters for List.
type TablesListOptions struct {
	// Include columns for each table
	IncludeColumns *bool
}

// List List Tables
func (s *TablesService) List(ctx context.Context, datasourceId string, opts *TablesListOptions) (*[]TagTable, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tables"
	query := url.Values{}
	if opts != nil && opts.IncludeColumns != nil {
		query.Set("include_columns", fmt.Sprintf("%v", *opts.IncludeColumns))
	}

	var result []TagTable
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

// UpdateDescriptions Update Table Descriptions
func (s *TablesService) UpdateDescriptions(ctx context.Context, datasourceId string, body UpdateTagTablesRequest) (*[]TagTable, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tables"

	var result []TagTable
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

// ListColumns List Columns
func (s *TablesService) ListColumns(ctx context.Context, datasourceId string, tableName string) (*[]TagColumn, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tables/" + fmt.Sprintf("%v", tableName) + "/columns"

	var result []TagColumn
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateColumnDescriptions Update Column Descriptions
func (s *TablesService) UpdateColumnDescriptions(ctx context.Context, datasourceId string, tableName string, body UpdateTagColumnsRequest) (*[]TagColumn, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/tables/" + fmt.Sprintf("%v", tableName) + "/columns"

	var result []TagColumn
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
