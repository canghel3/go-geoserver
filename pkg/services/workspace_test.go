package services

import (
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkspaces(t *testing.T) {
	t.Run("CREATE", func(t *testing.T) {
		t.Run("INVALID NAME", func(t *testing.T) {
			ws := &Workspaces{}

			var inputError *customerrors.InputError
			err := ws.Create(testdata.INVALID_WORKSPACE_NAME, false)
			assert.Error(t, err)
			assert.EqualError(t, err, "name can only contain alphanumerical characters")
			assert.ErrorAs(t, err, &inputError)
		})
	})

	t.Run("GET", func(t *testing.T) {
		t.Run("INVALID NAME", func(t *testing.T) {
			ws := &Workspaces{}

			var inputError *customerrors.InputError
			_, err := ws.Get(testdata.INVALID_WORKSPACE_NAME)
			assert.Error(t, err)
			assert.EqualError(t, err, "name can only contain alphanumerical characters")
			assert.ErrorAs(t, err, &inputError)
		})
	})

	t.Run("UPDATE", func(t *testing.T) {
		t.Run("INVALID OLD NAME", func(t *testing.T) {
			ws := &Workspaces{}

			var inputError *customerrors.InputError
			err := ws.Update(testdata.INVALID_WORKSPACE_NAME, "")
			assert.Error(t, err)
			assert.EqualError(t, err, "name can only contain alphanumerical characters")
			assert.ErrorAs(t, err, &inputError)
		})

		t.Run("INVALID NEW NAME", func(t *testing.T) {
			ws := &Workspaces{}

			var inputError *customerrors.InputError
			err := ws.Update(testdata.WORKSPACE, testdata.INVALID_WORKSPACE_NAME)
			assert.Error(t, err)
			assert.EqualError(t, err, "name can only contain alphanumerical characters")
			assert.ErrorAs(t, err, &inputError)
		})
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Run("INVALID NAME", func(t *testing.T) {
			ws := &Workspaces{}

			var inputError *customerrors.InputError
			err := ws.Delete(testdata.INVALID_WORKSPACE_NAME, false)
			assert.Error(t, err)
			assert.EqualError(t, err, "name can only contain alphanumerical characters")
			assert.ErrorAs(t, err, &inputError)
		})
	})
}
