package v2

import (
	"fmt"
)

// APIError is the base error type for all API errors.
type APIError struct {
	// Status is the HTTP status code.
	Status int
	// Message is a human-readable error message.
	Message string
	// Code is an optional error code from the API.
	Code string
	// Body is the raw response body.
	Body map[string]interface{}
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("[%d] %s: %s", e.Status, e.Code, e.Message)
	}
	return fmt.Sprintf("[%d] %s", e.Status, e.Message)
}

// AuthenticationError is returned when authentication fails (401).
type AuthenticationError struct {
	APIError
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("authentication failed: %s", e.Message)
}

// AuthorizationError is returned when authorization fails (403).
type AuthorizationError struct {
	APIError
}

func (e *AuthorizationError) Error() string {
	return fmt.Sprintf("authorization failed: %s", e.Message)
}

// NotFoundError is returned when a resource is not found (404).
type NotFoundError struct {
	APIError
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.Message)
}

// ValidationError is returned when request validation fails (422).
type ValidationError struct {
	APIError
	// Errors contains field-specific validation errors.
	Errors map[string][]string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s", e.Message)
}

// RateLimitError is returned when rate limit is exceeded (429).
type RateLimitError struct {
	APIError
	// RetryAfter is the number of seconds to wait before retrying.
	RetryAfter int
}

func (e *RateLimitError) Error() string {
	if e.RetryAfter > 0 {
		return fmt.Sprintf("rate limit exceeded, retry after %d seconds", e.RetryAfter)
	}
	return fmt.Sprintf("rate limit exceeded: %s", e.Message)
}

// ServerError is returned when the server encounters an error (5xx).
type ServerError struct {
	APIError
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("server error [%d]: %s", e.Status, e.Message)
}

// NetworkError is returned when a network error occurs.
type NetworkError struct {
	Err error
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error: %v", e.Err)
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}

// TimeoutError is returned when a request times out.
type TimeoutError struct {
	Err error
}

func (e *TimeoutError) Error() string {
	return "request timed out"
}

func (e *TimeoutError) Unwrap() error {
	return e.Err
}

// IsAuthenticationError returns true if the error is an AuthenticationError.
func IsAuthenticationError(err error) bool {
	_, ok := err.(*AuthenticationError)
	return ok
}

// IsAuthorizationError returns true if the error is an AuthorizationError.
func IsAuthorizationError(err error) bool {
	_, ok := err.(*AuthorizationError)
	return ok
}

// IsNotFoundError returns true if the error is a NotFoundError.
func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// IsValidationError returns true if the error is a ValidationError.
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// IsRateLimitError returns true if the error is a RateLimitError.
func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

// IsServerError returns true if the error is a ServerError.
func IsServerError(err error) bool {
	_, ok := err.(*ServerError)
	return ok
}

// IsNetworkError returns true if the error is a NetworkError.
func IsNetworkError(err error) bool {
	_, ok := err.(*NetworkError)
	return ok
}

// IsTimeoutError returns true if the error is a TimeoutError.
func IsTimeoutError(err error) bool {
	_, ok := err.(*TimeoutError)
	return ok
}
