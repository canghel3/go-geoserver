package client

import (
	"fmt"
	"testing"

	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/canghel3/go-geoserver/pkg/gwc"
	"github.com/canghel3/go-geoserver/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestGeoWebCacheIntegration_Status(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)
	addTestCoverage(t, formats.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Workspace In Layer Name", func(t *testing.T) {
			seedData := gwc.SeedData{
				Layer:       fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageGeoTiffName),
				Format:      formats.Png,
				Type:        types.Seed,
				ZoomStart:   0,
				ZoomStop:    10,
				ThreadCount: 1,
			}

			err := geoclient.Workspace(testdata.Workspace).GeoWebCache().Seed().Run(seedData)
			assert.NoError(t, err)

			status, err := geoclient.Workspace(testdata.Workspace).GeoWebCache().Seed().Status(seedData.Layer)
			assert.NoError(t, err)
			assert.NotNil(t, status)
		})

		t.Run("From Workspace", func(t *testing.T) {
			seedData := gwc.SeedData{
				Layer:       testdata.CoverageGeoTiffName,
				Format:      formats.Png,
				Type:        types.Seed,
				ZoomStart:   0,
				ZoomStop:    10,
				ThreadCount: 1,
			}

			err := geoclient.Workspace(testdata.Workspace).GeoWebCache().Seed().Run(seedData)
			assert.NoError(t, err)

			status, err := geoclient.Workspace(testdata.Workspace).GeoWebCache().Seed().Status(testdata.CoverageGeoTiffName)
			assert.NoError(t, err)
			assert.NotNil(t, status)
			//assert.Len(t, status.Info, 1)
		})
	})

	//t.Run("Input Error", func(t *testing.T) {
	//	expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.CoverageGeoTiffName)
	//	status, err := geoclient.GeoWebCache().Seed().Status(testdata.CoverageGeoTiffName)
	//	assert.Error(t, err)
	//	assert.Nil(t, status)
	//	assert.IsType(t, &customerrors.InputError{}, err)
	//	assert.EqualError(t, err, expectedError)
	//})
}

func TestGeoWebCacheIntegration_Seed(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)
	addTestCoverage(t, formats.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Workspace In Layer Name", func(t *testing.T) {
			seedData := gwc.SeedData{
				Layer:       fmt.Sprintf("%s:%s", testdata.Workspace, testdata.CoverageGeoTiffName),
				Format:      formats.Png,
				Type:        types.Seed,
				ZoomStart:   0,
				ZoomStop:    10,
				ThreadCount: 1,
			}

			err := geoclient.Workspace(testdata.Workspace).GeoWebCache().Seed().Run(seedData)
			assert.NoError(t, err)
		})

		t.Run("From Workspace", func(t *testing.T) {
			seedData := gwc.SeedData{
				Layer:       testdata.CoverageGeoTiffName,
				Format:      formats.Png,
				Type:        types.Seed,
				ZoomStart:   0,
				ZoomStop:    10,
				ThreadCount: 1,
			}

			err := geoclient.Workspace(testdata.Workspace).GeoWebCache().Seed().Run(seedData)
			assert.NoError(t, err)
		})
	})

	//t.Run("Input Error", func(t *testing.T) {
	//	seedData := gwc.SeedData{
	//		Layer:       testdata.CoverageGeoTiffName,
	//		Format:      formats.Png,
	//		Type:        types.Seed,
	//		ZoomStart:   0,
	//		ZoomStop:    10,
	//		ThreadCount: 1,
	//	}
	//
	//	expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.CoverageGeoTiffName)
	//	err := geoclient.GeoWebCache().Seed().Run(seedData)
	//	assert.Error(t, err)
	//	assert.IsType(t, &customerrors.InputError{}, err)
	//	assert.EqualError(t, err, expectedError)
	//})
}
