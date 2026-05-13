package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// TransformOption configures a document transform request.
type TransformOption func(*transformConfig)

type transformConfig struct {
	model          *string
	prompt         *string
	promptID       *string
	timeoutSeconds *int
}

// WithTransformModel sets the LLM model override.
func WithTransformModel(model string) TransformOption {
	return func(c *transformConfig) { c.model = &model }
}

// WithTransformPrompt sets extraction instructions override.
func WithTransformPrompt(prompt string) TransformOption {
	return func(c *transformConfig) { c.prompt = &prompt }
}

// WithTransformPromptID sets a prompt template reference.
func WithTransformPromptID(promptID string) TransformOption {
	return func(c *transformConfig) { c.promptID = &promptID }
}

// WithTransformTimeout sets the max wait time in seconds (sync only).
func WithTransformTimeout(seconds int) TransformOption {
	return func(c *transformConfig) { c.timeoutSeconds = &seconds }
}

// TransformDocument sends a document for AI extraction using a typed Go struct
// as the output schema. Uses reflect to generate JSON Schema from struct tags.
//
// Example:
//
//	type InvoiceOutput struct {
//	    VendorName  string  `json:"vendor_name" desc:"Name of the vendor"`
//	    TotalAmount float64 `json:"total_amount" desc:"Invoice total"`
//	}
//
//	result, err := TransformDocument[InvoiceOutput](client, "invoice.pdf")
//	fmt.Println(result.Data["vendor_name"])
func TransformDocument[T any](client *Client, file string, opts ...TransformOption) (*TransformDocumentResponse, error) {
	cfg := &transformConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	schema, err := structToTransformJsonSchema[T]()
	if err != nil {
		return nil, err
	}

	req := TransformDocumentRequest{
		File:           file,
		ArtifactSchema: schema,
	}
	if cfg.model != nil {
		req.Model = cfg.model
	}
	if cfg.prompt != nil {
		req.Prompt = cfg.prompt
	}
	if cfg.promptID != nil {
		req.PromptId = cfg.promptID
	}
	if cfg.timeoutSeconds != nil {
		req.TimeoutSeconds = cfg.timeoutSeconds
	}

	return client.Documents.Transform(context.Background(), req)
}

// TransformDocumentWithSchema sends a document for AI extraction using a
// pre-built schema value — either a string artifact schema ID or a
// map[string]interface{} JSON Schema.
//
// Example with artifact schema ID:
//
//	result, err := TransformDocumentWithSchema(client, "invoice.pdf", "artifact-schema-id-123")
//
// Example with raw JSON Schema:
//
//	schema := map[string]interface{}{
//	    "type": "object",
//	    "properties": map[string]interface{}{
//	        "vendor_name": map[string]interface{}{"type": "string"},
//	    },
//	}
//	result, err := TransformDocumentWithSchema(client, "invoice.pdf", schema)
func TransformDocumentWithSchema(client *Client, file string, schema interface{}, opts ...TransformOption) (*TransformDocumentResponse, error) {
	cfg := &transformConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	req := TransformDocumentRequest{
		File:           file,
		ArtifactSchema: schema,
	}
	if cfg.model != nil {
		req.Model = cfg.model
	}
	if cfg.prompt != nil {
		req.Prompt = cfg.prompt
	}
	if cfg.promptID != nil {
		req.PromptId = cfg.promptID
	}
	if cfg.timeoutSeconds != nil {
		req.TimeoutSeconds = cfg.timeoutSeconds
	}

	return client.Documents.Transform(context.Background(), req)
}

// SubmitDocumentTransform submits an async document transform using a typed Go struct.
// Returns immediately with an execution ID. Poll with client.Sessions.Get(executionId).
//
// Example:
//
//	job, err := SubmitDocumentTransform[InvoiceOutput](client, "invoice.pdf")
//	// ... do other work ...
//	session, _ := client.Sessions.Get(context.Background(), job.ExecutionId)
func SubmitDocumentTransform[T any](client *Client, file string, opts ...TransformOption) (*SubmitDocumentTransformResponse, error) {
	cfg := &transformConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	schema, err := structToTransformJsonSchema[T]()
	if err != nil {
		return nil, err
	}

	req := TransformDocumentRequest{
		File:           file,
		ArtifactSchema: schema,
	}
	if cfg.model != nil {
		req.Model = cfg.model
	}
	if cfg.prompt != nil {
		req.Prompt = cfg.prompt
	}
	if cfg.promptID != nil {
		req.PromptId = cfg.promptID
	}

	return client.Documents.SubmitTransform(context.Background(), req)
}

// SubmitDocumentTransformWithSchema submits an async document transform using
// a pre-built schema value — either a string artifact schema ID or a
// map[string]interface{} JSON Schema.
//
// Example:
//
//	job, err := SubmitDocumentTransformWithSchema(client, "invoice.pdf", "artifact-schema-id-123")
func SubmitDocumentTransformWithSchema(client *Client, file string, schema interface{}, opts ...TransformOption) (*SubmitDocumentTransformResponse, error) {
	cfg := &transformConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	req := TransformDocumentRequest{
		File:           file,
		ArtifactSchema: schema,
	}
	if cfg.model != nil {
		req.Model = cfg.model
	}
	if cfg.prompt != nil {
		req.Prompt = cfg.prompt
	}
	if cfg.promptID != nil {
		req.PromptId = cfg.promptID
	}

	return client.Documents.SubmitTransform(context.Background(), req)
}

// ---------------------------------------------------------------------------
// Internal: JSON Schema generation from struct
// ---------------------------------------------------------------------------

var transformTimeType = reflect.TypeOf(time.Time{})

func structToTransformJsonSchema[T any]() (interface{}, error) {
	var zero T
	t := reflect.TypeOf(zero)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("TransformDocument: expected struct type parameter, got %s", t.Kind())
	}

	schema, err := transformStructToSchema(t)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func transformStructToSchema(t reflect.Type) (map[string]interface{}, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", t.Kind())
	}

	properties := make(map[string]interface{})
	required := []string{}

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}

		name := transformFieldName(sf)
		jsonType := transformGoTypeToJsonSchema(sf.Type)

		if desc := sf.Tag.Get("desc"); desc != "" {
			jsonType["description"] = desc
		}

		properties[name] = jsonType

		if sf.Type.Kind() != reflect.Ptr {
			required = append(required, name)
		}
	}

	schema := map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}
	if len(required) > 0 {
		schema["required"] = required
	}
	return schema, nil
}

func transformGoTypeToJsonSchema(t reflect.Type) map[string]interface{} {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t == transformTimeType {
		return map[string]interface{}{"type": "string", "format": "date-time"}
	}

	switch t.Kind() {
	case reflect.String:
		return map[string]interface{}{"type": "string"}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return map[string]interface{}{"type": "integer"}
	case reflect.Float32, reflect.Float64:
		return map[string]interface{}{"type": "number"}
	case reflect.Bool:
		return map[string]interface{}{"type": "boolean"}
	case reflect.Slice:
		items := transformGoTypeToJsonSchema(t.Elem())
		return map[string]interface{}{"type": "array", "items": items}
	case reflect.Map:
		return map[string]interface{}{"type": "object"}
	case reflect.Struct:
		nested, err := transformStructToSchema(t)
		if err != nil {
			return map[string]interface{}{"type": "string"}
		}
		return nested
	default:
		return map[string]interface{}{"type": "string"}
	}
}

func transformFieldName(sf reflect.StructField) string {
	if jsonTag := sf.Tag.Get("json"); jsonTag != "" {
		parts := strings.SplitN(jsonTag, ",", 2)
		if parts[0] != "" && parts[0] != "-" {
			return parts[0]
		}
	}
	name := sf.Name
	return strings.ToLower(name[:1]) + name[1:]
}

// marshalTransformSchema converts the schema map to a JSON string for the API request.
func marshalTransformSchema(schema map[string]interface{}) (string, error) {
	b, err := json.Marshal(schema)
	if err != nil {
		return "", fmt.Errorf("failed to marshal transform schema: %w", err)
	}
	return string(b), nil
}
