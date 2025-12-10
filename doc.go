// Package drawthings provides a Go client library for the Draw Things API.
//
// This is an unofficial library and is not affiliated with, endorsed by, or
// connected to the makers of the Draw Things application. This library is a
// community-maintained project created independently to provide a Go client
// for the Draw Things API.
//
// Draw Things is an AI-assisted image generation application that allows users
// to create images from textual descriptions using advanced AI models. This
// library provides a simple and robust interface to interact with the Draw
// Things HTTP API server.
//
// Quick Start
//
//	import "github.com/drawthings_go"
//
//	client := drawthings.NewClient()
//	req := &drawthings.TextToImageRequest{
//		Prompt: "a beautiful sunset over mountains",
//		Steps:  20,
//	}
//
//	ctx := context.Background()
//	err := client.GenerateImageAndSave(ctx, req, "output.png")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// The client supports various configuration options:
//
//	client := drawthings.NewClient(
//		drawthings.WithBaseURL("http://custom-server:7860"),
//		drawthings.WithTimeout(10 * time.Minute),
//		drawthings.WithLogger(myLogger),
//	)
//
// # Error Handling
//
// The library provides specific error types for different failure scenarios:
//   - APIError: Errors returned by the API server
//   - ValidationError: Parameter validation failures
//   - NetworkError: Network-related errors (timeouts, connection issues)
//   - DecodeError: Errors during image decoding or processing
//
// Use the Is* functions to check error types:
//
//	if drawthings.IsAPIError(err) {
//		apiErr := err.(*drawthings.APIError)
//		fmt.Printf("API returned status %d\n", apiErr.StatusCode)
//	}
//
// For more information, see the documentation at:
// https://github.com/drawthings_go
package drawthings
