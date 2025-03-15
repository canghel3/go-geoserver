package main

import (
	"github.com/canghel3/go-geoserver/internal/misc"
	"github.com/canghel3/go-geoserver/pkg/client"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
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

	sampleWksp.DataStore("name").PublishFeature(featuretypes.FeatureType{
		Name:              "",
		NativeName:        "",
		Namespace:         featuretypes.Namespace{},
		Srs:               "",
		NativeBoundingBox: misc.BoundingBox{},
		ProjectionPolicy:  "",
		Keywords:          nil,
		Title:             "",
		Store:             featuretypes.Store{},
	})

}
