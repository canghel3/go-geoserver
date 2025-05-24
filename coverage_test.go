package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/models/coverages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageIntegration_Create(t *testing.T) {
	geoserverClient := NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)

	geoserverClient.Workspaces().Delete(testdata.Workspace, true)
	//create workspace
	geoserverClient.Workspaces().Create(testdata.Workspace, true)

	t.Run("200 OK", func(t *testing.T) {
		t.Run("GEOTIFF", func(t *testing.T) {
			err := geoserverClient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, testdata.FileGeoTiff)
			assert.NoError(t, err)

			err = geoserverClient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Publish(coverages.New(testdata.CoverageName, testdata.CoverageNativeName))
			assert.NoError(t, err)
		})
	})
}
