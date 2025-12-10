package drawthings

// TextToImageRequest represents a request to generate an image from text.
type TextToImageRequest struct {
	// Prompt is the textual description of the desired image (required).
	Prompt string `json:"prompt"`

	// NegativePrompt describes elements to exclude from the image (optional).
	NegativePrompt string `json:"negative_prompt,omitempty"`

	// Steps is the number of inference steps. Higher values may improve quality but take longer.
	// Range: 1-150, Recommended: 20-50, Default: 20
	Steps int `json:"steps,omitempty"`

	// GuidanceScale controls adherence to the prompt. Higher values result in images more closely matching the prompt.
	// Range: 1-20, Recommended: 4-7, Default: 4.0
	GuidanceScale float64 `json:"guidance_scale,omitempty"`

	// Width is the width of the generated image in pixels.
	// Common sizes: 512, 768, 1024, Default: 512
	Width int `json:"width,omitempty"`

	// Height is the height of the generated image in pixels.
	// Common sizes: 512, 768, 1024, Default: 512
	Height int `json:"height,omitempty"`

	// Seed is the random seed for image generation. Use -1 for a random seed.
	// Default: -1
	Seed int `json:"seed,omitempty"`
}

// SetDefaults sets default values for optional fields if they are zero values.
func (r *TextToImageRequest) SetDefaults() {
	if r.Steps == 0 {
		r.Steps = 20
	}
	if r.GuidanceScale == 0 {
		r.GuidanceScale = 4.0
	}
	if r.Width == 0 {
		r.Width = 512
	}
	if r.Height == 0 {
		r.Height = 512
	}
	if r.Seed == 0 {
		r.Seed = -1
	}
}

// TextToImageResponse represents the response from a text-to-image generation request.
type TextToImageResponse struct {
	// Images contains base64-encoded image data.
	Images []string `json:"images"`
}

