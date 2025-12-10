package validation

import (
	"fmt"
)

// ValidateTextToImageRequest validates request parameters.
// This is a helper function that works with the request struct fields directly.
func ValidateTextToImageRequest(prompt string, steps int, guidanceScale float64, width, height int) error {
	if prompt == "" {
		return fmt.Errorf("validation error for field 'prompt': prompt is required and cannot be empty")
	}

	if steps < 1 || steps > 150 {
		return fmt.Errorf("validation error for field 'steps': steps must be between 1 and 150, got %d", steps)
	}

	if guidanceScale < 1.0 || guidanceScale > 20.0 {
		return fmt.Errorf("validation error for field 'guidance_scale': guidance_scale must be between 1.0 and 20.0, got %.2f", guidanceScale)
	}

	if width < 64 || width > 4096 {
		return fmt.Errorf("validation error for field 'width': width must be between 64 and 4096 pixels, got %d", width)
	}

	if height < 64 || height > 4096 {
		return fmt.Errorf("validation error for field 'height': height must be between 64 and 4096 pixels, got %d", height)
	}

	return nil
}

