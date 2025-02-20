package service

import (
	"github.com/canghel3/go-geoserver/utils"
	"gotest.tools/v3/assert"
	"testing"
)

func TestWorkspace(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)

	t.Run("CREATE", func(t *testing.T) {
		t.Run("SIMPLE", func(t *testing.T) {
			assert.NilError(t, geoserverService.CreateWorkspace("init"))
		})

		t.Run("DUPLICATE", func(t *testing.T) {
			assert.Error(t, geoserverService.CreateWorkspace("init"), "workspace init already exists")
		})

		t.Run("NAME WITH SPECIAL CHARACTERS", func(t *testing.T) {
			assert.Error(t, geoserverService.CreateWorkspace("init!@#$%^&*()"), "name can only contain alphanumerical characters")
		})

		t.Run("WITH EMPTY NAME", func(t *testing.T) {
			assert.Error(t, geoserverService.CreateWorkspace(""), "empty workspace name")
		})
	})

	t.Run("GET", func(t *testing.T) {
		t.Run("WORKSPACE", func(t *testing.T) {
			workspace, err := geoserverService.GetWorkspace("init")
			assert.NilError(t, err)
			assert.Equal(t, workspace.Workspace.Name, "init")
		})

		t.Run("NON-EXISTENT WORKSPACE", func(t *testing.T) {
			_, err := geoserverService.GetWorkspace("nonexisting")
			assert.Error(t, err, "workspace nonexisting does not exist")
		})

		t.Run("WORKSPACES", func(t *testing.T) {
			workspaces, err := geoserverService.GetWorkspaces()
			assert.NilError(t, err)
			assert.Assert(t, len(workspaces.Workspaces.Workspace) == 1)
		})
	})

	t.Run("UPDATE", func(t *testing.T) {
		t.Skip()
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Run("WORKSPACE", func(t *testing.T) {
			assert.NilError(t, geoserverService.DeleteWorkspace("init", utils.RecurseOption(true)))
			_, err := geoserverService.GetWorkspace("init")
			assert.Error(t, err, "workspace init does not exist")
		})

		t.Run("NON-EXISTENT WORKSPACE", func(t *testing.T) {
			assert.Error(t, geoserverService.DeleteWorkspace("nonexisting"), "workspace nonexisting does not exist")
		})
	})
}
