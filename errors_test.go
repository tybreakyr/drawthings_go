package drawthings

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIError(t *testing.T) {
	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusBadRequest)
	httpResp := resp.Result()

	err := NewAPIError(httpResp, "invalid request")
	if err.StatusCode != http.StatusBadRequest {
		t.Errorf("StatusCode: got %d, want %d", err.StatusCode, http.StatusBadRequest)
	}

	if !IsAPIError(err) {
		t.Error("IsAPIError should return true for APIError")
	}

	if IsAPIError(nil) {
		t.Error("IsAPIError should return false for nil")
	}
}

func TestValidationError(t *testing.T) {
	err := NewValidationError("prompt", "prompt is required")
	if err.Field != "prompt" {
		t.Errorf("Field: got %q, want %q", err.Field, "prompt")
	}

	if !IsValidationError(err) {
		t.Error("IsValidationError should return true for ValidationError")
	}

	if IsValidationError(nil) {
		t.Error("IsValidationError should return false for nil")
	}
}

func TestNetworkError(t *testing.T) {
	underlyingErr := fmt.Errorf("connection refused")
	err := NewNetworkError("connection failed", underlyingErr)
	if err.Message != "connection failed" {
		t.Errorf("Message: got %q, want %q", err.Message, "connection failed")
	}

	if err.Unwrap() != underlyingErr {
		t.Error("Unwrap should return the underlying error")
	}

	if !IsNetworkError(err) {
		t.Error("IsNetworkError should return true for NetworkError")
	}

	if IsNetworkError(nil) {
		t.Error("IsNetworkError should return false for nil")
	}
}

func TestDecodeError(t *testing.T) {
	underlyingErr := fmt.Errorf("invalid base64 data")
	err := NewDecodeError("invalid base64", underlyingErr)
	if err.Message != "invalid base64" {
		t.Errorf("Message: got %q, want %q", err.Message, "invalid base64")
	}

	if err.Unwrap() != underlyingErr {
		t.Error("Unwrap should return the underlying error")
	}

	if !IsDecodeError(err) {
		t.Error("IsDecodeError should return true for DecodeError")
	}

	if IsDecodeError(nil) {
		t.Error("IsDecodeError should return false for nil")
	}
}

