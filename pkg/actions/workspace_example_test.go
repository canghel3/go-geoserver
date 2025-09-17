package actions

import (
	"fmt"

	"github.com/canghel3/go-geoserver/pkg/wms"
)

func ExampleWorkspaces_Create() {
	//the workspaces actions are generated from the client with geoclient.Workspaces()

	workspaces := Workspaces{}
	err := workspaces.Create("example-workspace", true)
	if err != nil {
		fmt.Println("Error creating workspace:", err)
		return
	}

	fmt.Println("Workspace created successfully")
}

func ExampleWorkspaces_Get() {
	//the workspaces actions are generated from the client with geoclient.Workspaces()

	workspaces := Workspaces{}

	// Get a specific workspace
	workspace, err := workspaces.Get("example-workspace")
	if err != nil {
		fmt.Println("Error retrieving workspace:", err)
		return
	}

	fmt.Println("Workspace name:", workspace.Name)
	// Output: Workspace name: example-workspace
}

func ExampleWorkspaces_GetAll() {
	//the workspaces actions are generated from the client with geoclient.Workspaces()

	workspaces := Workspaces{}

	// Get all workspaces
	spaces, err := workspaces.GetAll()
	if err != nil {
		fmt.Println("Error retrieving workspaces:", err)
		return
	}

	// Print the names of all workspaces
	fmt.Println("Number of workspaces:", len(spaces))
	for _, workspace := range spaces {
		fmt.Println("- Workspace:", workspace.Name)
	}
	// Output depends on the number of workspaces in your GeoServer instance
}

func ExampleWorkspaces_Update() {
	//the workspaces actions are generated from the client with geoclient.Workspaces()

	workspaces := Workspaces{}

	// Update a workspace (rename it)
	err := workspaces.Update("example-workspace", "new-workspace-name")
	if err != nil {
		fmt.Println("Error updating workspace:", err)
		return
	}

	fmt.Println("Workspace updated successfully")
	// Output: Workspace updated successfully
}

func ExampleWorkspaces_Delete() {
	//the workspaces actions are generated from the client with geoclient.Workspaces()

	workspaces := Workspaces{}

	// Delete a workspace
	err := workspaces.Delete("new-workspace-name", true)
	if err != nil {
		fmt.Println("Error deleting workspace:", err)
		return
	}

	fmt.Println("Workspace deleted successfully")
	// Output: Workspace deleted successfully
}

func ExampleWorkspaces_Use() {
	//the workspaces actions are generated from the client with geoclient.Workspaces()

	workspaces := Workspaces{}

	//Set which workspace to use
	workspaceActions := workspaces.Use("my-workspace")

	//workspaceActions contains available actions within a workspace
	workspaceActions.WMS(wms.Version130)
	workspaceActions.DataStores()
	workspaceActions.CoverageStores()
	workspaceActions.LayerGroups()
	workspaceActions.GeoWebCache()
}
