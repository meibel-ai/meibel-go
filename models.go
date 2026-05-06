package meibelgo

// Activity Activity
type Activity struct {
	Id interface{} `json:"id,omitempty"`
	BlueprintInstanceId string `json:"blueprint_instance_id"`
	ActivityType string `json:"activity_type"`
	Status interface{} `json:"status,omitempty"`
	StartTime interface{} `json:"start_time,omitempty"`
	EndTime interface{} `json:"end_time,omitempty"`
	InputData interface{} `json:"input_data,omitempty"`
	OutputData interface{} `json:"output_data,omitempty"`
	Error interface{} `json:"error,omitempty"`
	GroupId interface{} `json:"group_id,omitempty"`
	TaskMetadata interface{} `json:"task_metadata,omitempty"`
}

// ActivityStatus represents the possible values for ActivityStatus.
type ActivityStatus string

const (
	ActivityStatusPending ActivityStatus = "pending"
	ActivityStatusRunning ActivityStatus = "running"
	ActivityStatusWaiting ActivityStatus = "waiting"
	ActivityStatusCompleted ActivityStatus = "completed"
	ActivityStatusFailed ActivityStatus = "failed"
)

// AddActivityRequest AddActivityRequest
type AddActivityRequest struct {
	ActivityType string `json:"activity_type"`
	InputData interface{} `json:"input_data,omitempty"`
	OutputData interface{} `json:"output_data,omitempty"`
	GroupId interface{} `json:"group_id,omitempty"`
	TaskMetadata interface{} `json:"task_metadata,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// AddActivityResponse AddActivityResponse
type AddActivityResponse struct {
	Id string `json:"id"`
}

// AddBlueprintInstanceRequest AddBlueprintInstanceRequest
type AddBlueprintInstanceRequest struct {
	BlueprintId interface{} `json:"blueprint_id,omitempty"`
	WorkflowType interface{} `json:"workflow_type,omitempty"`
	TaskQueue interface{} `json:"task_queue,omitempty"`
	InstanceMetadata interface{} `json:"instance_metadata,omitempty"`
	ParentId interface{} `json:"parent_id,omitempty"`
}

// AddBlueprintInstanceResponse AddBlueprintInstanceResponse
type AddBlueprintInstanceResponse struct {
	Id string `json:"id"`
}

// AddBlueprintRequest AddBlueprintRequest
type AddBlueprintRequest struct {
	Name string `json:"name"`
	ExecutionMode interface{} `json:"execution_mode,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Description interface{} `json:"description,omitempty"`
	DslDefinition DslDefinition `json:"dsl_definition"`
	YamlSpecContent interface{} `json:"yaml_spec_content,omitempty"`
	JsonSpecContent interface{} `json:"json_spec_content,omitempty"`
	WorkflowType interface{} `json:"workflow_type,omitempty"`
	WorkflowTaskQueue interface{} `json:"workflow_task_queue,omitempty"`
	InitInput interface{} `json:"init_input,omitempty"`
}

// AddBlueprintResponse AddBlueprintResponse
type AddBlueprintResponse struct {
	Id string `json:"id"`
}

// AddBlueprintTaskRequest AddBlueprintTaskRequest
type AddBlueprintTaskRequest struct {
	Name string `json:"name"`
	Type interface{} `json:"type,omitempty"`
	Description interface{} `json:"description,omitempty"`
	InputSchema string `json:"input_schema"`
	OutputSchema string `json:"output_schema"`
	ConfigSchema interface{} `json:"config_schema,omitempty"`
	ToolSchema interface{} `json:"tool_schema,omitempty"`
}

// AddChunkingStrategyRequest AddChunkingStrategyRequest
type AddChunkingStrategyRequest struct {
	CodeSplitter interface{} `json:"code_splitter,omitempty"`
	HtmlNodeParser interface{} `json:"html_node_parser,omitempty"`
	JsonNodeParser interface{} `json:"json_node_parser,omitempty"`
	MarkdownNodeParser interface{} `json:"markdown_node_parser,omitempty"`
	SemanticSplitterNodeParser interface{} `json:"semantic_splitter_node_parser,omitempty"`
	SentenceSplitter interface{} `json:"sentence_splitter,omitempty"`
	TokenTextSplitter interface{} `json:"token_text_splitter,omitempty"`
}

// AddChunkingStrategyResponse AddChunkingStrategyResponse
type AddChunkingStrategyResponse struct {
	Message string `json:"message"`
}

// AddDataElementRequest AddDataElementRequest
type AddDataElementRequest struct {
	Description interface{} `json:"description"`
	Name string `json:"name"`
	Path string `json:"path"`
	MediaType string `json:"media_type"`
	DiscoveryRecord interface{} `json:"discovery_record"`
	ParentDataElementId interface{} `json:"parent_data_element_id,omitempty"`
}

// AddDataElementResponse AddDataElementResponse
type AddDataElementResponse struct {
	Id string `json:"id"`
}

// AddDatasourceRequest AddDatasourceRequest
type AddDatasourceRequest struct {
	CustomerId string `json:"customer_id"`
	ProjectId string `json:"project_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Recurrence string `json:"recurrence"`
	ObjectStorageConfig interface{} `json:"object_storage_config,omitempty"`
	WebConfig interface{} `json:"web_config,omitempty"`
	ConnectorConfig interface{} `json:"connector_config,omitempty"`
}

// AddDatasourceResponse AddDatasourceResponse
type AddDatasourceResponse struct {
	Id string `json:"id"`
}

// AddEventResponse AddEventResponse
type AddEventResponse struct {
	Id string `json:"id"`
}

// AddRagConfigRequest AddRagConfigRequest
type AddRagConfigRequest struct {
	Description interface{} `json:"description,omitempty"`
	CollectionId string `json:"collection_id"`
	ExtractorModel interface{} `json:"extractor_model,omitempty"`
	EmbeddingModel interface{} `json:"embedding_model,omitempty"`
	SparseEmbeddingModel interface{} `json:"sparse_embedding_model,omitempty"`
	CollectMetadata interface{} `json:"collect_metadata,omitempty"`
	MetadataOptions interface{} `json:"metadata_options,omitempty"`
}

// AddRagConfigResponse AddRagConfigResponse
type AddRagConfigResponse struct {
	Message string `json:"message"`
}

// AddTagColumnRequest AddTagColumnRequest
type AddTagColumnRequest struct {
	Description interface{} `json:"description,omitempty"`
	Dtype interface{} `json:"dtype,omitempty"`
	IsKey interface{} `json:"is_key,omitempty"`
	IsIndexed interface{} `json:"is_indexed,omitempty"`
	EngineeredFeatures interface{} `json:"engineered_features,omitempty"`
}

// AddTagColumnResponse AddTagColumnResponse
type AddTagColumnResponse struct {
	Message string `json:"message"`
}

// AddTagConfigRequest AddTagConfigRequest
type AddTagConfigRequest struct {
	Description interface{} `json:"description,omitempty"`
	LogicalGroupRegex interface{} `json:"logical_group_regex,omitempty"`
	WorkingBucket string `json:"working_bucket"`
	WorkingPlatform interface{} `json:"working_platform,omitempty"`
	DbPath interface{} `json:"db_path,omitempty"`
	DatabaseConfig interface{} `json:"database_config,omitempty"`
}

// AddTagConfigResponse AddTagConfigResponse
type AddTagConfigResponse struct {
	Message string `json:"message"`
}

// AddTagTableRequest AddTagTableRequest
type AddTagTableRequest struct {
	Description interface{} `json:"description,omitempty"`
}

// AddTagTableResponse AddTagTableResponse
type AddTagTableResponse struct {
	Message string `json:"message"`
}

// AirbyteConfig AirbyteConfig
type AirbyteConfig struct {
	SourceId string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	ConnectionId interface{} `json:"connection_id,omitempty"`
}

// AllowedDataElementFilterKeys represents the possible values for AllowedDataElementFilterKeys.
type AllowedDataElementFilterKeys string

const (
	AllowedDataElementFilterKeysId AllowedDataElementFilterKeys = "id"
	AllowedDataElementFilterKeysName AllowedDataElementFilterKeys = "name"
	AllowedDataElementFilterKeysPath AllowedDataElementFilterKeys = "path"
	AllowedDataElementFilterKeysMediaType AllowedDataElementFilterKeys = "media_type"
	AllowedDataElementFilterKeysParentDataElementId AllowedDataElementFilterKeys = "parent_data_element_id"
)

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

// BasicWebAuth BasicWebAuth
type BasicWebAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Blueprint Blueprint
type Blueprint struct {
	Id interface{} `json:"id,omitempty"`
	Name string `json:"name"`
	ExecutionMode interface{} `json:"execution_mode,omitempty"`
	ChatSignal interface{} `json:"chat_signal,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Description interface{} `json:"description,omitempty"`
	DslDefinition DslDefinition `json:"dsl_definition"`
	YamlSpecContent interface{} `json:"yaml_spec_content,omitempty"`
	JsonSpecContent interface{} `json:"json_spec_content,omitempty"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	UpdatedBy interface{} `json:"updated_by,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	CustomerId string `json:"customer_id"`
	ProjectId string `json:"project_id"`
	WorkflowType interface{} `json:"workflow_type,omitempty"`
	WorkflowTaskQueue interface{} `json:"workflow_task_queue,omitempty"`
	InitInput interface{} `json:"init_input,omitempty"`
}

// BlueprintExecutionMode represents the possible values for BlueprintExecutionMode.
type BlueprintExecutionMode string

const (
	BlueprintExecutionModeChat BlueprintExecutionMode = "CHAT"
	BlueprintExecutionModeWorkflow BlueprintExecutionMode = "WORKFLOW"
)

// BlueprintInstance BlueprintInstance
type BlueprintInstance struct {
	Id interface{} `json:"id,omitempty"`
	BlueprintId interface{} `json:"blueprint_id,omitempty"`
	WorkflowType interface{} `json:"workflow_type,omitempty"`
	TaskQueue interface{} `json:"task_queue,omitempty"`
	WorkflowRunId interface{} `json:"workflow_run_id,omitempty"`
	Status interface{} `json:"status,omitempty"`
	StartTime interface{} `json:"start_time,omitempty"`
	EndTime interface{} `json:"end_time,omitempty"`
	InstanceMetadata interface{} `json:"instance_metadata,omitempty"`
	ParentId interface{} `json:"parent_id,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Activities interface{} `json:"activities,omitempty"`
	Events interface{} `json:"events,omitempty"`
}

// BlueprintInstanceStatus represents the possible values for BlueprintInstanceStatus.
type BlueprintInstanceStatus string

const (
	BlueprintInstanceStatusCreated BlueprintInstanceStatus = "created"
	BlueprintInstanceStatusRunning BlueprintInstanceStatus = "running"
	BlueprintInstanceStatusCompleted BlueprintInstanceStatus = "completed"
	BlueprintInstanceStatusCancelled BlueprintInstanceStatus = "cancelled"
	BlueprintInstanceStatusFailed BlueprintInstanceStatus = "failed"
)

// CallToAction An action the user can take.
type CallToAction struct {
	Label string `json:"label"`
	Action string `json:"action"`
	ActionData interface{} `json:"action_data,omitempty"`
}

// ChatMessageRequest Request body for chat message endpoints.
type ChatMessageRequest struct {
	// The user's chat message
	UserMessage string `json:"user_message"`
	TimeoutSeconds interface{} `json:"timeout_seconds,omitempty"`
	IncludeThinking interface{} `json:"include_thinking,omitempty"`
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
	ToolActivity interface{} `json:"tool_activity,omitempty"`
	Thinking interface{} `json:"thinking,omitempty"`
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

// ClickhouseConfig ClickhouseConfig
type ClickhouseConfig struct {
	DatabaseType interface{} `json:"database_type,omitempty"`
	DatabaseName string `json:"database_name"`
}

// CodeChunking CodeChunking
type CodeChunking struct {
	ChunkLines interface{} `json:"chunk_lines,omitempty"`
	ChunkLinesOverlap interface{} `json:"chunk_lines_overlap,omitempty"`
	MaxChars interface{} `json:"max_chars,omitempty"`
}

// CompleteBlueprintInstanceRequest CompleteBlueprintInstanceRequest
type CompleteBlueprintInstanceRequest struct {
	Result interface{} `json:"result,omitempty"`
}

// ContentItem ContentItem
type ContentItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type ContentType `json:"type"`
	Size interface{} `json:"size,omitempty"`
	MediaType interface{} `json:"media_type,omitempty"`
	LastModified interface{} `json:"last_modified,omitempty"`
	Etag interface{} `json:"etag,omitempty"`
}

// ContentType represents the possible values for ContentType.
type ContentType string

const (
	ContentTypeFile ContentType = "file"
	ContentTypeDirectory ContentType = "directory"
)

// CustomEventRequest CustomEventRequest
type CustomEventRequest struct {
	ActivityId interface{} `json:"activity_id,omitempty"`
	// Name of the custom event being logged.
	EventName string `json:"event_name"`
	Details interface{} `json:"details,omitempty"`
	GroupId interface{} `json:"group_id,omitempty"`
	IsSignal interface{} `json:"is_signal,omitempty"`
	IsInternal interface{} `json:"is_internal,omitempty"`
	OriginatingSignalId interface{} `json:"originating_signal_id,omitempty"`
}

// DataElement DataElement
type DataElement struct {
	Id interface{} `json:"id"`
	DatasourceId string `json:"datasource_id"`
	Name string `json:"name"`
	Path string `json:"path"`
	MediaType string `json:"media_type"`
	DiscoveryRecord interface{} `json:"discovery_record"`
	Description interface{} `json:"description"`
	ParentDataElementId interface{} `json:"parent_data_element_id,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// DataElementCondition represents the possible values for DataElementCondition.
type DataElementCondition string

const (
	DataElementConditionEquals DataElementCondition = "equals"
	DataElementConditionContains DataElementCondition = "contains"
	DataElementConditionNotContains DataElementCondition = "not_contains"
	DataElementConditionStartsWith DataElementCondition = "starts_with"
	DataElementConditionEndsWith DataElementCondition = "ends_with"
	DataElementConditionIn DataElementCondition = "in"
	DataElementConditionRegexGroupMatch DataElementCondition = "regex_group_match"
)

// DataElementDiscoveryRecord DataElementDiscoveryRecord
type DataElementDiscoveryRecord struct {
	DiscoveryTime string `json:"discovery_time"`
	LastModifiedTime string `json:"last_modified_time"`
	Size interface{} `json:"size"`
	ElementHash string `json:"element_hash"`
	FileId interface{} `json:"file_id,omitempty"`
	FileCreatedAt interface{} `json:"file_created_at,omitempty"`
	FileModifiedAt interface{} `json:"file_modified_at,omitempty"`
	// Connector-specific extra metadata
	Extra interface{} `json:"extra,omitempty"`
}

// DataElementFilter DataElementFilter
type DataElementFilter struct {
	Key AllowedDataElementFilterKeys `json:"key"`
	Condition interface{} `json:"condition,omitempty"`
	Value interface{} `json:"value"`
}

// DataElementFilterRequest DataElementFilterRequest
type DataElementFilterRequest struct {
	Filters interface{} `json:"filters,omitempty"`
}

// DatabaseConfigInput DatabaseConfig
type DatabaseConfigInput struct {
	AnyofSchema_1Validator interface{} `json:"anyof_schema_1_validator,omitempty"`
	AnyofSchema_2Validator interface{} `json:"anyof_schema_2_validator,omitempty"`
	AnyofSchema_3Validator interface{} `json:"anyof_schema_3_validator,omitempty"`
	ActualInstance *string `json:"actual_instance,omitempty"`
	AnyOfSchemas []string `json:"any_of_schemas,omitempty"`
}

// DatabaseConfigOutput DatabaseConfig
type DatabaseConfigOutput struct {
	AnyofSchema_1Validator interface{} `json:"anyof_schema_1_validator,omitempty"`
	AnyofSchema_2Validator interface{} `json:"anyof_schema_2_validator,omitempty"`
	AnyofSchema_3Validator interface{} `json:"anyof_schema_3_validator,omitempty"`
	ActualInstance *string `json:"actual_instance,omitempty"`
	AnyOfSchemas []string `json:"any_of_schemas,omitempty"`
}

// DatabaseType represents the possible values for DatabaseType.
type DatabaseType string

const (
	DatabaseTypeClickhouse DatabaseType = "clickhouse"
	DatabaseTypeDuckdb DatabaseType = "duckdb"
	DatabaseTypePostgres DatabaseType = "postgres"
)

// Datasource Datasource
type Datasource struct {
	Id interface{} `json:"id"`
	CustomerId string `json:"customer_id"`
	ProjectId string `json:"project_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Recurrence interface{} `json:"recurrence"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedBy string `json:"updated_by"`
	UpdatedAt string `json:"updated_at"`
	ObjectStorageConfig interface{} `json:"object_storage_config"`
	WebConfig interface{} `json:"web_config"`
	ConnectorConfig interface{} `json:"connector_config"`
}

// DatasourceConnectorConfig DatasourceConnectorConfig
type DatasourceConnectorConfig struct {
	ConnectorId string `json:"connector_id"`
	AirbyteConfig interface{} `json:"airbyte_config,omitempty"`
	SourceConfig interface{} `json:"source_config,omitempty"`
}

// DatasourceWebConfig DatasourceWebConfig
type DatasourceWebConfig struct {
	BaseUrl string `json:"baseURL"`
	JavascriptRender bool `json:"javascriptRender"`
	WaitForSelector interface{} `json:"waitForSelector"`
	Domains interface{} `json:"domains"`
	Authentication interface{} `json:"authentication"`
}

// DeleteChunkingStrategyResponse DeleteChunkingStrategyResponse
type DeleteChunkingStrategyResponse struct {
	Message string `json:"message"`
}

// DeleteContentResponse DeleteContentResponse
type DeleteContentResponse struct {
	Success bool `json:"success"`
	DeletedPaths []string `json:"deleted_paths"`
}

// DeleteDataElementResponse DeleteDataElementResponse
type DeleteDataElementResponse struct {
	Id string `json:"id"`
}

// DeleteDatasourceResponse DeleteDatasourceResponse
type DeleteDatasourceResponse struct {
	Id string `json:"id"`
}

// DeleteTagTableResponse DeleteTagTableResponse
type DeleteTagTableResponse struct {
	Message string `json:"message"`
}

// DslDefinition represents the possible values for DslDefinition.
type DslDefinition string

const (
	DslDefinitionServerlessWorkflowV1_0_0 DslDefinition = "serverless-workflow_v1.0.0"
)

// DuckDbConfig DuckDBConfig
type DuckDbConfig struct {
	DatabaseType interface{} `json:"database_type,omitempty"`
	DatabaseFilepath interface{} `json:"database_filepath,omitempty"`
	DatabaseName interface{} `json:"database_name,omitempty"`
	DatabaseSchema interface{} `json:"database_schema,omitempty"`
}

// EmbeddingModel EmbeddingModel
type EmbeddingModel struct {
	Name string `json:"name"`
	Endpoint string `json:"endpoint"`
	Dimensions int64 `json:"dimensions"`
}

// Event Event
type Event struct {
	Id interface{} `json:"id,omitempty"`
	ActivityId interface{} `json:"activity_id,omitempty"`
	BlueprintInstanceId string `json:"blueprint_instance_id"`
	EventType interface{} `json:"event_type,omitempty"`
	Timestamp interface{} `json:"timestamp,omitempty"`
	Details interface{} `json:"details,omitempty"`
	GroupId interface{} `json:"group_id,omitempty"`
	IsSignal interface{} `json:"is_signal,omitempty"`
	IsInternal interface{} `json:"is_internal,omitempty"`
	OriginatingSignalId interface{} `json:"originating_signal_id,omitempty"`
}

// EventType represents the possible values for EventType.
type EventType string

const (
	EventTypeInstanceStarted EventType = "instance_started"
	EventTypeInstanceCompleted EventType = "instance_completed"
	EventTypeInstanceFailed EventType = "instance_failed"
	EventTypeInstanceCancelled EventType = "instance_cancelled"
	EventTypeActivityStarted EventType = "activity_started"
	EventTypeActivityCompleted EventType = "activity_completed"
	EventTypeActivityFailed EventType = "activity_failed"
	EventTypeUserInput EventType = "user_input"
	EventTypeAgentOutput EventType = "agent_output"
	EventTypeSignalReceived EventType = "signal_received"
	EventTypeQueryReceived EventType = "query_received"
	EventTypeWorkflowStarted EventType = "workflow_started"
	EventTypeWorkflowCompleted EventType = "workflow_completed"
	EventTypeWorkflowFailed EventType = "workflow_failed"
	EventTypeChildWorkflowStarted EventType = "child_workflow_started"
	EventTypeChildWorkflowCompleted EventType = "child_workflow_completed"
	EventTypeChildWorkflowFailed EventType = "child_workflow_failed"
	EventTypeAgentThinking EventType = "agent_thinking"
	EventTypeAgentToolUsage EventType = "agent_tool_usage"
	EventTypeAgentStateChange EventType = "agent_state_change"
	EventTypeAgentDecision EventType = "agent_decision"
	EventTypeAgentSubtaskStarted EventType = "agent_subtask_started"
	EventTypeAgentSubtaskCompleted EventType = "agent_subtask_completed"
	EventTypeAgentError EventType = "agent_error"
	EventTypeAgentLog EventType = "agent_log"
	EventTypeCustomEvent EventType = "custom_event"
)

// ExecuteBlueprintRequest ExecuteBlueprintRequest
type ExecuteBlueprintRequest struct {
	InitInput interface{} `json:"init_input,omitempty"`
	EnableStreaming interface{} `json:"enable_streaming,omitempty"`
}

// ExtractorModel ExtractorModel
type ExtractorModel struct {
	Name string `json:"name"`
	Endpoint string `json:"endpoint"`
}

// FailBlueprintInstanceRequest FailBlueprintInstanceRequest
type FailBlueprintInstanceRequest struct {
	// Error message for failure
	Error interface{} `json:"error,omitempty"`
	ErrorDetails interface{} `json:"error_details,omitempty"`
}

// GetActivitiesResponse GetActivitiesResponse
type GetActivitiesResponse struct {
	Data []Activity `json:"data"`
}

// GetAllDatasourceIdsResponse GetAllDatasourceIdsResponse
type GetAllDatasourceIdsResponse struct {
	DatasourceIds []string `json:"datasource_ids"`
}

// GetBlueprintInstancesResponse GetBlueprintInstancesResponse
type GetBlueprintInstancesResponse struct {
	Data []BlueprintInstance `json:"data"`
}

// GetBlueprintsResponse GetBlueprintsResponse
type GetBlueprintsResponse struct {
	Data []Blueprint `json:"data"`
}

// GetContentResponse GetContentResponse
type GetContentResponse struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type ContentType `json:"type"`
	Size interface{} `json:"size,omitempty"`
	MediaType interface{} `json:"media_type,omitempty"`
	LastModified interface{} `json:"last_modified,omitempty"`
	Etag interface{} `json:"etag,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// GetEventsResponse GetEventsResponse
type GetEventsResponse struct {
	Data []Event `json:"data"`
}

// HtmlChunking HTMLChunking
type HtmlChunking struct {
	IncludeMetadata bool `json:"include_metadata"`
	IncludePrevNextRel bool `json:"include_prev_next_rel"`
	Tags interface{} `json:"tags,omitempty"`
}

// HttpValidationError represents the HTTPValidationError type.
type HttpValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}

// JsonNodeChunking JSONNodeChunking
type JsonNodeChunking struct {
	IncludeMetadata bool `json:"include_metadata"`
	IncludePrevNextRel bool `json:"include_prev_next_rel"`
}

// ListContentResponse ListContentResponse
type ListContentResponse struct {
	Items []ContentItem `json:"items"`
	ContinuationToken interface{} `json:"continuation_token,omitempty"`
}

// MarkdownNodeChunking MarkdownNodeChunking
type MarkdownNodeChunking struct {
	IncludeMetadata bool `json:"include_metadata"`
	IncludePrevNextRel bool `json:"include_prev_next_rel"`
	HeaderPathSeparator string `json:"header_path_separator"`
}

// MetadataOptions MetadataOptions
type MetadataOptions struct {
	CreateTitle interface{} `json:"create_title,omitempty"`
	ExtractQuestionsAnswers interface{} `json:"extract_questions_answers,omitempty"`
	ExtractSummary interface{} `json:"extract_summary,omitempty"`
	HasConsumerContent interface{} `json:"has_consumer_content,omitempty"`
	GetBibliographicalInformation interface{} `json:"get_bibliographical_information,omitempty"`
}

// ObjectStorageConfig ObjectStorageConfig
type ObjectStorageConfig struct {
	Bucket string `json:"bucket"`
	Prefix interface{} `json:"prefix"`
	Filters interface{} `json:"filters,omitempty"`
	GcsConfig interface{} `json:"gcs_config,omitempty"`
	S3Config interface{} `json:"s3_config,omitempty"`
}

// ObjectStorageFilters ObjectStorageFilters
type ObjectStorageFilters struct {
	IncludedPrefixes interface{} `json:"included_prefixes,omitempty"`
	IncludedFileTypes interface{} `json:"included_file_types,omitempty"`
	RecursivePrefixes interface{} `json:"recursive_prefixes,omitempty"`
	ModifiedDateStart interface{} `json:"modified_date_start,omitempty"`
	ModifiedDateEnd interface{} `json:"modified_date_end,omitempty"`
	MinFileSize interface{} `json:"min_file_size,omitempty"`
	MaxFileSize interface{} `json:"max_file_size,omitempty"`
}

// Platform represents the possible values for Platform.
type Platform string

const (
	PlatformAzure Platform = "azure"
	PlatformGcs Platform = "gcs"
	PlatformR2 Platform = "r2"
	PlatformS3 Platform = "s3"
)

// PostgreSqlConfig PostgreSQLConfig
type PostgreSqlConfig struct {
	DatabaseType interface{} `json:"database_type,omitempty"`
	DatabaseName string `json:"database_name"`
}

// RagChunkingStrategy RagChunkingStrategy
type RagChunkingStrategy struct {
	DatasourceId string `json:"datasource_id"`
	CodeSplitter interface{} `json:"code_splitter,omitempty"`
	HtmlNodeParser interface{} `json:"html_node_parser,omitempty"`
	JsonNodeParser interface{} `json:"json_node_parser,omitempty"`
	MarkdownNodeParser interface{} `json:"markdown_node_parser,omitempty"`
	SemanticSplitterNodeParser interface{} `json:"semantic_splitter_node_parser,omitempty"`
	SentenceSplitter interface{} `json:"sentence_splitter,omitempty"`
	TokenTextSplitter interface{} `json:"token_text_splitter,omitempty"`
}

// RagConfig RagConfig
type RagConfig struct {
	DatasourceId string `json:"datasource_id"`
	Description string `json:"description"`
	CollectionId string `json:"collection_id"`
	ExtractorModel interface{} `json:"extractor_model,omitempty"`
	EmbeddingModel interface{} `json:"embedding_model,omitempty"`
	SparseEmbeddingModel interface{} `json:"sparse_embedding_model,omitempty"`
	CollectMetadata interface{} `json:"collect_metadata,omitempty"`
	MetadataOptions interface{} `json:"metadata_options,omitempty"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	UpdatedBy interface{} `json:"updated_by,omitempty"`
}

// S3Config S3Config
type S3Config struct {
	RoleArn string `json:"role_arn"`
	Region string `json:"region"`
}

// SemanticChunking SemanticChunking
type SemanticChunking struct {
	BufferSize interface{} `json:"buffer_size,omitempty"`
	IncludeMetadata bool `json:"include_metadata"`
	IncludePrevNextRel bool `json:"include_prev_next_rel"`
	BreakpointPercentileThreshold interface{} `json:"breakpoint_percentile_threshold,omitempty"`
}

// SentenceChunking SentenceChunking
type SentenceChunking struct {
	ChunkSize interface{} `json:"chunk_size,omitempty"`
	ChunkOverlap interface{} `json:"chunk_overlap,omitempty"`
	Separator interface{} `json:"separator,omitempty"`
	ParagraphSeparator interface{} `json:"paragraph_separator,omitempty"`
	SecondaryChunkingRegex interface{} `json:"secondary_chunking_regex,omitempty"`
}

// Source A source/citation in the response.
type Source struct {
	Title string `json:"title"`
	Url interface{} `json:"url,omitempty"`
	Snippet interface{} `json:"snippet,omitempty"`
	DataElementId interface{} `json:"data_element_id,omitempty"`
	RelevanceScore interface{} `json:"relevance_score,omitempty"`
}

// SparseEmbeddingModel SparseEmbeddingModel
type SparseEmbeddingModel struct {
	Name string `json:"name"`
	Endpoint string `json:"endpoint"`
}

// StartBlueprintInstanceRequest StartBlueprintInstanceRequest
type StartBlueprintInstanceRequest struct {
	WorkflowArgs interface{} `json:"workflow_args,omitempty"`
	WorkflowKwargs interface{} `json:"workflow_kwargs,omitempty"`
	// Enable streaming responses to Redis for chat workflows. When True, chat responses are streamed to Redis streams that can be consumed via the /chat/stream endpoint.
	EnableStreaming interface{} `json:"enable_streaming,omitempty"`
}

// TagColumnInfo TagColumnInfo
type TagColumnInfo struct {
	DatasourceId string `json:"datasource_id"`
	TableName string `json:"table_name"`
	Name string `json:"name"`
	Description interface{} `json:"description"`
	Dtype interface{} `json:"dtype"`
	IsKey interface{} `json:"is_key"`
	IsIndexed interface{} `json:"is_indexed"`
	EngineeredFeatures interface{} `json:"engineered_features"`
}

// TagConfig TagConfig
type TagConfig struct {
	DatasourceId string `json:"datasource_id"`
	Description interface{} `json:"description"`
	LogicalGroupRegex interface{} `json:"logical_group_regex"`
	WorkingBucket string `json:"working_bucket"`
	WorkingPlatform interface{} `json:"working_platform,omitempty"`
	DbPath interface{} `json:"db_path"`
	DatabaseConfig interface{} `json:"database_config,omitempty"`
	CreatedBy interface{} `json:"created_by,omitempty"`
	UpdatedBy interface{} `json:"updated_by,omitempty"`
}

// TagTableInfo TagTableInfo
type TagTableInfo struct {
	DatasourceId string `json:"datasource_id"`
	Name string `json:"name"`
	Description interface{} `json:"description"`
}

// TokenTextChunking TokenTextChunking
type TokenTextChunking struct {
	ChunkSize interface{} `json:"chunk_size,omitempty"`
	ChunkOverlap interface{} `json:"chunk_overlap,omitempty"`
	Separator interface{} `json:"separator,omitempty"`
	BackupSeparators interface{} `json:"backup_separators,omitempty"`
	KeepWhitespaces interface{} `json:"keep_whitespaces,omitempty"`
}

// ToolActivity Record of a tool call and its result.
type ToolActivity struct {
	ToolId string `json:"tool_id"`
	ToolName string `json:"tool_name"`
	Arguments string `json:"arguments"`
	Result interface{} `json:"result,omitempty"`
	Timestamp string `json:"timestamp"`
}

// UpdateBlueprintRequest UpdateBlueprintRequest
type UpdateBlueprintRequest struct {
	Name interface{} `json:"name,omitempty"`
	ExecutionMode interface{} `json:"execution_mode,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Description interface{} `json:"description,omitempty"`
	DslDefinition interface{} `json:"dsl_definition,omitempty"`
	YamlSpecContent interface{} `json:"yaml_spec_content,omitempty"`
	JsonSpecContent interface{} `json:"json_spec_content,omitempty"`
	InitInput interface{} `json:"init_input,omitempty"`
}

// UpdateBlueprintTaskRequest UpdateBlueprintTaskRequest
type UpdateBlueprintTaskRequest struct {
	Name interface{} `json:"name,omitempty"`
	Description interface{} `json:"description,omitempty"`
	InputSchema interface{} `json:"input_schema,omitempty"`
	OutputSchema interface{} `json:"output_schema,omitempty"`
	ConfigSchema interface{} `json:"config_schema,omitempty"`
	ToolSchema interface{} `json:"tool_schema,omitempty"`
}

// UpdateChunkingStrategyRequest UpdateChunkingStrategyRequest
type UpdateChunkingStrategyRequest struct {
	CodeSplitter interface{} `json:"code_splitter,omitempty"`
	HtmlNodeParser interface{} `json:"html_node_parser,omitempty"`
	JsonNodeParser interface{} `json:"json_node_parser,omitempty"`
	MarkdownNodeParser interface{} `json:"markdown_node_parser,omitempty"`
	SemanticSplitterNodeParser interface{} `json:"semantic_splitter_node_parser,omitempty"`
	SentenceSplitter interface{} `json:"sentence_splitter,omitempty"`
	TokenTextSplitter interface{} `json:"token_text_splitter,omitempty"`
}

// UpdateChunkingStrategyResponse UpdateChunkingStrategyResponse
type UpdateChunkingStrategyResponse struct {
	Message string `json:"message"`
}

// UpdateDataElementRequest UpdateDataElementRequest
type UpdateDataElementRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Path interface{} `json:"path,omitempty"`
	MediaType interface{} `json:"media_type,omitempty"`
	DiscoveryRecord interface{} `json:"discovery_record,omitempty"`
	ParentDataElementId interface{} `json:"parent_data_element_id,omitempty"`
}

// UpdateDataElementResponse UpdateDataElementResponse
type UpdateDataElementResponse struct {
	Id string `json:"id"`
}

// UpdateDatasourceRequest UpdateDatasourceRequest
type UpdateDatasourceRequest struct {
	Name interface{} `json:"name,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Recurrence interface{} `json:"recurrence,omitempty"`
	ObjectStorageConfig interface{} `json:"object_storage_config,omitempty"`
	WebConfig interface{} `json:"web_config,omitempty"`
	ConnectorConfig interface{} `json:"connector_config,omitempty"`
}

// UpdateDatasourceResponse UpdateDatasourceResponse
type UpdateDatasourceResponse struct {
	Id string `json:"id"`
}

// UpdateRagConfigRequest UpdateRagConfigRequest
type UpdateRagConfigRequest struct {
	Description interface{} `json:"description,omitempty"`
	CollectionId interface{} `json:"collection_id,omitempty"`
	ExtractorModel interface{} `json:"extractor_model,omitempty"`
	EmbeddingModel interface{} `json:"embedding_model,omitempty"`
	SparseEmbeddingModel interface{} `json:"sparse_embedding_model,omitempty"`
	CollectMetadata interface{} `json:"collect_metadata,omitempty"`
	MetadataOptions interface{} `json:"metadata_options,omitempty"`
}

// UpdateRagConfigResponse UpdateRagConfigResponse
type UpdateRagConfigResponse struct {
	Message string `json:"message"`
}

// UpdateTagColumnRequest UpdateTagColumnRequest
type UpdateTagColumnRequest struct {
	Description interface{} `json:"description,omitempty"`
	Dtype interface{} `json:"dtype,omitempty"`
	IsKey interface{} `json:"is_key,omitempty"`
	IsIndexed interface{} `json:"is_indexed,omitempty"`
	EngineeredFeatures interface{} `json:"engineered_features,omitempty"`
}

// UpdateTagColumnResponse UpdateTagColumnResponse
type UpdateTagColumnResponse struct {
	Message string `json:"message"`
}

// UpdateTagConfigRequest UpdateTagConfigRequest
type UpdateTagConfigRequest struct {
	Description interface{} `json:"description,omitempty"`
	LogicalGroupRegex interface{} `json:"logical_group_regex,omitempty"`
	WorkingBucket interface{} `json:"working_bucket,omitempty"`
	WorkingPlatform interface{} `json:"working_platform,omitempty"`
	DbPath interface{} `json:"db_path,omitempty"`
	DatabaseConfig interface{} `json:"database_config,omitempty"`
}

// UpdateTagConfigResponse UpdateTagConfigResponse
type UpdateTagConfigResponse struct {
	Message string `json:"message"`
}

// UpdateTagTableRequest UpdateTagTableRequest
type UpdateTagTableRequest struct {
	Description interface{} `json:"description,omitempty"`
}

// UpdateTagTableResponse UpdateTagTableResponse
type UpdateTagTableResponse struct {
	Message string `json:"message"`
}

// UploadContentResponse UploadContentResponse
type UploadContentResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	UploadId string `json:"upload_id"`
	SseUrl string `json:"sse_url"`
	EstimatedFiles int64 `json:"estimated_files"`
	EstimatedSize int64 `json:"estimated_size"`
}

// WebDomain WebDomain
type WebDomain struct {
	Domain string `json:"domain"`
	LimitPattern string `json:"limitPattern"`
	ExcludePattern string `json:"excludePattern"`
	Ingestible bool `json:"ingestible"`
	Expandable bool `json:"expandable"`
}

// HeartbeatEvent A server-sent event indicating that the server is still processing the request
type HeartbeatEvent struct {
	Event string `json:"event"`
}

// ChatEvent A server-sent event containing chat completion content
type ChatEvent struct {
	Id string `json:"id"`
	Event string `json:"event"`
	Data string `json:"data"`
}

// SseEvent Generic Server-Sent Event
type SseEvent struct {
	// Event type (e.g., connected, progress, stream_complete, error, keepalive)
	Event string `json:"event"`
	// Event data payload
	Data *string `json:"data,omitempty"`
}
