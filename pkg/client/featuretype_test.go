package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeatureTypeIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	geoserverClient.Workspaces().Delete(testdata.Workspace, true)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("POSTGIS", func(t *testing.T) {
			//create datastore
			err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
				Host:     testdata.PostgisHost,
				Database: testdata.PostgisDb,
				User:     testdata.PostgisUsername,
				Password: testdata.PostgisPassword,
				Port:     testdata.PostgisPort,
				SSL:      testdata.PostgisSsl,
			})
			assert.NoError(t, err)

			t.Run("WITHOUT ANY OPTIONS", func(t *testing.T) {
				feature := featuretypes.New(testdata.FeatureTypePostgis, testdata.FeatureTypePostgisNativeName)

				err = geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Publish(feature)
				assert.NoError(t, err)

				get, err := geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis)
				assert.NoError(t, err)
				assert.Equal(t, get.Name, testdata.FeatureTypePostgis)
				assert.Equal(t, get.NativeName, testdata.FeatureTypePostgisNativeName)
				assert.Equal(t, get.Srs, "EPSG:4326")
				assert.Equal(t, get.Keywords.Keywords, []string{"features", "init"})
			})

			t.Run("WITH BBOX OPTION", func(t *testing.T) {
				var featureName = testdata.FeatureTypePostgis + "_WITH_BBOX"
				var bbox = [4]float64{-180.0, -90.0, 180.0, 90.0}
				var bboxSrs = "EPSG:4326"

				feature := featuretypes.New(featureName, testdata.FeatureTypePostgisNativeName, options.FeatureType.BBOX(bbox, bboxSrs))

				err = geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Publish(feature)
				assert.NoError(t, err)

				get, err := geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(featureName)
				assert.NoError(t, err)
				assert.Equal(t, get.Name, featureName)
				assert.Equal(t, get.NativeName, testdata.FeatureTypePostgisNativeName)
				assert.Equal(t, get.Srs, "EPSG:4326")
				assert.Equal(t, get.NativeBoundingBox.MinX, bbox[0])
				assert.Equal(t, get.NativeBoundingBox.MinY, bbox[1])
				assert.Equal(t, get.NativeBoundingBox.MaxX, bbox[2])
				assert.Equal(t, get.NativeBoundingBox.MaxY, bbox[3])
				assert.Equal(t, get.NativeBoundingBox.CRS, bboxSrs)
			})
		})
	})

	err := geoserverClient.Workspaces().Delete(testdata.Workspace, true)
	assert.NoError(t, err)
}

func TestFeatureTypeIntegration_Get(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)

	//create datastore
	err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
		Host:     testdata.PostgisHost,
		Database: testdata.PostgisDb,
		User:     testdata.PostgisUsername,
		Password: testdata.PostgisPassword,
		Port:     testdata.PostgisPort,
		SSL:      testdata.PostgisSsl,
	})
	assert.NoError(t, err)

	//create feature type
	feature := featuretypes.New(testdata.FeatureTypePostgis, testdata.FeatureTypePostgisNativeName)
	err = geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Publish(feature)
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		get, err := geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis)
		assert.NoError(t, err)
		assert.NotNil(t, get)
		assert.Equal(t, get.Name, testdata.FeatureTypePostgis)
		assert.Equal(t, get.NativeName, testdata.FeatureTypePostgisNativeName)
		assert.Equal(t, get.Srs, "EPSG:4326")
	})

	t.Run("404 Not Found", func(t *testing.T) {
		get, err := geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis + "_DOES_NOT_EXIST")
		assert.Error(t, err)
		assert.Nil(t, get)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	err = geoserverClient.Workspaces().Delete(testdata.Workspace, true)
	assert.NoError(t, err)
}

func TestFeatureTypeIntegration_Delete(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)

	//create datastore
	err := geoserverClient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
		Host:     testdata.PostgisHost,
		Database: testdata.PostgisDb,
		User:     testdata.PostgisUsername,
		Password: testdata.PostgisPassword,
		Port:     testdata.PostgisPort,
		SSL:      testdata.PostgisSsl,
	})
	assert.NoError(t, err)

	//create feature type
	feature := featuretypes.New(testdata.FeatureTypePostgis, testdata.FeatureTypePostgisNativeName)
	err = geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Publish(feature)
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		err = geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Delete(testdata.FeatureTypePostgis, true)
		assert.NoError(t, err)

		//try to retrieve the feature type
		get, err := geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis)
		assert.Nil(t, get)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err = geoserverClient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Delete(testdata.FeatureTypePostgis, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	err = geoserverClient.Workspaces().Delete(testdata.Workspace, true)
	assert.NoError(t, err)
}
