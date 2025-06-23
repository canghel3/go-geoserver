package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataStoreIntegration_Create(t *testing.T) {
	addTestWorkspace(t)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Generic Options", func(t *testing.T) {
			t.Run("Description", func(t *testing.T) {
				var description = "generic description"
				var name = testdata.DatastorePostgis + "WITH_DESCRIPTION"
				err := geoclient.Workspace(testdata.Workspace).DataStores().Create(options.DataStore.Description(description)).PostGIS(name, postgis.ConnectionParams{
					Host:     testdata.PostgisHost,
					Database: testdata.PostgisDb,
					User:     testdata.PostgisUsername,
					Password: testdata.PostgisPassword,
					Port:     testdata.PostgisPort,
					SSL:      testdata.PostgisSsl,
				})
				assert.NoError(t, err)

				ds, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(name)
				assert.NoError(t, err)
				assert.NotNil(t, ds)
				assert.Equal(t, ds.Description, description)
				assert.Equal(t, ds.DisableConnectionOnFailure, false)
			})

			t.Run("DISABLE CONNECTION ON FAILURE", func(t *testing.T) {
				var name = testdata.DatastorePostgis + "WITH_DISABLE_CONNECTION_ON_FAILURE"
				err := geoclient.Workspace(testdata.Workspace).DataStores().Create(options.DataStore.DisableConnectionOnFailure(true)).PostGIS(name, postgis.ConnectionParams{
					Host:     testdata.PostgisHost,
					Database: testdata.PostgisDb,
					User:     testdata.PostgisUsername,
					Password: testdata.PostgisPassword,
					Port:     testdata.PostgisPort,
					SSL:      testdata.PostgisSsl,
				})
				assert.NoError(t, err)

				ds, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(name)
				assert.NoError(t, err)
				assert.NotNil(t, ds)
				assert.Equal(t, ds.Description, "")
				assert.Equal(t, ds.DisableConnectionOnFailure, true)
			})
		})

		t.Run("PostGIS", func(t *testing.T) {
			addTestDataStore(t, types.PostGIS)

			store, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("WITH OPTIONS", func(t *testing.T) {
				t.Run("VALIDATE CONNECTIONS", func(t *testing.T) {
					var suffix = "_WITH_VALIDATE_CONNECTIONS"
					err := geoclient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis+suffix, postgis.ConnectionParams{
						Host:     testdata.PostgisHost,
						Database: testdata.PostgisDb,
						User:     testdata.PostgisUsername,
						Password: testdata.PostgisPassword,
						Port:     testdata.PostgisPort,
						SSL:      testdata.PostgisSsl,
					}, options.PostGIS.ValidateConnections())
					assert.NoError(t, err)

					store, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis + suffix)
					assert.NoError(t, err)
					assert.NotNil(t, store)

					v, ok := store.ConnectionParameters.Get("validate connections")
					assert.True(t, ok)
					assert.Equal(t, "true", v)
				})
			})
		})

		t.Run("Shapefile", func(t *testing.T) {
			addTestDataStore(t, types.Shapefile)

			store, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastoreShapefile)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("Directory Of Shapefiles", func(t *testing.T) {
			addTestDataStore(t, types.DirOfShapefiles)

			store, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastoreShapefile)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("GeoPackage", func(t *testing.T) {
			addTestDataStore(t, types.GeoPackage)

			store, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastoreShapefile)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("Csv", func(t *testing.T) {
			t.Run("Lat Lon", func(t *testing.T) {

			})

			t.Run("Wkt", func(t *testing.T) {

			})
		})

		t.Run("WebFeatureService", func(t *testing.T) {
			addTestDataStore(t, types.WebFeatureService)

			store, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastoreWebFeatureService)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})
	})

	//NOTE: there is no 409 test because geoserver responds with 500 for a conflict error (:

	t.Run("500 Internal Server Error", func(t *testing.T) {
		geoclient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
			Host:     testdata.PostgisHost,
			Database: testdata.PostgisDb,
			User:     testdata.PostgisUsername,
			Password: testdata.PostgisPassword,
			Port:     testdata.PostgisPort,
			SSL:      testdata.PostgisSsl,
		})
		err := geoclient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
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
}

func TestDataStoreIntegration_Get(t *testing.T) {
	addTestWorkspace(t)

	addTestDataStore(t, types.PostGIS)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("POSTGIS", func(t *testing.T) {
			ds, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis)
			assert.NoError(t, err)
			assert.NotNil(t, ds)
			assert.Equal(t, ds.Name, testdata.DatastorePostgis)
			assert.Equal(t, ds.Workspace.Name, testdata.Workspace)
		})
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ds, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis + "_DOES_NOT_EXIST")
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
	addTestWorkspace(t)

	addTestDataStore(t, types.PostGIS)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("POSTGIS", func(t *testing.T) {
			err := geoclient.Workspace(testdata.Workspace).DataStores().Delete(testdata.DatastorePostgis, true)
			assert.NoError(t, err)

			//try to retrieve the workspace
			ds, err := geoclient.Workspace(testdata.Workspace).DataStores().Get(testdata.DatastorePostgis)
			assert.Nil(t, ds)
			assert.Error(t, err)
			assert.IsType(t, err, &customerrors.NotFoundError{})
		})
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).DataStores().Delete(testdata.DatastorePostgis, true)
		assert.Error(t, err)
		assert.IsType(t, err, &customerrors.NotFoundError{})
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DatastorePostgis))
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		//TODO: try to delete store that contains a feature inside with recurse set to false
	})
}

func TestDataStoreIntegration_Update(t *testing.T) {
	//TODO: implement
}

func TestDataStoreIntegration_Reset(t *testing.T) {
	//TODO: implement
}
