package actions

import (
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/options"
)

func ExampleWorkspace_DataStores() {
	//TODO: write example
}

func ExampleDataStores_Create() {
	//the workspace actions are generated from the client with e.g. geoclient.Workspace("my-workspace")
	workspace := Workspace{}

	//without data store options
	dataStoreList := workspace.DataStores().Create()

	//with description option
	dataStoreList = workspace.DataStores().Create(options.DataStore.Description("example description"))

	//with disable conn option
	dataStoreList = workspace.DataStores().Create(options.DataStore.DisableConnectionOnFailure(true))

	_ = dataStoreList
}

func ExampleDataStores_Get() {

}

func ExampleDataStoreList_PostGIS() {
	workspace := Workspace{}
	workspace.DataStores().Create().PostGIS("postgis-store", postgis.ConnectionParams{
		Host:     "localhost",
		Database: "db",
		User:     "user",
		Password: "secret",
		Port:     "5432",
		SSL:      "disable",
	})
}
