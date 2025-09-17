package actions

import (
	"fmt"

	"github.com/canghel3/go-geoserver/pkg/datastores"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/options"
)

func ExampleDataStores_Create() {
	//the workspace actions are generated from the client with geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	//without data store options
	dataStoreList := workspace.DataStores().Create()

	//with description option
	dataStoreList = workspace.DataStores().Create(options.GenericStore.Description("example description"))

	//with disable conn option
	dataStoreList = workspace.DataStores().Create(options.GenericStore.AutoDisableOnConnFailure())

	_ = dataStoreList
}

func ExampleDataStores_Get() {
	//the workspace actions are generated from the client with geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	//without data store options
	store, err := workspace.DataStores().Get("store")
	if err != nil {
		fmt.Println("Error getting store:", err)
		return
	}

	fmt.Println("DataStore name:", store.Name)
}

func ExampleDataStores_Update() {
	//the workspace actions are generated from the client with geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	err := workspace.DataStores().Update("store-to-update", datastores.DataStore{})
	if err != nil {
		fmt.Println("Error updating store:", err)
		return
	}

	fmt.Println("DataStore updated")
}

func ExampleDataStores_Delete() {
	//the workspace actions are generated from the client with geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	err := workspace.DataStores().Delete("store-name", true)
	if err != nil {
		fmt.Println("Error deleting store:", err)
	}

	fmt.Println("DataStore deleted")
}

func ExampleDataStores_GetAll() {
	//the workspace actions are generated from the client with geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	dataStores, err := workspace.DataStores().GetAll()
	if err != nil {
		fmt.Println("Error getting all datastores:", err)
	}

	fmt.Println("All data stores:", dataStores)
}

func ExampleDataStores_Reset() {
	//the workspace actions are generated from the client with geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	err := workspace.DataStores().Reset("store-name")
	if err != nil {
		fmt.Println("Error resetting store:", err)
		return
	}

	fmt.Println("DataStore resetted")
}

func ExampleDataStores_Use() {
	//the workspace actions are generated from the client with geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	featureActions := workspace.DataStores().Use("store-name")
	featureActions.GetAll()
	featureActions.Reset("feature")
}

func ExampleDataStoreList_PostGIS() {
	workspace := Workspace{}
	err := workspace.DataStores().Create().PostGIS("postgis-store", postgis.ConnectionParams{
		Host:     "localhost",
		Database: "db",
		User:     "user",
		Password: "secret",
		Port:     "5432",
		SSL:      "disable",
	})
	if err != nil {
		fmt.Println("Error creating PostGIS datastore:", err)
		return
	}
}
