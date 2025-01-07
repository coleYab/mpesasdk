package errors

import (
    "fmt"
)

type SDKError struct {
    // Might be originator conversation id to check if this request is failed
    RequestId    string
    code    string
    message string
}

func (e *SDKError) Error() string {
    return fmt.Sprintf("%v: %v", e.code, e.message)
}

func NewSDKError(code, message string) *SDKError {
    return &SDKError{
        code: code,
        message: message,
    }
}

var (
    // General Errors
    NetworkError        = func(msg string) *SDKError { return NewSDKError("NETWORK_ERROR", msg) }
    AuthenticationError = func(msg string) *SDKError { return NewSDKError("AUTH_ERROR", msg) }
    ValidationError     = func(msg string) *SDKError { return NewSDKError("VALIDATION_ERROR", msg) }
    ProcessingError     = func(msg string) *SDKError { return NewSDKError("PROCESSING_ERROR", msg) }
    EnvironmentError    = func(msg string) *SDKError { return NewSDKError("ENVIRONMENT_ERROR", msg) }
    TimeoutError        = func(msg string) *SDKError { return NewSDKError("TIMEOUT_ERROR", msg) }

    // Server Errors
    InternalServerError = func(msg string) *SDKError { return NewSDKError("INTERNAL_SERVER_ERROR", msg) }
    ServiceUnavailable  = func(msg string) *SDKError { return NewSDKError("SERVICE_UNAVAILABLE", msg) }

    // Request/Response Errors
    BadRequestError      = func(msg string) *SDKError { return NewSDKError("BAD_REQUEST_ERROR", msg) }
    UnauthorizedError    = func(msg string) *SDKError { return NewSDKError("UNAUTHORIZED_ERROR", msg) }
    ForbiddenError       = func(msg string) *SDKError { return NewSDKError("FORBIDDEN_ERROR", msg) }
    NotFoundError        = func(msg string) *SDKError { return NewSDKError("NOT_FOUND_ERROR", msg) }

    // Customizable Errors
    CustomError = func(code, msg string) *SDKError { return NewSDKError(code, msg) }
)
