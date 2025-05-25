package examples

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/client"
)

func Create() {
	geoclient := client.NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)
	err := geoclient.Workspaces().Create(testdata.Workspace, true)
	if err != nil {
		panic(err)
	}
}

func Retrieve() {
	geoclient := client.NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)
	wksp, err := geoclient.Workspaces().Get(testdata.Workspace)
	if err != nil {
		panic(err)
	}
	_ = wksp

	all, err := geoclient.Workspaces().GetAll()
	if err != nil {
		panic(err)
	}
	_ = all
}

func Update() {
	geoclient := client.NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)
	err := geoclient.Workspaces().Update(testdata.Workspace, "new_name")
	if err != nil {
		panic(err)
	}
}

func Delete() {
	geoclient := client.NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)
	err := geoclient.Workspaces().Delete(testdata.Workspace, true)
	if err != nil {
		panic(err)
	}
}
