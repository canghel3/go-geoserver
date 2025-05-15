//go:build integration

package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataStoreIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	geoserverClient.Workspaces().Delete(testdata.Workspace, true)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("GENERIC OPTIONS", func(t *testing.T) {
			t.Run("DESCRIPTION", func(t *testing.T) {
				var description = "generic description"
				var name = testdata.DatastorePostgis + "WITH_DESCRIPTION"
				err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create(options.Datastore.Description(description)).PostGIS(name, postgis.ConnectionParams{
					Host:     testdata.PostgisHost,
					Database: testdata.PostgisDb,
					User:     testdata.PostgisUsername,
					Password: testdata.PostgisPassword,
					Port:     testdata.PostgisPort,
					SSL:      testdata.PostgisSsl,
				})
				assert.NoError(t, err)

				ds, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(name)
				assert.NoError(t, err)
				assert.NotNil(t, ds)
				assert.Equal(t, ds.Description, description)
				assert.Equal(t, ds.DisableConnectionOnFailure, false)
			})

			t.Run("DISABLE CONNECTION ON FAILURE", func(t *testing.T) {
				var name = testdata.DatastorePostgis + "WITH_DISABLE_CONNECTION_ON_FAILURE"
				err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create(options.Datastore.DisableConnectionOnFailure(true)).PostGIS(name, postgis.ConnectionParams{
					Host:     testdata.PostgisHost,
					Database: testdata.PostgisDb,
					User:     testdata.PostgisUsername,
					Password: testdata.PostgisPassword,
					Port:     testdata.PostgisPort,
					SSL:      testdata.PostgisSsl,
				})
				assert.NoError(t, err)

				ds, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(name)
				assert.NoError(t, err)
				assert.NotNil(t, ds)
				assert.Equal(t, ds.Description, "")
				assert.Equal(t, ds.DisableConnectionOnFailure, true)
			})
		})

		t.Run("POSTGIS", func(t *testing.T) {
			err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
				Host:     testdata.PostgisHost,
				Database: testdata.PostgisDb,
				User:     testdata.PostgisUsername,
				Password: testdata.PostgisPassword,
				Port:     testdata.PostgisPort,
				SSL:      testdata.PostgisSsl,
			})
			assert.NoError(t, err)

			store, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("WITH OPTIONS", func(t *testing.T) {
				t.Run("VALIDATE CONNECTIONS", func(t *testing.T) {
					var suffix = "_WITH_VALIDATE_CONNECTIONS"
					err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis+suffix, postgis.ConnectionParams{
						Host:     testdata.PostgisHost,
						Database: testdata.PostgisDb,
						User:     testdata.PostgisUsername,
						Password: testdata.PostgisPassword,
						Port:     testdata.PostgisPort,
						SSL:      testdata.PostgisSsl,
					}, options.PostGIS.ValidateConnections())
					assert.NoError(t, err)

					store, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis + suffix)
					assert.NoError(t, err)
					assert.NotNil(t, store)

					v, ok := store.ConnectionParameters.Get("validate connections")
					assert.True(t, ok)
					assert.Equal(t, "true", v)
				})
			})
		})

		t.Run("SHAPEFILE", func(t *testing.T) {
			err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().Shapefile(testdata.DatastoreShapefile, testdata.Shapefile)
			assert.NoError(t, err)

			store, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastoreShapefile)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("GEOPACKAGE", func(t *testing.T) {
			err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().GeoPackage(testdata.DatastoreGeoPackage, testdata.GeoPackage)
			assert.NoError(t, err)

			store, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastoreShapefile)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})
	})

	//NOTE: there is no 409 test because geoserver responds with 500 for a conflict error (:

	t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
		geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
			Host:     testdata.PostgisHost,
			Database: testdata.PostgisDb,
			User:     testdata.PostgisUsername,
			Password: testdata.PostgisPassword,
			Port:     testdata.PostgisPort,
			SSL:      testdata.PostgisSsl,
		})
		err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
			Host:     testdata.PostgisHost,
			Database: testdata.PostgisDb,
			User:     testdata.PostgisUsername,
			Password: testdata.PostgisPassword,
			Port:     testdata.PostgisPort,
			SSL:      testdata.PostgisSsl,
		})
		assert.Error(t, err)
		assert.IsType(t, err, &customerrors.GeoserverError{})
		//yes, geoserver actually responds with 500 for a conflict error
		assert.ErrorContains(t, err, fmt.Sprintf(`Store '%s' already exists in workspace '%s'`, testdata.DatastorePostgis, testdata.Workspace))
	})

	err := geoserverClient.Workspaces().Delete(testdata.Workspace, true)
	assert.NoError(t, err)
}

func TestDataStoreIntegration_Get(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)
	defer geoserverClient.Workspaces().Delete(testdata.Workspace, true)
	err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
		Host:     testdata.PostgisHost,
		Database: testdata.PostgisDb,
		User:     testdata.PostgisUsername,
		Password: testdata.PostgisPassword,
		Port:     testdata.PostgisPort,
		SSL:      testdata.PostgisSsl,
	})
	assert.NoError(t, err)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("POSTGIS", func(t *testing.T) {
			ds, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis)
			assert.NoError(t, err)
			assert.NotNil(t, ds)
			assert.Equal(t, ds.Name, testdata.DatastorePostgis)
			assert.Equal(t, ds.Workspace.Name, testdata.Workspace)
		})
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		ds, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis + "_DOES_NOT_EXIST")
		assert.Error(t, err)
		assert.Nil(t, ds)
		assert.IsType(t, err, &customerrors.NotFoundError{})
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DatastorePostgis+"_DOES_NOT_EXIST"))
	})
}

func TestDataStoreIntegration_GetAll(t *testing.T) {
	//TODO: implement and test functionality
}

func TestDataStoreIntegration_Delete(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)
	err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
		Host:     testdata.PostgisHost,
		Database: testdata.PostgisDb,
		User:     testdata.PostgisUsername,
		Password: testdata.PostgisPassword,
		Port:     testdata.PostgisPort,
		SSL:      testdata.PostgisSsl,
	})
	assert.NoError(t, err)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("POSTGIS", func(t *testing.T) {
			err = geoserverClient.Workspace(testdata.Workspace).DataStores().Delete(testdata.DatastorePostgis, true)
			assert.NoError(t, err)

			//try to retrieve the workspace
			ds, err := geoserverClient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis)
			assert.Nil(t, ds)
			assert.Error(t, err)
			assert.IsType(t, err, &customerrors.NotFoundError{})
		})
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		err = geoserverClient.Workspace(testdata.Workspace).DataStores().Delete(testdata.DatastorePostgis, true)
		assert.Error(t, err)
		assert.IsType(t, err, &customerrors.NotFoundError{})
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DatastorePostgis))
	})

	t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
		//TODO: try to delete store that contains a feature inside with recurse set to false
	})

	err = geoserverClient.Workspaces().Delete(testdata.Workspace, true)
	assert.NoError(t, err)
}

func TestDataStoreIntegration_Update(t *testing.T) {
	//TODO: implement
}

func TestDataStoreIntegration_Reset(t *testing.T) {
	//TODO: implement
}
