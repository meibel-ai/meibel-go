package meibelgo

import (
	"context"
	"fmt"
)

// SessionsService handles Sessions operations.
type SessionsService struct {
	client *MeibelClient
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
