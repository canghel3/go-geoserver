//go:build integration

package main

import (
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkspaceIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword, testdata.GeoserverDatadir)

	geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)

	t.Run("201 CREATED", func(t *testing.T) {
		err := geoserverClient.Workspaces().Create(testdata.WORKSPACE, false)
		assert.NoError(t, err)
	})

	t.Run("409 CONFLICT", func(t *testing.T) {
		err := geoserverClient.Workspaces().Create(testdata.WORKSPACE, false)
		assert.IsType(t, &customerrors.ConflictError{}, err)
		assert.EqualError(t, err, "workspace already exists")
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(testdata.WORKSPACE, false)
	assert.NoError(t, err)
}

func TestWorkspaceIntegration_Get(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword, testdata.GeoserverDatadir)

	//create required resources
	geoserverClient.Workspaces().Create(testdata.WORKSPACE, false)

	t.Run("200 OK", func(t *testing.T) {
		wksp, err := geoserverClient.Workspaces().Get(testdata.WORKSPACE)
		assert.NoError(t, err)
		assert.Equal(t, testdata.WORKSPACE, wksp.Workspace.Name)
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		var suffix = "_NOT_FOUND"
		wksp, err := geoserverClient.Workspaces().Get(testdata.WORKSPACE + suffix)
		assert.Nil(t, wksp)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.WORKSPACE+suffix))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(testdata.WORKSPACE, false)
	assert.NoError(t, err)
}

func TestWorkspaceIntegration_Update(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword, testdata.GeoserverDatadir)

	//create required resources
	geoserverClient.Workspaces().Create(testdata.WORKSPACE, false)
	var toRemove string

	t.Run("200 OK", func(t *testing.T) {
		var suffix = "_UPDATED"
		toRemove = testdata.WORKSPACE + suffix
		err := geoserverClient.Workspaces().Update(testdata.WORKSPACE, testdata.WORKSPACE+suffix)
		assert.NoError(t, err)
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		var suffix = "_NOT_FOUND"
		err := geoserverClient.Workspaces().Update(testdata.WORKSPACE, testdata.WORKSPACE+suffix)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.WORKSPACE))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(toRemove, false)
	assert.NoError(t, err)
}

func TestWorkspaceIntegration_Delete(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword, testdata.GeoserverDatadir)

	//create required resources
	geoserverClient.Workspaces().Create(testdata.WORKSPACE, false)

	t.Run("200 OK", func(t *testing.T) {
		err := geoserverClient.Workspaces().Delete(testdata.WORKSPACE, false)
		assert.NoError(t, err)
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		err := geoserverClient.Workspaces().Delete(testdata.WORKSPACE, false)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.WORKSPACE))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestWorkspaceIntegration_GetAll(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword, testdata.GeoserverDatadir)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("EMPTY", func(t *testing.T) {
			wksp, err := geoserverClient.Workspaces().GetAll()
			assert.Nil(t, wksp.Workspaces.Workspace)
			assert.NoError(t, err)
		})

		t.Run("SINGLE WORKSPACE", func(t *testing.T) {
			//create required resource
			err := geoserverClient.Workspaces().Create(testdata.WORKSPACE, false)
			assert.NoError(t, err)

			wksp, err := geoserverClient.Workspaces().GetAll()
			assert.NoError(t, err)
			assert.Len(t, wksp.Workspaces.Workspace, 1)
			assert.Equal(t, testdata.WORKSPACE, wksp.Workspaces.Workspace[0].Name)
		})

		t.Run("MULTIPLE WORKSPACES", func(t *testing.T) {
			//create another temporary workspace
			err := geoserverClient.Workspaces().Create(testdata.WORKSPACE+"_2", false)
			assert.NoError(t, err)

			wksp, err := geoserverClient.Workspaces().GetAll()
			assert.NoError(t, err)
			assert.Len(t, wksp.Workspaces.Workspace, 2)
			assert.Equal(t, testdata.WORKSPACE+"_2", wksp.Workspaces.Workspace[0].Name)
			assert.Equal(t, testdata.WORKSPACE, wksp.Workspaces.Workspace[1].Name)
		})
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(testdata.WORKSPACE, false)
	assert.NoError(t, err)

	err = geoserverClient.Workspaces().Delete(testdata.WORKSPACE+"_2", false)
	assert.NoError(t, err)
}
