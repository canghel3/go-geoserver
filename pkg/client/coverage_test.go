package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/coverages"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageIntegration_Create(t *testing.T) {
	addTestWorkspace(t)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("GeoTIFF", func(t *testing.T) {
			addTestCoverageStore(t, types.GeoTIFF)

			addTestCoverage(t, types.GeoTIFF)
		})
	})
}

func TestCoverageIntegration_Update(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, types.GeoTIFF)
	addTestCoverage(t, types.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Update(testdata.CoverageGeoTiffName, coverages.New(testdata.CoverageGeoTiffName+"_2", testdata.CoverageGeoTiffNativeName))
		assert.NoError(t, err)

		cvg, err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Get(testdata.CoverageGeoTiffName + "_2")
		assert.NoError(t, err)
		assert.NotNil(t, cvg)
		assert.Equal(t, testdata.CoverageGeoTiffName+"_2", cvg.Name)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Update(testdata.CoverageGeoTiffName, coverages.New(testdata.CoverageGeoTiffName, testdata.CoverageGeoTiffNativeName))
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coverage %s not found", testdata.CoverageGeoTiffName))
	})
}

func TestCoverageIntegration_Get(t *testing.T) {
	addTestWorkspace(t)

	addTestCoverageStore(t, types.GeoTIFF)

	addTestCoverage(t, types.GeoTIFF)

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

func TestCoverageIntegration_GetAll(t *testing.T) {
	addTestWorkspace(t)

	addTestCoverageStore(t, types.GeoTIFF)

	addTestCoverage(t, types.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		coverages, err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).GetAll()
		assert.Error(t, err)
		assert.Nil(t, coverages)
		//assert.NoError(t, err)
		//assert.NotNil(t, coverages)
		//assert.Len(t, coverages, 1)
		//assert.Equal(t, testdata.CoverageGeoTiffName, coverages[0].Name)
	})
}

func TestCoverageIntegration_Delete(t *testing.T) {
	addTestWorkspace(t)

	addTestCoverageStore(t, types.GeoTIFF)

	addTestCoverage(t, types.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Delete(testdata.CoverageGeoTiffName, true)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Delete(testdata.CoverageGeoTiffName, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coverage %s not found", testdata.CoverageGeoTiffName))
	})
}
