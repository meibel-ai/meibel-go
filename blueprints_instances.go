package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// BlueprintsInstancesService handles blueprints.instances operations.
type BlueprintsInstancesService struct {
	client *MeibelgoClient
}

// GetAllBlueprintInstancesOptions contains optional parameters for GetAllBlueprintInstances.
type GetAllBlueprintInstancesOptions struct {
	IncludeChildren *bool
	IncludeActivities *bool
	IncludeEvents *bool
}

// GetBlueprintInstanceOptions contains optional parameters for GetBlueprintInstance.
type GetBlueprintInstanceOptions struct {
	IncludeChildren *bool
	IncludeActivities *bool
	IncludeEvents *bool
}

// UpdateBlueprintInstanceStatusOptions contains optional parameters for UpdateBlueprintInstanceStatus.
type UpdateBlueprintInstanceStatusOptions struct {
	WorkflowRunId interface{}
}

// GetAllBlueprintInstances Get All Blueprint Instances
func (s *BlueprintsInstancesService) GetAllBlueprintInstances(ctx context.Context, opts *GetAllBlueprintInstancesOptions) (*GetBlueprintInstancesResponse, error) {
	path := "/blueprint-instance/"
	query := url.Values{}
	if opts != nil && opts.IncludeChildren != nil {
		query.Set("include_children", fmt.Sprintf("%v", *opts.IncludeChildren))
	}
	if opts != nil && opts.IncludeActivities != nil {
		query.Set("include_activities", fmt.Sprintf("%v", *opts.IncludeActivities))
	}
	if opts != nil && opts.IncludeEvents != nil {
		query.Set("include_events", fmt.Sprintf("%v", *opts.IncludeEvents))
	}

	var result GetBlueprintInstancesResponse
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

// AddBlueprintInstance Add Blueprint Instance
func (s *BlueprintsInstancesService) AddBlueprintInstance(ctx context.Context, body AddBlueprintInstanceRequest) (*AddBlueprintInstanceResponse, error) {
	path := "/blueprint-instance/"

	var result AddBlueprintInstanceResponse
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

// GetBlueprintInstance Get Blueprint Instance
func (s *BlueprintsInstancesService) GetBlueprintInstance(ctx context.Context, blueprintInstanceId string, opts *GetBlueprintInstanceOptions) (*GetBlueprintInstancesResponse, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId)
	query := url.Values{}
	if opts != nil && opts.IncludeChildren != nil {
		query.Set("include_children", fmt.Sprintf("%v", *opts.IncludeChildren))
	}
	if opts != nil && opts.IncludeActivities != nil {
		query.Set("include_activities", fmt.Sprintf("%v", *opts.IncludeActivities))
	}
	if opts != nil && opts.IncludeEvents != nil {
		query.Set("include_events", fmt.Sprintf("%v", *opts.IncludeEvents))
	}

	var result GetBlueprintInstancesResponse
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

// DeleteBlueprintInstance Delete Blueprint Instance
func (s *BlueprintsInstancesService) DeleteBlueprintInstance(ctx context.Context, blueprintInstanceId string) error {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId)

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBlueprintInstanceStatus Update Blueprint Instance Status
func (s *BlueprintsInstancesService) UpdateBlueprintInstanceStatus(ctx context.Context, blueprintInstanceId string, updatedStatusValue BlueprintInstanceStatus, opts *UpdateBlueprintInstanceStatusOptions) error {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/status"
	query := url.Values{}
	query.Set("updated_status_value", fmt.Sprintf("%v", updatedStatusValue))
	if opts != nil && opts.WorkflowRunId != nil {
		query.Set("workflow_run_id", fmt.Sprintf("%v", opts.WorkflowRunId))
	}

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "PUT",
		Path:   path,
		Query:  query,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// CompleteBlueprintInstance Complete a Blueprint Instance
//
// This endpoint is used to mark a Blueprint Instance as completed. It will update the status of the Blueprint Instance to 'COMPLETED' and log the completion event.
func (s *BlueprintsInstancesService) CompleteBlueprintInstance(ctx context.Context, blueprintInstanceId string, body *interface{}) (*string, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/complete-instance"

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

// FailBlueprintInstance Fail a Blueprint Instance
//
// This endpoint is used to mark a Blueprint Instance as failed. It will update the status of the Blueprint Instance to 'FAILED' and log the failure event.
func (s *BlueprintsInstancesService) FailBlueprintInstance(ctx context.Context, blueprintInstanceId string, body FailBlueprintInstanceRequest) (*string, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/fail-instance"

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

// AddActivityByBlueprintInstance Add Activity By Blueprint Instance
func (s *BlueprintsInstancesService) AddActivityByBlueprintInstance(ctx context.Context, blueprintInstanceId string, body AddActivityRequest) (*AddActivityResponse, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/activity"

	var result AddActivityResponse
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

// GetActivityByBlueprintInstance Get Activity By Blueprint Instance
func (s *BlueprintsInstancesService) GetActivityByBlueprintInstance(ctx context.Context, blueprintInstanceId string, activityId string) (*GetActivitiesResponse, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/activity/" + fmt.Sprintf("%v", activityId)

	var result GetActivitiesResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetActivitiesByBlueprintInstance Get Activities By Blueprint Instance
func (s *BlueprintsInstancesService) GetActivitiesByBlueprintInstance(ctx context.Context, blueprintInstanceId string) (*GetActivitiesResponse, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/activities"

	var result GetActivitiesResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateActivityStatus Update Activity Status
func (s *BlueprintsInstancesService) UpdateActivityStatus(ctx context.Context, blueprintInstanceId string, activityId string, updatedStatusValue ActivityStatus) error {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/activity/" + fmt.Sprintf("%v", activityId) + "/status"
	query := url.Values{}
	query.Set("updated_status_value", fmt.Sprintf("%v", updatedStatusValue))

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "PUT",
		Path:   path,
		Query:  query,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetEventByBlueprintInstance Get Event By Blueprint Instance
func (s *BlueprintsInstancesService) GetEventByBlueprintInstance(ctx context.Context, blueprintInstanceId string, eventId string) (*GetEventsResponse, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/event/" + fmt.Sprintf("%v", eventId)

	var result GetEventsResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateEventByBlueprintInstanceId Create Event By Blueprint Instance Id
func (s *BlueprintsInstancesService) CreateEventByBlueprintInstanceId(ctx context.Context, blueprintInstanceId string, body CustomEventRequest) (*AddEventResponse, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/event"

	var result AddEventResponse
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

// GetEventsByBlueprintInstance Get Events By Blueprint Instance
func (s *BlueprintsInstancesService) GetEventsByBlueprintInstance(ctx context.Context, blueprintInstanceId string) (*GetEventsResponse, error) {
	path := "/blueprint-instance/" + fmt.Sprintf("%v", blueprintInstanceId) + "/events"

	var result GetEventsResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
