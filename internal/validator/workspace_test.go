package validator

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkspaceValidator_Name(t *testing.T) {
	tests := []struct {
		name          string
		workspaceName string
		wantErr       bool
		errorMessage  string
	}{
		{
			name:          "Valid workspace name",
			workspaceName: testdata.Workspace,
			wantErr:       false,
		},
		{
			name:          "Empty workspace name",
			workspaceName: "",
			wantErr:       true,
			errorMessage:  "empty workspace name",
		},
		{
			name:          "Invalid workspace name with special characters",
			workspaceName: testdata.InvalidWorkspaceName,
			wantErr:       true,
			errorMessage:  "name can only contain alphanumerical characters",
		},
		{
			name:          "Valid workspace name with underscore",
			workspaceName: "valid_workspace",
			wantErr:       false,
		},
		{
			name:          "Valid workspace name with hyphen",
			workspaceName: "valid-workspace",
			wantErr:       false,
		},
		{
			name:          "Valid workspace name with numbers",
			workspaceName: "workspace123",
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wv := WorkspaceValidator{}
			err := wv.Name(tt.workspaceName)

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
