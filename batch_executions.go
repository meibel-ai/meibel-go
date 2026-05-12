package v2

import (
	"context"
	"fmt"
	"net/url"
)

// BatchExecutionsService handles BatchExecutions operations.
type BatchExecutionsService struct {
	client *MeibelClient
}

// BatchExecutionsListOptions contains optional parameters for List.
type BatchExecutionsListOptions struct {
	// Filter by input datasource ID
	InputDatasourceId interface{}
	Offset *int64
	Limit interface{}
	// Field to sort by: start_time, status
	SortBy *string
	SortOrder *string
}

// List List Batch Executions
func (s *BatchExecutionsService) List(ctx context.Context, opts *BatchExecutionsListOptions) *PageIterator[BatchExecutionResponse] {
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

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[BatchExecutionResponse], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Data []BatchExecutionResponse `json:"data"`
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

		return &Page[BatchExecutionResponse]{
			Items:      resp.Data,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// Create Create Batch Execution
func (s *BatchExecutionsService) Create(ctx context.Context, body CreateBatchExecutionRequest) (*BatchExecutionResponse, error) {
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
func (s *BatchExecutionsService) GetById(ctx context.Context, executionId string) (*BatchExecutionResponse, error) {
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
func (s *BatchExecutionsService) UpdateById(ctx context.Context, executionId string, body UpdateBatchExecutionRequest) (*BatchExecutionResponse, error) {
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

// GetBatchRealtimeProgress Get Batch Realtime Progress
func (s *BatchExecutionsService) GetBatchRealtimeProgress(ctx context.Context, executionId string) (*string, error) {
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
func (s *BatchExecutionsService) RetryFailedItems(ctx context.Context, executionId string) (*BatchExecutionResponse, error) {
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
func (s *BatchExecutionsService) Cancel(ctx context.Context, executionId string) (*BatchExecutionResponse, error) {
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
