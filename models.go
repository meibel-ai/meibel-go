package meibelgo

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
	Tools []string `json:"tools"`
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
	Name interface{} `json:"name,omitempty"`
	DisplayName string `json:"display_name"`
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
	SchemaDef string `json:"schema_def"`
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
	SchemaDef string `json:"schema_def"`
	// Maximum artifact size in bytes
	MaxSizeBytes interface{} `json:"max_size_bytes,omitempty"`
	// Storage strategy (inline, gcs, auto)
	StorageStrategy interface{} `json:"storage_strategy,omitempty"`
	AdditionalProperties *string `json:"additional_properties,omitempty"`
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
	AdditionalProperties *string `json:"additional_properties,omitempty"`
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

// CreateDatasourceRequest represents the CreateDatasourceRequest type.
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

// DataElementListResponse represents the DataElementListResponse type.
type DataElementListResponse struct {
	Items []DataElementResponse `json:"items"`
	NextCursor interface{} `json:"next_cursor,omitempty"`
	HasNext *bool `json:"has_next,omitempty"`
}

// DataElementResponse represents the DataElementResponse type.
type DataElementResponse struct {
	Id string `json:"id"`
	DatasourceId string `json:"datasource_id"`
	Name string `json:"name"`
	Description interface{} `json:"description,omitempty"`
	MediaType interface{} `json:"media_type,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
}

// DataElementSearchRequest represents the DataElementSearchRequest type.
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

// DatasourceListResponse represents the DatasourceListResponse type.
type DatasourceListResponse struct {
	Datasources []DatasourceResponse `json:"datasources"`
}

// DatasourceResponse represents the DatasourceResponse type.
type DatasourceResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Connector ConnectorConfig `json:"connector"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	LastSyncAt interface{} `json:"last_sync_at,omitempty"`
	LastSyncStatus interface{} `json:"last_sync_status,omitempty"`
	TotalIngestedFiles interface{} `json:"total_ingested_files,omitempty"`
	MetadataConfig interface{} `json:"metadata_config,omitempty"`
	Files interface{} `json:"files,omitempty"`
	IngestCounts interface{} `json:"ingest_counts,omitempty"`
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

// DownloadJobRequest represents the DownloadJobRequest type.
type DownloadJobRequest struct {
	// Content to include: files, parsed_content, or files_and_parsed_content
	Content interface{} `json:"content,omitempty"`
	// Specific data element IDs to include
	DataElementIds interface{} `json:"data_element_ids,omitempty"`
}

// DownloadJobResponse represents the DownloadJobResponse type.
type DownloadJobResponse struct {
	JobId string `json:"job_id"`
	// Current job status
	Status string `json:"status"`
	// Stream progress events from this SSE URL
	StatusUrl string `json:"status_url"`
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

// FileUploadSyncResponse represents the FileUploadSyncResponse type.
type FileUploadSyncResponse struct {
	DatasourceId string `json:"datasource_id"`
	Items []ContentItem `json:"items"`
	ContinuationToken interface{} `json:"continuation_token,omitempty"`
	IngestUrl interface{} `json:"ingest_url,omitempty"`
}

// FilesSummaryResponse represents the FilesSummaryResponse type.
type FilesSummaryResponse struct {
	Total int64 `json:"total"`
	Deleted interface{} `json:"deleted,omitempty"`
}

// HttpValidationError represents the HTTPValidationError type.
type HttpValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}

// IngestCountsResponse represents the IngestCountsResponse type.
type IngestCountsResponse struct {
	Rag interface{} `json:"rag,omitempty"`
	Tag interface{} `json:"tag,omitempty"`
	RefGraph interface{} `json:"ref_graph,omitempty"`
}

// IngestMethodCountsResponse represents the IngestMethodCountsResponse type.
type IngestMethodCountsResponse struct {
	Total int64 `json:"total"`
	New interface{} `json:"new,omitempty"`
	Updated interface{} `json:"updated,omitempty"`
}

// IngestMethodSummary represents the IngestMethodSummary type.
type IngestMethodSummary struct {
	Method string `json:"method"`
	TotalFiles *int64 `json:"total_files,omitempty"`
	ProcessedFiles *int64 `json:"processed_files,omitempty"`
	Adds *int64 `json:"adds,omitempty"`
	Updates *int64 `json:"updates,omitempty"`
	Errors *int64 `json:"errors,omitempty"`
	Warnings *int64 `json:"warnings,omitempty"`
}

// IngestStatusResponse represents the IngestStatusResponse type.
type IngestStatusResponse struct {
	DatasourceId string `json:"datasource_id"`
	Status string `json:"status"`
	StartedAt interface{} `json:"started_at,omitempty"`
	CompletedAt interface{} `json:"completed_at,omitempty"`
	Methods []IngestMethodSummary `json:"methods,omitempty"`
}

// JudgeConfig Configuration for judge-based confidence scoring (LLM-as-judge patterns).
type JudgeConfig struct {
	Prompt string `json:"prompt"`
	TemperatureMax interface{} `json:"temperature_max,omitempty"`
	TemperatureStep interface{} `json:"temperature_step,omitempty"`
}

// ListMetadataModelCatalogResponse ListMetadataModelCatalogResponse
type ListMetadataModelCatalogResponse struct {
	Models []MetadataModelCatalogEntry `json:"models"`
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

// MetadataConfigResponse represents the MetadataConfigResponse type.
type MetadataConfigResponse struct {
	Type string `json:"type"`
	ModelId interface{} `json:"model_id,omitempty"`
	Fields []MetadataField `json:"fields"`
}

// MetadataConfigResponseType represents the possible values for type.
type MetadataConfigResponseType string

const (
	MetadataConfigResponseTypeCatalog MetadataConfigResponseType = "catalog"
	MetadataConfigResponseTypeCustom MetadataConfigResponseType = "custom"
	MetadataConfigResponseTypeDefault MetadataConfigResponseType = "default"
)

// MetadataField represents the MetadataField type.
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

// MetadataModelCatalogEntry MetadataModelCatalogEntry
type MetadataModelCatalogEntry struct {
	ModelId string `json:"model_id"`
	Name string `json:"name"`
	Description interface{} `json:"description,omitempty"`
	Scope string `json:"scope"`
	CustomerId interface{} `json:"customer_id,omitempty"`
	ProjectId interface{} `json:"project_id,omitempty"`
	Fields []MetadataModelField `json:"fields"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	UpdatedBy interface{} `json:"updated_by,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
}

// MetadataModelField MetadataModelField
type MetadataModelField struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	Index interface{} `json:"index,omitempty"`
}

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
	NliModelConfig string `json:"nli_model_config"`
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
	Result []string `json:"result,omitempty"`
}

// Source A source/citation in the response.
type Source struct {
	Title string `json:"title"`
	Url interface{} `json:"url,omitempty"`
	Snippet interface{} `json:"snippet,omitempty"`
	DataElementId interface{} `json:"data_element_id,omitempty"`
	RelevanceScore interface{} `json:"relevance_score,omitempty"`
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

// TableDescriptionUpdate represents the TableDescriptionUpdate type.
type TableDescriptionUpdate struct {
	TableName string `json:"table_name"`
	Description interface{} `json:"description,omitempty"`
	Columns interface{} `json:"columns,omitempty"`
}

// TableSummaryResponse represents the TableSummaryResponse type.
type TableSummaryResponse struct {
	Name string `json:"name"`
	Description interface{} `json:"description,omitempty"`
	ColumnCount int64 `json:"column_count"`
}

// TagColumn represents the TagColumn type.
type TagColumn struct {
	ColumnName string `json:"column_name"`
	Type interface{} `json:"type,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// TagColumnUpdateItem represents the TagColumnUpdateItem type.
type TagColumnUpdateItem struct {
	ColumnName string `json:"column_name"`
	Description string `json:"description"`
}

// TagTable represents the TagTable type.
type TagTable struct {
	TableName string `json:"table_name"`
	Description interface{} `json:"description,omitempty"`
	Columns interface{} `json:"columns,omitempty"`
}

// TagTableUpdateItem represents the TagTableUpdateItem type.
type TagTableUpdateItem struct {
	TableName string `json:"table_name"`
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
	Arguments string `json:"arguments"`
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

// UpdatePromptResponse represents the UpdatePromptResponse type.
type UpdatePromptResponse struct {
	Id string `json:"id"`
	Version string `json:"version"`
}

// WebCrawlConnector Connect to a website for crawling.
type WebCrawlConnector struct {
	// Starting URL for the crawl
	BaseUrl string `json:"base_url"`
	// Enable JavaScript rendering
	JavascriptRender *bool `json:"javascript_render,omitempty"`
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

// ContentItem represents the ContentItem type.
type ContentItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type interface{} `json:"type,omitempty"`
	Size interface{} `json:"size,omitempty"`
	MediaType interface{} `json:"media_type,omitempty"`
	LastModified interface{} `json:"last_modified,omitempty"`
	Etag interface{} `json:"etag,omitempty"`
}

// ListContentResponse represents the ListContentResponse type.
type ListContentResponse struct {
	Items []ContentItem `json:"items"`
	ContinuationToken interface{} `json:"continuation_token,omitempty"`
}

// UploadContentResponse represents the UploadContentResponse type.
type UploadContentResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	DatasourceId string `json:"datasource_id"`
	UploadId string `json:"upload_id"`
	SseUrl string `json:"sse_url"`
	EstimatedFiles interface{} `json:"estimated_files,omitempty"`
	EstimatedSize interface{} `json:"estimated_size,omitempty"`
	IngestUrl interface{} `json:"ingest_url,omitempty"`
}

// UpdateDataElementRequest represents the UpdateDataElementRequest type.
type UpdateDataElementRequest struct {
	// Updated name
	Name interface{} `json:"name,omitempty"`
	// Updated description
	Description interface{} `json:"description,omitempty"`
	// Metadata key-value pairs — replaces all existing metadata
	Metadata interface{} `json:"metadata,omitempty"`
}

// UpdateDatasourceRequest represents the UpdateDatasourceRequest type.
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

// WebDomain represents the WebDomain type.
type WebDomain struct {
	// Domain to crawl (e.g. example.com)
	Domain string `json:"domain"`
	// Regex URL pattern to include
	IncludePattern string `json:"include_pattern"`
	// Regex URL pattern to exclude
	ExcludePattern *string `json:"exclude_pattern,omitempty"`
}
