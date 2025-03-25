package main

import (
	"github.com/canghel3/go-geoserver/datastores/postgis"
	"github.com/canghel3/go-geoserver/featuretypes"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeatureTypeIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword, testdata.GeoserverDatadir)

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

	t.Run("200 OK", func(t *testing.T) {
		feature := featuretypes.CreateFeatureType{
			Name:       testdata.FEATURE_TYPE_NAME,
			NativeName: testdata.FEATURE_TYPE_NATIVE_NAME,
			Title:      testdata.FEATURE_TYPE_TITLE,
		}

		err = geoserverClient.Workspace(testdata.WORKSPACE).DataStore(testdata.DATASTORE_POSTGIS).PublishFeature(feature)
		assert.NoError(t, err)
	})

	err = geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)
	assert.NoError(t, err)
}
