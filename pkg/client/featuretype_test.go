package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeatureTypeIntegration_Create(t *testing.T) {
	addTestWorkspace(t)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("POSTGIS", func(t *testing.T) {
			addTestDataStore(t, formats.PostGIS)

			t.Run("Without Any Options", func(t *testing.T) {
				addTestFeatureType(t, formats.PostGIS)

				get, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis)
				assert.NoError(t, err)
				assert.Equal(t, get.Name, testdata.FeatureTypePostgis)
				assert.Equal(t, get.NativeName, testdata.FeatureTypePostgisNativeName)
				assert.Equal(t, get.Srs, "EPSG:4326")
				assert.Equal(t, get.Keywords.Keywords, []string{"features", "init"})
			})

			t.Run("With Bbox Option", func(t *testing.T) {
				var featureName = testdata.FeatureTypePostgis + "_WITH_BBOX"
				var bbox = [4]float64{-180.0, -90.0, 180.0, 90.0}
				var bboxSrs = "EPSG:4326"

				feature := featuretypes.New(featureName, testdata.FeatureTypePostgisNativeName, options.FeatureType.BBOX(bbox, bboxSrs))

				err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Publish(feature)
				assert.NoError(t, err)

				get, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(featureName)
				assert.NoError(t, err)
				assert.Equal(t, get.Name, featureName)
				assert.Equal(t, get.NativeName, testdata.FeatureTypePostgisNativeName)
				assert.Equal(t, get.Srs, "EPSG:4326")
				assert.Equal(t, get.NativeBoundingBox.MinX, bbox[0])
				assert.Equal(t, get.NativeBoundingBox.MinY, bbox[1])
				assert.Equal(t, get.NativeBoundingBox.MaxX, bbox[2])
				assert.Equal(t, get.NativeBoundingBox.MaxY, bbox[3])
				assert.Equal(t, get.NativeBoundingBox.CRS.Value, bboxSrs)
			})
		})
	})
}

func TestFeatureTypeIntegration_Get(t *testing.T) {
	addTestWorkspace(t)
	addTestDataStore(t, formats.PostGIS)
	addTestFeatureType(t, formats.PostGIS)

	t.Run("200 Ok", func(t *testing.T) {
		get, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis)
		assert.NoError(t, err)
		assert.NotNil(t, get)
		assert.Equal(t, get.Name, testdata.FeatureTypePostgis)
		assert.Equal(t, get.NativeName, testdata.FeatureTypePostgisNativeName)
		assert.Equal(t, get.Srs, "EPSG:4326")
	})

	t.Run("404 Not Found", func(t *testing.T) {
		get, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis + "_DOES_NOT_EXIST")
		assert.Error(t, err)
		assert.Nil(t, get)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestFeatureTypeIntegration_GetAll(t *testing.T) {
	addTestWorkspace(t)
	addTestDataStore(t, formats.PostGIS)
	addTestFeatureType(t, formats.PostGIS)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Single FeatureType", func(t *testing.T) {
			ft, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).GetAll()
			assert.NoError(t, err)
			assert.NotNil(t, ft)
			assert.Len(t, ft.Entries, 1)
		})

		t.Run("No FeatureType", func(t *testing.T) {
			addTestWorkspace(t)
			addTestDataStore(t, formats.PostGIS)

			ft, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).GetAll()
			assert.NoError(t, err)
			assert.Nil(t, ft.Entries)
		})
	})
}

func TestFeatureTypeIntegration_Update(t *testing.T) {
	addTestWorkspace(t)
	addTestDataStore(t, formats.PostGIS)
	addTestFeatureType(t, formats.PostGIS)

	t.Run("200 Ok", func(t *testing.T) {
		ft, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis)
		assert.NoError(t, err)
		assert.NotNil(t, ft)

		ft.Name = "changed"

		err = geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Update(testdata.FeatureTypePostgis, *ft)
		assert.NoError(t, err)

		fts, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get("changed")
		assert.NoError(t, err)
		assert.NotNil(t, fts)
		assert.Equal(t, "changed", fts.Name)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ft, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get("changed")
		assert.NoError(t, err)
		assert.NotNil(t, ft)

		err = geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Update("does-not-exist", *ft)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("featuretype %s not found", "does-not-exist"))
	})

	t.Run("Invalid Previous Name", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Update(testdata.InvalidName, featuretypes.FeatureType{Name: "some"})
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})

	t.Run("Invalid New Name", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Update("some", featuretypes.FeatureType{Name: testdata.InvalidName})
		assert.Error(t, err)
		assert.IsType(t, &customerrors.InputError{}, err)
		assert.EqualError(t, err, "name can only contain alphanumerical characters")
	})
}

func TestFeatureTypeIntegration_Delete(t *testing.T) {
	addTestWorkspace(t)
	addTestDataStore(t, formats.PostGIS)
	addTestFeatureType(t, formats.PostGIS)

	t.Run("200 Ok", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Delete(testdata.FeatureTypePostgis, true)
		assert.NoError(t, err)

		//try to retrieve the feature type
		get, err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Get(testdata.FeatureTypePostgis)
		assert.Nil(t, get)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Delete(testdata.FeatureTypePostgis, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestFeatureTypeIntegration_Reset(t *testing.T) {
	addTestWorkspace(t)
	addTestDataStore(t, formats.PostGIS)
	addTestFeatureType(t, formats.PostGIS)

	t.Run("200 Ok", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Reset(testdata.FeatureTypePostgis)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Reset("does-not-exist")
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("featuretype %s not found", "does-not-exist"))
	})
}
