package main

import (
	"github.com/canghel3/go-geoserver/datastores/postgis"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeatureTypeIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GEOSERVER_URL, testdata.GEOSERVER_USERNAME, testdata.GEOSERVER_PASSWORD, testdata.GEOSERVER_DATADIR)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.WORKSPACE, true)

	//create datastore
	err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Create().PostGIS(testdata.DATASTORE_POSTGIS, postgis.ConnectionParams{
		Host:     testdata.POSTGIS_HOST,
		Database: testdata.POSTGIS_DB,
		User:     testdata.POSTGIS_USERNAME,
		Password: testdata.POSTGIS_PASSWORD,
		Port:     testdata.POSTGIS_PORT,
		SSL:      testdata.POSTGIS_SSL,
	})
	assert.NoError(t, err)
}
