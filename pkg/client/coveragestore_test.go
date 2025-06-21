package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageStoreIntegration_Create(t *testing.T) {
	addTestWorkspace(t)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("GENERIC OPTIONS", func(t *testing.T) {
			t.Skip()
		})

		t.Run("GEOTIFF", func(t *testing.T) {
			addTestCoverageStore(t, types.GeoTIFF)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoTiff)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("EHdr", func(t *testing.T) {
			addTestCoverageStore(t, types.EHdr)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreEHdr)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("ENVIHdr", func(t *testing.T) {
			addTestCoverageStore(t, types.ENVIHdr)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreENVIHdr)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("GeoPackage", func(t *testing.T) {

		})

		t.Run("NITF", func(t *testing.T) {

		})

		t.Run("RST", func(t *testing.T) {

		})

		t.Run("VRT", func(t *testing.T) {})
	})
}

func TestCoverageStoreIntegration_Delete(t *testing.T) {
	addTestWorkspace(t)

	addTestCoverageStore(t, types.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreGeoTiff, true)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStores().Delete("unknown", true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}

func TestCoverageStoreIntegration_Get(t *testing.T) {
	addTestWorkspace(t)

	addTestCoverageStore(t, types.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoTiff)
		assert.NoError(t, err)
		assert.NotNil(t, store)
		assert.Equal(t, testdata.CoverageStoreGeoTiff, store.Name)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get("unknown")
		assert.Error(t, err)
		assert.Nil(t, store)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
	})
}
