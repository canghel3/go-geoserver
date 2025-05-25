//go:build integration

package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	customerrors2 "github.com/canghel3/go-geoserver/pkg/models/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkspaceIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	geoserverClient.Workspaces().Delete(testdata.Workspace, true)

	t.Run("201 CREATED", func(t *testing.T) {
		err := geoserverClient.Workspaces().Create(testdata.Workspace, false)
		assert.NoError(t, err)
	})

	t.Run("409 CONFLICT", func(t *testing.T) {
		err := geoserverClient.Workspaces().Create(testdata.Workspace, false)
		assert.IsType(t, &customerrors2.ConflictError{}, err)
		assert.EqualError(t, err, "workspace already exists")
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(testdata.Workspace, false)
	assert.NoError(t, err)
}

func TestWorkspaceIntegration_Get(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	//create required resources
	geoserverClient.Workspaces().Create(testdata.Workspace, false)

	t.Run("200 OK", func(t *testing.T) {
		wksp, err := geoserverClient.Workspaces().Get(testdata.Workspace)
		assert.NoError(t, err)
		assert.Equal(t, testdata.Workspace, wksp.Name)
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		var suffix = "_NOT_FOUND"
		wksp, err := geoserverClient.Workspaces().Get(testdata.Workspace + suffix)
		assert.Nil(t, wksp)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.Workspace+suffix))
		assert.IsType(t, &customerrors2.NotFoundError{}, err)
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(testdata.Workspace, false)
	assert.NoError(t, err)
}

func TestWorkspaceIntegration_Update(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	//create required resources
	geoserverClient.Workspaces().Create(testdata.Workspace, false)
	var toRemove string

	t.Run("200 OK", func(t *testing.T) {
		var suffix = "_UPDATED"
		toRemove = testdata.Workspace + suffix
		err := geoserverClient.Workspaces().Update(testdata.Workspace, testdata.Workspace+suffix)
		assert.NoError(t, err)
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		var suffix = "_NOT_FOUND"
		err := geoserverClient.Workspaces().Update(testdata.Workspace, testdata.Workspace+suffix)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.Workspace))
		assert.IsType(t, &customerrors2.NotFoundError{}, err)
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(toRemove, false)
	assert.NoError(t, err)
}

func TestWorkspaceIntegration_Delete(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	//create required resources
	geoserverClient.Workspaces().Create(testdata.Workspace, false)

	t.Run("200 OK", func(t *testing.T) {
		err := geoserverClient.Workspaces().Delete(testdata.Workspace, false)
		assert.NoError(t, err)
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		err := geoserverClient.Workspaces().Delete(testdata.Workspace, false)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.Workspace))
		assert.IsType(t, &customerrors2.NotFoundError{}, err)
	})
}

func TestWorkspaceIntegration_GetAll(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("EMPTY", func(t *testing.T) {
			wksp, err := geoserverClient.Workspaces().GetAll()
			assert.Nil(t, wksp)
			assert.NoError(t, err)
		})

		t.Run("SINGLE Workspace", func(t *testing.T) {
			//create required resource
			err := geoserverClient.Workspaces().Create(testdata.Workspace, false)
			assert.NoError(t, err)

			wksp, err := geoserverClient.Workspaces().GetAll()
			assert.NoError(t, err)
			assert.Len(t, wksp, 1)
			assert.Equal(t, testdata.Workspace, wksp[0].Name)
		})

		t.Run("MULTIPLE WORKSPACES", func(t *testing.T) {
			//create another temporary workspace
			err := geoserverClient.Workspaces().Create(testdata.Workspace+"_2", false)
			assert.NoError(t, err)

			wksp, err := geoserverClient.Workspaces().GetAll()
			assert.NoError(t, err)
			assert.Len(t, wksp, 2)
			assert.Equal(t, testdata.Workspace+"_2", wksp[0].Name)
			assert.Equal(t, testdata.Workspace, wksp[1].Name)
		})
	})

	//revert changes made in the test
	err := geoserverClient.Workspaces().Delete(testdata.Workspace, false)
	assert.NoError(t, err)

	err = geoserverClient.Workspaces().Delete(testdata.Workspace+"_2", false)
	assert.NoError(t, err)
}
