package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// TableDescriptionsService handles TableDescriptions operations.
type TableDescriptionsService struct {
	client *MeibelClient
}

// ListTablesOptions contains optional parameters for ListTables.
type ListTablesOptions struct {
	// Include columns for each table
	IncludeColumns *bool
}

// ListTables List Tables
func (s *TableDescriptionsService) ListTables(ctx context.Context, datasourceId string, opts *ListTablesOptions) (*[]TagTable, error) {
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

// UpdateTableDescriptions Update Table Descriptions
func (s *TableDescriptionsService) UpdateTableDescriptions(ctx context.Context, datasourceId string, body []TagTableUpdateItem) (*[]TagTable, error) {
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
func (s *TableDescriptionsService) ListColumns(ctx context.Context, datasourceId string, tableName string) (*[]TagColumn, error) {
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
func (s *TableDescriptionsService) UpdateColumnDescriptions(ctx context.Context, datasourceId string, tableName string, body []TagColumnUpdateItem) (*[]TagColumn, error) {
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
