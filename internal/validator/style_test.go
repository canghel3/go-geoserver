package validator

import (
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStyleValidator_Name(t *testing.T) {
	tests := []struct {
		name         string
		styleName    string
		wantErr      bool
		errorMessage string
	}{
		{
			name:      "Valid style name",
			styleName: "validStyle",
			wantErr:   false,
		},
		{
			name:         "Empty style name",
			styleName:    "",
			wantErr:      true,
			errorMessage: "empty style name",
		},
		{
			name:         "Invalid style name with special characters",
			styleName:    "invalid!@#$%^&*()",
			wantErr:      true,
			errorMessage: "name can only contain alphanumerical characters",
		},
		{
			name:      "Valid style name with underscore",
			styleName: "valid_style",
			wantErr:   false,
		},
		{
			name:      "Valid style name with hyphen",
			styleName: "valid-style",
			wantErr:   false,
		},
		{
			name:      "Valid style name with numbers",
			styleName: "style123",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := StyleValidator{}
			err := sv.Name(tt.styleName)

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
