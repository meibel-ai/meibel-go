package meibelgo

import (
	"context"
	"fmt"
)

// BlueprintsService handles blueprints operations.
type BlueprintsService struct {
	client *MeibelgoClient
	Executions *ExecutionsService
	Instances *InstancesService
}

// GetBlueprints Get Blueprints
func (s *BlueprintsService) GetBlueprints(ctx context.Context) (*GetBlueprintsResponse, error) {
	path := "/blueprint/"

	var result GetBlueprintsResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateBlueprint Create Blueprint
func (s *BlueprintsService) CreateBlueprint(ctx context.Context, body AddBlueprintRequest) (*AddBlueprintResponse, error) {
	path := "/blueprint/"

	var result AddBlueprintResponse
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

// GetBlueprint Get Blueprint
func (s *BlueprintsService) GetBlueprint(ctx context.Context, blueprintId string) (*Blueprint, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId)

	var result Blueprint
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateBlueprint Update Blueprint
func (s *BlueprintsService) UpdateBlueprint(ctx context.Context, blueprintId string, body UpdateBlueprintRequest) (*Blueprint, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId)

	var result Blueprint
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

// DeleteBlueprint Delete Blueprint
func (s *BlueprintsService) DeleteBlueprint(ctx context.Context, blueprintId string) (*string, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId)

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ExecuteBlueprint Execute Blueprint
func (s *BlueprintsService) ExecuteBlueprint(ctx context.Context, blueprintId string, body *ExecuteBlueprintRequest) (*string, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId) + "/execute"

	var result string
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

// GetBlueprintTasks Get Blueprint Tasks
func (s *BlueprintsService) GetBlueprintTasks(ctx context.Context, blueprintId string) (*string, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId) + "/task"

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

// CreateBlueprintTask Create Blueprint Task
func (s *BlueprintsService) CreateBlueprintTask(ctx context.Context, blueprintId string, body AddBlueprintTaskRequest) (*string, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId) + "/task"

	var result string
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

// UpdateBlueprintTask Update Blueprint Task
func (s *BlueprintsService) UpdateBlueprintTask(ctx context.Context, blueprintId string, taskId string, body UpdateBlueprintTaskRequest) (*string, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId) + "/task/" + fmt.Sprintf("%v", taskId)

	var result string
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

// DeleteBlueprintTask Delete Blueprint Task
func (s *BlueprintsService) DeleteBlueprintTask(ctx context.Context, blueprintId string, taskId string) (*string, error) {
	path := "/blueprint/" + fmt.Sprintf("%v", blueprintId) + "/task/" + fmt.Sprintf("%v", taskId)

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
