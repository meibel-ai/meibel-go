package v2

import (
	"time"
)

// AgentDetailResponse represents the AgentDetailResponse type.
type AgentDetailResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	CatalogUrn string `json:"catalog_urn"`
	Version string `json:"version"`
	ParentVersion interface{} `json:"parent_version,omitempty"`
	Type string `json:"type"`
	Description interface{} `json:"description,omitempty"`
	LlmModel string `json:"llm_model"`
	FallbackModels []string `json:"fallback_models"`
	Datasources []string `json:"datasources"`
	Instructions string `json:"instructions"`
	Tools []map[string]interface{} `json:"tools"`
	Artifacts []string `json:"artifacts"`
	ConfidenceConfigs []string `json:"confidence_configs"`
	Temperature interface{} `json:"temperature"`
	MaxTokens interface{} `json:"max_tokens,omitempty"`
	Tags []string `json:"tags"`
	Icon interface{} `json:"icon,omitempty"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	LastExecutionStatus interface{} `json:"last_execution_status,omitempty"`
	LastExecutionTime interface{} `json:"last_execution_time,omitempty"`
}

// AgentExecutionDetailsResponse AgentExecutionDetailsResponse
type AgentExecutionDetailsResponse struct {
	AgentId interface{} `json:"agent_id"`
	AgentName interface{} `json:"agent_name"`
	Version interface{} `json:"version"`
	Status string `json:"status"`
	Messages []MessageEntry `json:"messages"`
	ToolActivity []ToolActivityEntry `json:"tool_activity"`
	TokenUsage []interface{} `json:"token_usage"`
	FileParsing []FileParseEntry `json:"file_parsing"`
	Result []ArtifactEntry `json:"result"`
}

// AgentIdentityContext Identity context for agent execution.  Contains only immutable identity fields that answer: - WHO: customer_id, project_id (tenant identity) - WHAT: agent_name, agent_version, agent_execution_id (agent identity, optional) - WHERE: agent_workflow_name, agent_workflow_version, agent_workflow_execution_id (parent workflow, optional) - WHICH TOOL: tool_id, tool_instance_id, tool_execution_id (tool identity, optional)  This model is FLAT - no inheritance, all fields in one model. Agent, workflow, and tool fields are optional, making this suitable for all execution contexts.  This model does NOT contain: - Configuration (see AgentExecutionConfig in agent-platform) - Runtime state (see AgentExecutionState in agent-platform)  Design Pattern - Progressive Enhancement: - Callers provide only tenant/project identity - FSMWorkflow fills in agent_workflow_* fields from AgentWorkflowSpec - ReactAgent fills in agent_* fields from AgentSpec and workflow.info().workflow_id - Tool activities add tool_* fields via model_copy() - Context gains fields as it flows through the system  Examples:     # Starting FSMWorkflow (caller provides minimal context)     context = AgentIdentityContext(         customer_id="cust_123",         project_id="proj_456"     )     # FSM fills in workflow identity     context = context.model_copy(update={         "agent_workflow_name": "support_fsm",         "agent_workflow_version": "3.0.0",         "agent_workflow_execution_id": workflow.info().workflow_id     })      # Starting ReactAgent standalone (caller provides minimal context)     context = AgentIdentityContext(         customer_id="cust_123",         project_id="proj_456"     )     # ReactAgent fills in agent identity     context = context.model_copy(update={         "agent_name": "sales_assistant",         "agent_version": "2.0.0",         "agent_execution_id": workflow.info().workflow_id     })      # ReactAgent as FSM child (inherits workflow context, adds agent identity)     child_context = parent_context.model_copy(update={         "agent_name": "router",         "agent_version": "1.0.0",         "agent_execution_id": workflow.info().workflow_id         # agent_workflow_* fields inherited from parent     })      # Tool execution (adds tool identity to agent context)     tool_context = context.model_copy(update={         'tool_id': "tool_xyz",         'tool_instance_id': "tool_inst_123",         'tool_execution_id': "tool_exec_456"     })
type AgentIdentityContext struct {
	// Customer/tenant identifier
	CustomerId string `json:"customer_id"`
	// Project identifier
	ProjectId string `json:"project_id"`
	AgentName interface{} `json:"agent_name,omitempty"`
	AgentVersion interface{} `json:"agent_version,omitempty"`
	AgentExecutionId interface{} `json:"agent_execution_id,omitempty"`
	AgentWorkflowName interface{} `json:"agent_workflow_name,omitempty"`
	AgentWorkflowVersion interface{} `json:"agent_workflow_version,omitempty"`
	AgentWorkflowExecutionId interface{} `json:"agent_workflow_execution_id,omitempty"`
	ToolId interface{} `json:"tool_id,omitempty"`
	ToolInstanceId interface{} `json:"tool_instance_id,omitempty"`
	ToolExecutionId interface{} `json:"tool_execution_id,omitempty"`
}

// AgentListResponse represents the AgentListResponse type.
type AgentListResponse struct {
	Data []AgentSummary `json:"data"`
	Total int64 `json:"total"`
}

// AgentSummary represents the AgentSummary type.
type AgentSummary struct {
	Id string `json:"id"`
	DisplayName string `json:"display_name"`
	Description interface{} `json:"description,omitempty"`
	LlmModel string `json:"llm_model"`
	ToolCount int64 `json:"tool_count"`
	DatasourceCount int64 `json:"datasource_count"`
	LastExecutionStatus interface{} `json:"last_execution_status,omitempty"`
	LastExecutionTime interface{} `json:"last_execution_time,omitempty"`
}

// AgentToolDefinition AgentToolDefinition
type AgentToolDefinition struct {
	// Instance name - what the LLM sees and calls
	Name string `json:"name"`
	// Tool type: rag_search, database_query, etc.
	Type string `json:"type"`
	// Description shown to LLM
	Description interface{} `json:"description,omitempty"`
	// Tool config passed to activity via tool_context (datasource_id, base_prompt, etc.)
	Config interface{} `json:"config,omitempty"`
	// Optional override for the tool's parameters schema
	ParametersSchema interface{} `json:"parameters_schema,omitempty"`
	// When to use this tool (injected into system prompt)
	UseFor interface{} `json:"use_for,omitempty"`
	// When NOT to use this tool (injected into system prompt)
	AvoidFor interface{} `json:"avoid_for,omitempty"`
	// If true, workflow pauses for human approval before executing this tool
	RequireApproval interface{} `json:"require_approval,omitempty"`
	// Message to display when requesting approval (supports {{variable}} templates)
	ApprovalMessage interface{} `json:"approval_message,omitempty"`
}

// AgentVersionListResponse represents the AgentVersionListResponse type.
type AgentVersionListResponse struct {
	Data []AgentVersionSummary `json:"data"`
	Total int64 `json:"total"`
}

// AgentVersionSummary represents the AgentVersionSummary type.
type AgentVersionSummary struct {
	Id string `json:"id"`
	DisplayName string `json:"display_name"`
	Version string `json:"version"`
	ParentVersion interface{} `json:"parent_version,omitempty"`
	Description interface{} `json:"description,omitempty"`
	LlmModel string `json:"llm_model"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	IsPublished bool `json:"is_published"`
	PublishedAt interface{} `json:"published_at,omitempty"`
	CommitMessage interface{} `json:"commit_message,omitempty"`
}

// Artifact A generated artifact/file from the chat agent.
type Artifact struct {
	ArtifactId string `json:"artifact_id"`
	Filename string `json:"filename"`
	MimeType string `json:"mime_type"`
	Content interface{} `json:"content,omitempty"`
	StorageUrl interface{} `json:"storage_url,omitempty"`
	SizeBytes interface{} `json:"size_bytes,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
}

// ArtifactEntry ArtifactEntry
type ArtifactEntry struct {
	Name string `json:"name"`
	Content interface{} `json:"content,omitempty"`
	FileType interface{} `json:"file_type"`
}

// ArtifactSchemaListResponse represents the ArtifactSchemaListResponse type.
type ArtifactSchemaListResponse struct {
	Data []ArtifactSchemaSummary `json:"data"`
	Total int64 `json:"total"`
}

// ArtifactSchemaResponse represents the ArtifactSchemaResponse type.
type ArtifactSchemaResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Version string `json:"version"`
	ParentVersion interface{} `json:"parent_version,omitempty"`
	Type string `json:"type"`
	Description string `json:"description"`
	Required bool `json:"required"`
	SchemaDef map[string]interface{} `json:"schema_def"`
	MaxSizeBytes interface{} `json:"max_size_bytes,omitempty"`
	StorageStrategy string `json:"storage_strategy"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
}

// ArtifactSchemaSummary represents the ArtifactSchemaSummary type.
type ArtifactSchemaSummary struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	FieldsSummary []FieldSummary `json:"fields_summary"`
}

// ArtifactStorageStrategy represents the possible values for ArtifactStorageStrategy.
type ArtifactStorageStrategy string

const (
	ArtifactStorageStrategyInline ArtifactStorageStrategy = "inline"
	ArtifactStorageStrategyGcs ArtifactStorageStrategy = "gcs"
	ArtifactStorageStrategyAuto ArtifactStorageStrategy = "auto"
)

// ArtifactType represents the possible values for ArtifactType.
type ArtifactType string

const (
	ArtifactTypeJson ArtifactType = "json"
	ArtifactTypeMarkdown ArtifactType = "markdown"
	ArtifactTypeCsv ArtifactType = "csv"
	ArtifactTypeYaml ArtifactType = "yaml"
	ArtifactTypeText ArtifactType = "text"
	ArtifactTypeHtml ArtifactType = "html"
	ArtifactTypePdf ArtifactType = "pdf"
)

// BatchDefinitionFilters Recipe-level filters. element_ids belongs here; per-execution overrides use BatchInputOverrides on the execution row.
type BatchDefinitionFilters struct {
	// Filter Data Elements by name pattern (regex)
	Regex interface{} `json:"regex,omitempty"`
	// Filter Data Elements by content type
	MediaTypes interface{} `json:"media_types,omitempty"`
	// Recipe-pinned subset of Data Element IDs.
	ElementIds interface{} `json:"element_ids,omitempty"`
}

// BatchDefinitionResponse Full BatchDefinition snapshot.
type BatchDefinitionResponse struct {
	Id string `json:"id"`
	CustomerId string `json:"customer_id"`
	ProjectId string `json:"project_id"`
	Name string `json:"name"`
	Version string `json:"version"`
	ParentVersion interface{} `json:"parent_version"`
	CatalogUrn string `json:"catalog_urn"`
	AgentUrn string `json:"agent_urn"`
	AgentSpecJson map[string]interface{} `json:"agent_spec_json"`
	InputDatasourceId string `json:"input_datasource_id"`
	// Optional override for the tool's parameters schema
	Filters interface{} `json:"filters,omitempty"`
	OutputDatasourceId interface{} `json:"output_datasource_id,omitempty"`
	UserMessage interface{} `json:"user_message,omitempty"`
	Concurrency int64 `json:"concurrency"`
	RetryLimit int64 `json:"retry_limit"`
	RecurrenceCron interface{} `json:"recurrence_cron,omitempty"`
	Description interface{} `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string `json:"created_by"`
	DeletedAt interface{} `json:"deleted_at,omitempty"`
}

// BatchExecutionResponse Response shape for a single batch execution. The legacy `batch_spec_json` / `agent_spec_json` / `agent_urn` / `input_datasource_id` fields are kept for client compatibility (DEL-1376 §5.5) — they are reconstructed from the linked BatchDefinition by the router, not stored on the execution row.
type BatchExecutionResponse struct {
	// Execution ID — also the Temporal workflow ID for direct queries
	Id string `json:"id"`
	// FK to the BatchDefinition this execution ran against
	BatchDefinitionId string `json:"batch_definition_id"`
	CustomerId string `json:"customer_id"`
	ProjectId string `json:"project_id"`
	AgentUrn interface{} `json:"agent_urn,omitempty"`
	// Optional override for the tool's parameters schema
	BatchSpecJson interface{} `json:"batch_spec_json,omitempty"`
	AgentSpecJson interface{} `json:"agent_spec_json,omitempty"`
	InputDatasourceId interface{} `json:"input_datasource_id,omitempty"`
	OutputDatasourceId interface{} `json:"output_datasource_id,omitempty"`
	// Optional override for the tool's parameters schema
	InputOverrides interface{} `json:"input_overrides,omitempty"`
	TotalItems interface{} `json:"total_items,omitempty"`
	Succeeded interface{} `json:"succeeded,omitempty"`
	Failed interface{} `json:"failed,omitempty"`
	StartTime time.Time `json:"start_time"`
	EndTime interface{} `json:"end_time,omitempty"`
	Status string `json:"status"`
	// Overall error message
	Error interface{} `json:"error,omitempty"`
	// Per-item results (populated on completion by status callback)
	Items interface{} `json:"items,omitempty"`
}

// BatchItemResult Per-item result from the Temporal workflow.
type BatchItemResult struct {
	InputDataElementId string `json:"input_data_element_id"`
	Filename string `json:"filename"`
	Status string `json:"status"`
	Error interface{} `json:"error,omitempty"`
	OutputArtifacts interface{} `json:"output_artifacts,omitempty"`
	Attempts interface{} `json:"attempts,omitempty"`
}

// BodySendChatMessageStream represents the Body_sendChatMessageStream type.
type BodySendChatMessageStream struct {
	UserMessage interface{} `json:"user_message,omitempty"`
	TimeoutSeconds interface{} `json:"timeout_seconds,omitempty"`
	IncludeThinking interface{} `json:"include_thinking,omitempty"`
	IncludeToolActivity interface{} `json:"include_tool_activity,omitempty"`
	Files interface{} `json:"files,omitempty"`
}

// BoundingBox represents the BoundingBox type.
type BoundingBox struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Page int64 `json:"page"`
}

// CallToAction An action the user can take.
type CallToAction struct {
	Label string `json:"label"`
	Action string `json:"action"`
	// Optional override for the tool's parameters schema
	ActionData interface{} `json:"action_data,omitempty"`
}

// ChatMessageRequest Request body for chat message endpoints.
type ChatMessageRequest struct {
	// The user's chat message
	UserMessage string `json:"user_message"`
	// Maximum time to wait for response (seconds)
	TimeoutSeconds interface{} `json:"timeout_seconds,omitempty"`
	// Whether to include thinking content in response
	IncludeThinking interface{} `json:"include_thinking,omitempty"`
	// Whether to include tool call/result activity
	IncludeToolActivity interface{} `json:"include_tool_activity,omitempty"`
}

// ChatMessageResponse Response from the non-streaming chat endpoint.
type ChatMessageResponse struct {
	// Unique ID for this message exchange
	SignalId string `json:"signal_id"`
	// The structured response
	Response ChatResponse `json:"response"`
	// The assistant response in text-format
	AssistantResponse string `json:"assistant_response"`
	// Tool calls made during response generation
	ToolActivity interface{} `json:"tool_activity,omitempty"`
	// LLM thinking/reasoning content
	Thinking interface{} `json:"thinking,omitempty"`
	// Token usage statistics
	TokenUsage interface{} `json:"token_usage,omitempty"`
}

// ChatResponse The structured chat response.
type ChatResponse struct {
	Message interface{} `json:"message,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
	FollowUpQuestions interface{} `json:"follow_up_questions,omitempty"`
	CallToActions interface{} `json:"call_to_actions,omitempty"`
	Artifacts interface{} `json:"artifacts,omitempty"`
}

// ChatWithDatasourceRequest represents the ChatWithDatasourceRequest type.
type ChatWithDatasourceRequest struct {
	// Datasources to query
	DatasourceIds []string `json:"datasource_ids"`
	// User question
	Message string `json:"message"`
	// LLM model override
	Model interface{} `json:"model,omitempty"`
}

// CloudStorageConnector Connect to a cloud storage bucket.
type CloudStorageConnector struct {
	// Cloud storage provider
	Provider string `json:"provider"`
	// Bucket name
	Bucket string `json:"bucket"`
	// Key prefix to scope the datasource
	Prefix interface{} `json:"prefix,omitempty"`
	// AWS IAM role ARN (S3 only)
	RoleArn interface{} `json:"role_arn,omitempty"`
	// AWS region (S3 only)
	Region interface{} `json:"region,omitempty"`
}

// CloudStorageConnectorProvider represents the possible values for provider.
type CloudStorageConnectorProvider string

const (
	CloudStorageConnectorProviderS3 CloudStorageConnectorProvider = "s3"
	CloudStorageConnectorProviderGcs CloudStorageConnectorProvider = "gcs"
)

// ConfidenceScoringConfig Simplified configuration wrapper that separates module name from config.  This model is shared between confidence-scoring-service and confidence-framework to ensure type consistency without OpenAPI Generator wrapper issues.
type ConfidenceScoringConfig struct {
	Module string `json:"module"`
	Config Config `json:"config"`
}

// Config Config
type Config struct {
	AnyofSchema_1Validator interface{} `json:"anyof_schema_1_validator,omitempty"`
	AnyofSchema_2Validator interface{} `json:"anyof_schema_2_validator,omitempty"`
	AnyofSchema_3Validator interface{} `json:"anyof_schema_3_validator,omitempty"`
	AnyofSchema_4Validator interface{} `json:"anyof_schema_4_validator,omitempty"`
	ActualInstance *string `json:"actual_instance,omitempty"`
	AnyOfSchemas []string `json:"any_of_schemas,omitempty"`
}

// ConnectorConfig Datasource connection configuration. Exactly one connector type must be set.
type ConnectorConfig struct {
	// Connector type — set the matching config object: 'database' → database, 'cloud_storage' → cloud_storage, 'web_crawl' → web_crawl
	Type string `json:"type"`
	Database interface{} `json:"database,omitempty"`
	CloudStorage interface{} `json:"cloud_storage,omitempty"`
	WebCrawl interface{} `json:"web_crawl,omitempty"`
}

// ConnectorConfigType represents the possible values for type.
type ConnectorConfigType string

const (
	ConnectorConfigTypeDatabase ConnectorConfigType = "database"
	ConnectorConfigTypeCloudStorage ConnectorConfigType = "cloud_storage"
	ConnectorConfigTypeWebCrawl ConnectorConfigType = "web_crawl"
)

// CreateAgentArtifactRequest Request model for creating a new agent artifact.
type CreateAgentArtifactRequest struct {
	// Human-readable name of the artifact (letters, numbers, and spaces only). Converted to kebab-case internally.
	DisplayName string `json:"display_name"`
	// Artifact type (json, markdown, csv, yaml, text, html, pdf)
	Type ArtifactType `json:"type"`
	// Description of the artifact
	Description interface{} `json:"description,omitempty"`
	// Whether agent must produce this artifact
	Required interface{} `json:"required,omitempty"`
	// Schema definition
	SchemaDef map[string]interface{} `json:"schema_def"`
	// Maximum artifact size in bytes
	MaxSizeBytes interface{} `json:"max_size_bytes,omitempty"`
	// Storage strategy (inline, gcs, auto)
	StorageStrategy interface{} `json:"storage_strategy,omitempty"`
	AdditionalProperties map[string]interface{} `json:"additional_properties,omitempty"`
}

// CreateAgentDefinitionRequest Request model for creating a new agent definition.
type CreateAgentDefinitionRequest struct {
	// Human-readable name of the agent (letters, numbers, and spaces only). Converted to kebab-case internally.
	DisplayName string `json:"display_name"`
	// System prompt/instructions for the agent
	Instructions string `json:"instructions"`
	// Agent type
	Type interface{} `json:"type,omitempty"`
	// Description of the agent
	Description interface{} `json:"description,omitempty"`
	// LLM model to use
	LlmModel interface{} `json:"llm_model,omitempty"`
	// List of fallback models
	FallbackModels interface{} `json:"fallback_models,omitempty"`
	// Datasource IDs the agent has access to
	Datasources interface{} `json:"datasources,omitempty"`
	// Tools configuration
	Tools interface{} `json:"tools,omitempty"`
	// Catalog URNs of artifacts the agent produces
	Artifacts interface{} `json:"artifacts,omitempty"`
	// Confidence scoring module names to apply during execution
	ConfidenceConfigs interface{} `json:"confidence_configs,omitempty"`
	// LLM temperature
	Temperature interface{} `json:"temperature,omitempty"`
	// Maximum tokens in response
	MaxTokens interface{} `json:"max_tokens,omitempty"`
	// Tags for categorization
	Tags interface{} `json:"tags,omitempty"`
	// UI icon identifier
	Icon interface{} `json:"icon,omitempty"`
	AdditionalProperties map[string]interface{} `json:"additional_properties,omitempty"`
}

// CreateAgentPromptRequest Request model for creating a new agent prompt.
type CreateAgentPromptRequest struct {
	// Human-readable name of the prompt (letters, numbers, and spaces only). Converted to kebab-case internally.
	DisplayName string `json:"display_name"`
	// Prompt text
	Prompt string `json:"prompt"`
}

// CreateAgentResponse represents the CreateAgentResponse type.
type CreateAgentResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Version string `json:"version"`
}

// CreateArtifactSchemaResponse represents the CreateArtifactSchemaResponse type.
type CreateArtifactSchemaResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Version string `json:"version"`
}

// CreateBatchDefinitionRequest Create a new BatchDefinition lineage.
type CreateBatchDefinitionRequest struct {
	// Kebab-case label (non-unique within tenant)
	Name string `json:"name"`
	// AgentDefinition ID; resolved + pinned at creation time
	AgentId string `json:"agent_id"`
	// Datasource holding the input Data Elements
	InputDatasourceId string `json:"input_datasource_id"`
	Filters interface{} `json:"filters,omitempty"`
	// Pinned output sink. NULL = workflow auto-creates per execution.
	OutputDatasourceId interface{} `json:"output_datasource_id,omitempty"`
	UserMessage interface{} `json:"user_message,omitempty"`
	Concurrency interface{} `json:"concurrency,omitempty"`
	RetryLimit interface{} `json:"retry_limit,omitempty"`
	// Cron expression validated by croniter; not yet scheduled in DEL-1376.
	RecurrenceCron interface{} `json:"recurrence_cron,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// CreateBatchDefinitionResponse Compact post-create payload mirroring CreateAgentDefinitionResponse.
type CreateBatchDefinitionResponse struct {
	Id string `json:"id"`
	CatalogUrn string `json:"catalog_urn"`
	Name string `json:"name"`
	Version string `json:"version"`
}

// CreateBatchExecutionRequest Legacy request body for POST /batch-execution/ (pre-DEL-1376 compat shim).
type CreateBatchExecutionRequest struct {
	BatchSpecJson LegacyBatchSpecJson `json:"batch_spec_json"`
}

// CreateDatasourceRequest Body for creating a new datasource.
type CreateDatasourceRequest struct {
	// Human-readable datasource name
	Name string `json:"name"`
	// What this datasource contains
	Description *string `json:"description,omitempty"`
	// Connection configuration
	Connector ConnectorConfig `json:"connector"`
	// Optional metadata extraction config to apply after creation
	MetadataConfig interface{} `json:"metadata_config,omitempty"`
}

// CreatePromptResponse represents the CreatePromptResponse type.
type CreatePromptResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Version string `json:"version"`
}

// CreateSessionRequest represents the CreateSessionRequest type.
type CreateSessionRequest struct {
	Prompt interface{} `json:"prompt,omitempty"`
	InitialContext interface{} `json:"initial_context,omitempty"`
	MaxIterationsPerUserMessage interface{} `json:"max_iterations_per_user_message,omitempty"`
}

// CreateSessionResponse represents the CreateSessionResponse type.
type CreateSessionResponse struct {
	SessionId string `json:"session_id"`
}

// DataElementListResponse Paginated list of data elements.
type DataElementListResponse struct {
	// Page of data elements
	Items []DataElementResponse `json:"items"`
	// Pass to the next request as `cursor` to fetch the next page. Null when there are no more pages
	NextCursor interface{} `json:"next_cursor,omitempty"`
	// True if more pages are available
	HasNext *bool `json:"has_next,omitempty"`
}

// DataElementResponse A single data element on a datasource.
type DataElementResponse struct {
	// Unique data element ID
	Id string `json:"id"`
	// ID of the datasource this element belongs to
	DatasourceId string `json:"datasource_id"`
	// Data element name
	Name string `json:"name"`
	// Human-authored description
	Description interface{} `json:"description,omitempty"`
	// MIME type of the underlying content
	MediaType interface{} `json:"media_type,omitempty"`
	// Arbitrary metadata key-value pairs
	Metadata interface{} `json:"metadata,omitempty"`
	// ISO 8601 creation timestamp
	CreatedAt interface{} `json:"created_at,omitempty"`
	// ISO 8601 last-update timestamp
	UpdatedAt interface{} `json:"updated_at,omitempty"`
}

// DataElementSearchRequest Body for searching data elements on a datasource.
type DataElementSearchRequest struct {
	// Regex pattern to filter by name
	RegexFilter interface{} `json:"regex_filter,omitempty"`
	// Filter by MIME types
	MediaTypeFilters interface{} `json:"media_type_filters,omitempty"`
}

// DatabaseConnector Connect to a relational database.
type DatabaseConnector struct {
	// Database host address
	Host string `json:"host"`
	// Database port
	Port int64 `json:"port"`
	// Database name
	Database string `json:"database"`
	// Schema name (defaults to public)
	SchemaName interface{} `json:"schema_name,omitempty"`
}

// DatasourceListResponse List of datasources visible to the caller.
type DatasourceListResponse struct {
	// Datasources in the caller's project
	Datasources []DatasourceResponse `json:"datasources"`
}

// DatasourceResponse A datasource with its latest sync/ingest state and (optionally) table details.
type DatasourceResponse struct {
	// Unique datasource ID
	Id string `json:"id"`
	// Human-readable datasource name
	Name string `json:"name"`
	// What this datasource contains
	Description string `json:"description"`
	// Connection configuration
	Connector ConnectorConfig `json:"connector"`
	// ISO 8601 creation timestamp
	CreatedAt string `json:"created_at"`
	// ISO 8601 last-update timestamp
	UpdatedAt string `json:"updated_at"`
	// ISO 8601 timestamp of the most recent ingest run
	LastSyncAt interface{} `json:"last_sync_at,omitempty"`
	// Status of the most recent ingest run (e.g. 'completed', 'failed')
	LastSyncStatus interface{} `json:"last_sync_status,omitempty"`
	// Total number of files ingested across all runs
	TotalIngestedFiles interface{} `json:"total_ingested_files,omitempty"`
	// Current metadata extraction configuration
	MetadataConfig interface{} `json:"metadata_config,omitempty"`
	// File counts for the datasource
	Files interface{} `json:"files,omitempty"`
	// Per-method counts from the latest ingest run
	IngestCounts interface{} `json:"ingest_counts,omitempty"`
	// Tables discovered on a structured datasource — only populated when include_tables=true
	Tables interface{} `json:"tables,omitempty"`
}

// DocumentChild Child document from container (ZIP/TAR/EML).
type DocumentChild struct {
	JobId string `json:"job_id"`
	Filename string `json:"filename"`
	Status string `json:"status"`
	MediaType string `json:"media_type"`
}

// DocumentElement A structural element in a parsed document.
type DocumentElement struct {
	// heading | paragraph | table | list_item | image | code_block | ...
	Type string `json:"type"`
	Text interface{} `json:"text,omitempty"`
	// Heading level (1-6)
	Level interface{} `json:"level,omitempty"`
	Table interface{} `json:"table,omitempty"`
	Bbox interface{} `json:"bbox,omitempty"`
	Confidence interface{} `json:"confidence,omitempty"`
	Page interface{} `json:"page,omitempty"`
}

// DocumentStatus Returned from GET /documents/{job_id}.
type DocumentStatus struct {
	JobId string `json:"job_id"`
	// queued | processing | completed | failed
	Status string `json:"status"`
	// meibel | markdown | docling
	Format string `json:"format"`
	Pages interface{} `json:"pages,omitempty"`
	Elements interface{} `json:"elements,omitempty"`
	Tables interface{} `json:"tables,omitempty"`
	Confidence interface{} `json:"confidence,omitempty"`
	ProcessingTimeMs interface{} `json:"processing_time_ms,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

// DownloadJobRequest Body for creating a download job. Omit both fields to download all files.
type DownloadJobRequest struct {
	// Content to include: files, parsed_content, or files_and_parsed_content
	Content interface{} `json:"content,omitempty"`
	// Specific data element IDs to include
	DataElementIds interface{} `json:"data_element_ids,omitempty"`
}

// DownloadJobResponse Result of creating a download job.
type DownloadJobResponse struct {
	// Identifier for the job — use with the status_url to track progress
	JobId string `json:"job_id"`
	// Current job status
	Status string `json:"status"`
	// Stream progress events from this SSE URL
	StatusUrl string `json:"status_url"`
}

// ExecuteBatchDefinitionResponse ExecuteBatchDefinitionResponse
type ExecuteBatchDefinitionResponse struct {
	ExecutionId string `json:"execution_id"`
	WorkflowId string `json:"workflow_id"`
}

// FieldSummary represents the FieldSummary type.
type FieldSummary struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// FileParseCompleteInfo FileParseCompleteInfo
type FileParseCompleteInfo struct {
	Status interface{} `json:"status"`
	Error interface{} `json:"error,omitempty"`
	BboxCount interface{} `json:"bbox_count"`
	PageCount interface{} `json:"page_count"`
	ContentType interface{} `json:"content_type"`
	Timestamp interface{} `json:"timestamp"`
}

// FileParseEntry FileParseEntry
type FileParseEntry struct {
	FileId string `json:"file_id"`
	Filename interface{} `json:"filename"`
	ParseStart interface{} `json:"parse_start"`
	ParseComplete interface{} `json:"parse_complete"`
}

// FileParseStartInfo FileParseStartInfo
type FileParseStartInfo struct {
	Attempt interface{} `json:"attempt"`
	Timestamp interface{} `json:"timestamp"`
}

// FileUploadSyncResponse Result of a synchronous upload — waits until files are persisted, optionally triggers ingest, and returns the resulting content listing.
type FileUploadSyncResponse struct {
	// ID of the datasource the files were uploaded to
	DatasourceId string `json:"datasource_id"`
	// Content items present on the datasource after the upload completes
	Items []ContentItem `json:"items"`
	// Set when the listing is truncated — pass to GET /datasources/{id}/content to fetch the rest
	ContinuationToken interface{} `json:"continuation_token,omitempty"`
	// URL to poll for ingest status. Only set when `trigger_ingest=true` was supplied
	IngestUrl interface{} `json:"ingest_url,omitempty"`
}

// FilesSummaryResponse File-count summary across the datasource's content.
type FilesSummaryResponse struct {
	// Total files currently tracked on the datasource
	Total int64 `json:"total"`
	// Files that have been removed since the last ingest
	Deleted interface{} `json:"deleted,omitempty"`
}

// GetBatchDefinitionsResponse GetBatchDefinitionsResponse
type GetBatchDefinitionsResponse struct {
	Data []BatchDefinitionResponse `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// GetBatchExecutionsResponse Response model for listing batch executions.
type GetBatchExecutionsResponse struct {
	Data []BatchExecutionResponse `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// HttpValidationError represents the HTTPValidationError type.
type HttpValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}

// IngestCountsResponse File counts broken down by ingest method for the latest run.
type IngestCountsResponse struct {
	// Counts for the RAG ingest method
	Rag interface{} `json:"rag,omitempty"`
	// Counts for the TAG (tables/columns) ingest method
	Tag interface{} `json:"tag,omitempty"`
	// Counts for the reference-graph ingest method
	RefGraph interface{} `json:"ref_graph,omitempty"`
}

// IngestMethodCountsResponse Per-method file counts produced by the latest ingest run.
type IngestMethodCountsResponse struct {
	// Total files processed by this ingest method
	Total int64 `json:"total"`
	// Files newly added in this run
	New interface{} `json:"new,omitempty"`
	// Files re-processed because they changed
	Updated interface{} `json:"updated,omitempty"`
}

// IngestMethodSummary Per-method aggregate counts for the current ingest run.
type IngestMethodSummary struct {
	// Ingest method name (e.g. 'rag', 'tag', 'ref_graph')
	Method string `json:"method"`
	// Total files this method intends to process
	TotalFiles *int64 `json:"total_files,omitempty"`
	// Files this method has finished processing
	ProcessedFiles *int64 `json:"processed_files,omitempty"`
	// Files added in this run
	Adds *int64 `json:"adds,omitempty"`
	// Files re-processed because they changed
	Updates *int64 `json:"updates,omitempty"`
	// Files that errored during processing
	Errors *int64 `json:"errors,omitempty"`
	// Files that produced warnings during processing
	Warnings *int64 `json:"warnings,omitempty"`
}

// IngestStatus represents the possible values for IngestStatus.
type IngestStatus string

const (
	IngestStatusNotStarted IngestStatus = "not_started"
	IngestStatusRunning IngestStatus = "running"
	IngestStatusCompleted IngestStatus = "completed"
	IngestStatusFailed IngestStatus = "failed"
	IngestStatusCanceled IngestStatus = "canceled"
	IngestStatusTerminated IngestStatus = "terminated"
	IngestStatusTimedOut IngestStatus = "timed_out"
	IngestStatusUnknown IngestStatus = "unknown"
)

// IngestStatusResponse Status of the most recent ingest run for a datasource.
type IngestStatusResponse struct {
	// ID of the datasource
	DatasourceId string `json:"datasource_id"`
	// Overall run status
	Status IngestStatus `json:"status"`
	// ISO 8601 timestamp when this run started
	StartedAt interface{} `json:"started_at,omitempty"`
	// ISO 8601 timestamp when this run finished — null while still running
	CompletedAt interface{} `json:"completed_at,omitempty"`
	// Per-method progress and counts for this run
	Methods []IngestMethodSummary `json:"methods,omitempty"`
}

// JudgeConfig Configuration for judge-based confidence scoring (LLM-as-judge patterns).
type JudgeConfig struct {
	Prompt string `json:"prompt"`
	TemperatureMax interface{} `json:"temperature_max,omitempty"`
	TemperatureStep interface{} `json:"temperature_step,omitempty"`
}

// LegacyBatchExecutionParams LegacyBatchExecutionParams
type LegacyBatchExecutionParams struct {
	Concurrency interface{} `json:"concurrency,omitempty"`
	RetryLimit interface{} `json:"retry_limit,omitempty"`
}

// LegacyBatchInputConfig LegacyBatchInputConfig
type LegacyBatchInputConfig struct {
	DatasourceId string `json:"datasource_id"`
	Filters interface{} `json:"filters,omitempty"`
}

// LegacyBatchInputFilters LegacyBatchInputFilters
type LegacyBatchInputFilters struct {
	Regex interface{} `json:"regex,omitempty"`
	MediaTypes interface{} `json:"media_types,omitempty"`
	ElementIds interface{} `json:"element_ids,omitempty"`
	AdditionalProperties map[string]interface{} `json:"additional_properties,omitempty"`
}

// LegacyBatchOutputConfig LegacyBatchOutputConfig
type LegacyBatchOutputConfig struct {
	DatasourceId interface{} `json:"datasource_id,omitempty"`
}

// LegacyBatchSpecJson LegacyBatchSpecJson
type LegacyBatchSpecJson struct {
	Name string `json:"name"`
	Version interface{} `json:"version,omitempty"`
	// AgentDefinition ID
	Agent string `json:"agent"`
	UserMessage interface{} `json:"user_message,omitempty"`
	Input LegacyBatchInputConfig `json:"input"`
	Output interface{} `json:"output,omitempty"`
	Execution interface{} `json:"execution,omitempty"`
	AdditionalProperties map[string]interface{} `json:"additional_properties,omitempty"`
}

// MeibelDocumentResult Full structured parse result (meibel format).
type MeibelDocumentResult struct {
	Elements []DocumentElement `json:"elements"`
	Pages int64 `json:"pages"`
	Tables int64 `json:"tables"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// MessageEntry MessageEntry
type MessageEntry struct {
	Role string `json:"role"`
	Message string `json:"message"`
	SignalId interface{} `json:"signal_id"`
	Timestamp time.Time `json:"timestamp"`
}

// MetadataConfigRequest Configure automatic metadata extraction from documents on ingest.
type MetadataConfigRequest struct {
	// Use 'catalog' to select a pre-built extraction model (set model_id); use 'custom' to define your own fields (set fields)
	Type string `json:"type"`
	// Pre-built model ID from the metadata model catalog — required when type is 'catalog'
	ModelId interface{} `json:"model_id,omitempty"`
	// Custom field definitions to extract — required when type is 'custom'
	Fields interface{} `json:"fields,omitempty"`
}

// MetadataConfigRequestType represents the possible values for type.
type MetadataConfigRequestType string

const (
	MetadataConfigRequestTypeCatalog MetadataConfigRequestType = "catalog"
	MetadataConfigRequestTypeCustom MetadataConfigRequestType = "custom"
)

// MetadataConfigResponse Current metadata extraction configuration on a datasource.
type MetadataConfigResponse struct {
	// 'catalog' = using a pre-built model, 'custom' = user-defined fields, 'default' = no extraction configured
	Type string `json:"type"`
	// Catalog model ID — only set when type is 'catalog'
	ModelId interface{} `json:"model_id,omitempty"`
	// Resolved field definitions in effect. Empty when type is 'default'
	Fields []MetadataField `json:"fields"`
}

// MetadataConfigResponseType represents the possible values for type.
type MetadataConfigResponseType string

const (
	MetadataConfigResponseTypeCatalog MetadataConfigResponseType = "catalog"
	MetadataConfigResponseTypeCustom MetadataConfigResponseType = "custom"
	MetadataConfigResponseTypeDefault MetadataConfigResponseType = "default"
)

// MetadataField A single field extracted from documents during ingest.
type MetadataField struct {
	// Field name (snake_case)
	Name string `json:"name"`
	// Data type of the field
	Type string `json:"type"`
	// What this field captures
	Description string `json:"description"`
	// Whether this field is indexed for filtering
	Index *bool `json:"index,omitempty"`
}

// MetadataFieldType represents the possible values for type.
type MetadataFieldType string

const (
	MetadataFieldTypeString MetadataFieldType = "string"
	MetadataFieldTypeInteger MetadataFieldType = "integer"
	MetadataFieldTypeFloat MetadataFieldType = "float"
	MetadataFieldTypeBoolean MetadataFieldType = "boolean"
	MetadataFieldTypeDatetime MetadataFieldType = "datetime"
	MetadataFieldTypeUuid MetadataFieldType = "uuid"
	MetadataFieldTypeGeo MetadataFieldType = "geo"
	MetadataFieldTypeListString MetadataFieldType = "list[string]"
)

// NBootstraps NBootstraps
type NBootstraps struct {
	AnyofSchema_1Validator interface{} `json:"anyof_schema_1_validator,omitempty"`
	AnyofSchema_2Validator interface{} `json:"anyof_schema_2_validator,omitempty"`
	ActualInstance *string `json:"actual_instance,omitempty"`
	AnyOfSchemas []string `json:"any_of_schemas,omitempty"`
}

// OcConfig Configuration for Observed Consistency confidence scoring.
type OcConfig struct {
	NCompletions interface{} `json:"n_completions,omitempty"`
	MaxTokens interface{} `json:"max_tokens,omitempty"`
	Temperature interface{} `json:"temperature,omitempty"`
	Models interface{} `json:"models,omitempty"`
	NliModelConfig map[string]interface{} `json:"nli_model_config"`
	NBootstraps interface{} `json:"n_bootstraps,omitempty"`
	TokenLimit interface{} `json:"token_limit,omitempty"`
	OriginalCompletion interface{} `json:"original_completion,omitempty"`
	ComparisonCompletions interface{} `json:"comparison_completions,omitempty"`
}

// OcrConfig Configuration for OCR confidence scoring.
type OcrConfig struct {
	CalibrationModel interface{} `json:"calibration_model,omitempty"`
	OcrConfidenceScores interface{} `json:"ocr_confidence_scores,omitempty"`
}

// PaginationMeta Pagination metadata included in list responses.
type PaginationMeta struct {
	// Total number of items matching the query
	Total int64 `json:"total"`
	// Number of items skipped
	Offset int64 `json:"offset"`
	// Maximum number of items returned (None means no limit applied)
	Limit interface{} `json:"limit,omitempty"`
}

// ParseDocumentResponse Returned from POST /documents (async).
type ParseDocumentResponse struct {
	JobId string `json:"job_id"`
	// Job status, e.g. 'queued'
	Status string `json:"status"`
}

// ProcessDocumentResponse Returned from POST /documents/process (sync).
type ProcessDocumentResponse struct {
	JobId string `json:"job_id"`
	// 'completed'
	Status string `json:"status"`
	// MeibelDocumentResult for meibel format, str for markdown
	Result interface{} `json:"result"`
}

// PromptListResponse represents the PromptListResponse type.
type PromptListResponse struct {
	Data []PromptSummary `json:"data"`
}

// PromptResponse represents the PromptResponse type.
type PromptResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Version string `json:"version"`
	ParentVersion interface{} `json:"parent_version,omitempty"`
	Prompt string `json:"prompt"`
	Description interface{} `json:"description,omitempty"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
}

// PromptSummary represents the PromptSummary type.
type PromptSummary struct {
	Id string `json:"id"`
	DisplayName string `json:"display_name"`
	Version string `json:"version"`
	Preview string `json:"preview"`
}

// PublishAgentDefinitionRequest Request model for publishing the current draft of an agent.
type PublishAgentDefinitionRequest struct {
	// User-provided description of what changed in this version
	CommitMessage string `json:"commit_message"`
}

// PublishAgentDefinitionResponse Response model for a publish event.
type PublishAgentDefinitionResponse struct {
	// Registry entry ID
	Id string `json:"id"`
	// Catalog URN of the published AgentDefinition version
	AgentDefinitionUrn string `json:"agent_definition_urn"`
	// Agent name
	AgentName string `json:"agent_name"`
	// Published version slug
	Version string `json:"version"`
	// Display name of the published version
	DisplayName string `json:"display_name"`
	// User-provided description of what changed in this version
	CommitMessage string `json:"commit_message"`
	// Timestamp of the publish event
	PublishedAt time.Time `json:"published_at"`
	// User who published
	PublishedBy interface{} `json:"published_by,omitempty"`
}

// ScoreSummary Aggregated summary of scoring jobs matching one or two AgentIdentityContext filters.  With one level (primary only): flat aggregate of all jobs matching primary_field=primary_value. With two levels (primary + secondary): both constraints are applied; primary is the higher level and secondary is the lower level.
type ScoreSummary struct {
	PrimaryField string `json:"primary_field"`
	PrimaryValue string `json:"primary_value"`
	SecondaryField interface{} `json:"secondary_field,omitempty"`
	SecondaryValue interface{} `json:"secondary_value,omitempty"`
	Status interface{} `json:"status,omitempty"`
	AggregateScore interface{} `json:"aggregate_score,omitempty"`
	ModuleScores interface{} `json:"module_scores,omitempty"`
	NJobsPerModule interface{} `json:"n_jobs_per_module,omitempty"`
	Jobs interface{} `json:"jobs,omitempty"`
}

// ScoringJobRecord ScoringJobRecord
type ScoringJobRecord struct {
	JobId string `json:"job_id"`
	AgentIdentityContext AgentIdentityContext `json:"agent_identity_context"`
	Module string `json:"module"`
	ScoringConfig ConfidenceScoringConfig `json:"scoring_config"`
	InputValue string `json:"input_value"`
	OutputValue string `json:"output_value"`
	Status ScoringStatus `json:"status"`
	Score interface{} `json:"score,omitempty"`
}

// ScoringStatus represents the possible values for ScoringStatus.
type ScoringStatus string

const (
	ScoringStatusSubmitted ScoringStatus = "submitted"
	ScoringStatusInProgress ScoringStatus = "in_progress"
	ScoringStatusCompleted ScoringStatus = "completed"
	ScoringStatusFailed ScoringStatus = "failed"
)

// SessionListResponse represents the SessionListResponse type.
type SessionListResponse struct {
	Data []SessionSummary `json:"data"`
	Total int64 `json:"total"`
}

// SessionMessageItem represents the SessionMessageItem type.
type SessionMessageItem struct {
	Type string `json:"type"`
	Timestamp interface{} `json:"timestamp,omitempty"`
	Message interface{} `json:"message,omitempty"`
	SignalId interface{} `json:"signal_id,omitempty"`
	ToolId interface{} `json:"tool_id,omitempty"`
	ToolName interface{} `json:"tool_name,omitempty"`
	Arguments interface{} `json:"arguments,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// SessionMessagesResponse represents the SessionMessagesResponse type.
type SessionMessagesResponse struct {
	AgentId interface{} `json:"agent_id,omitempty"`
	AgentName interface{} `json:"agent_name,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Messages []SessionMessageItem `json:"messages"`
}

// SessionSummary represents the SessionSummary type.
type SessionSummary struct {
	SessionId string `json:"session_id"`
	Status string `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime interface{} `json:"end_time,omitempty"`
	AgentName interface{} `json:"agent_name,omitempty"`
	AgentVersion interface{} `json:"agent_version,omitempty"`
	MessagesCount *int64 `json:"messages_count,omitempty"`
	TokenUsage interface{} `json:"token_usage,omitempty"`
	Result []map[string]interface{} `json:"result,omitempty"`
}

// Source A source/citation in the response.
type Source struct {
	Title string `json:"title"`
	Url interface{} `json:"url,omitempty"`
	Snippet interface{} `json:"snippet,omitempty"`
	DataElementId interface{} `json:"data_element_id,omitempty"`
	RelevanceScore interface{} `json:"relevance_score,omitempty"`
}

// SubmitDocumentTransformResponse represents the SubmitDocumentTransformResponse type.
type SubmitDocumentTransformResponse struct {
	// Poll via client.sessions.get(execution_id)
	ExecutionId string `json:"execution_id"`
}

// Table represents the Table type.
type Table struct {
	Cells []TableCell `json:"cells"`
	Rows int64 `json:"rows"`
	Cols int64 `json:"cols"`
	Bbox interface{} `json:"bbox,omitempty"`
}

// TableCell represents the TableCell type.
type TableCell struct {
	Text string `json:"text"`
	Row int64 `json:"row"`
	Col int64 `json:"col"`
	RowSpan *int64 `json:"row_span,omitempty"`
	ColSpan *int64 `json:"col_span,omitempty"`
	Bbox interface{} `json:"bbox,omitempty"`
}

// TableDescriptionUpdate A nested table-update entry used inside UpdateDatasourceRequest.tables.
type TableDescriptionUpdate struct {
	// Name of the table to update
	TableName string `json:"table_name"`
	// Updated description for the table (omit to leave unchanged)
	Description interface{} `json:"description,omitempty"`
	// Optional list of column-description updates for this table
	Columns interface{} `json:"columns,omitempty"`
}

// TableSummaryResponse Summary of a single table discovered on a structured datasource.
type TableSummaryResponse struct {
	// Table name
	Name string `json:"name"`
	// Human-authored description of the table
	Description interface{} `json:"description,omitempty"`
	// Number of columns on the table
	ColumnCount int64 `json:"column_count"`
}

// TagColumn A column on a structured-datasource table, with its description.
type TagColumn struct {
	// Column name as defined in the source table
	ColumnName string `json:"column_name"`
	// SQL data type of the column (e.g. 'varchar', 'integer')
	Type interface{} `json:"type,omitempty"`
	// Human-authored description of what this column represents
	Description interface{} `json:"description,omitempty"`
}

// TagColumnUpdateItem A single column-description update entry within an UpdateTagColumnsRequest.
type TagColumnUpdateItem struct {
	// Name of the column to update
	ColumnName string `json:"column_name"`
	// New description for the column
	Description string `json:"description"`
}

// TagTable A table on a structured datasource, with its description and optionally its columns.
type TagTable struct {
	// Table name as defined on the datasource
	TableName string `json:"table_name"`
	// Human-authored description of what this table represents
	Description interface{} `json:"description,omitempty"`
	// Columns on the table — only populated when explicitly requested via include_columns
	Columns interface{} `json:"columns,omitempty"`
}

// TagTableUpdateItem A single table-description update entry within an UpdateTagTablesRequest.
type TagTableUpdateItem struct {
	// Name of the table to update
	TableName string `json:"table_name"`
	// New description for the table
	Description string `json:"description"`
}

// TokenConfig Configuration for token-based confidence scoring (TF-IDF).
type TokenConfig struct {
	Model interface{} `json:"model,omitempty"`
	RemoveStopWords interface{} `json:"remove_stop_words,omitempty"`
	LowerCase interface{} `json:"lower_case,omitempty"`
	MaxNgrams interface{} `json:"max_ngrams,omitempty"`
	NInfluencers interface{} `json:"n_influencers,omitempty"`
}

// ToolActivity Record of a tool call and its result.
type ToolActivity struct {
	ToolId string `json:"tool_id"`
	ToolName string `json:"tool_name"`
	Arguments map[string]interface{} `json:"arguments"`
	// Optional override for the tool's parameters schema
	Result interface{} `json:"result,omitempty"`
	Timestamp string `json:"timestamp"`
}

// ToolActivityEntry ToolActivityEntry
type ToolActivityEntry struct {
	ToolId string `json:"tool_id"`
	ToolCall interface{} `json:"tool_call"`
	ToolResult interface{} `json:"tool_result"`
}

// ToolCallInfo ToolCallInfo
type ToolCallInfo struct {
	ToolName interface{} `json:"tool_name"`
	// Optional override for the tool's parameters schema
	Arguments interface{} `json:"arguments"`
	Sequence interface{} `json:"sequence"`
	Timestamp interface{} `json:"timestamp"`
}

// ToolResultInfo ToolResultInfo
type ToolResultInfo struct {
	ToolName interface{} `json:"tool_name"`
	Result interface{} `json:"result,omitempty"`
	Sequence interface{} `json:"sequence"`
	Timestamp interface{} `json:"timestamp"`
}

// TransformDocumentRequest represents the TransformDocumentRequest type.
type TransformDocumentRequest struct {
	// File path, URL, or GCS URI to transform
	File string `json:"file"`
	// Schema name/ID or inline JSON Schema
	ArtifactSchema interface{} `json:"artifact_schema"`
	// LLM model override
	Model interface{} `json:"model,omitempty"`
	// Extraction instructions override
	Prompt interface{} `json:"prompt,omitempty"`
	// Prompt template reference
	PromptId interface{} `json:"prompt_id,omitempty"`
	// Max wait time in seconds (sync only)
	TimeoutSeconds interface{} `json:"timeout_seconds,omitempty"`
}

// TransformDocumentResponse represents the TransformDocumentResponse type.
type TransformDocumentResponse struct {
	// Execution ID for debugging/tracing
	ExecutionId string `json:"execution_id"`
	// Extracted artifact data
	Data map[string]interface{} `json:"data"`
	// LLM token consumption
	TokenUsage interface{} `json:"token_usage,omitempty"`
}

// TriggerIngestResponse Acknowledgement that ingest was kicked off for a datasource.
type TriggerIngestResponse struct {
	// Human-readable confirmation message
	Message string `json:"message"`
	// ID of the datasource ingest was triggered on
	DatasourceId string `json:"datasource_id"`
}

// UpdateAgentArtifactRequest Request model for updating an agent artifact. Name is intentionally excluded as it serves as the stable identifier for a version chain and cannot be changed.
type UpdateAgentArtifactRequest struct {
	// Human-readable name of the artifact
	DisplayName interface{} `json:"display_name,omitempty"`
	// Artifact type
	Type interface{} `json:"type,omitempty"`
	// Description of the artifact
	Description interface{} `json:"description,omitempty"`
	// Whether agent must produce this artifact
	Required interface{} `json:"required,omitempty"`
	// Optional override for the tool's parameters schema
	SchemaDef interface{} `json:"schema_def,omitempty"`
	// Maximum artifact size in bytes
	MaxSizeBytes interface{} `json:"max_size_bytes,omitempty"`
	// Storage strategy
	StorageStrategy interface{} `json:"storage_strategy,omitempty"`
}

// UpdateAgentDefinitionRequest Request model for updating an agent definition. Name is intentionally excluded as it serves as the stable identifier for a version chain and cannot be changed.
type UpdateAgentDefinitionRequest struct {
	// Human-readable name of the agent
	DisplayName interface{} `json:"display_name,omitempty"`
	// System prompt/instructions
	Instructions interface{} `json:"instructions,omitempty"`
	// Agent type
	Type interface{} `json:"type,omitempty"`
	// Description of the agent
	Description interface{} `json:"description,omitempty"`
	// LLM model to use
	LlmModel interface{} `json:"llm_model,omitempty"`
	// List of fallback models
	FallbackModels interface{} `json:"fallback_models,omitempty"`
	// Datasource IDs the agent has access to
	Datasources interface{} `json:"datasources,omitempty"`
	// Tools configuration
	Tools interface{} `json:"tools,omitempty"`
	// Catalog URNs of artifacts the agent produces
	Artifacts interface{} `json:"artifacts,omitempty"`
	// Confidence scoring module names to apply during execution
	ConfidenceConfigs interface{} `json:"confidence_configs,omitempty"`
	// LLM temperature
	Temperature interface{} `json:"temperature,omitempty"`
	// Maximum tokens in response
	MaxTokens interface{} `json:"max_tokens,omitempty"`
	// Tags for categorization
	Tags interface{} `json:"tags,omitempty"`
	// UI icon identifier
	Icon interface{} `json:"icon,omitempty"`
}

// UpdateAgentDefinitionResponse Response model for updating an agent definition.
type UpdateAgentDefinitionResponse struct {
	// New agent definition ID
	Id string `json:"id"`
	// Catalog URN for the new version
	CatalogUrn string `json:"catalog_urn"`
	// New version number
	Version string `json:"version"`
}

// UpdateAgentPromptRequest Request model for updating an agent prompt. Name is intentionally excluded as it serves as the stable identifier for a version chain and cannot be changed.
type UpdateAgentPromptRequest struct {
	// Human-readable name of the prompt
	DisplayName interface{} `json:"display_name,omitempty"`
	// Prompt text
	Prompt interface{} `json:"prompt,omitempty"`
}

// UpdateArtifactSchemaResponse represents the UpdateArtifactSchemaResponse type.
type UpdateArtifactSchemaResponse struct {
	Id string `json:"id"`
	Version string `json:"version"`
}

// UpdateBatchDefinitionRequest Patch a BatchDefinition; the service forks a new version row.
type UpdateBatchDefinitionRequest struct {
	Name interface{} `json:"name,omitempty"`
	// If set, re-resolves and re-pins the agent spec
	AgentId interface{} `json:"agent_id,omitempty"`
	InputDatasourceId interface{} `json:"input_datasource_id,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
	OutputDatasourceId interface{} `json:"output_datasource_id,omitempty"`
	UserMessage interface{} `json:"user_message,omitempty"`
	Concurrency interface{} `json:"concurrency,omitempty"`
	RetryLimit interface{} `json:"retry_limit,omitempty"`
	RecurrenceCron interface{} `json:"recurrence_cron,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// UpdateBatchDefinitionResponse New version metadata returned after a successful update fork.
type UpdateBatchDefinitionResponse struct {
	Id string `json:"id"`
	CatalogUrn string `json:"catalog_urn"`
	Version string `json:"version"`
}

// UpdateBatchExecutionRequest Runtime-only patch fields. Identity, definition link, and overrides are immutable.
type UpdateBatchExecutionRequest struct {
	// Execution status
	Status interface{} `json:"status,omitempty"`
	// Execution end time
	EndTime interface{} `json:"end_time,omitempty"`
	// Total items in batch
	TotalItems interface{} `json:"total_items,omitempty"`
	// Number of succeeded items
	Succeeded interface{} `json:"succeeded,omitempty"`
	// Number of failed items
	Failed interface{} `json:"failed,omitempty"`
	// Output datasource ID
	OutputDatasourceId interface{} `json:"output_datasource_id,omitempty"`
	// Per-item results
	Items interface{} `json:"items,omitempty"`
	// Overall error message
	Error interface{} `json:"error,omitempty"`
	AdditionalProperties map[string]interface{} `json:"additional_properties,omitempty"`
}

// UpdatePromptResponse represents the UpdatePromptResponse type.
type UpdatePromptResponse struct {
	Id string `json:"id"`
	Version string `json:"version"`
}

// UpdateTagColumnsRequest Bulk update of column descriptions on a single table.
type UpdateTagColumnsRequest struct {
	// One entry per column to update on the target table
	Columns []TagColumnUpdateItem `json:"columns"`
}

// UpdateTagTablesRequest Bulk update of table descriptions on a datasource.
type UpdateTagTablesRequest struct {
	// One entry per table to update
	Tables []TagTableUpdateItem `json:"tables"`
}

// WebCrawlConnector Connect to a website for crawling.
type WebCrawlConnector struct {
	// Starting URL for the crawl
	BaseUrl string `json:"base_url"`
	// Enable JavaScript rendering
	JavascriptRender *bool `json:"javascript_render,omitempty"`
	// Per-domain include/exclude rules. If omitted, the crawler stays on the base_url's domain
	Domains interface{} `json:"domains,omitempty"`
}

// ConnectedEvent A server-sent event indicating the stream connection has been established
type ConnectedEvent struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

// StatusEvent A server-sent event containing an agent status update
type StatusEvent struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

// ToolCallEvent A server-sent event indicating the agent is calling a tool
type ToolCallEvent struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

// ToolResultEvent A server-sent event containing the result of a tool call
type ToolResultEvent struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

// PartialResponseEvent A server-sent event containing an incremental response from the agent
type PartialResponseEvent struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

// CompletionEvent A server-sent event containing the final complete response from the agent, sent once at the end of the stream
type CompletionEvent struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

// ContentItem A single file in a datasource's content store.
type ContentItem struct {
	// Filename
	Name string `json:"name"`
	// Object-storage path to the file relative to the datasource
	Path string `json:"path"`
	// Object kind reported by storage (e.g. 'file', 'directory')
	Type interface{} `json:"type,omitempty"`
	// File size in bytes
	Size interface{} `json:"size,omitempty"`
	// MIME type of the file
	MediaType interface{} `json:"media_type,omitempty"`
	// ISO 8601 timestamp of last modification in object storage
	LastModified interface{} `json:"last_modified,omitempty"`
	// Object-storage ETag for the file
	Etag interface{} `json:"etag,omitempty"`
}

// ListContentResponse Paginated list of files in a datasource's content store.
type ListContentResponse struct {
	// Page of content items
	Items []ContentItem `json:"items"`
	// Pass to the next request as `continuation_token` to fetch the next page. Null when there are no more pages
	ContinuationToken interface{} `json:"continuation_token,omitempty"`
}

// UploadContentResponse Result of an async upload — files are accepted and streamed asynchronously.
type UploadContentResponse struct {
	// True if the upload was accepted for processing
	Success bool `json:"success"`
	// Human-readable status message
	Message string `json:"message"`
	// ID of the datasource the files were uploaded to (created on the fly if `name` was supplied)
	DatasourceId string `json:"datasource_id"`
	// Identifier for this upload batch — use with the SSE stream to track progress
	UploadId string `json:"upload_id"`
	// Server-sent-events URL to stream upload progress until 'stream_complete'
	SseUrl string `json:"sse_url"`
	// Number of files the server expects to process for this upload
	EstimatedFiles interface{} `json:"estimated_files,omitempty"`
	// Total estimated size of the upload in bytes
	EstimatedSize interface{} `json:"estimated_size,omitempty"`
}

// UpdateDataElementRequest Body for updating a data element. Omit a field to leave it unchanged.
type UpdateDataElementRequest struct {
	// Updated name
	Name interface{} `json:"name,omitempty"`
	// Updated description
	Description interface{} `json:"description,omitempty"`
	// Metadata key-value pairs — replaces all existing metadata
	Metadata interface{} `json:"metadata,omitempty"`
}

// DeleteDatasourceResponse Result of deleting a datasource.
type DeleteDatasourceResponse struct {
	// ID of the datasource that was deleted
	Id string `json:"id"`
}

// UpdateDatasourceRequest Body for updating a datasource. Omit a field to leave it unchanged.
type UpdateDatasourceRequest struct {
	// Updated datasource name
	Name interface{} `json:"name,omitempty"`
	// Updated description
	Description interface{} `json:"description,omitempty"`
	// Updated connection configuration
	Connector interface{} `json:"connector,omitempty"`
	// Metadata extraction config — if changed, re-extraction triggers automatically
	MetadataConfig interface{} `json:"metadata_config,omitempty"`
	// Table and column descriptions to update (structured datasources only)
	Tables interface{} `json:"tables,omitempty"`
}

// WebDomain An allowed domain for a web-crawl datasource, with include/exclude URL patterns.
type WebDomain struct {
	// Domain to crawl (e.g. example.com)
	Domain string `json:"domain"`
	// Regex URL pattern to include
	IncludePattern string `json:"include_pattern"`
	// Regex URL pattern to exclude
	ExcludePattern *string `json:"exclude_pattern,omitempty"`
}

// ListMetadataModelCatalogResponse List of available metadata-extraction models in the catalog.
type ListMetadataModelCatalogResponse struct {
	// Catalog entries visible to the caller, filtered by the optional scope query param
	Models []MetadataModelCatalogEntry `json:"models"`
}

// MetadataModelCatalogEntry A pre-built metadata extraction model from the catalog, selectable by model_id.
type MetadataModelCatalogEntry struct {
	// Stable ID used to reference this model when configuring a datasource
	ModelId string `json:"model_id"`
	// Human-readable model name
	Name string `json:"name"`
	// What this model is designed to extract
	Description interface{} `json:"description,omitempty"`
	// Visibility of the model (e.g. 'global', 'customer', 'project')
	Scope string `json:"scope"`
	// Field definitions this model extracts
	Fields []MetadataField `json:"fields"`
}

// BodyUploadContent represents the Body_uploadContent type.
type BodyUploadContent struct {
	// One or more files to upload
	Files [][]byte `json:"files"`
	// ID of an existing datasource to upload to. Provide this or name.
	DatasourceId *string `json:"datasource_id,omitempty"`
	// Name for a new datasource to create. Provide this or datasource_id.
	Name *string `json:"name,omitempty"`
	// Description of the new datasource (only used when creating with name).
	Description *string `json:"description,omitempty"`
	MetadataConfig *MetadataConfigRequest `json:"metadata_config,omitempty"`
}

// BodyUploadAndListContent represents the Body_uploadAndListContent type.
type BodyUploadAndListContent struct {
	// One or more files to upload
	Files [][]byte `json:"files"`
	// ID of an existing datasource to upload to. Provide this or name.
	DatasourceId *string `json:"datasource_id,omitempty"`
	// Name for a new datasource to create. Provide this or datasource_id.
	Name *string `json:"name,omitempty"`
	// Description of the new datasource (only used when creating with name).
	Description *string `json:"description,omitempty"`
	MetadataConfig *MetadataConfigRequest `json:"metadata_config,omitempty"`
	// Start ingestion after upload completes. Returns ingest_url to poll for status.
	TriggerIngest *bool `json:"trigger_ingest,omitempty"`
}

// BodyParseDocument represents the Body_parseDocument type.
type BodyParseDocument struct {
	// The document file to parse
	File []byte `json:"file"`
}

// BodyProcessDocument represents the Body_processDocument type.
type BodyProcessDocument struct {
	// The document file to process
	File []byte `json:"file"`
}
