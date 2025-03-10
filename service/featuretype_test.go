package service

import (
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/datastore/postgis"
	"github.com/canghel3/go-geoserver/internal"
	"gotest.tools/v3/assert"
	"testing"
)

func TestFeatureType(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)

	bbox := [4]float64{}
	connectionParams := postgis.ConnectionParams{
		Host:     host,
		Database: databaseName,
		User:     databaseUser,
		Password: databasePassword,
		Port:     databasePort,
		SSL:      "disable",
	}

	t.Run("CREATE", func(t *testing.T) {
		t.Run("SIMPLE", func(t *testing.T) {
			assert.NilError(t, geoserverService.CreateWorkspace("init"))
			assert.NilError(t, geoserverService.CreatePostGISDataStore("init", "init", connectionParams))
			assert.NilError(t, geoserverService.CreateFeatureType("init", "init", "init", "init", "EPSG:4326", bbox))
		})

		t.Run("DUPLICATE", func(t *testing.T) {
			err := geoserverService.CreateFeatureType("init", "init", "init", "init", "EPSG:4326", bbox)
			assert.ErrorType(t, err, &customerrors.ConflictError{})
			assert.Error(t, err, "featuretype init already exists")
		})

		t.Run("IN NON-EXISTENT WORKSPACE", func(t *testing.T) {
			err := geoserverService.CreateFeatureType("_", "init", "init", "init", "EPSG:4326", bbox)
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "workspace _ does not exist")
		})

		t.Run("IN NON-EXISTENT STORE", func(t *testing.T) {
			err := geoserverService.CreateFeatureType("init", "_", "init", "init", "EPSG:4326", bbox)
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "datastore _ does not exist")
		})

		t.Run("EMPTY LAYER NAME", func(t *testing.T) {
			err := geoserverService.CreateFeatureType("init", "init", "", "init", "EPSG:4326", bbox)
			assert.ErrorType(t, err, &customerrors.InputError{})
			assert.Error(t, err, "empty layer name")
		})

		t.Run("WITH KEYWORDS OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.CreateFeatureType("init", "init", "with_keywords", "init", "EPSG:4326", bbox, internal.KeywordsOption([]string{"with", "keywords"})))
		})

		t.Run("WITH TITLE OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.CreateFeatureType("init", "init", "with_title", "init", "EPSG:4326", bbox, internal.TitleOption("titlu misto")))
		})

		t.Run("WITH PROJECTION POLICY OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.CreateFeatureType("init", "init", "with_projection", "init", "EPSG:4326", bbox, internal.ProjectionPolicyOption("FORCE_DECLARED")))
		})
	})

	t.Run("GET", func(t *testing.T) {
		t.Run("SIMPLE", func(t *testing.T) {
			featureType, err := geoserverService.GetFeatureType("init", "init", "init")
			assert.NilError(t, err)

			assert.Equal(t, "init", featureType.FeatureType.Name)
			assert.Equal(t, "init", featureType.FeatureType.Title)
			assert.Equal(t, "EPSG:4326", featureType.FeatureType.Srs)
			assert.Equal(t, "NONE", featureType.FeatureType.ProjectionPolicy)
		})

		t.Run("NON-EXISTENT", func(t *testing.T) {
			_, err := geoserverService.GetFeatureType("init", "init", "does_not_exist")
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "featuretype does_not_exist does not exist")
		})

		t.Run("WITH KEYWORDS", func(t *testing.T) {
			featureType, err := geoserverService.GetFeatureType("init", "init", "with_keywords")
			assert.NilError(t, err)

			assert.Equal(t, "with_keywords", featureType.FeatureType.Name)
			assert.Equal(t, "with_keywords", featureType.FeatureType.Title)
			assert.Equal(t, "EPSG:4326", featureType.FeatureType.Srs)
			assert.Equal(t, "NONE", featureType.FeatureType.ProjectionPolicy)
			assert.DeepEqual(t, []string{"with", "keywords"}, featureType.FeatureType.Keywords.Keywords)
		})

		t.Run("WITH TITLE", func(t *testing.T) {
			featureType, err := geoserverService.GetFeatureType("init", "init", "with_title")
			assert.NilError(t, err)

			assert.Equal(t, "with_title", featureType.FeatureType.Name)
			assert.Equal(t, "titlu misto", featureType.FeatureType.Title)
			assert.Equal(t, "EPSG:4326", featureType.FeatureType.Srs)
			assert.Equal(t, "NONE", featureType.FeatureType.ProjectionPolicy)
		})

		t.Run("WITH PROJECTION POLICY", func(t *testing.T) {
			featureType, err := geoserverService.GetFeatureType("init", "init", "with_projection")
			assert.NilError(t, err)

			assert.Equal(t, "with_projection", featureType.FeatureType.Name)
			assert.Equal(t, "with_projection", featureType.FeatureType.Title)
			assert.Equal(t, "EPSG:4326", featureType.FeatureType.Srs)
			assert.Equal(t, "FORCE_DECLARED", featureType.FeatureType.ProjectionPolicy)
		})
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Run("WITHOUT RECURSE", func(t *testing.T) {
			err := geoserverService.DeleteFeatureType("init", "init", "init", internal.RecurseOption(false))
			assert.ErrorType(t, err, &customerrors.GeoserverError{})
			assert.Error(t, err, "layer is being referenced by other resources")
		})

		t.Run("WITH RECURSE", func(t *testing.T) {
			assert.NilError(t, geoserverService.DeleteFeatureType("init", "init", "init", internal.RecurseOption(true)))
		})

		t.Run("NON-EXISTENT", func(t *testing.T) {
			err := geoserverService.DeleteFeatureType("init", "init", "non_existent")
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "featuretype non_existent does not exist")
		})
	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", internal.RecurseOption(true)))
}
