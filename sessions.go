package meibelgo

import (
	"context"
	"fmt"
)

// SessionsService handles Sessions operations.
type SessionsService struct {
	client *MeibelgoClient
}

// GetSession Get Session
func (s *SessionsService) GetSession(ctx context.Context, sessionId string) (*AgentExecutionDetailsResponse, error) {
	path := "/sessions/" + fmt.Sprintf("%v", sessionId)

	var result AgentExecutionDetailsResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetSessionMessages Get Session Messages
func (s *SessionsService) GetSessionMessages(ctx context.Context, sessionId string) (*SessionMessagesResponse, error) {
	path := "/sessions/" + fmt.Sprintf("%v", sessionId) + "/messages"

	var result SessionMessagesResponse
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
func (s *SessionsService) SendChatMessage(ctx context.Context, sessionId string, body ChatMessageRequest) (*ChatMessageResponse, error) {
	path := "/sessions/" + fmt.Sprintf("%v", sessionId) + "/chat"

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
func (s *SessionsService) SendChatMessageStream(ctx context.Context, sessionId string, body ChatMessageRequest) (*EventStream[interface{}], error) {
	path := "/sessions/" + fmt.Sprintf("%v", sessionId) + "/chat/stream"

	resp, err := s.client.http.DoStream(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	return JSONEventStream[interface{}](resp), nil
}
