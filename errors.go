package drawthings

import (
	"fmt"
	"net/http"
)

// APIError represents an error returned by the Draw Things API.
type APIError struct {
	StatusCode int
	Message    string
	Body       string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Body)
}

// IsAPIError checks if an error is an APIError.
func IsAPIError(err error) bool {
	_, ok := err.(*APIError)
	return ok
}

// ValidationError represents a parameter validation failure.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// IsValidationError checks if an error is a ValidationError.
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// NetworkError represents a network-related error (timeout, connection refused, etc.).
type NetworkError struct {
	Message string
	Err     error
}

func (e *NetworkError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("network error: %s: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("network error: %s", e.Message)
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}

// IsNetworkError checks if an error is a NetworkError.
func IsNetworkError(err error) bool {
	_, ok := err.(*NetworkError)
	return ok
}

// DecodeError represents an error during base64 decoding or image processing.
type DecodeError struct {
	Message string
	Err     error
}

func (e *DecodeError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("decode error: %s: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("decode error: %s", e.Message)
}

func (e *DecodeError) Unwrap() error {
	return e.Err
}

// IsDecodeError checks if an error is a DecodeError.
func IsDecodeError(err error) bool {
	_, ok := err.(*DecodeError)
	return ok
}

// NewAPIError creates a new APIError from an HTTP response.
func NewAPIError(resp *http.Response, body string) *APIError {
	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    resp.Status,
		Body:       body,
	}
}

// NewValidationError creates a new ValidationError.
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// NewNetworkError creates a new NetworkError.
func NewNetworkError(message string, err error) *NetworkError {
	return &NetworkError{
		Message: message,
		Err:     err,
	}
}

// NewDecodeError creates a new DecodeError.
func NewDecodeError(message string, err error) *DecodeError {
	return &DecodeError{
		Message: message,
		Err:     err,
	}
}

