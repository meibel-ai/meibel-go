package v2

import (
	"context"
	"fmt"
)

// ExecutionPoliciesService handles ExecutionPolicies operations.
type ExecutionPoliciesService struct {
	client *MeibelClient
}

// List List Execution Policies
func (s *ExecutionPoliciesService) List(ctx context.Context) (*GetExecutionPoliciesResponse, error) {
	path := "/execution-policies"

	var result GetExecutionPoliciesResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Create Create Execution Policy
func (s *ExecutionPoliciesService) Create(ctx context.Context, body CreateExecutionPolicyRequest) (*ExecutionPolicyResponse, error) {
	path := "/execution-policies"

	var result ExecutionPolicyResponse
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

// Get Get Execution Policy
func (s *ExecutionPoliciesService) Get(ctx context.Context, policyId string) (*ExecutionPolicyResponse, error) {
	path := "/execution-policies/" + fmt.Sprintf("%v", policyId)

	var result ExecutionPolicyResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update Update Execution Policy
func (s *ExecutionPoliciesService) Update(ctx context.Context, policyId string, body UpdateExecutionPolicyRequest) (*ExecutionPolicyResponse, error) {
	path := "/execution-policies/" + fmt.Sprintf("%v", policyId)

	var result ExecutionPolicyResponse
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

// Delete Delete Execution Policy
func (s *ExecutionPoliciesService) Delete(ctx context.Context, policyId string) error {
	path := "/execution-policies/" + fmt.Sprintf("%v", policyId)

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
