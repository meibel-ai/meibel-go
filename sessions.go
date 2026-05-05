package meibelgo

import (
	"context"
	"fmt"
)

// SessionsService handles Sessions operations.
type SessionsService struct {
	client *MeibelClient
}

// SendChatMessage Send Chat Message
func (s *SessionsService) SendChatMessage(ctx context.Context, blueprintInstanceId string, body ChatMessageRequest) (*ChatMessageResponse, error) {
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
func (s *SessionsService) SendChatMessageStream(ctx context.Context, blueprintInstanceId string, body ChatMessageRequest) error {
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
