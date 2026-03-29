package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// ConfidenceScoringService handles confidence_scoring operations.
type ConfidenceScoringService struct {
	client *MeibelgoClient
}

// GetAllScoringJobsOptions contains optional parameters for GetAllScoringJobs.
type GetAllScoringJobsOptions struct {
	AgentName interface{}
	AgentVersion interface{}
	AgentExecutionId interface{}
	AgentWorkflowName interface{}
	AgentWorkflowVersion interface{}
	AgentWorkflowExecutionId interface{}
	ToolId interface{}
	ToolInstanceId interface{}
	ToolExecutionId interface{}
}

// GetScoringJobsSummaryOptions contains optional parameters for GetScoringJobsSummary.
type GetScoringJobsSummaryOptions struct {
	Secondary interface{}
}

// GetScoringJob Get Scoring Job
//
// Get a scoring job by ID. Returns 403 if the job does not belong to the caller's customer.
func (s *ConfidenceScoringService) GetScoringJob(ctx context.Context, jobId string) (*string, error) {
	path := "/confidence-scoring/job/" + fmt.Sprintf("%v", jobId)

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

// GetAllScoringJobs Get All Scoring Jobs
//
// Get all scoring jobs for the caller's customer.
func (s *ConfidenceScoringService) GetAllScoringJobs(ctx context.Context, opts *GetAllScoringJobsOptions) (*string, error) {
	path := "/confidence-scoring/jobs"
	query := url.Values{}
	if opts != nil && opts.AgentName != nil {
		query.Set("agent_name", fmt.Sprintf("%v", opts.AgentName))
	}
	if opts != nil && opts.AgentVersion != nil {
		query.Set("agent_version", fmt.Sprintf("%v", opts.AgentVersion))
	}
	if opts != nil && opts.AgentExecutionId != nil {
		query.Set("agent_execution_id", fmt.Sprintf("%v", opts.AgentExecutionId))
	}
	if opts != nil && opts.AgentWorkflowName != nil {
		query.Set("agent_workflow_name", fmt.Sprintf("%v", opts.AgentWorkflowName))
	}
	if opts != nil && opts.AgentWorkflowVersion != nil {
		query.Set("agent_workflow_version", fmt.Sprintf("%v", opts.AgentWorkflowVersion))
	}
	if opts != nil && opts.AgentWorkflowExecutionId != nil {
		query.Set("agent_workflow_execution_id", fmt.Sprintf("%v", opts.AgentWorkflowExecutionId))
	}
	if opts != nil && opts.ToolId != nil {
		query.Set("tool_id", fmt.Sprintf("%v", opts.ToolId))
	}
	if opts != nil && opts.ToolInstanceId != nil {
		query.Set("tool_instance_id", fmt.Sprintf("%v", opts.ToolInstanceId))
	}
	if opts != nil && opts.ToolExecutionId != nil {
		query.Set("tool_execution_id", fmt.Sprintf("%v", opts.ToolExecutionId))
	}

	var result string
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

// GetScoringJobsSummary Get Scoring Jobs Summary
//
// Get aggregated scoring summary for the caller's customer.
// 
// primary: Required filter in format 'field:value' (e.g., 'agent_execution_id:exec_123').
// secondary: Optional secondary filter in format 'field:value' (e.g., 'agent_name:my_agent').
// Results are always scoped to the caller's customer_id.
func (s *ConfidenceScoringService) GetScoringJobsSummary(ctx context.Context, primary string, opts *GetScoringJobsSummaryOptions) (*ScoreSummary, error) {
	path := "/confidence-scoring/summary"
	query := url.Values{}
	query.Set("primary", fmt.Sprintf("%v", primary))
	if opts != nil && opts.Secondary != nil {
		query.Set("secondary", fmt.Sprintf("%v", opts.Secondary))
	}

	var result ScoreSummary
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
