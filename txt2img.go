package drawthings

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	httpclient "github.com/drawthings_go/internal/http"
	"github.com/drawthings_go/internal/validation"
)

// GenerateImage generates an image from a text prompt using the Draw Things API.
func (c *Client) GenerateImage(ctx context.Context, req *TextToImageRequest) (*TextToImageResponse, error) {
	// Set defaults for optional fields
	req.SetDefaults()

	// Validate request parameters
	if err := validation.ValidateTextToImageRequest(req.Prompt, req.Steps, req.GuidanceScale, req.Width, req.Height); err != nil {
		// Convert to ValidationError
		return nil, NewValidationError("", err.Error())
	}

	// Build the API endpoint URL
	url := fmt.Sprintf("%s/sdapi/v1/txt2img", c.baseURL)

	// Make the API request
	resp, err := c.httpClient.PostJSON(ctx, url, req)
	if err != nil {
		return nil, NewNetworkError("API request failed", err)
	}

	// Decode the response
	var apiResp TextToImageResponse
	if err := c.httpClient.DecodeJSONResponse(resp, &apiResp); err != nil {
		// Check if it's an HTTP error
		if httpErr, ok := err.(*httpclient.HTTPError); ok {
			return nil, NewAPIError(&http.Response{
				StatusCode: httpErr.StatusCode,
				Status:     httpErr.Status,
			}, httpErr.Body)
		}
		return nil, NewNetworkError("failed to decode response", err)
	}

	// Validate response
	if len(apiResp.Images) == 0 {
		return nil, NewDecodeError("no images in response", nil)
	}

	return &apiResp, nil
}

// GenerateImageAndSave generates an image and saves it to the specified file path.
func (c *Client) GenerateImageAndSave(ctx context.Context, req *TextToImageRequest, outputPath string) error {
	resp, err := c.GenerateImage(ctx, req)
	if err != nil {
		return err
	}

	if len(resp.Images) == 0 {
		return NewDecodeError("no images in response", nil)
	}

	// Decode the first image
	imageData, err := base64.StdEncoding.DecodeString(resp.Images[0])
	if err != nil {
		return NewDecodeError("failed to decode base64 image data", err)
	}

	// Ensure the output directory exists
	dir := filepath.Dir(outputPath)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Write the image to file
	if err := os.WriteFile(outputPath, imageData, 0644); err != nil {
		return fmt.Errorf("failed to write image file: %w", err)
	}

	return nil
}

