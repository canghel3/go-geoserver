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
			name:      "Valid layer name",
			layerName: "validLayer",
			wantErr:   false,
		},
		{
			name:         "Empty layer name",
			layerName:    "",
			wantErr:      true,
			errorMessage: "layer store name",
		},
		{
			name:         "Invalid layer name with special characters",
			layerName:    "invalid!@#$%^&*()",
			wantErr:      true,
			errorMessage: "name can only contain alphanumerical characters",
		},
		{
			name:      "Valid layer name with underscore",
			layerName: "valid_layer",
			wantErr:   false,
		},
		{
			name:      "Valid layer name with hyphen",
			layerName: "valid-layer",
			wantErr:   false,
		},
		{
			name:      "Valid layer name with numbers",
			layerName: "layer123",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lv := LayerValidator{}
			err := lv.Name(tt.layerName)

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
