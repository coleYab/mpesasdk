package errors

import (
	"fmt"

	"github.com/coleYab/mpesasdk/service"
)

type SDKError struct {
	code    string
	message string
}

func (e *SDKError) Error() string {
    return fmt.Sprintf("%v Error: %v", e.code, e.message)
}

func NewSDKError(code, message string) *SDKError {
    return &SDKError{
        code: code,
        message: message,
    }
}

var (
	NetworkingError     = func(msg string) *SDKError { return NewSDKError("NETWORK_ERROR", msg) }
	AuthenticationError = func(msg string) *SDKError { return NewSDKError("AUTH_ERROR", msg) }
	ValidationError     = func(msg string) *SDKError { return NewSDKError("VALIDATION_ERROR", msg) }
	ProcessingError     = func(msg string) *SDKError { return NewSDKError("PROCESSING_ERROR", msg) }
)
