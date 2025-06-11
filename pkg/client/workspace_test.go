package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ExampleWorkspaceCreate demonstrates how to create a workspace
func ExampleGeoserverClient_Workspaces_create() {
	// Initialize the client
	geoclient := NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	// Create a workspace
	// The second parameter (true) indicates that the workspace should be set as default
	err := geoclient.Workspaces().Create("example-workspace", true)
	if err != nil {
		fmt.Println("Error creating workspace:", err)
		return
	}

	fmt.Println("Workspace created successfully")
	// Output: Workspace created successfully
}

func ExampleGeoserverClient_Workspaces_get() {
	// Initialize the client
	geoserverClient := NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	// Get a specific workspace
	workspace, err := geoserverClient.Workspaces().Get("example-workspace")
	if err != nil {
		fmt.Println("Error retrieving workspace:", err)
		return
	}

	fmt.Println("Workspace name:", workspace.Name)
	// Output: Workspace name: example-workspace
}

func ExampleGeoserverClient_Workspaces_update() {
	// Initialize the client
	geoserverClient := NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	// Update a workspace (rename it)
	err := geoserverClient.Workspaces().Update("example-workspace", "new-workspace-name")
	if err != nil {
		fmt.Println("Error updating workspace:", err)
		return
	}

	fmt.Println("Workspace updated successfully")
	// Output: Workspace updated successfully
}

// ExampleWorkspaceDelete demonstrates how to delete a workspace
func ExampleGeoserverClient_Workspaces_delete() {
	// Initialize the client
	geoserverClient := NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	// Delete a workspace
	// The second parameter (true) indicates that all resources in the workspace should be deleted as well
	err := geoserverClient.Workspaces().Delete("new-workspace-name", true)
	if err != nil {
		fmt.Println("Error deleting workspace:", err)
		return
	}

	fmt.Println("Workspace deleted successfully")
	// Output: Workspace deleted successfully
}

func ExampleGeoserverClient_Workspace_use() {
	// Initialize the client
	geoserverClient := NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	//Set which workspace to use
	workspaceActions := geoserverClient.Workspaces().Use("my-workspace")

	//workspaceActions contains available actions within a workspace
	workspaceActions.DataStores().Create()
}

func ExampleGeoserverClient_Workspaces_getAll() {
	// Initialize the client
	geoserverClient := NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	// Get all workspaces
	workspaces, err := geoserverClient.Workspaces().GetAll()
	if err != nil {
		fmt.Println("Error retrieving workspaces:", err)
		return
	}

	// Print the names of all workspaces
	fmt.Println("Number of workspaces:", len(workspaces))
	for _, workspace := range workspaces {
		fmt.Println("- Workspace:", workspace.Name)
	}
	// Output depends on the number of workspaces in your GeoServer instance
}

func TestWorkspaceIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	geoserverClient.Workspaces().Delete(testdata.Workspace, true)

	t.Run("201 CREATED", func(t *testing.T) {
		err := geoserverClient.Workspaces().Create(testdata.Workspace, false)
		assert.NoError(t, err)
	})

	t.Run("409 CONFLICT", func(t *testing.T) {
		err := geoserverClient.Workspaces().Create(testdata.Workspace, false)
		assert.IsType(t, &customerrors.ConflictError{}, err)
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
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace+suffix))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
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
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
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
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace))
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestWorkspaceIntegration_GetAll(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("EMPTY", func(t *testing.T) {
			wksp, err := geoserverClient.Workspaces().GetAll()
			assert.Empty(t, wksp)
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
