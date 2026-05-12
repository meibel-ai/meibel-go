package v2

import (
	"context"
	"fmt"
)

// SessionsService handles Sessions operations.
type SessionsService struct {
	client *MeibelClient
}

// Get Get Session
func (s *SessionsService) Get(ctx context.Context, sessionId string) (*AgentExecutionDetailsResponse, error) {
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

// GetMessages Get Session Messages
func (s *SessionsService) GetMessages(ctx context.Context, sessionId string) (*SessionMessagesResponse, error) {
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
