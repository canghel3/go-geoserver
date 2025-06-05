package validator

import (
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreValidator_Name(t *testing.T) {
	tests := []struct {
		name         string
		storeName    string
		wantErr      bool
		errorMessage string
	}{
		{
			name:      "Valid store name",
			storeName: "validStore",
			wantErr:   false,
		},
		{
			name:         "Empty store name",
			storeName:    "",
			wantErr:      true,
			errorMessage: "empty store name",
		},
		{
			name:         "Invalid store name with special characters",
			storeName:    "invalid!@#$%^&*()",
			wantErr:      true,
			errorMessage: "name can only contain alphanumerical characters",
		},
		{
			name:      "Valid store name with underscore",
			storeName: "valid_store",
			wantErr:   false,
		},
		{
			name:      "Valid store name with hyphen",
			storeName: "valid-store",
			wantErr:   false,
		},
		{
			name:      "Valid store name with numbers",
			storeName: "store123",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := StoreValidator{}
			err := sv.Name(tt.storeName)

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
