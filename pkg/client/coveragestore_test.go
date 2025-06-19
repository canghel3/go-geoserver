package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageStoreIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	addTestWorkspace(t)
	geoserverClient.Workspaces().Delete(testdata.Workspace, true)

	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("GENERIC OPTIONS", func(t *testing.T) {
			t.Skip()
		})

		t.Run("GEOTIFF", func(t *testing.T) {
			err := geoserverClient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, testdata.FileGeoTiff)
			assert.NoError(t, err)

			store, err := geoserverClient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoTiff)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})
	})
}

func TestCoverageStoreIntegration_Delete(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {

	})
}

func TestCoverageStoreIntegration_Get(t *testing.T) {

}
