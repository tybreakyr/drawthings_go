# API Reference

Complete API documentation for the Draw Things Go client library.

## Client

### NewClient

Creates a new Draw Things API client with the provided options.

```go
func NewClient(opts ...Option) *Client
```

**Options:**
- `WithBaseURL(baseURL string)` - Set the base URL (default: `http://127.0.0.1:7860`)
- `WithTimeout(timeout time.Duration)` - Set HTTP client timeout (default: 5 minutes)
- `WithLogger(logger Logger)` - Set a logger for request/response logging

**Example:**
```go
client := drawthings.NewClient(
    drawthings.WithBaseURL("http://custom-server:7860"),
    drawthings.WithTimeout(10 * time.Minute),
)
```

### NewClientWithDefaults

Creates a new client with default settings.

```go
func NewClientWithDefaults() *Client
```

**Example:**
```go
client := drawthings.NewClientWithDefaults()
```

### GenerateImage

Generates an image from a text prompt and returns the response.

```go
func (c *Client) GenerateImage(ctx context.Context, req *TextToImageRequest) (*TextToImageResponse, error)
```

**Parameters:**
- `ctx`: Context for cancellation and timeouts
- `req`: Text-to-image request parameters

**Returns:**
- `*TextToImageResponse`: Response containing base64-encoded images
- `error`: Error if generation fails

**Example:**
```go
req := &drawthings.TextToImageRequest{
    Prompt: "a beautiful sunset",
    Steps:  20,
}

resp, err := client.GenerateImage(ctx, req)
if err != nil {
    log.Fatal(err)
}

// Access base64-encoded image data
if len(resp.Images) > 0 {
    imageData := resp.Images[0]
}
```

### GenerateImageAndSave

Generates an image and saves it to the specified file path.

```go
func (c *Client) GenerateImageAndSave(ctx context.Context, req *TextToImageRequest, outputPath string) error
```

**Parameters:**
- `ctx`: Context for cancellation and timeouts
- `req`: Text-to-image request parameters
- `outputPath`: Path where the image will be saved

**Returns:**
- `error`: Error if generation or saving fails

**Example:**
```go
req := &drawthings.TextToImageRequest{
    Prompt: "a beautiful sunset",
}

err := client.GenerateImageAndSave(ctx, req, "sunset.png")
if err != nil {
    log.Fatal(err)
}
```

### BaseURL

Returns the base URL of the client.

```go
func (c *Client) BaseURL() string
```

## Types

### TextToImageRequest

Request structure for text-to-image generation.

```go
type TextToImageRequest struct {
    Prompt         string  `json:"prompt"`
    NegativePrompt string  `json:"negative_prompt,omitempty"`
    Steps          int     `json:"steps,omitempty"`
    GuidanceScale  float64 `json:"guidance_scale,omitempty"`
    Width          int     `json:"width,omitempty"`
    Height         int     `json:"height,omitempty"`
    Seed           int     `json:"seed,omitempty"`
}
```

**Fields:**
- `Prompt` (string, required): Textual description of the desired image
- `NegativePrompt` (string, optional): Elements to exclude from the image
- `Steps` (int, optional): Number of inference steps (1-150, default: 20)
- `GuidanceScale` (float64, optional): Prompt adherence (1.0-20.0, default: 4.0)
- `Width` (int, optional): Image width in pixels (default: 512)
- `Height` (int, optional): Image height in pixels (default: 512)
- `Seed` (int, optional): Random seed (-1 for random, default: -1)

**Methods:**
- `SetDefaults()`: Sets default values for optional fields

### TextToImageResponse

Response structure from text-to-image generation.

```go
type TextToImageResponse struct {
    Images []string `json:"images"`
}
```

**Fields:**
- `Images` ([]string): Array of base64-encoded image data

## Error Types

### APIError

Error returned by the Draw Things API.

```go
type APIError struct {
    StatusCode int
    Message    string
    Body       string
}
```

**Check:**
```go
if drawthings.IsAPIError(err) {
    apiErr := err.(*drawthings.APIError)
    fmt.Printf("Status: %d\n", apiErr.StatusCode)
}
```

### ValidationError

Parameter validation failure.

```go
type ValidationError struct {
    Field   string
    Message string
}
```

**Check:**
```go
if drawthings.IsValidationError(err) {
    valErr := err.(*drawthings.ValidationError)
    fmt.Printf("Field: %s, Error: %s\n", valErr.Field, valErr.Message)
}
```

### NetworkError

Network-related error (timeout, connection refused, etc.).

```go
type NetworkError struct {
    Message string
    Err     error
}
```

**Check:**
```go
if drawthings.IsNetworkError(err) {
    netErr := err.(*drawthings.NetworkError)
    fmt.Printf("Network error: %s\n", netErr.Message)
}
```

### DecodeError

Error during base64 decoding or image processing.

```go
type DecodeError struct {
    Message string
    Err     error
}
```

**Check:**
```go
if drawthings.IsDecodeError(err) {
    decodeErr := err.(*drawthings.DecodeError)
    fmt.Printf("Decode error: %s\n", decodeErr.Message)
}
```

## Constants

```go
const (
    DefaultBaseURL  = "http://127.0.0.1:7860"
    DefaultTimeout  = 5 * time.Minute
)
```

## Interfaces

### Logger

Interface for request/response logging.

```go
type Logger interface {
    Logf(format string, args ...interface{})
}
```

**Example:**
```go
type myLogger struct{}

func (l *myLogger) Logf(format string, args ...interface{}) {
    fmt.Printf("[LOG] "+format+"\n", args...)
}

client := drawthings.NewClient(
    drawthings.WithLogger(&myLogger{}),
)
```

## Parameter Guidelines

### Steps
- **Range**: 1-150
- **Recommended**: 20-50 for good balance of quality and speed
- **Higher values**: Better quality but slower generation

### Guidance Scale
- **Range**: 1.0-20.0
- **Recommended**: 4-7 for most use cases
- **Lower values**: More creative, less prompt adherence
- **Higher values**: More prompt adherence, potentially over-saturated

### Image Dimensions
- **Common sizes**: 512x512, 768x768, 1024x1024
- **Note**: Larger images require more memory and processing time
- **Aspect ratios**: Maintain power-of-2 dimensions for best results

### Seed
- **-1**: Random seed (different image each time)
- **Positive integer**: Reproducible results (same seed = same image with same prompt)
- **Useful for**: Iterating on a specific image or reproducing results

