# Getting Started

This guide will help you get started with the Draw Things Go client library.

## Prerequisites

Before using this library, you need:

1. **Draw Things Application**: Install the Draw Things app on your device
   - macOS: Available on the Mac App Store
   - iOS/iPadOS: Available on the App Store

2. **Go Environment**: Make sure you have Go installed (version 1.21 or later)
   ```bash
   go version
   ```

## Installation

### Install the Library

```bash
go get github.com/drawthings_go
```

### Install the CLI Tool

```bash
go install github.com/drawthings_go/cmd/drawthings@latest
```

## Setting Up Draw Things API Server

1. **Launch Draw Things**: Open the Draw Things application on your device

2. **Enable the HTTP Server**:
   - Navigate to the application's settings
   - Locate the "API Server" option
   - Enable the "HTTP" server
   - Note the server address (default: `http://127.0.0.1:7860`)

3. **Verify the Server**:
   ```bash
   curl http://127.0.0.1:7860/sdapi/v1/options
   ```
   If the server is running, you should get a JSON response.

**Important**: Ensure that port 7860 is available and not blocked by any firewall or security software.

## Your First Image

### Using the Library

Create a file `main.go`:

```go
package main

import (
    "context"
    "log"
    
    "github.com/drawthings_go"
)

func main() {
    // Create a client with default settings
    client := drawthings.NewClient()
    
    // Create a request
    req := &drawthings.TextToImageRequest{
        Prompt: "a beautiful sunset over mountains, digital art",
        Steps:  20,
    }
    
    // Generate and save the image
    ctx := context.Background()
    err := client.GenerateImageAndSave(ctx, req, "my_first_image.png")
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("Image saved successfully!")
}
```

Run it:

```bash
go run main.go
```

### Using the CLI

```bash
drawthings -prompt "a beautiful sunset over mountains, digital art" -output my_first_image.png
```

## Next Steps

- Read the [API Reference](api-reference.md) for detailed documentation
- Check out [Examples](examples.md) for more use cases
- See [Troubleshooting](troubleshooting.md) if you encounter issues

## Configuration

### Default Settings

The client uses these defaults:
- **Base URL**: `http://127.0.0.1:7860`
- **Timeout**: 5 minutes
- **Steps**: 20
- **Guidance Scale**: 4.0
- **Width**: 512 pixels
- **Height**: 512 pixels
- **Seed**: -1 (random)

### Custom Configuration

```go
client := drawthings.NewClient(
    drawthings.WithBaseURL("http://custom-server:7860"),
    drawthings.WithTimeout(10 * time.Minute),
)
```

## Common Patterns

### Basic Generation

```go
req := &drawthings.TextToImageRequest{
    Prompt: "your prompt here",
}
err := client.GenerateImageAndSave(ctx, req, "output.png")
```

### High-Quality Generation

```go
req := &drawthings.TextToImageRequest{
    Prompt:         "your prompt",
    NegativePrompt: "blurry, low quality",
    Steps:          50,
    GuidanceScale:  7.0,
    Width:          768,
    Height:         768,
}
```

### Reproducible Results

```go
req := &drawthings.TextToImageRequest{
    Prompt: "your prompt",
    Seed:   42, // Fixed seed for reproducibility
}
```

## Need Help?

- Check the [Troubleshooting Guide](troubleshooting.md)
- Review the [API Reference](api-reference.md)
- See [Examples](examples.md) for code samples

