package meibelgo

import (
	"context"
	"fmt"
)

// BlueprintsExecutionsService handles blueprints.executions operations.
type BlueprintsExecutionsService struct {
	client *MeibelgoClient
}

// StartBlueprintInstance Start Blueprint Instance
func (s *BlueprintsExecutionsService) StartBlueprintInstance(ctx context.Context, blueprintInstanceId string, body *interface{}) (*string, error) {
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
func (s *BlueprintsExecutionsService) CancelBlueprintInstance(ctx context.Context, blueprintInstanceId string) (*string, error) {
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
func (s *BlueprintsExecutionsService) SendSignal(ctx context.Context, blueprintInstanceId string, signalName string, body *interface{}) (*string, error) {
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
func (s *BlueprintsExecutionsService) QueryWorkflow(ctx context.Context, blueprintInstanceId string, queryName string, body *interface{}) (*string, error) {
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
func (s *BlueprintsExecutionsService) GetBlueprintInstanceWorkflowStatus(ctx context.Context, blueprintInstanceId string) (*string, error) {
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
func (s *BlueprintsExecutionsService) SendChatMessage(ctx context.Context, blueprintInstanceId string, body ChatMessageRequest) (*ChatMessageResponse, error) {
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

// SendChatMessageStream Send a chat message and stream the response via SSE
//
// Send a chat message to a running chat agent workflow and stream the response as Server-Sent Events.
func (s *BlueprintsExecutionsService) SendChatMessageStream(ctx context.Context, blueprintInstanceId string, body ChatMessageRequest) error {
	path := "/" + fmt.Sprintf("%v", blueprintInstanceId) + "/chat/stream"

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
