package drawthings

import (
	"time"

	httpclient "github.com/drawthings_go/internal/http"
)

const (
	// DefaultBaseURL is the default base URL for the Draw Things API.
	DefaultBaseURL = "http://127.0.0.1:7860"
	// DefaultTimeout is the default HTTP client timeout.
	DefaultTimeout = 5 * time.Minute
)

// Logger interface for request/response logging.
type Logger interface {
	Logf(format string, args ...interface{})
}

// Client is the main client for interacting with the Draw Things API.
type Client struct {
	baseURL    string
	httpClient *httpclient.Client
	timeout    time.Duration
	logger     Logger
}

// Option is a function that configures a Client.
type Option func(*Client)

// WithBaseURL sets the base URL for the API client.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
		c.httpClient = httpclient.NewClient(timeout, c.logger)
	}
}

// WithLogger sets a logger for request/response logging.
func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
		timeout := c.timeout
		if timeout == 0 {
			timeout = DefaultTimeout
		}
		c.httpClient = httpclient.NewClient(timeout, logger)
	}
}

// NewClient creates a new Draw Things API client with the provided options.
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL: DefaultBaseURL,
		timeout: DefaultTimeout,
		logger:  nil,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Initialize HTTP client if not already set
	if c.httpClient == nil {
		c.httpClient = httpclient.NewClient(c.timeout, c.logger)
	}

	return c
}

// NewClientWithDefaults creates a new client with default settings.
func NewClientWithDefaults() *Client {
	return NewClient()
}

// BaseURL returns the base URL of the client.
func (c *Client) BaseURL() string {
	return c.baseURL
}

