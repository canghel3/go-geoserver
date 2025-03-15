package main

import (
	"github.com/canghel3/go-geoserver/pkg/client"
	"github.com/canghel3/go-geoserver/pkg/datastore/postgis"
)

func main() {
	geoclient := client.New()
	sampleWksp := geoclient.Workspaces().Use("sample")
	geoclient.Workspace("").DataStores()
	//sampleWksp.DataStores().Get()
	err := sampleWksp.DataStores().Create().PostGIS("name", postgis.ConnectionParams{
		Host:     "",
		Database: "",
		User:     "",
		Password: "",
		Port:     "",
		SSL:      "",
	})
	if err != nil {
		return
	}
}
