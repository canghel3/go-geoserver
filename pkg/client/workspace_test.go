package client

import (
	"fmt"
	"testing"

	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceIntegration_Create(t *testing.T) {
	err := geoclient.Workspaces().Delete(testdata.Workspace, true)
	assert.NoError(t, err)

	t.Run("201 Created", func(t *testing.T) {
		err := geoclient.Workspaces().Create(testdata.Workspace, false)
		assert.NoError(t, err)
	})

	t.Run("Already Exists", func(t *testing.T) {
		err := geoclient.Workspaces().Create(testdata.Workspace, false)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.ErrorContains(t, err, "received status code 409 from geoserver")
	})

	t.Run("Invalid Name", func(t *testing.T) {
		err = geoclient.Workspaces().Create(testdata.InvalidName, false)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})
}

func TestWorkspaceIntegration_Get(t *testing.T) {
	addTestWorkspace(t)

	t.Run("200 Ok", func(t *testing.T) {
		wksp, err := geoclient.Workspaces().Get(testdata.Workspace)
		assert.NoError(t, err)
		assert.Equal(t, testdata.Workspace, wksp.Name)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		var suffix = "_NOT_FOUND"
		wksp, err := geoclient.Workspaces().Get(testdata.Workspace + suffix)
		assert.Nil(t, wksp)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace+suffix))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestWorkspaceIntegration_Update(t *testing.T) {
	addTestWorkspace(t)
	var toRemove string

	t.Run("200 Ok", func(t *testing.T) {
		var suffix = "_UPDATED"
		toRemove = testdata.Workspace + suffix
		err := geoclient.Workspaces().Update(testdata.Workspace, testdata.Workspace+suffix)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		var suffix = "_NOT_FOUND"
		err := geoclient.Workspaces().Update(testdata.Workspace, testdata.Workspace+suffix)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	t.Run("Invalid Previous Name", func(t *testing.T) {
		err := geoclient.Workspaces().Update(testdata.InvalidName, "")
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	t.Run("Invalid New Name", func(t *testing.T) {
		err := geoclient.Workspaces().Update(testdata.Workspace, testdata.InvalidName)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	//revert changes made in the test
	err := geoclient.Workspaces().Delete(toRemove, false)
	assert.NoError(t, err)
}

func TestWorkspaceIntegration_Delete(t *testing.T) {
	addTestWorkspace(t)

	t.Run("200 Ok", func(t *testing.T) {
		err := geoclient.Workspaces().Delete(testdata.Workspace, false)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspaces().Delete(testdata.Workspace, false)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestWorkspaceIntegration_GetAll(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		t.Run("EMPTY", func(t *testing.T) {
			wksp, err := geoclient.Workspaces().GetAll()
			assert.Empty(t, wksp)
			assert.NoError(t, err)
		})

		t.Run("SINGLE Workspace", func(t *testing.T) {
			//create required resource
			err := geoclient.Workspaces().Create(testdata.Workspace, false)
			assert.NoError(t, err)

			wksp, err := geoclient.Workspaces().GetAll()
			assert.NoError(t, err)
			assert.Len(t, wksp, 1)
			assert.Equal(t, testdata.Workspace, wksp[0].Name)
		})

		t.Run("MULTIPLE WORKSPACES", func(t *testing.T) {
			//create another temporary workspace
			err := geoclient.Workspaces().Create(testdata.Workspace+"_2", false)
			assert.NoError(t, err)

			wksp, err := geoclient.Workspaces().GetAll()
			assert.NoError(t, err)
			assert.Len(t, wksp, 2)
		})
	})

	err := geoclient.Workspaces().Delete(testdata.Workspace+"_2", false)
	assert.NoError(t, err)
}
