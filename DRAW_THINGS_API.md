# Draw Things REST API Documentation

## Overview

Draw Things is an AI-assisted image generation application available for iPhone, iPad, and Mac. It allows users to create images from textual descriptions using advanced AI models. The application supports running an HTTP server, enabling programmatic access to its image generation capabilities through a REST API.

**Official Website**: [drawthings.ai](https://drawthings.ai)

## Prerequisites

- **Draw Things Application**: Install the Draw Things app on your device (macOS, iPhone, or iPad)
- **API Server Activation**: Enable the HTTP server within the Draw Things application
- **Network Configuration**: By default, the API server runs on `localhost` (127.0.0.1) and listens on port `7860`

## Setting Up the API Server

To utilize the Draw Things REST API, you need to start the application's HTTP server:

1. **Launch Draw Things**: Open the Draw Things application on your device
2. **Enable the HTTP Server**:
   - Navigate to the application's settings
   - Locate the "API Server" option
   - Enable the "HTTP" server
   - Note the server address (default: `http://127.0.0.1:7860`)

**Important**: Ensure that port 7860 is available and not blocked by any firewall or security software.

## Base URL

```
http://127.0.0.1:7860
```

## API Endpoints

### Text-to-Image Generation

Generate an image from a textual prompt.

- **Endpoint**: `POST /sdapi/v1/txt2img`
- **Content-Type**: `application/json`

#### Request Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `prompt` | string | Yes | The textual description of the desired image |
| `negative_prompt` | string | No | Descriptions of elements to exclude from the image |
| `steps` | integer | No | Number of inference steps (default: 20). Higher values may improve image quality but take longer |
| `guidance_scale` | float | No | Controls adherence to the prompt (default: 4). Higher values result in images more closely matching the prompt |
| `width` | integer | No | Width of the generated image in pixels (default: 512) |
| `height` | integer | No | Height of the generated image in pixels (default: 512) |
| `seed` | integer | No | Random seed for image generation. Use `-1` for a random seed (default: -1) |

#### Request Example

```json
{
  "prompt": "a beautiful sunset over mountains, digital art",
  "negative_prompt": "blurry, low quality",
  "steps": 20,
  "guidance_scale": 4,
  "width": 512,
  "height": 512,
  "seed": -1
}
```

#### cURL Example

```bash
curl -X POST http://127.0.0.1:7860/sdapi/v1/txt2img \
     -H "Content-Type: application/json" \
     -d '{
           "prompt": "a beautiful sunset over mountains, digital art",
           "negative_prompt": "",
           "steps": 20,
           "guidance_scale": 4,
           "width": 512,
           "height": 512,
           "seed": -1
         }'
```

#### Response Format

The API responds with a JSON object containing the generated image(s) encoded in base64 format:

```json
{
  "images": ["base64-encoded-image-data"]
}
```

#### Saving the Image

To save the generated image to a file, you need to:
1. Extract the base64 string from the `images` array
2. Decode the base64 string
3. Write the decoded data to an image file (typically PNG format)

**Python Example**:
```python
import base64
import json

# Assuming response is the JSON response from the API
response_data = json.loads(response)
image_data = response_data["images"][0]

# Decode and save
image_bytes = base64.b64decode(image_data)
with open("generated_image.png", "wb") as f:
    f.write(image_bytes)
```

**Bash Example**:
```bash
# Extract base64 and decode
curl -X POST http://127.0.0.1:7860/sdapi/v1/txt2img \
     -H "Content-Type: application/json" \
     -d '{"prompt": "a beautiful sunset"}' \
     | jq -r '.images[0]' \
     | base64 -d > output.png
```

## API Compatibility

The Draw Things REST API follows a structure similar to **Stable Diffusion's API**, indicating compatibility with existing tools and scripts designed for Stable Diffusion. This means:

- Endpoints follow the `/sdapi/v1/` path structure
- Request/response formats are compatible
- Many Stable Diffusion API clients may work with Draw Things

### Potential Additional Endpoints

Since Draw Things is compatible with Stable Diffusion's API structure, the following endpoints may also be available (verify with your installation):

- `GET /sdapi/v1/options` - Get current options
- `POST /sdapi/v1/options` - Set options
- `POST /sdapi/v1/img2img` - Image-to-image generation
- `GET /sdapi/v1/samplers` - List available samplers
- `GET /sdapi/v1/models` - List available models
- `GET /sdapi/v1/sd-models` - List Stable Diffusion models
- `GET /sdapi/v1/upscalers` - List available upscalers

**Note**: These endpoints should be verified by testing with your Draw Things installation.

## Parameter Guidelines

### Steps
- **Range**: Typically 1-150
- **Recommended**: 20-50 for good balance of quality and speed
- **Higher values**: Better quality but slower generation

### Guidance Scale
- **Range**: Typically 1-20
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

## Tips and Best Practices

1. **Prompt Engineering**: 
   - Be specific and descriptive
   - Use commas to separate concepts
   - Include style keywords (e.g., "digital art", "photorealistic", "watercolor")

2. **Negative Prompts**:
   - Specify what you don't want in the image
   - Common negative prompts: "blurry", "low quality", "distorted", "ugly"

3. **Performance**:
   - Start with lower steps (20) to test prompts quickly
   - Increase steps for final high-quality images
   - Monitor memory usage with larger dimensions

4. **Reproducibility**:
   - Use fixed seeds when iterating on a prompt
   - Save successful parameter combinations

5. **Error Handling**:
   - Check that the API server is running before making requests
   - Handle network timeouts for longer generation times
   - Validate image dimensions and parameters before sending

## Troubleshooting

### Common Issues

1. **Connection Refused**
   - Ensure Draw Things is running
   - Verify the API server is enabled in settings
   - Check that port 7860 is not blocked

2. **Timeout Errors**
   - Increase timeout settings for requests
   - Reduce image dimensions or steps
   - Check system resources (CPU, memory)

3. **Invalid Parameters**
   - Verify parameter types (integers vs strings)
   - Check parameter ranges
   - Ensure required parameters are included

4. **Base64 Decoding Errors**
   - Verify the response format
   - Check that the image data is complete
   - Ensure proper base64 decoding

## Additional Resources

### Official Resources
- **Draw Things Website**: [drawthings.ai](https://drawthings.ai)
- **Draw Things Wiki**: [wiki.drawthings.ai](https://wiki.drawthings.ai/wiki/Help)
- **Draw Things Community Repository**: [github.com/drawthingsai/draw-things-community](https://github.com/drawthingsai/draw-things-community)

### Example Clients
- **TinyDTClient**: A minimal macOS SwiftUI application demonstrating basic client functions for the Draw Things API
  - Repository: [github.com/S1D1T1/TinyDTClient](https://github.com/S1D1T1/TinyDTClient)
  - Shows how to connect to Draw Things server and send image generation requests

### Related Documentation
- **Stable Diffusion API**: Since Draw Things is compatible, Stable Diffusion API documentation may be helpful
- **Stable Diffusion WebUI API**: Reference for additional endpoint structures

## Example Workflows

### Basic Image Generation

```bash
# Generate a simple image
curl -X POST http://127.0.0.1:7860/sdapi/v1/txt2img \
     -H "Content-Type: application/json" \
     -d '{
           "prompt": "a cat wearing sunglasses",
           "steps": 20,
           "width": 512,
           "height": 512
         }'
```

### Reproducible Generation

```bash
# Use a fixed seed for reproducibility
curl -X POST http://127.0.0.1:7860/sdapi/v1/txt2img \
     -H "Content-Type: application/json" \
     -d '{
           "prompt": "a cat wearing sunglasses",
           "seed": 42,
           "steps": 30,
           "width": 512,
           "height": 512
         }'
```

### High-Quality Generation

```bash
# Higher steps and guidance for better quality
curl -X POST http://127.0.0.1:7860/sdapi/v1/txt2img \
     -H "Content-Type: application/json" \
     -d '{
           "prompt": "a beautiful landscape, photorealistic, 4k",
           "negative_prompt": "blurry, low quality, distorted",
           "steps": 50,
           "guidance_scale": 7,
           "width": 768,
           "height": 768
         }'
```

## Notes

- The Draw Things application must be running with the API server enabled
- The API structure is compatible with Stable Diffusion's API
- Images are returned as base64-encoded PNG data
- The API runs locally by default (localhost)
- For production use, consider security implications of exposing the API

## Version Information

This documentation is based on information available as of the creation date. API endpoints and parameters may vary by Draw Things version. Always verify endpoint availability and parameter support with your specific installation.

---

**Last Updated**: Based on available information as of documentation creation

**Contributing**: If you discover additional endpoints or parameters, please update this documentation.

