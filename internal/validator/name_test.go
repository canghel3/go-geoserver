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

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
