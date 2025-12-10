# Troubleshooting

Common issues and solutions when using the Draw Things Go client library.

## Connection Issues

### Connection Refused

**Error:**
```
network error: request failed: dial tcp 127.0.0.1:7860: connect: connection refused
```

**Solutions:**
1. Ensure Draw Things is running
2. Verify the API server is enabled in Draw Things settings
3. Check that port 7860 is not blocked by firewall
4. Verify the base URL is correct:
   ```go
   client := drawthings.NewClient(
       drawthings.WithBaseURL("http://127.0.0.1:7860"),
   )
   ```

### Connection Timeout

**Error:**
```
network error: request failed: context deadline exceeded
```

**Solutions:**
1. Increase the timeout:
   ```go
   client := drawthings.NewClient(
       drawthings.WithTimeout(10 * time.Minute),
   )
   ```
2. Check if the server is responding:
   ```bash
   curl http://127.0.0.1:7860/sdapi/v1/options
   ```
3. Reduce image dimensions or steps to speed up generation

## Validation Errors

### Missing Prompt

**Error:**
```
validation error for field 'prompt': prompt is required and cannot be empty
```

**Solution:**
Always provide a prompt:
```go
req := &drawthings.TextToImageRequest{
    Prompt: "your prompt here", // Required!
    Steps:  20,
}
```

### Invalid Parameter Ranges

**Error:**
```
validation error for field 'steps': steps must be between 1 and 150, got 200
```

**Solutions:**
- **Steps**: Must be between 1 and 150
- **Guidance Scale**: Must be between 1.0 and 20.0
- **Width/Height**: Must be between 64 and 4096 pixels

```go
req := &drawthings.TextToImageRequest{
    Prompt:        "test",
    Steps:         50,        // Valid: 1-150
    GuidanceScale:  7.0,       // Valid: 1.0-20.0
    Width:         768,        // Valid: 64-4096
    Height:        768,         // Valid: 64-4096
}
```

## API Errors

### 500 Internal Server Error

**Error:**
```
API error (status 500): Internal Server Error
```

**Solutions:**
1. Check Draw Things application logs
2. Verify the model is loaded in Draw Things
3. Try reducing image dimensions or steps
4. Restart the Draw Things application

### 400 Bad Request

**Error:**
```
API error (status 400): Bad Request
```

**Solutions:**
1. Verify all parameters are valid
2. Check the request format matches API expectations
3. Review the Draw Things API documentation

## Decode Errors

### Base64 Decoding Failed

**Error:**
```
decode error: failed to decode base64 image data: illegal base64 data
```

**Solutions:**
1. Verify the API response contains valid image data
2. Check if the response format matches expectations
3. Enable logging to inspect the raw response:
   ```go
   client := drawthings.NewClient(
       drawthings.WithLogger(&myLogger{}),
   )
   ```

### No Images in Response

**Error:**
```
decode error: no images in response
```

**Solutions:**
1. Check the API response structure
2. Verify the generation completed successfully
3. Review Draw Things application status

## Performance Issues

### Slow Generation

**Symptoms:**
- Generation takes a very long time
- Timeout errors occur

**Solutions:**
1. Reduce steps (e.g., from 50 to 20):
   ```go
   req.Steps = 20
   ```
2. Reduce image dimensions:
   ```go
   req.Width = 512
   req.Height = 512
   ```
3. Increase timeout for large images:
   ```go
   client := drawthings.NewClient(
       drawthings.WithTimeout(15 * time.Minute),
   )
   ```

### Memory Issues

**Symptoms:**
- Application crashes
- Out of memory errors

**Solutions:**
1. Use smaller image dimensions
2. Reduce batch sizes if generating multiple images
3. Close unused resources properly

## Debugging

### Enable Logging

Add a logger to see request/response details:

```go
type debugLogger struct{}

func (l *debugLogger) Logf(format string, args ...interface{}) {
    fmt.Printf("[DEBUG] "+format+"\n", args...)
}

client := drawthings.NewClient(
    drawthings.WithLogger(&debugLogger{}),
)
```

### Check Client Configuration

```go
fmt.Printf("Base URL: %s\n", client.BaseURL())
```

### Verify Request Parameters

```go
req := &drawthings.TextToImageRequest{
    Prompt: "test",
    Steps:  20,
}
req.SetDefaults() // Ensure defaults are set
fmt.Printf("Request: %+v\n", req)
```

## Common Mistakes

### Forgetting to Set Context

**Problem:**
```go
// Missing context
err := client.GenerateImage(nil, req)
```

**Solution:**
```go
ctx := context.Background()
err := client.GenerateImage(ctx, req)
```

### Not Handling Errors

**Problem:**
```go
client.GenerateImageAndSave(ctx, req, "output.png")
// Error ignored!
```

**Solution:**
```go
err := client.GenerateImageAndSave(ctx, req, "output.png")
if err != nil {
    log.Fatalf("Failed: %v", err)
}
```

### Incorrect File Paths

**Problem:**
```go
// Missing directory
err := client.GenerateImageAndSave(ctx, req, "/nonexistent/path/output.png")
```

**Solution:**
The library creates directories automatically, but ensure you have write permissions.

## Getting Help

If you're still experiencing issues:

1. **Check the Logs**: Enable logging to see detailed request/response information
2. **Verify Setup**: Ensure Draw Things API server is running and accessible
3. **Review Examples**: Check the [examples documentation](examples.md)
4. **API Reference**: Consult the [API reference](api-reference.md)
5. **Draw Things Documentation**: See the [official Draw Things wiki](https://wiki.drawthings.ai)

## Error Type Reference

### Identifying Error Types

```go
if err != nil {
    switch {
    case drawthings.IsValidationError(err):
        // Parameter validation failed
    case drawthings.IsAPIError(err):
        // API returned an error
    case drawthings.IsNetworkError(err):
        // Network/connection issue
    case drawthings.IsDecodeError(err):
        // Image decoding failed
    default:
        // Unknown error
    }
}
```

### Error Details

Each error type provides specific information:

- **ValidationError**: Field name and validation message
- **APIError**: HTTP status code and response body
- **NetworkError**: Network error message and underlying error
- **DecodeError**: Decode error message and underlying error

Use these details to diagnose and fix issues.

