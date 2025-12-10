package drawthings

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}
	if client.baseURL != DefaultBaseURL {
		t.Errorf("BaseURL: got %q, want %q", client.baseURL, DefaultBaseURL)
	}
}

func TestNewClientWithDefaults(t *testing.T) {
	client := NewClientWithDefaults()
	if client == nil {
		t.Fatal("NewClientWithDefaults() returned nil")
	}
	if client.baseURL != DefaultBaseURL {
		t.Errorf("BaseURL: got %q, want %q", client.baseURL, DefaultBaseURL)
	}
}

func TestWithBaseURL(t *testing.T) {
	customURL := "http://example.com:8080"
	client := NewClient(WithBaseURL(customURL))
	if client.baseURL != customURL {
		t.Errorf("BaseURL: got %q, want %q", client.baseURL, customURL)
	}
}

func TestWithTimeout(t *testing.T) {
	customTimeout := 10 * time.Minute
	client := NewClient(WithTimeout(customTimeout))
	if client.timeout != customTimeout {
		t.Errorf("Timeout: got %v, want %v", client.timeout, customTimeout)
	}
}

func TestBaseURL(t *testing.T) {
	client := NewClient()
	if got := client.BaseURL(); got != DefaultBaseURL {
		t.Errorf("BaseURL() = %q, want %q", got, DefaultBaseURL)
	}
}

func TestClientOptions(t *testing.T) {
	customURL := "http://example.com:8080"
	customTimeout := 15 * time.Minute
	client := NewClient(
		WithBaseURL(customURL),
		WithTimeout(customTimeout),
	)

	if client.baseURL != customURL {
		t.Errorf("BaseURL: got %q, want %q", client.baseURL, customURL)
	}
	if client.timeout != customTimeout {
		t.Errorf("Timeout: got %v, want %v", client.timeout, customTimeout)
	}
}

