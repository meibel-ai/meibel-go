package v2

import (
	"context"
	"fmt"
	"net/url"
)

// ExecutionsService handles Executions operations.
type ExecutionsService struct {
	client *MeibelClient
}

// ExecutionsListOptions contains optional parameters for List.
type ExecutionsListOptions struct {
	// Filter by input datasource ID
	InputDatasourceId interface{}
	Offset *int64
	Limit interface{}
	// Field to sort by: start_time, status
	SortBy *string
	SortOrder *string
}

// List List Batch Executions
func (s *ExecutionsService) List(ctx context.Context, opts *ExecutionsListOptions) (*GetBatchExecutionsResponse, error) {
	path := "/batch-executions/"
	query := url.Values{}
	if opts != nil && opts.InputDatasourceId != nil {
		query.Set("input_datasource_id", fmt.Sprintf("%v", opts.InputDatasourceId))
	}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", opts.Limit))
	}
	if opts != nil && opts.SortBy != nil {
		query.Set("sort_by", fmt.Sprintf("%v", *opts.SortBy))
	}
	if opts != nil && opts.SortOrder != nil {
		query.Set("sort_order", fmt.Sprintf("%v", *opts.SortOrder))
	}

	var result GetBatchExecutionsResponse
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

// Create Create Batch Execution
func (s *ExecutionsService) Create(ctx context.Context, body CreateBatchExecutionRequest) (*BatchExecutionResponse, error) {
	path := "/batch-executions/"

	var result BatchExecutionResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetById Get Batch Execution By Id
func (s *ExecutionsService) GetById(ctx context.Context, executionId string) (*BatchExecutionResponse, error) {
	path := "/batch-executions/id/" + fmt.Sprintf("%v", executionId)

	var result BatchExecutionResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateById Update Batch Execution By Id
func (s *ExecutionsService) UpdateById(ctx context.Context, executionId string, body UpdateBatchExecutionRequest) (*BatchExecutionResponse, error) {
	path := "/batch-executions/id/" + fmt.Sprintf("%v", executionId)

	var result BatchExecutionResponse
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

// GetRealtimeProgress Get Batch Realtime Progress
func (s *ExecutionsService) GetRealtimeProgress(ctx context.Context, executionId string) (*string, error) {
	path := "/batch-executions/id/" + fmt.Sprintf("%v", executionId) + "/realtime-progress"

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

// RetryFailedItems Retry Failed Items
func (s *ExecutionsService) RetryFailedItems(ctx context.Context, executionId string) (*BatchExecutionResponse, error) {
	path := "/batch-executions/id/" + fmt.Sprintf("%v", executionId) + "/retry-failed"

	var result BatchExecutionResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Cancel Cancel Batch Execution
func (s *ExecutionsService) Cancel(ctx context.Context, executionId string) (*BatchExecutionResponse, error) {
	path := "/batch-executions/id/" + fmt.Sprintf("%v", executionId) + "/cancel"

	var result BatchExecutionResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
