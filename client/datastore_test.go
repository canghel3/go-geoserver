//go:build integration

package client

import (
	"github.com/canghel3/go-geoserver/datastores/postgis"
	"github.com/canghel3/go-geoserver/options"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataStoreClient_Create(t *testing.T) {
	geoserverClient := New(testdata.GEOSERVER_URL, testdata.GEOSERVER_USERNAME, testdata.GEOSERVER_PASSWORD, testdata.GEOSERVER_DATADIR)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.WORKSPACE, true)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("POSTGIS", func(t *testing.T) {
			err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Create().PostGIS(testdata.DATASTORE_POSTGIS, postgis.ConnectionParams{
				Host:     testdata.POSTGIS_HOST,
				Database: testdata.POSTGIS_DB,
				User:     testdata.POSTGIS_USERNAME,
				Password: testdata.POSTGIS_PASSWORD,
				Port:     testdata.POSTGIS_PORT,
				SSL:      testdata.POSTGIS_SSL,
			})
			assert.NoError(t, err)

			t.Run("WITH OPTIONS", func(t *testing.T) {
				t.Run("VALIDATE CONNECTIONS", func(t *testing.T) {
					var suffix = "_WITH_VALIDATE_CONNECTIONS"
					err = geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Create().PostGIS(testdata.DATASTORE_POSTGIS+suffix, postgis.ConnectionParams{
						Host:     testdata.POSTGIS_HOST,
						Database: testdata.POSTGIS_DB,
						User:     testdata.POSTGIS_USERNAME,
						Password: testdata.POSTGIS_PASSWORD,
						Port:     testdata.POSTGIS_PORT,
						SSL:      testdata.POSTGIS_SSL,
					}, options.PostGIS.ValidateConnections())
					assert.NoError(t, err)

					store, err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Get(testdata.DATASTORE_POSTGIS + suffix)
					assert.NoError(t, err)
					assert.NotNil(t, store)

					v, ok := store.ConnectionParameters.Get("validate connections")
					assert.True(t, ok)
					assert.Equal(t, "true", v)
				})
			})
			//geoserverClient.Workspace(testdata.WORKSPACE).DataStore(testdata.DATASTORE_POSTGIS).PublishFeature()
		})
	})

	err := geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)
	assert.NoError(t, err)
}
