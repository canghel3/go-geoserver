package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageIntegration_Create(t *testing.T) {
	err := addTestWorkspace()
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("GeoTIFF", func(t *testing.T) {
			err = addTestCoverageStore(types.GeoTIFF)
			assert.NoError(t, err)

			err = addTestCoverage(types.GeoTIFF)
			assert.NoError(t, err)
		})
	})
}

func TestCoverageIntegration_Get(t *testing.T) {
	err := addTestWorkspace()
	assert.NoError(t, err)

	err = addTestCoverageStore(types.GeoTIFF)
	assert.NoError(t, err)

	err = addTestCoverage(types.GeoTIFF)
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("GeoTIFF", func(t *testing.T) {
			coverage, err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Get(testdata.CoverageGeoTiffName)
			assert.NoError(t, err)
			assert.NotNil(t, coverage)
			assert.Equal(t, testdata.CoverageGeoTiffName, coverage.Name)
			assert.NotNil(t, coverage.Srs)
			assert.Equal(t, "EPSG:32631", *coverage.Srs)
		})
	})
}

func TestCoverageIntegration_Delete(t *testing.T) {
	err := addTestWorkspace()
	assert.NoError(t, err)

	err = addTestCoverageStore(types.GeoTIFF)
	assert.NoError(t, err)

	err = addTestCoverage(types.GeoTIFF)
	assert.NoError(t, err)

	t.Run("200 Ok", func(t *testing.T) {
		err = geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Delete(testdata.CoverageGeoTiffName, true)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err = geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Delete(testdata.CoverageGeoTiffName, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coverage %s not found", testdata.CoverageGeoTiffName))
	})
}
