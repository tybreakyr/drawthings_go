package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Logger interface for request/response logging.
type Logger interface {
	Logf(format string, args ...interface{})
}

// Client wraps an HTTP client with additional functionality.
type Client struct {
	httpClient *http.Client
	logger     Logger
}

// NewClient creates a new HTTP client wrapper.
func NewClient(timeout time.Duration, logger Logger) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		logger: logger,
	}
}

// PostJSON sends a POST request with JSON body and returns the response.
func (c *Client) PostJSON(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)

		if c.logger != nil {
			c.logger.Logf("POST %s\nRequest body: %s", url, string(jsonData))
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if c.logger != nil {
		c.logger.Logf("Response status: %s", resp.Status)
	}

	return resp, nil
}

// HTTPError represents an HTTP error response.
type HTTPError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Status)
}

// DecodeJSONResponse decodes a JSON response body into the provided value.
func (c *Client) DecodeJSONResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if c.logger != nil {
		c.logger.Logf("Response body: %s", string(body))
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(body),
		}
	}

	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

