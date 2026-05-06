package meibelgo

import (
	"context"
	"fmt"
)

// ExecutionsService handles executions operations.
type ExecutionsService struct {
	client *MeibelgoClient
}

// StartBlueprintInstance Start Blueprint Instance
func (s *ExecutionsService) StartBlueprintInstance(ctx context.Context, blueprintInstanceId string, body *interface{}) (*string, error) {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/start-instance"

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

// CancelBlueprintInstance Cancel Blueprint Instance
func (s *ExecutionsService) CancelBlueprintInstance(ctx context.Context, blueprintInstanceId string) (*string, error) {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/cancel-instance"

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

// SendSignal Send Signal
func (s *ExecutionsService) SendSignal(ctx context.Context, blueprintInstanceId string, signalName string, body *interface{}) (*string, error) {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/signals/" + fmt.Sprintf("%v", signalName)

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

// QueryWorkflow Query Workflow
func (s *ExecutionsService) QueryWorkflow(ctx context.Context, blueprintInstanceId string, queryName string, body *interface{}) (*string, error) {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/queries/" + fmt.Sprintf("%v", queryName)

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

// GetBlueprintInstanceWorkflowStatus Get Blueprint Instance Workflow Status
func (s *ExecutionsService) GetBlueprintInstanceWorkflowStatus(ctx context.Context, blueprintInstanceId string) (*string, error) {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/workflow-status"

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

// SendChatMessage Send Chat Message
func (s *ExecutionsService) SendChatMessage(ctx context.Context, blueprintInstanceId string, body ChatMessageRequest) (*ChatMessageResponse, error) {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/chat"

	var result ChatMessageResponse
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

// SendChatMessageStreamBlueprintInstanceIdChatStreamPost Send a chat message and stream the response via SSE
//
// Send a chat message to a running chat agent workflow and stream the response as Server-Sent Events.
func (s *ExecutionsService) SendChatMessageStreamBlueprintInstanceIdChatStreamPost(ctx context.Context, blueprintInstanceId string, body ChatMessageRequest) (*string, error) {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/chat/stream"

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
