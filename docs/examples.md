# Examples

This document provides various examples demonstrating how to use the Draw Things Go client library.

## Table of Contents

- [Basic Usage](#basic-usage)
- [Error Handling](#error-handling)
- [Custom Configuration](#custom-configuration)
- [Batch Generation](#batch-generation)
- [Working with Image Data](#working-with-image-data)
- [Advanced Patterns](#advanced-patterns)

## Basic Usage

### Simple Image Generation

```go
package main

import (
    "context"
    "log"
    
    "github.com/drawthings_go"
)

func main() {
    client := drawthings.NewClient()
    
    req := &drawthings.TextToImageRequest{
        Prompt: "a beautiful sunset over mountains",
        Steps:  20,
    }
    
    ctx := context.Background()
    err := client.GenerateImageAndSave(ctx, req, "sunset.png")
    if err != nil {
        log.Fatal(err)
    }
}
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
    Seed:           42,
}

err := client.GenerateImageAndSave(ctx, req, "landscape.png")
```

### Reproducible Results

```go
// Use a fixed seed for reproducibility
req := &drawthings.TextToImageRequest{
    Prompt: "a cat wearing sunglasses",
    Seed:   42, // Same seed + same prompt = same image
    Steps:  30,
}

err := client.GenerateImageAndSave(ctx, req, "cat.png")
```

## Error Handling

### Comprehensive Error Handling

```go
resp, err := client.GenerateImage(ctx, req)
if err != nil {
    switch {
    case drawthings.IsValidationError(err):
        valErr := err.(*drawthings.ValidationError)
        fmt.Printf("Validation error in field '%s': %s\n", 
            valErr.Field, valErr.Message)
        
    case drawthings.IsAPIError(err):
        apiErr := err.(*drawthings.APIError)
        fmt.Printf("API error (status %d): %s\n", 
            apiErr.StatusCode, apiErr.Message)
        
    case drawthings.IsNetworkError(err):
        fmt.Printf("Network error: %v\n", err)
        
    case drawthings.IsDecodeError(err):
        fmt.Printf("Decode error: %v\n", err)
        
    default:
        fmt.Printf("Unexpected error: %v\n", err)
    }
    return
}
```

### Retry Logic

```go
func generateWithRetry(client *drawthings.Client, req *drawthings.TextToImageRequest, maxRetries int) error {
    ctx := context.Background()
    
    for i := 0; i < maxRetries; i++ {
        err := client.GenerateImageAndSave(ctx, req, "output.png")
        if err == nil {
            return nil
        }
        
        if drawthings.IsNetworkError(err) {
            time.Sleep(time.Second * time.Duration(i+1))
            continue
        }
        
        // Don't retry on validation or API errors
        return err
    }
    
    return fmt.Errorf("failed after %d retries", maxRetries)
}
```

## Custom Configuration

### Custom Base URL and Timeout

```go
client := drawthings.NewClient(
    drawthings.WithBaseURL("http://192.168.1.100:7860"),
    drawthings.WithTimeout(10 * time.Minute),
)
```

### With Logging

```go
type fileLogger struct {
    file *os.File
}

func (l *fileLogger) Logf(format string, args ...interface{}) {
    fmt.Fprintf(l.file, format+"\n", args...)
}

logger := &fileLogger{file: logFile}
client := drawthings.NewClient(
    drawthings.WithLogger(logger),
)
```

### Multiple Configurations

```go
client := drawthings.NewClient(
    drawthings.WithBaseURL("http://custom-server:7860"),
    drawthings.WithTimeout(15 * time.Minute),
    drawthings.WithLogger(&consoleLogger{}),
)
```

## Batch Generation

### Generate Multiple Images

```go
prompts := []string{
    "a beautiful sunset",
    "a futuristic city",
    "a peaceful forest",
}

for i, prompt := range prompts {
    req := &drawthings.TextToImageRequest{
        Prompt: prompt,
        Steps:  20,
    }
    
    outputPath := fmt.Sprintf("image_%d.png", i+1)
    err := client.GenerateImageAndSave(ctx, req, outputPath)
    if err != nil {
        log.Printf("Failed to generate image %d: %v", i+1, err)
        continue
    }
    
    fmt.Printf("Generated: %s\n", outputPath)
}
```

### Generate Variations

```go
basePrompt := "a beautiful landscape"
seeds := []int{42, 123, 456, 789}

for i, seed := range seeds {
    req := &drawthings.TextToImageRequest{
        Prompt: basePrompt,
        Seed:   seed,
        Steps:  30,
    }
    
    outputPath := fmt.Sprintf("variation_%d.png", i+1)
    err := client.GenerateImageAndSave(ctx, req, outputPath)
    if err != nil {
        log.Printf("Failed to generate variation %d: %v", i+1, err)
    }
}
```

## Working with Image Data

### Decode and Process Image

```go
resp, err := client.GenerateImage(ctx, req)
if err != nil {
    log.Fatal(err)
}

if len(resp.Images) == 0 {
    log.Fatal("no images in response")
}

// Decode base64 image
imageData, err := base64.StdEncoding.DecodeString(resp.Images[0])
if err != nil {
    log.Fatal(err)
}

// Process image data as needed
// For example, save to custom location or process with image library
err = os.WriteFile("custom_output.png", imageData, 0644)
```

### Save Multiple Images from Response

```go
resp, err := client.GenerateImage(ctx, req)
if err != nil {
    log.Fatal(err)
}

for i, imageData := range resp.Images {
    data, err := base64.StdEncoding.DecodeString(imageData)
    if err != nil {
        log.Printf("Failed to decode image %d: %v", i, err)
        continue
    }
    
    outputPath := fmt.Sprintf("image_%d.png", i+1)
    err = os.WriteFile(outputPath, data, 0644)
    if err != nil {
        log.Printf("Failed to save image %d: %v", i, err)
    }
}
```

## Advanced Patterns

### Context with Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

err := client.GenerateImageAndSave(ctx, req, "output.png")
if err != nil {
    if err == context.DeadlineExceeded {
        log.Println("Generation timed out")
    } else {
        log.Printf("Error: %v", err)
    }
}
```

### Concurrent Generation

```go
func generateConcurrent(prompts []string) {
    var wg sync.WaitGroup
    client := drawthings.NewClient()
    
    for i, prompt := range prompts {
        wg.Add(1)
        go func(idx int, p string) {
            defer wg.Done()
            
            req := &drawthings.TextToImageRequest{
                Prompt: p,
                Steps:  20,
            }
            
            outputPath := fmt.Sprintf("concurrent_%d.png", idx+1)
            err := client.GenerateImageAndSave(context.Background(), req, outputPath)
            if err != nil {
                log.Printf("Failed to generate %d: %v", idx+1, err)
            } else {
                log.Printf("Generated: %s", outputPath)
            }
        }(i, prompt)
    }
    
    wg.Wait()
}
```

### Parameter Presets

```go
type Preset struct {
    Name          string
    Steps         int
    GuidanceScale float64
    Width         int
    Height        int
}

var presets = map[string]Preset{
    "fast": {
        Steps:         20,
        GuidanceScale: 4.0,
        Width:         512,
        Height:        512,
    },
    "quality": {
        Steps:         50,
        GuidanceScale: 7.0,
        Width:         768,
        Height:        768,
    },
    "ultra": {
        Steps:         100,
        GuidanceScale: 9.0,
        Width:         1024,
        Height:        1024,
    },
}

func generateWithPreset(presetName, prompt string) error {
    preset, ok := presets[presetName]
    if !ok {
        return fmt.Errorf("unknown preset: %s", presetName)
    }
    
    req := &drawthings.TextToImageRequest{
        Prompt:        prompt,
        Steps:         preset.Steps,
        GuidanceScale: preset.GuidanceScale,
        Width:         preset.Width,
        Height:        preset.Height,
    }
    
    return client.GenerateImageAndSave(context.Background(), req, "output.png")
}
```

### Progress Tracking

```go
type progressLogger struct {
    total   int
    current int
}

func (l *progressLogger) Logf(format string, args ...interface{}) {
    // Filter and log only relevant messages
    if strings.Contains(format, "Response status") {
        l.current++
        fmt.Printf("Progress: %d/%d\n", l.current, l.total)
    }
}

func generateWithProgress(prompts []string) {
    logger := &progressLogger{total: len(prompts)}
    client := drawthings.NewClient(drawthings.WithLogger(logger))
    
    for _, prompt := range prompts {
        req := &drawthings.TextToImageRequest{Prompt: prompt}
        // Generate images...
    }
}
```

## More Examples

See the [examples directory](../examples/) for complete, runnable examples.

