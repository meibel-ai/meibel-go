package v2

import (
	"context"
	"fmt"
)

// PromptsService handles Prompts operations.
type PromptsService struct {
	client *MeibelClient
}

// List List Prompts
func (s *PromptsService) List(ctx context.Context) (*PromptListResponse, error) {
	path := "/prompts/"

	var result PromptListResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Create Create Prompt
func (s *PromptsService) Create(ctx context.Context, body CreateAgentPromptRequest) (*CreatePromptResponse, error) {
	path := "/prompts/"

	var result CreatePromptResponse
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

// Get Get Prompt
func (s *PromptsService) Get(ctx context.Context, promptId string) (*PromptResponse, error) {
	path := "/prompts/" + fmt.Sprintf("%v", promptId)

	var result PromptResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update Update Prompt
func (s *PromptsService) Update(ctx context.Context, promptId string, body UpdateAgentPromptRequest) (*UpdatePromptResponse, error) {
	path := "/prompts/" + fmt.Sprintf("%v", promptId)

	var result UpdatePromptResponse
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

// Delete Delete Prompt
func (s *PromptsService) Delete(ctx context.Context, promptId string) error {
	path := "/prompts/" + fmt.Sprintf("%v", promptId)

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
