package drawthings

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerateImage(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/sdapi/v1/txt2img" {
			t.Errorf("expected path /sdapi/v1/txt2img, got %s", r.URL.Path)
		}

		// Decode request
		var req TextToImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("failed to decode request: %v", err)
		}

		// Create a minimal valid PNG (1x1 transparent)
		pngData := []byte{
			0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
			0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
			0x08, 0x06, 0x00, 0x00, 0x00, 0x1F, 0x15, 0xC4,
			0x89, 0x00, 0x00, 0x00, 0x0A, 0x49, 0x44, 0x41,
			0x54, 0x78, 0x9C, 0x63, 0x00, 0x01, 0x00, 0x00,
			0x05, 0x00, 0x01, 0x0D, 0x0A, 0x2D, 0xB4, 0x00,
			0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, 0x44, 0xAE,
			0x42, 0x60, 0x82,
		}
		encoded := base64.StdEncoding.EncodeToString(pngData)

		response := TextToImageResponse{
			Images: []string{encoded},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(WithBaseURL(server.URL))

	req := &TextToImageRequest{
		Prompt: "a test image",
		Steps:  20,
	}

	ctx := context.Background()
	resp, err := client.GenerateImage(ctx, req)
	if err != nil {
		t.Fatalf("GenerateImage() error = %v", err)
	}

	if len(resp.Images) == 0 {
		t.Error("expected at least one image in response")
	}
}

func TestGenerateImage_ValidationError(t *testing.T) {
	client := NewClient()
	req := &TextToImageRequest{
		// Missing prompt
		Steps: 20,
	}

	ctx := context.Background()
	_, err := client.GenerateImage(ctx, req)
	if err == nil {
		t.Fatal("expected validation error")
	}
	if !IsValidationError(err) {
		t.Errorf("expected ValidationError, got %T", err)
	}
}

func TestGenerateImage_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}))
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))
	req := &TextToImageRequest{
		Prompt: "test",
	}

	ctx := context.Background()
	_, err := client.GenerateImage(ctx, req)
	if err == nil {
		t.Fatal("expected error")
	}
	if !IsAPIError(err) {
		t.Errorf("expected APIError, got %T", err)
	}
}

func TestGenerateImageAndSave(t *testing.T) {
	// Create a temporary directory for test output
	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "test_output.png")

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a minimal valid PNG
		pngData := []byte{
			0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
			0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
			0x08, 0x06, 0x00, 0x00, 0x00, 0x1F, 0x15, 0xC4,
			0x89, 0x00, 0x00, 0x00, 0x0A, 0x49, 0x44, 0x41,
			0x54, 0x78, 0x9C, 0x63, 0x00, 0x01, 0x00, 0x00,
			0x05, 0x00, 0x01, 0x0D, 0x0A, 0x2D, 0xB4, 0x00,
			0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, 0x44, 0xAE,
			0x42, 0x60, 0x82,
		}
		encoded := base64.StdEncoding.EncodeToString(pngData)

		response := TextToImageResponse{
			Images: []string{encoded},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))
	req := &TextToImageRequest{
		Prompt: "a test image",
		Steps:  20,
	}

	ctx := context.Background()
	err := client.GenerateImageAndSave(ctx, req, outputPath)
	if err != nil {
		t.Fatalf("GenerateImageAndSave() error = %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatal("output file was not created")
	}

	// Verify file is not empty
	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("failed to stat output file: %v", err)
	}
	if info.Size() == 0 {
		t.Error("output file is empty")
	}
}

func TestGenerateImage_Timeout(t *testing.T) {
	// Create a slow server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(
		WithBaseURL(server.URL),
		WithTimeout(100*time.Millisecond),
	)

	req := &TextToImageRequest{
		Prompt: "test",
	}

	ctx := context.Background()
	_, err := client.GenerateImage(ctx, req)
	if err == nil {
		t.Fatal("expected timeout error")
	}
	if !IsNetworkError(err) {
		t.Errorf("expected NetworkError, got %T: %v", err, err)
	}
}

