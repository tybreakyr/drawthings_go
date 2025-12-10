package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/drawthings_go"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		prompt         = flag.String("prompt", "", "Textual description of the desired image (required)")
		negativePrompt = flag.String("negative-prompt", "", "Descriptions of elements to exclude from the image")
		steps          = flag.Int("steps", 20, "Number of inference steps (1-150, default: 20)")
		guidanceScale  = flag.Float64("guidance-scale", 4.0, "Controls adherence to the prompt (1.0-20.0, default: 4.0)")
		width          = flag.Int("width", 512, "Width of the generated image in pixels (default: 512)")
		height         = flag.Int("height", 512, "Height of the generated image in pixels (default: 512)")
		seed           = flag.Int("seed", -1, "Random seed for image generation (-1 for random, default: -1)")
		output         = flag.String("output", "output.png", "Output file path for the generated image")
		baseURL        = flag.String("base-url", drawthings.DefaultBaseURL, "Base URL of the Draw Things API server")
		timeout        = flag.Duration("timeout", drawthings.DefaultTimeout, "HTTP client timeout")
		showVersion    = flag.Bool("version", false, "Show version information")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Generate images using the Draw Things API.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -prompt \"a beautiful sunset\"\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -prompt \"a cat\" -steps 30 -width 768 -height 768 -output cat.png\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -prompt \"landscape\" -seed 42 -guidance-scale 7.0\n", os.Args[0])
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("drawthings version %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("built: %s\n", date)
		return nil
	}

	if *prompt == "" {
		flag.Usage()
		return fmt.Errorf("prompt is required")
	}

	// Create client
	client := drawthings.NewClient(
		drawthings.WithBaseURL(*baseURL),
		drawthings.WithTimeout(*timeout),
	)

	// Create request
	req := &drawthings.TextToImageRequest{
		Prompt:         *prompt,
		NegativePrompt: *negativePrompt,
		Steps:          *steps,
		GuidanceScale:  *guidanceScale,
		Width:          *width,
		Height:         *height,
		Seed:           *seed,
	}

	// Generate and save image
	fmt.Printf("Generating image with prompt: %q\n", *prompt)
	fmt.Printf("Parameters: steps=%d, guidance_scale=%.2f, width=%d, height=%d, seed=%d\n",
		*steps, *guidanceScale, *width, *height, *seed)

	ctx := context.Background()
	if err := client.GenerateImageAndSave(ctx, req, *output); err != nil {
		return fmt.Errorf("failed to generate image: %w", err)
	}

	fmt.Printf("Image saved to: %s\n", *output)
	return nil
}

