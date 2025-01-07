// Package errors provides a structured way to handle and represent errors in the SDK.
// It defines various error types to encapsulate different failure scenarios,
// making it easier to debug and handle issues effectively.
package errors

import (
	"fmt"
)

// SDKError represents a custom error type used throughout the SDK.
// It includes a unique code and descriptive message to provide detailed information about the error.
//
// Fields:
//   - RequestId: An optional field, typically used to store the originator conversation ID for failed requests.
//   - code: A short, unique identifier for the type of error.
//   - message: A detailed description of the error.
type SDKError struct {
	RequestId string
	code      string
	message   string
}

// Error implements the error interface for SDKError.
// It formats the error as a string containing the code and message.
//
// Returns:
//   - A string representation of the error in the format: "<code>: <message>".
func (e *SDKError) Error() string {
	return fmt.Sprintf("%v: %v", e.code, e.message)
}

// NewSDKError creates a new instance of SDKError with the given code and message.
//
// Parameters:
//   - code: A short, unique identifier for the error type.
//   - message: A descriptive message explaining the error.
//
// Returns:
//   - A pointer to the newly created SDKError.
func NewSDKError(code, message string) *SDKError {
	return &SDKError{
		code:    code,
		message: message,
	}
}

// Predefined error generators for common scenarios. Each generator returns an *SDKError
// with a specific error code and a customizable message.

// General Errors:
//
// NetworkError creates an error for network-related issues.
// Code: "NETWORK_ERROR".
//
// AuthenticationError creates an error for authentication failures.
// Code: "AUTH_ERROR".
//
// ValidationError creates an error for validation failures.
// Code: "VALIDATION_ERROR".
//
// ProcessingError creates an error for internal processing errors.
// Code: "PROCESSING_ERROR".
//
// EnvironmentError creates an error for environment configuration issues.
// Code: "ENVIRONMENT_ERROR".
//
// TimeoutError creates an error for timeout-related issues.
// Code: "TIMEOUT_ERROR".
var (
	NetworkError        = func(msg string) *SDKError { return NewSDKError("NETWORK_ERROR", msg) }
	AuthenticationError = func(msg string) *SDKError { return NewSDKError("AUTH_ERROR", msg) }
	ValidationError     = func(msg string) *SDKError { return NewSDKError("VALIDATION_ERROR", msg) }
	ProcessingError     = func(msg string) *SDKError { return NewSDKError("PROCESSING_ERROR", msg) }
	EnvironmentError    = func(msg string) *SDKError { return NewSDKError("ENVIRONMENT_ERROR", msg) }
	TimeoutError        = func(msg string) *SDKError { return NewSDKError("TIMEOUT_ERROR", msg) }
)

// Server Errors:
//
// InternalServerError creates an error for server-side internal issues.
// Code: "INTERNAL_SERVER_ERROR".
//
// ServiceUnavailable creates an error for service unavailability.
// Code: "SERVICE_UNAVAILABLE".
var (
	InternalServerError = func(msg string) *SDKError { return NewSDKError("INTERNAL_SERVER_ERROR", msg) }
	ServiceUnavailable  = func(msg string) *SDKError { return NewSDKError("SERVICE_UNAVAILABLE", msg) }
)

// Request/Response Errors:
//
// BadRequestError creates an error for invalid requests sent to the server.
// Code: "BAD_REQUEST_ERROR".
//
// UnauthorizedError creates an error for requests lacking valid authentication.
// Code: "UNAUTHORIZED_ERROR".
//
// ForbiddenError creates an error for requests that are explicitly denied.
// Code: "FORBIDDEN_ERROR".
//
// NotFoundError creates an error for requests targeting a non-existent resource.
// Code: "NOT_FOUND_ERROR".
var (
	BadRequestError   = func(msg string) *SDKError { return NewSDKError("BAD_REQUEST_ERROR", msg) }
	UnauthorizedError = func(msg string) *SDKError { return NewSDKError("UNAUTHORIZED_ERROR", msg) }
	ForbiddenError    = func(msg string) *SDKError { return NewSDKError("FORBIDDEN_ERROR", msg) }
	NotFoundError     = func(msg string) *SDKError { return NewSDKError("NOT_FOUND_ERROR", msg) }
)

// Custom Errors:
//
// CustomError allows creation of a custom error with any specified code and message.
//
// Parameters:
//   - code: A unique identifier for the custom error.
//   - msg: A descriptive message explaining the error.
//
// Returns:
//   - An *SDKError with the provided code and message.
var CustomError = func(code, msg string) *SDKError { return NewSDKError(code, msg) }

