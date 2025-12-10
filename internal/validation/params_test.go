package validation

import (
	"strings"
	"testing"
)

func TestValidateTextToImageRequest(t *testing.T) {
	tests := []struct {
		name          string
		prompt        string
		steps         int
		guidanceScale float64
		width         int
		height        int
		wantErr       bool
	}{
		{
			name:          "valid request",
			prompt:        "a beautiful sunset",
			steps:         20,
			guidanceScale: 4.0,
			width:         512,
			height:        512,
			wantErr:       false,
		},
		{
			name:          "missing prompt",
			prompt:        "",
			steps:         20,
			guidanceScale: 4.0,
			width:         512,
			height:        512,
			wantErr:       true,
		},
		{
			name:          "steps too low",
			prompt:        "test",
			steps:         0,
			guidanceScale: 4.0,
			width:         512,
			height:        512,
			wantErr:       true,
		},
		{
			name:          "steps too high",
			prompt:        "test",
			steps:         200,
			guidanceScale: 4.0,
			width:         512,
			height:        512,
			wantErr:       true,
		},
		{
			name:          "guidance_scale too low",
			prompt:        "test",
			steps:         20,
			guidanceScale: 0.5,
			width:         512,
			height:        512,
			wantErr:       true,
		},
		{
			name:          "guidance_scale too high",
			prompt:        "test",
			steps:         20,
			guidanceScale: 25.0,
			width:         512,
			height:        512,
			wantErr:       true,
		},
		{
			name:          "width too low",
			prompt:        "test",
			steps:         20,
			guidanceScale: 4.0,
			width:         32,
			height:        512,
			wantErr:       true,
		},
		{
			name:          "width too high",
			prompt:        "test",
			steps:         20,
			guidanceScale: 4.0,
			width:         8192,
			height:        512,
			wantErr:       true,
		},
		{
			name:          "height too low",
			prompt:        "test",
			steps:         20,
			guidanceScale: 4.0,
			width:         512,
			height:        32,
			wantErr:       true,
		},
		{
			name:          "height too high",
			prompt:        "test",
			steps:         20,
			guidanceScale: 4.0,
			width:         512,
			height:        8192,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTextToImageRequest(tt.prompt, tt.steps, tt.guidanceScale, tt.width, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTextToImageRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), "validation error") {
					t.Errorf("expected validation error message, got %q", err.Error())
				}
			}
		})
	}
}

