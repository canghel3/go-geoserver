package validator

import (
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLayerValidator_Name(t *testing.T) {
	tests := []struct {
		name         string
		layerName    string
		wantErr      bool
		errorMessage string
	}{
		{
			name:      "Valid name",
			layerName: "validLayer",
			wantErr:   false,
		},
		{
			name:         "Empty name",
			layerName:    "",
			wantErr:      true,
			errorMessage: "empty name",
		},
		{
			name:         "Blank name",
			layerName:    "   ",
			wantErr:      true,
			errorMessage: "empty name",
		},
		{
			name:         "Invalid name with special characters",
			layerName:    "invalid!@#$%^&*()",
			wantErr:      true,
			errorMessage: "name can only contain alphanumerical characters",
		},
		{
			name:      "Valid name with underscore",
			layerName: "valid_layer",
			wantErr:   false,
		},
		{
			name:      "Valid name with hyphen",
			layerName: "valid-layer",
			wantErr:   false,
		},
		{
			name:      "Valid name with numbers",
			layerName: "layer123",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Name(tt.layerName)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWorkspaceLayerFormat(t *testing.T) {
	tests := []struct {
		name         string
		workspace    string
		layer        string
		wantErr      bool
		errorMessage string
	}{
		{
			name:      "No error",
			workspace: "wksp",
			layer:     "la",
			wantErr:   false,
		},
		{
			name:         "No workspace included in layer name",
			workspace:    "",
			layer:        "some",
			wantErr:      true,
			errorMessage: "unspecified workspace in layer name some. format the layer name as <workspace>:some",
		},
		{
			name:         "No workspace included in layer name, but layer name resembles workspace:layer format",
			workspace:    "",
			layer:        ":some",
			wantErr:      true,
			errorMessage: "unspecified workspace in layer name some. format the layer name as <workspace>:some",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WorkspaceLayerFormat(tt.workspace, tt.layer)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "String with only spaces",
			input:    "   ",
			expected: true,
		},
		{
			name:     "String with only tabs",
			input:    "\t\t\t",
			expected: true,
		},
		{
			name:     "String with only newlines",
			input:    "\n\n\n",
			expected: true,
		},
		{
			name:     "String with mixed whitespace",
			input:    " \t \n ",
			expected: true,
		},
		{
			name:     "Non-empty string",
			input:    "test",
			expected: false,
		},
		{
			name:     "String with whitespace and non-whitespace",
			input:    "  test  ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Empty(tt.input)
			assert.Equal(t, tt.expected, result, "Empty() returned unexpected result for input: %q", tt.input)
		})
	}
}

func TestValidateAlphaNumerical(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid alphanumerical string",
			input:   "validString123",
			wantErr: false,
		},
		{
			name:    "Valid string with underscore",
			input:   "valid_string",
			wantErr: false,
		},
		{
			name:    "Valid string with hyphen",
			input:   "valid-string",
			wantErr: false,
		},
		{
			name:         "Invalid string with special characters",
			input:        "invalid!@#$%^&*()",
			wantErr:      true,
			errorMessage: "name can only contain alphanumerical characters",
		},
		{
			name:         "Invalid string with spaces",
			input:        "invalid string",
			wantErr:      true,
			errorMessage: "name can only contain alphanumerical characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAlphaNumerical(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
