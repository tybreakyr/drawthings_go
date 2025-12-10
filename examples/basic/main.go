package main

import (
	"context"
	"fmt"
	"log"

	drawthings "github.com/drawthings_go"
)

func main() {
	// Create a client with default settings
	client := drawthings.NewClient()

	// Example 1: Basic image generation
	fmt.Println("Example 1: Basic image generation")
	req1 := &drawthings.TextToImageRequest{
		Prompt: "a beautiful sunset over mountains, digital art",
		Steps:  20,
	}

	ctx := context.Background()
	err := client.GenerateImageAndSave(ctx, req1, "example1_basic.png")
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}
	fmt.Println("✓ Image saved to example1_basic.png")
	fmt.Println()

	// Example 2: High-quality generation with more parameters
	fmt.Println("Example 2: High-quality generation")
	req2 := &drawthings.TextToImageRequest{
		Prompt:         "a beautiful landscape, photorealistic, 4k",
		NegativePrompt: "blurry, low quality, distorted",
		Steps:          50,
		GuidanceScale:  7.0,
		Width:          768,
		Height:         768,
		Seed:           42, // Fixed seed for reproducibility
	}

	err = client.GenerateImageAndSave(ctx, req2, "example2_high_quality.png")
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}
	fmt.Println("✓ Image saved to example2_high_quality.png")
	fmt.Println()

	// Example 3: Get image data without saving
	fmt.Println("Example 3: Get image data")
	req3 := &drawthings.TextToImageRequest{
		Prompt: "a cat wearing sunglasses",
		Steps:  20,
	}

	resp, err := client.GenerateImage(ctx, req3)
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}

	if len(resp.Images) > 0 {
		fmt.Printf("✓ Generated %d image(s)\n", len(resp.Images))
		fmt.Printf("  First image base64 length: %d characters\n", len(resp.Images[0]))
	}

	// Example 4: Custom client configuration
	fmt.Println("\nExample 4: Custom client configuration")
	customClient := drawthings.NewClient(
		drawthings.WithBaseURL("http://127.0.0.1:7860"),
		drawthings.WithLogger(&consoleLogger{}),
	)

	req4 := &drawthings.TextToImageRequest{
		Prompt: "a futuristic city at night",
		Steps:  30,
	}

	err = customClient.GenerateImageAndSave(ctx, req4, "example4_custom.png")
	if err != nil {
		log.Fatalf("Failed to generate image: %v", err)
	}
	fmt.Println("✓ Image saved to example4_custom.png")
	fmt.Println()

	// Example 5: Error handling
	fmt.Println("Example 5: Error handling demonstration")
	req5 := &drawthings.TextToImageRequest{
		// Missing prompt - will cause validation error
		Steps: 20,
	}

	_, err = client.GenerateImage(ctx, req5)
	if err != nil {
		if drawthings.IsValidationError(err) {
			fmt.Printf("✓ Caught validation error: %v\n", err)
		} else {
			fmt.Printf("Unexpected error type: %v\n", err)
		}
	}

	fmt.Println("\nAll examples completed!")
}

// consoleLogger is a simple logger implementation
type consoleLogger struct{}

func (l *consoleLogger) Logf(format string, args ...interface{}) {
	fmt.Printf("[LOG] "+format+"\n", args...)
}
