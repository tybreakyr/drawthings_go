package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

)

func TestClient_PostJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("failed to decode body: %v", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, nil)
	ctx := context.Background()

	reqBody := map[string]string{
		"test": "value",
	}

	resp, err := client.PostJSON(ctx, server.URL, reqBody)
	if err != nil {
		t.Fatalf("PostJSON() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestClient_PostJSON_ErrorHandling(t *testing.T) {
	client := NewClient(100*time.Millisecond, nil)
	ctx := context.Background()

	// Test with invalid URL
	_, err := client.PostJSON(ctx, "http://invalid-url-that-does-not-exist.local", nil)
	if err == nil {
		t.Fatal("expected error for invalid URL")
	}
	// Just check that we got an error
	if err == nil {
		t.Error("expected error")
	}
}

func TestClient_DecodeJSONResponse(t *testing.T) {
	type TestResponse struct {
		Status string `json:"status"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, nil)
	ctx := context.Background()

	resp, err := client.PostJSON(ctx, server.URL, nil)
	if err != nil {
		t.Fatalf("PostJSON() error = %v", err)
	}

	var result TestResponse
	if err := client.DecodeJSONResponse(resp, &result); err != nil {
		t.Fatalf("DecodeJSONResponse() error = %v", err)
	}

	if result.Status != "success" {
		t.Errorf("expected status 'success', got %q", result.Status)
	}
}

func TestClient_DecodeJSONResponse_ErrorStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, nil)
	ctx := context.Background()

	resp, err := client.PostJSON(ctx, server.URL, nil)
	if err != nil {
		t.Fatalf("PostJSON() error = %v", err)
	}

	var result map[string]interface{}
	err = client.DecodeJSONResponse(resp, &result)
	if err == nil {
		t.Fatal("expected error for bad status code")
	}

	// Check that we got an HTTPError
	httpErr, ok := err.(*HTTPError)
	if !ok {
		t.Errorf("expected HTTPError, got %T", err)
	}
	if httpErr.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", httpErr.StatusCode)
	}
}

func TestClient_WithLogger(t *testing.T) {
	var loggedMessages []string
	logger := &testLogger{
		logFunc: func(format string, args ...interface{}) {
			loggedMessages = append(loggedMessages, format)
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, logger)
	ctx := context.Background()

	reqBody := map[string]string{"test": "value"}
	resp, err := client.PostJSON(ctx, server.URL, reqBody)
	if err != nil {
		t.Fatalf("PostJSON() error = %v", err)
	}

	var result map[string]interface{}
	if err := client.DecodeJSONResponse(resp, &result); err != nil {
		t.Fatalf("DecodeJSONResponse() error = %v", err)
	}

	// Verify logging occurred
	if len(loggedMessages) == 0 {
		t.Error("expected log messages, got none")
	}
}

type testLogger struct {
	logFunc func(format string, args ...interface{})
}

func (l *testLogger) Logf(format string, args ...interface{}) {
	l.logFunc(format, args...)
}

