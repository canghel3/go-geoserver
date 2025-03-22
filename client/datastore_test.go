//go:build integration

package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
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
		t.Run("GENERIC OPTIONS", func(t *testing.T) {
			t.Run("DESCRIPTION", func(t *testing.T) {
				var description = "generic description"
				var name = testdata.DATASTORE_POSTGIS + "WITH_DESCRIPTION"
				err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Create(options.Datastore.Description(description)).PostGIS(name, postgis.ConnectionParams{
					Host:     testdata.POSTGIS_HOST,
					Database: testdata.POSTGIS_DB,
					User:     testdata.POSTGIS_USERNAME,
					Password: testdata.POSTGIS_PASSWORD,
					Port:     testdata.POSTGIS_PORT,
					SSL:      testdata.POSTGIS_SSL,
				})
				assert.NoError(t, err)

				ds, err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Get(name)
				assert.NoError(t, err)
				assert.NotNil(t, ds)
				assert.Equal(t, ds.Description, description)
			})

			t.Run("DISABLE CONNECTION ON FAILURE", func(t *testing.T) {

			})
		})

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
				//TODO: test with description and disable on connection failure options

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
		})
	})

	//NOTE: there is no 409 test because geoserver responds with 500 for a conflict error (:

	t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
		err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Create().PostGIS(testdata.DATASTORE_POSTGIS, postgis.ConnectionParams{
			Host:     testdata.POSTGIS_HOST,
			Database: testdata.POSTGIS_DB,
			User:     testdata.POSTGIS_USERNAME,
			Password: testdata.POSTGIS_PASSWORD,
			Port:     testdata.POSTGIS_PORT,
			SSL:      testdata.POSTGIS_SSL,
		})
		assert.Error(t, err)
		assert.IsType(t, err, &customerrors.GeoserverError{})
		//yes, geoserver actually responds with 500 for a conflict error
		assert.ErrorContains(t, err, fmt.Sprintf(`Store '%s' already exists in workspace '%s'`, testdata.DATASTORE_POSTGIS, testdata.WORKSPACE))
	})

	err := geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)
	assert.NoError(t, err)
}

func TestDataStoreClient_Get(t *testing.T) {
	geoserverClient := New(testdata.GEOSERVER_URL, testdata.GEOSERVER_USERNAME, testdata.GEOSERVER_PASSWORD, testdata.GEOSERVER_DATADIR)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.WORKSPACE, true)
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
		t.Run("POSTGIS", func(t *testing.T) {
			ds, err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Get(testdata.DATASTORE_POSTGIS)
			assert.NoError(t, err)
			assert.NotNil(t, ds)
			assert.Equal(t, ds.Name, testdata.DATASTORE_POSTGIS)
			assert.Equal(t, ds.Workspace.Name, testdata.WORKSPACE)
		})
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		ds, err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Get(testdata.DATASTORE_POSTGIS + "_DOES_NOT_EXIST")
		assert.Error(t, err)
		assert.Nil(t, ds)
		assert.IsType(t, err, &customerrors.NotFoundError{})
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DATASTORE_POSTGIS+"_DOES_NOT_EXIST"))
	})

	err = geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)
	assert.NoError(t, err)
}

func TestDataStoreClient_GetAll(t *testing.T) {
	//TODO: implement and test functionality
}

func TestDataStoreClient_Delete(t *testing.T) {
	geoserverClient := New(testdata.GEOSERVER_URL, testdata.GEOSERVER_USERNAME, testdata.GEOSERVER_PASSWORD, testdata.GEOSERVER_DATADIR)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.WORKSPACE, true)
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
		t.Run("POSTGIS", func(t *testing.T) {
			err = geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Delete(testdata.DATASTORE_POSTGIS, true)
			assert.NoError(t, err)

			//try to retrieve the workspace
			ds, err := geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Get(testdata.DATASTORE_POSTGIS)
			assert.Nil(t, ds)
			assert.Error(t, err)
			assert.IsType(t, err, &customerrors.NotFoundError{})
		})
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		err = geoserverClient.Workspace(testdata.WORKSPACE).DataStores().Delete(testdata.DATASTORE_POSTGIS, true)
		assert.Error(t, err)
		assert.IsType(t, err, &customerrors.NotFoundError{})
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DATASTORE_POSTGIS))
	})

	t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
		//TODO: try to delete store that contains a feature inside with recurse set to false
	})
}
