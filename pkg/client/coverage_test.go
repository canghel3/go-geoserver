package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/coverages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageIntegration_Create(t *testing.T) {
	err := addTestWorkspace()
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("GeoTIFF", func(t *testing.T) {
			err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, testdata.FileGeoTiff)
			assert.NoError(t, err)

			err = geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Publish(coverages.New(testdata.CoverageName, testdata.CoverageNativeName))
			assert.NoError(t, err)
		})
	})
}

func TestCoverageIntegration_Get(t *testing.T) {
	err := addTestWorkspace()
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("GeoTIFF", func(t *testing.T) {

		})
	})
}
