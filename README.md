# Unofficial Draw Things Go Client Library

A robust, well-tested Go client library for the [Draw Things](https://drawthings.ai) API. Draw Things is an AI-assisted image generation application that allows users to create images from textual descriptions using advanced AI models.

> **⚠️ Disclaimer**: This is an **unofficial** library and is **not affiliated with, endorsed by, or connected to** the makers of the Draw Things application. This library is a solo developer maintained project created independently to provide a Go client for the Draw Things API.

## Features

- 🚀 Simple and intuitive API
- 🛡️ Robust error handling with specific error types
- ✅ Comprehensive parameter validation
- 🧪 Extensive test coverage
- 📝 Well-documented with examples
- 🔧 Configurable client (timeouts, logging, custom base URL)
- 💻 Command-line interface (CLI) tool included
- 📚 Complete documentation and wiki

## Installation

```bash
go get github.com/drawthings_go
```

Or if you prefer to use the CLI tool:

```bash
go install github.com/drawthings_go/cmd/drawthings@latest
```

## Quick Start

### Using the Library

```go
package main

import (
    "context"
    "log"
    
    "github.com/drawthings_go"
)

func main() {
    // Create a client
    client := drawthings.NewClient()
    
    // Create a request
    req := &drawthings.TextToImageRequest{
        Prompt: "a beautiful sunset over mountains, digital art",
        Steps:  20,
    }
    
    // Generate and save the image
    ctx := context.Background()
    err := client.GenerateImageAndSave(ctx, req, "output.png")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Using the CLI

```bash
# Basic usage
drawthings -prompt "a beautiful sunset"

# With custom parameters
drawthings -prompt "a cat wearing sunglasses" \
           -steps 30 \
           -width 768 \
           -height 768 \
           -guidance-scale 7.0 \
           -seed 42 \
           -output cat.png

# Show help
drawthings -help
```

## Usage Examples

### Basic Image Generation

```go
client := drawthings.NewClient()
req := &drawthings.TextToImageRequest{
    Prompt: "a beautiful sunset over mountains",
    Steps:  20,
}

ctx := context.Background()
err := client.GenerateImageAndSave(ctx, req, "sunset.png")
```

### High-Quality Generation

```go
req := &drawthings.TextToImageRequest{
    Prompt:         "a beautiful landscape, photorealistic, 4k",
    NegativePrompt: "blurry, low quality, distorted",
    Steps:          50,
    GuidanceScale:  7.0,
    Width:          768,
    Height:         768,
    Seed:           42, // Fixed seed for reproducibility
}

err := client.GenerateImageAndSave(ctx, req, "landscape.png")
```

### Get Image Data Without Saving

```go
resp, err := client.GenerateImage(ctx, req)
if err != nil {
    log.Fatal(err)
}

// Access base64-encoded image data
if len(resp.Images) > 0 {
    imageData := resp.Images[0]
    // Decode and process as needed
}
```

### Custom Client Configuration

```go
// Custom base URL and timeout
client := drawthings.NewClient(
    drawthings.WithBaseURL("http://custom-server:7860"),
    drawthings.WithTimeout(10 * time.Minute),
)

// With logging
type myLogger struct{}
func (l *myLogger) Logf(format string, args ...interface{}) {
    fmt.Printf("[LOG] "+format+"\n", args...)
}

client := drawthings.NewClient(
    drawthings.WithLogger(&myLogger{}),
)
```

### Error Handling

```go
resp, err := client.GenerateImage(ctx, req)
if err != nil {
    if drawthings.IsValidationError(err) {
        // Handle validation errors
        fmt.Printf("Invalid parameters: %v\n", err)
    } else if drawthings.IsAPIError(err) {
        // Handle API errors
        apiErr := err.(*drawthings.APIError)
        fmt.Printf("API returned status %d: %s\n", apiErr.StatusCode, apiErr.Message)
    } else if drawthings.IsNetworkError(err) {
        // Handle network errors
        fmt.Printf("Network error: %v\n", err)
    } else {
        // Handle other errors
        fmt.Printf("Error: %v\n", err)
    }
    return
}
```

## API Reference

### Client

- `NewClient(opts ...Option) *Client` - Create a new client with options
- `NewClientWithDefaults() *Client` - Create a client with default settings
- `GenerateImage(ctx context.Context, req *TextToImageRequest) (*TextToImageResponse, error)` - Generate image and return response
- `GenerateImageAndSave(ctx context.Context, req *TextToImageRequest, outputPath string) error` - Generate image and save to file

### Request Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `prompt` | string | Yes | Textual description of the desired image |
| `negative_prompt` | string | No | Elements to exclude from the image |
| `steps` | int | No | Number of inference steps (1-150, default: 20) |
| `guidance_scale` | float64 | No | Prompt adherence (1.0-20.0, default: 4.0) |
| `width` | int | No | Image width in pixels (default: 512) |
| `height` | int | No | Image height in pixels (default: 512) |
| `seed` | int | No | Random seed (-1 for random, default: -1) |

## CLI Usage

The CLI tool supports all request parameters:

```bash
drawthings [options]

Options:
  -prompt string
        Textual description of the desired image (required)
  -negative-prompt string
        Descriptions of elements to exclude from the image
  -steps int
        Number of inference steps (1-150, default: 20)
  -guidance-scale float
        Controls adherence to the prompt (1.0-20.0, default: 4.0)
  -width int
        Width of the generated image in pixels (default: 512)
  -height int
        Height of the generated image in pixels (default: 512)
  -seed int
        Random seed for image generation (-1 for random, default: -1)
  -output string
        Output file path for the generated image (default: "output.png")
  -base-url string
        Base URL of the Draw Things API server (default: "http://127.0.0.1:7860")
  -timeout duration
        HTTP client timeout (default: 5m0s)
  -version
        Show version information
```

## Prerequisites

- **Draw Things Application**: Install the Draw Things app on your device (macOS, iPhone, or iPad)
- **API Server**: Enable the HTTP server within the Draw Things application
- **Network**: The API server runs on `localhost:7860` by default

See the [Getting Started Guide](docs/getting-started.md) for detailed setup instructions.

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Documentation

- [Getting Started](docs/getting-started.md) - Installation and setup guide
- [API Reference](docs/api-reference.md) - Complete API documentation
- [Examples](docs/examples.md) - Code examples and use cases
- [Troubleshooting](docs/troubleshooting.md) - Common issues and solutions

## Error Types

The library provides specific error types for better error handling:

- **APIError**: Errors returned by the API server (includes status code)
- **ValidationError**: Parameter validation failures
- **NetworkError**: Network-related errors (timeouts, connection issues)
- **DecodeError**: Errors during image decoding or processing

Use the `Is*` functions to check error types programmatically.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Related Resources

- [Draw Things Website](https://drawthings.ai)
- [Draw Things Wiki](https://wiki.drawthings.ai/wiki/Help)
- [Draw Things Community Repository](https://github.com/drawthingsai/draw-things-community)

## Acknowledgments

- This library is built for the Draw Things API, which follows a structure similar to Stable Diffusion's API. 
- I acknowledge the entire human race for creating all the data that was scraped, processed, and fed into the AI models that I used to "write" this with.

## How I "wrote" this
I had an ai search the web for information on interfacing with the Draw Things Application and create a markdown file with the findings that could be used by a copilot to create a client. 
I then had Cursor, in plan mode, create a plan with this prompt:

Create a plan for a new client library written in go in a new folder called "drawthings_go".
- Use @DRAW_THINGS_API.md  as your source of truth.
- The library should be built for opensource.
- It must be robust and easy to debug.
- It must have tests to cover exposed functions/types/features
- It must come with a simple to use cli
- It must come with it's own wiki using markdown
- It must have a README.md file.
- Use Go best practices @Go for project structure. 

The only changes I made were to the documentation where Cursor was giving out acknowledgements and added the Unofficial label. 
