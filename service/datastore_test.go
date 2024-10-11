package service

import (
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/datastore/postgis"
	"github.com/canghel3/go-geoserver/utils"
	"gotest.tools/v3/assert"
	"testing"
)

func TestDataStore(t *testing.T) {
	geoserverService := NewGeoserverService(target, username, password)

	t.Run("PostGIS", func(t *testing.T) {
		t.Run("CREATE", func(t *testing.T) {
			t.Run("SIMPLE", func(t *testing.T) {
				connectionParams := postgis.ConnectionParams{
					Host:     host,
					Database: databaseName,
					User:     databaseUser,
					Password: databasePassword,
					Port:     databasePort,
					SSL:      "disable",
				}

				assert.NilError(t, geoserverService.CreateWorkspace("init"))
				assert.NilError(t, geoserverService.CreatePostGISDataStore("init", "init", connectionParams))
			})

			t.Run("DUPLICATE", func(t *testing.T) {
				connectionParams := postgis.ConnectionParams{
					Host:     host,
					Database: databaseName,
					User:     databaseUser,
					Password: databasePassword,
					Port:     databasePort,
					SSL:      "disable",
				}

				assert.Error(t, geoserverService.CreatePostGISDataStore("init", "init", connectionParams), "datastore init already exists")
			})

			t.Run("INVALID PARAMETERS", func(t *testing.T) {
				connectionParams := postgis.ConnectionParams{
					Host:     "",
					Database: "",
					User:     "",
					Password: databasePassword,
					Port:     databasePort,
					SSL:      "disable",
				}

				err := geoserverService.CreatePostGISDataStore("init", "fail", connectionParams)
				assert.ErrorType(t, err, &customerrors.InputError{})
				assert.Error(t, err, "host cannot be empty")
			})
		})

		t.Run("GET", func(t *testing.T) {
			t.Run("SIMPLE", func(t *testing.T) {
				ds, err := geoserverService.GetDataStore("init", "init")
				assert.NilError(t, err)

				assert.Equal(t, ds.Name, "init")
				assert.Equal(t, ds.Workspace.Name, "init")
				assert.Equal(t, ds.Default, false)
			})

			t.Run("NON-EXISTENT", func(t *testing.T) {
				_, err := geoserverService.GetDataStore("init", "none")
				assert.ErrorType(t, err, &customerrors.NotFoundError{})
				assert.Error(t, err, "datastore none does not exist")
			})
		})

		t.Run("DELETE", func(t *testing.T) {
			t.Run("WITHOUT RECURSE", func(t *testing.T) {
				assert.NilError(t, geoserverService.DeleteDataStore("init", "init", utils.RecurseOption(false)))
			})

			t.Run("WITH RECURSE", func(t *testing.T) {
				connectionParams := postgis.ConnectionParams{
					Host:     host,
					Database: databaseName,
					User:     databaseUser,
					Password: databasePassword,
					Port:     databasePort,
					SSL:      "disable",
				}

				assert.NilError(t, geoserverService.CreatePostGISDataStore("init", "init", connectionParams))
				assert.NilError(t, geoserverService.DeleteDataStore("init", "init", utils.RecurseOption(false)))
			})

			t.Run("NON-EXISTENT", func(t *testing.T) {
				err := geoserverService.DeleteDataStore("init", "none")
				assert.ErrorType(t, err, &customerrors.NotFoundError{})
				assert.Error(t, err, "datastore none does not exist")
			})
		})
	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", utils.RecurseOption(true)))
}
