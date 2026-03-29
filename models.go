package meibelgo

// AgentIdentityContext Identity context for agent execution.  Contains only immutable identity fields that answer: - WHO: customer_id, project_id (tenant identity) - WHAT: agent_name, agent_version, agent_execution_id (agent identity, optional) - WHERE: agent_workflow_name, agent_workflow_version, agent_workflow_execution_id (parent workflow, optional) - WHICH TOOL: tool_id, tool_instance_id, tool_execution_id (tool identity, optional)  This model is FLAT - no inheritance, all fields in one model. Agent, workflow, and tool fields are optional, making this suitable for all execution contexts.  This model does NOT contain: - Configuration (see AgentExecutionConfig in agent-platform) - Runtime state (see AgentExecutionState in agent-platform)  Design Pattern - Progressive Enhancement: - Callers provide only tenant/project identity - FSMWorkflow fills in agent_workflow_* fields from AgentWorkflowSpec - ReactAgent fills in agent_* fields from AgentSpec and workflow.info().workflow_id - Tool activities add tool_* fields via model_copy() - Context gains fields as it flows through the system  Examples:     # Starting FSMWorkflow (caller provides minimal context)     context = AgentIdentityContext(         customer_id="cust_123",         project_id="proj_456"     )     # FSM fills in workflow identity     context = context.model_copy(update={         "agent_workflow_name": "support_fsm",         "agent_workflow_version": "3.0.0",         "agent_workflow_execution_id": workflow.info().workflow_id     })      # Starting ReactAgent standalone (caller provides minimal context)     context = AgentIdentityContext(         customer_id="cust_123",         project_id="proj_456"     )     # ReactAgent fills in agent identity     context = context.model_copy(update={         "agent_name": "sales_assistant",         "agent_version": "2.0.0",         "agent_execution_id": workflow.info().workflow_id     })      # ReactAgent as FSM child (inherits workflow context, adds agent identity)     child_context = parent_context.model_copy(update={         "agent_name": "router",         "agent_version": "1.0.0",         "agent_execution_id": workflow.info().workflow_id         # agent_workflow_* fields inherited from parent     })      # Tool execution (adds tool identity to agent context)     tool_context = context.model_copy(update={         'tool_id': "tool_xyz",         'tool_instance_id': "tool_inst_123",         'tool_execution_id': "tool_exec_456"     })
type AgentIdentityContext struct {
	// Customer/tenant identifier
	CustomerId string `json:"customer_id"`
	// Project identifier
	ProjectId                string      `json:"project_id"`
	AgentName                interface{} `json:"agent_name,omitempty"`
	AgentVersion             interface{} `json:"agent_version,omitempty"`
	AgentExecutionId         interface{} `json:"agent_execution_id,omitempty"`
	AgentWorkflowName        interface{} `json:"agent_workflow_name,omitempty"`
	AgentWorkflowVersion     interface{} `json:"agent_workflow_version,omitempty"`
	AgentWorkflowExecutionId interface{} `json:"agent_workflow_execution_id,omitempty"`
	ToolId                   interface{} `json:"tool_id,omitempty"`
	ToolInstanceId           interface{} `json:"tool_instance_id,omitempty"`
	ToolExecutionId          interface{} `json:"tool_execution_id,omitempty"`
}

// BoundingBox represents the BoundingBox type.
type BoundingBox struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Page   int64   `json:"page"`
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
}

// CloudStorageConnectorProvider represents the possible values for provider.
type CloudStorageConnectorProvider string

const (
	CloudStorageConnectorProviderS3  CloudStorageConnectorProvider = "s3"
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
	ActualInstance         *string     `json:"actual_instance,omitempty"`
	AnyOfSchemas           []string    `json:"any_of_schemas,omitempty"`
}

// ConnectorConfigInput Datasource connection configuration. Exactly one connector type must be set.
type ConnectorConfigInput struct {
	Type         string      `json:"type"`
	Database     interface{} `json:"database,omitempty"`
	CloudStorage interface{} `json:"cloud_storage,omitempty"`
	WebCrawl     interface{} `json:"web_crawl,omitempty"`
}

// ConnectorConfigInputType represents the possible values for type.
type ConnectorConfigInputType string

const (
	ConnectorConfigInputTypeDatabase     ConnectorConfigInputType = "database"
	ConnectorConfigInputTypeCloudStorage ConnectorConfigInputType = "cloud_storage"
	ConnectorConfigInputTypeWebCrawl     ConnectorConfigInputType = "web_crawl"
)

// ConnectorConfigOutput Datasource connection configuration. Exactly one connector type must be set.
type ConnectorConfigOutput struct {
	Type         string      `json:"type"`
	Database     interface{} `json:"database,omitempty"`
	CloudStorage interface{} `json:"cloud_storage,omitempty"`
	WebCrawl     interface{} `json:"web_crawl,omitempty"`
}

// ConnectorConfigOutputType represents the possible values for type.
type ConnectorConfigOutputType string

const (
	ConnectorConfigOutputTypeDatabase     ConnectorConfigOutputType = "database"
	ConnectorConfigOutputTypeCloudStorage ConnectorConfigOutputType = "cloud_storage"
	ConnectorConfigOutputTypeWebCrawl     ConnectorConfigOutputType = "web_crawl"
)

// CreateDataElementRequest represents the CreateDataElementRequest type.
type CreateDataElementRequest struct {
	// Data element name
	Name string `json:"name"`
	// MIME type of the data element
	MediaType interface{} `json:"media_type,omitempty"`
	// Arbitrary metadata
	Metadata interface{} `json:"metadata,omitempty"`
}

// CreateDatasourceRequest represents the CreateDatasourceRequest type.
type CreateDatasourceRequest struct {
	// Human-readable datasource name
	Name string `json:"name"`
	// What this datasource contains
	Description *string `json:"description,omitempty"`
	// Connection configuration
	Connector ConnectorConfigInput `json:"connector"`
}

// DataElementMetadata Metadata key-value pairs on a data element.
type DataElementMetadata struct {
	// Arbitrary key-value metadata
	Metadata string `json:"metadata"`
}

// DataElementResponse represents the DataElementResponse type.
type DataElementResponse struct {
	Id           string      `json:"id"`
	DatasourceId string      `json:"datasource_id"`
	Name         string      `json:"name"`
	MediaType    interface{} `json:"media_type,omitempty"`
	Metadata     interface{} `json:"metadata,omitempty"`
	CreatedAt    interface{} `json:"created_at,omitempty"`
	UpdatedAt    interface{} `json:"updated_at,omitempty"`
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
	Id          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Connector   ConnectorConfigOutput `json:"connector"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
}

// DocumentChild Child document from container (ZIP/TAR/EML).
type DocumentChild struct {
	JobId     string `json:"job_id"`
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	MediaType string `json:"media_type"`
}

// DocumentElement A structural element in a parsed document.
type DocumentElement struct {
	// heading | paragraph | table | list_item | image | code_block | ...
	Type string      `json:"type"`
	Text interface{} `json:"text,omitempty"`
	// Heading level (1-6)
	Level      interface{} `json:"level,omitempty"`
	Table      interface{} `json:"table,omitempty"`
	Bbox       interface{} `json:"bbox,omitempty"`
	Confidence interface{} `json:"confidence,omitempty"`
	Page       interface{} `json:"page,omitempty"`
}

// DocumentStatus Returned from GET /documents/{job_id}.
type DocumentStatus struct {
	JobId string `json:"job_id"`
	// queued | processing | completed | failed
	Status string `json:"status"`
	// meibel | markdown | docling
	Format           string      `json:"format"`
	Pages            interface{} `json:"pages,omitempty"`
	Elements         interface{} `json:"elements,omitempty"`
	Tables           interface{} `json:"tables,omitempty"`
	Confidence       interface{} `json:"confidence,omitempty"`
	ProcessingTimeMs interface{} `json:"processing_time_ms,omitempty"`
	Error            interface{} `json:"error,omitempty"`
}

// HttpValidationError represents the HTTPValidationError type.
type HttpValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}

// JudgeConfig Configuration for judge-based confidence scoring (LLM-as-judge patterns).
type JudgeConfig struct {
	Prompt          string      `json:"prompt"`
	TemperatureMax  interface{} `json:"temperature_max,omitempty"`
	TemperatureStep interface{} `json:"temperature_step,omitempty"`
}

// ListMetadataModelCatalogResponse ListMetadataModelCatalogResponse
type ListMetadataModelCatalogResponse struct {
	Models []MetadataModelCatalogEntry `json:"models"`
}

// MeibelDocumentResult Full structured parse result (meibel format).
type MeibelDocumentResult struct {
	Elements []DocumentElement `json:"elements"`
	Pages    int64             `json:"pages"`
	Tables   int64             `json:"tables"`
	Metadata interface{}       `json:"metadata,omitempty"`
}

// MetadataConfigRequest Configure automatic metadata extraction from documents on ingest.
type MetadataConfigRequest struct {
	Type string `json:"type"`
	// Required when type='catalog'
	ModelId interface{} `json:"model_id,omitempty"`
	// Required when type='custom'
	Fields interface{} `json:"fields,omitempty"`
}

// MetadataConfigRequestType represents the possible values for type.
type MetadataConfigRequestType string

const (
	MetadataConfigRequestTypeCatalog MetadataConfigRequestType = "catalog"
	MetadataConfigRequestTypeCustom  MetadataConfigRequestType = "custom"
)

// MetadataConfigResponse represents the MetadataConfigResponse type.
type MetadataConfigResponse struct {
	Type    string          `json:"type"`
	ModelId interface{}     `json:"model_id,omitempty"`
	Fields  []MetadataField `json:"fields"`
}

// MetadataConfigResponseType represents the possible values for type.
type MetadataConfigResponseType string

const (
	MetadataConfigResponseTypeCatalog MetadataConfigResponseType = "catalog"
	MetadataConfigResponseTypeCustom  MetadataConfigResponseType = "custom"
	MetadataConfigResponseTypeDefault MetadataConfigResponseType = "default"
)

// MetadataField represents the MetadataField type.
type MetadataField struct {
	// Field name (snake_case)
	Name string `json:"name"`
	Type string `json:"type"`
	// What this field captures
	Description string `json:"description"`
}

// MetadataFieldType represents the possible values for type.
type MetadataFieldType string

const (
	MetadataFieldTypeString     MetadataFieldType = "string"
	MetadataFieldTypeNumber     MetadataFieldType = "number"
	MetadataFieldTypeBoolean    MetadataFieldType = "boolean"
	MetadataFieldTypeListString MetadataFieldType = "list[string]"
)

// MetadataModelCatalogEntry MetadataModelCatalogEntry
type MetadataModelCatalogEntry struct {
	ModelId     string               `json:"model_id"`
	Name        string               `json:"name"`
	Description interface{}          `json:"description,omitempty"`
	Scope       string               `json:"scope"`
	CustomerId  interface{}          `json:"customer_id,omitempty"`
	ProjectId   interface{}          `json:"project_id,omitempty"`
	Fields      []MetadataModelField `json:"fields"`
	CreatedBy   interface{}          `json:"created_by,omitempty"`
	UpdatedBy   interface{}          `json:"updated_by,omitempty"`
	CreatedAt   interface{}          `json:"created_at,omitempty"`
	UpdatedAt   interface{}          `json:"updated_at,omitempty"`
}

// MetadataModelField MetadataModelField
type MetadataModelField struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// NBootstraps NBootstraps
type NBootstraps struct {
	AnyofSchema_1Validator interface{} `json:"anyof_schema_1_validator,omitempty"`
	AnyofSchema_2Validator interface{} `json:"anyof_schema_2_validator,omitempty"`
	ActualInstance         *string     `json:"actual_instance,omitempty"`
	AnyOfSchemas           []string    `json:"any_of_schemas,omitempty"`
}

// OcConfig Configuration for Observed Consistency confidence scoring.
type OcConfig struct {
	NCompletions          interface{} `json:"n_completions,omitempty"`
	MaxTokens             interface{} `json:"max_tokens,omitempty"`
	Temperature           interface{} `json:"temperature,omitempty"`
	Models                interface{} `json:"models,omitempty"`
	NliModelConfig        string      `json:"nli_model_config"`
	NBootstraps           interface{} `json:"n_bootstraps,omitempty"`
	TokenLimit            interface{} `json:"token_limit,omitempty"`
	OriginalCompletion    interface{} `json:"original_completion,omitempty"`
	ComparisonCompletions interface{} `json:"comparison_completions,omitempty"`
}

// OcrConfig Configuration for OCR confidence scoring.
type OcrConfig struct {
	CalibrationModel    interface{} `json:"calibration_model,omitempty"`
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

// ScoreSummary Aggregated summary of scoring jobs matching one or two AgentIdentityContext filters.  With one level (primary only): flat aggregate of all jobs matching primary_field=primary_value. With two levels (primary + secondary): both constraints are applied; primary is the higher level and secondary is the lower level.
type ScoreSummary struct {
	PrimaryField   string      `json:"primary_field"`
	PrimaryValue   string      `json:"primary_value"`
	SecondaryField interface{} `json:"secondary_field,omitempty"`
	SecondaryValue interface{} `json:"secondary_value,omitempty"`
	Status         interface{} `json:"status,omitempty"`
	AggregateScore interface{} `json:"aggregate_score,omitempty"`
	ModuleScores   interface{} `json:"module_scores,omitempty"`
	NJobsPerModule interface{} `json:"n_jobs_per_module,omitempty"`
	Jobs           interface{} `json:"jobs,omitempty"`
}

// ScoringJobRecord ScoringJobRecord
type ScoringJobRecord struct {
	JobId                string                  `json:"job_id"`
	AgentIdentityContext AgentIdentityContext    `json:"agent_identity_context"`
	Module               string                  `json:"module"`
	ScoringConfig        ConfidenceScoringConfig `json:"scoring_config"`
	InputValue           string                  `json:"input_value"`
	OutputValue          string                  `json:"output_value"`
	Status               ScoringStatus           `json:"status"`
	Score                interface{}             `json:"score,omitempty"`
}

// ScoringStatus represents the possible values for ScoringStatus.
type ScoringStatus string

const (
	ScoringStatusSubmitted  ScoringStatus = "submitted"
	ScoringStatusInProgress ScoringStatus = "in_progress"
	ScoringStatusCompleted  ScoringStatus = "completed"
	ScoringStatusFailed     ScoringStatus = "failed"
)

// Table represents the Table type.
type Table struct {
	Cells []TableCell `json:"cells"`
	Rows  int64       `json:"rows"`
	Cols  int64       `json:"cols"`
	Bbox  interface{} `json:"bbox,omitempty"`
}

// TableCell represents the TableCell type.
type TableCell struct {
	Text    string      `json:"text"`
	Row     int64       `json:"row"`
	Col     int64       `json:"col"`
	RowSpan *int64      `json:"row_span,omitempty"`
	ColSpan *int64      `json:"col_span,omitempty"`
	Bbox    interface{} `json:"bbox,omitempty"`
}

// TagColumn represents the TagColumn type.
type TagColumn struct {
	ColumnName  string      `json:"column_name"`
	Description interface{} `json:"description,omitempty"`
}

// TagTable represents the TagTable type.
type TagTable struct {
	TableName   string      `json:"table_name"`
	Description interface{} `json:"description,omitempty"`
}

// TokenConfig Configuration for token-based confidence scoring (TF-IDF).
type TokenConfig struct {
	Model           interface{} `json:"model,omitempty"`
	RemoveStopWords interface{} `json:"remove_stop_words,omitempty"`
	LowerCase       interface{} `json:"lower_case,omitempty"`
	MaxNgrams       interface{} `json:"max_ngrams,omitempty"`
	NInfluencers    interface{} `json:"n_influencers,omitempty"`
}

// UpdateTagDescriptionRequest represents the UpdateTagDescriptionRequest type.
type UpdateTagDescriptionRequest struct {
	// Description for AI context
	Description string `json:"description"`
}

// WebCrawlConnector Connect to a website for crawling.
type WebCrawlConnector struct {
	// Starting URL for the crawl
	BaseUrl string `json:"base_url"`
	// Enable JavaScript rendering
	JavascriptRender *bool       `json:"javascript_render,omitempty"`
	Domains          interface{} `json:"domains,omitempty"`
}

// UpdateDataElementRequest represents the UpdateDataElementRequest type.
type UpdateDataElementRequest struct {
	Name     interface{} `json:"name,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// UpdateDatasourceRequest represents the UpdateDatasourceRequest type.
type UpdateDatasourceRequest struct {
	Name        interface{} `json:"name,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Connector   interface{} `json:"connector,omitempty"`
}

// WebDomain represents the WebDomain type.
type WebDomain struct {
	Domain string `json:"domain"`
	// URL pattern to include
	IncludePattern string `json:"include_pattern"`
	// URL pattern to exclude
	ExcludePattern *string `json:"exclude_pattern,omitempty"`
}
