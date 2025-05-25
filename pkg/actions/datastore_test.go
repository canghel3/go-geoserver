package actions

import (
	"github.com/canghel3/go-geoserver/pkg/client"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/options"
)

func ExampleDataStores_Create() {
	// initialize the client
	geoclient := client.NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	//without data store options
	dataStoreList := geoclient.Workspace("example-workspace").DataStores().Create()

	//with description option
	dataStoreList = geoclient.Workspace("example-workspace").DataStores().Create(options.DataStore.Description("example description"))

	//with disable conn option
	dataStoreList = geoclient.Workspace("example-workspace").DataStores().Create(options.DataStore.DisableConnectionOnFailure(true))

	_ = dataStoreList
}

func ExampleDataStores_Get() {

}

func ExampleDataStoreList_PostGIS() {
	// initialize the client
	geoclient := client.NewGeoserverClient(
		"http://localhost:1111",
		"admin",
		"geoserver",
	)

	geoclient.Workspace("example-workspace").DataStores().Create().PostGIS("postgis-store", postgis.ConnectionParams{
		Host:     "localhost",
		Database: "db",
		User:     "user",
		Password: "secret",
		Port:     "5432",
		SSL:      "disable",
	})
}
