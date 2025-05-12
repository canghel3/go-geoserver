//go:build integration

package client

import (
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeatureTypeIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

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
		t.Run("WITHOUT ANY OPTIONS", func(t *testing.T) {
			feature := featuretypes.New(testdata.FEATURE_TYPE_NAME, testdata.FEATURE_TYPE_NATIVE_NAME)

			err = geoserverClient.Workspace(testdata.WORKSPACE).DataStore(testdata.DATASTORE_POSTGIS).PublishFeature(feature)
			assert.NoError(t, err)

			get, err := geoserverClient.Workspace(testdata.WORKSPACE).FeatureTypes().GetFeature(testdata.FEATURE_TYPE_NAME)
			assert.NoError(t, err)
			assert.Equal(t, get.FeatureType.Name, testdata.FEATURE_TYPE_NAME)
			assert.Equal(t, get.FeatureType.NativeName, testdata.FEATURE_TYPE_NATIVE_NAME)
			assert.Equal(t, get.FeatureType.Srs, "EPSG:4326")
			assert.Equal(t, get.FeatureType.Keywords.Keywords, []string{"features", "init"})
		})

		t.Run("WITH BBOX OPTION", func(t *testing.T) {
			var featureName = testdata.FEATURE_TYPE_NAME + "_WITH_BBOX"
			var bbox = [4]float64{-180.0, -90.0, 180.0, 90.0}
			var bboxSrs = "EPSG:4326"

			feature := featuretypes.New(featureName, testdata.FEATURE_TYPE_NATIVE_NAME, featuretypes.Options.BBOX(bbox, bboxSrs))

			err = geoserverClient.Workspace(testdata.WORKSPACE).DataStore(testdata.DATASTORE_POSTGIS).PublishFeature(feature)
			assert.NoError(t, err)

			get, err := geoserverClient.Workspace(testdata.WORKSPACE).FeatureTypes().GetFeature(featureName)
			assert.NoError(t, err)
			assert.Equal(t, get.FeatureType.Name, featureName)
			assert.Equal(t, get.FeatureType.NativeName, testdata.FEATURE_TYPE_NATIVE_NAME)
			assert.Equal(t, get.FeatureType.Srs, "EPSG:4326")
			assert.Equal(t, get.FeatureType.NativeBoundingBox.MinX, bbox[0])
			assert.Equal(t, get.FeatureType.NativeBoundingBox.MinY, bbox[1])
			assert.Equal(t, get.FeatureType.NativeBoundingBox.MaxX, bbox[2])
			assert.Equal(t, get.FeatureType.NativeBoundingBox.MaxY, bbox[3])
			assert.Equal(t, get.FeatureType.NativeBoundingBox.CRS, bboxSrs)
		})
	})

	err = geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)
	assert.NoError(t, err)
}

func TestFeatureTypeIntegration_Get(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

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

	//create feature type
	feature := featuretypes.New(testdata.FEATURE_TYPE_NAME, testdata.FEATURE_TYPE_NATIVE_NAME)
	err = geoserverClient.Workspace(testdata.WORKSPACE).DataStore(testdata.DATASTORE_POSTGIS).PublishFeature(feature)
	assert.NoError(t, err)

	t.Run("200 OK", func(t *testing.T) {
		get, err := geoserverClient.Workspace(testdata.WORKSPACE).FeatureTypes().GetFeature(testdata.FEATURE_TYPE_NAME)
		assert.NoError(t, err)
		assert.NotNil(t, get)
		assert.Equal(t, get.FeatureType.Name, testdata.FEATURE_TYPE_NAME)
		assert.Equal(t, get.FeatureType.NativeName, testdata.FEATURE_TYPE_NATIVE_NAME)
		assert.Equal(t, get.FeatureType.Srs, "EPSG:4326")
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		get, err := geoserverClient.Workspace(testdata.WORKSPACE).FeatureTypes().GetFeature(testdata.FEATURE_TYPE_NAME + "_DOES_NOT_EXIST")
		assert.Error(t, err)
		assert.Nil(t, get)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	err = geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)
	assert.NoError(t, err)
}

func TestFeatureTypeIntegration_Delete(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

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

	//create feature type
	feature := featuretypes.New(testdata.FEATURE_TYPE_NAME, testdata.FEATURE_TYPE_NATIVE_NAME)
	err = geoserverClient.Workspace(testdata.WORKSPACE).DataStore(testdata.DATASTORE_POSTGIS).PublishFeature(feature)
	assert.NoError(t, err)

	t.Run("200 OK", func(t *testing.T) {
		err = geoserverClient.Workspace(testdata.WORKSPACE).FeatureTypes().DeleteFeature(testdata.FEATURE_TYPE_NAME, true)
		assert.NoError(t, err)

		//try to retrieve the feature type
		get, err := geoserverClient.Workspace(testdata.WORKSPACE).FeatureTypes().GetFeature(testdata.FEATURE_TYPE_NAME)
		assert.Nil(t, get)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	t.Run("404 NOT FOUND", func(t *testing.T) {
		err = geoserverClient.Workspace(testdata.WORKSPACE).FeatureTypes().DeleteFeature(testdata.FEATURE_TYPE_NAME, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	err = geoserverClient.Workspaces().Delete(testdata.WORKSPACE, true)
	assert.NoError(t, err)
}
