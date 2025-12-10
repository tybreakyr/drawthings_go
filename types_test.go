package drawthings

import (
	"testing"
)

func TestTextToImageRequest_SetDefaults(t *testing.T) {
	tests := []struct {
		name     string
		req      *TextToImageRequest
		expected *TextToImageRequest
	}{
		{
			name: "all defaults",
			req: &TextToImageRequest{
				Prompt: "test",
			},
			expected: &TextToImageRequest{
				Prompt:        "test",
				Steps:         20,
				GuidanceScale: 4.0,
				Width:         512,
				Height:        512,
				Seed:          -1,
			},
		},
		{
			name: "no defaults needed",
			req: &TextToImageRequest{
				Prompt:        "test",
				Steps:         30,
				GuidanceScale: 7.0,
				Width:         768,
				Height:        768,
				Seed:          42,
			},
			expected: &TextToImageRequest{
				Prompt:        "test",
				Steps:         30,
				GuidanceScale: 7.0,
				Width:         768,
				Height:        768,
				Seed:          42,
			},
		},
		{
			name: "partial defaults",
			req: &TextToImageRequest{
				Prompt: "test",
				Steps:  50,
				Width:  1024,
			},
			expected: &TextToImageRequest{
				Prompt:        "test",
				Steps:         50,
				GuidanceScale: 4.0,
				Width:         1024,
				Height:        512,
				Seed:          -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.req.SetDefaults()
			if tt.req.Prompt != tt.expected.Prompt {
				t.Errorf("Prompt: got %q, want %q", tt.req.Prompt, tt.expected.Prompt)
			}
			if tt.req.Steps != tt.expected.Steps {
				t.Errorf("Steps: got %d, want %d", tt.req.Steps, tt.expected.Steps)
			}
			if tt.req.GuidanceScale != tt.expected.GuidanceScale {
				t.Errorf("GuidanceScale: got %f, want %f", tt.req.GuidanceScale, tt.expected.GuidanceScale)
			}
			if tt.req.Width != tt.expected.Width {
				t.Errorf("Width: got %d, want %d", tt.req.Width, tt.expected.Width)
			}
			if tt.req.Height != tt.expected.Height {
				t.Errorf("Height: got %d, want %d", tt.req.Height, tt.expected.Height)
			}
			if tt.req.Seed != tt.expected.Seed {
				t.Errorf("Seed: got %d, want %d", tt.req.Seed, tt.expected.Seed)
			}
		})
	}
}

