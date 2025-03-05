package service

import (
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"gotest.tools/v3/assert"
	"testing"
)

func TestCoverageStore(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)
	assert.NilError(t, geoserverService.CreateWorkspace("init"))

	t.Run("GeoTIFF", func(t *testing.T) {
		t.Run("CREATE", func(t *testing.T) {
			t.Run("SIMPLE", func(t *testing.T) {
				assert.NilError(t, geoserverService.CreateCoverageStore("init", "init", "file:/opt/geoserver/data/shipments_2_geocoded.tif", "GeoTIFF"))
			})

			t.Run("DUPLICATE", func(t *testing.T) {
				assert.ErrorType(t, geoserverService.CreateCoverageStore("init", "init", "shipments_2_geocoded.tif", "GeoTIFF"), &customerrors.ConflictError{})
			})

			t.Run("WITHOUT FILE PREFIX IN URL", func(t *testing.T) {
				geoserverService = NewGeoserverService(geoserverURL, username, password, GeoserverServiceDataDirOption("/opt/geoserver/data"))
				assert.Equal(t, geoserverService.isDataDirectorySet(), true)
				assert.NilError(t, geoserverService.CreateCoverageStore("init", "init_with_option", "shipments_2_geocoded.tif", "GeoTIFF"))
			})

			t.Run("FILE DOES NOT END WITH .TIF", func(t *testing.T) {
				err := geoserverService.CreateCoverageStore("init", "new", "file:/opt/geoserver/data/shipments_2_geocoded", "geotiff")
				assert.ErrorType(t, err, &customerrors.InputError{})
				assert.Error(t, err, "file must be of type .tif, got ")
			})
		})

		t.Run("GET", func(t *testing.T) {
			t.Run("SIMPLE", func(t *testing.T) {
				cs, err := geoserverService.GetCoverageStore("init", "init_with_option")
				assert.NilError(t, err)
				assert.Equal(t, cs.Name, "init_with_option")
				assert.Equal(t, cs.URL, "file:/opt/geoserver/data/shipments_2_geocoded.tif")
				assert.Equal(t, cs.Type, "GeoTIFF")
			})

			t.Run("GET NON-EXISTENT", func(t *testing.T) {
				_, err := geoserverService.GetCoverageStore("init", "no")
				assert.ErrorType(t, err, &customerrors.NotFoundError{})
				assert.Error(t, err, "coverage store no does not exist")
			})
		})

		//TODO: test delete if coverage is in layer group
		t.Run("DELETE", func(t *testing.T) {
			t.Run("WITHOUT RECURSE", func(t *testing.T) {
				assert.NilError(t, geoserverService.CreateCoverageStore("init", "init2", "shipments_2_geocoded.tif", "GeoTIFF"))
				assert.NilError(t, geoserverService.DeleteCoverageStore("init", "init2", internal.RecurseOption(false)))
			})

			t.Run("WITH RECURSE", func(t *testing.T) {
				assert.NilError(t, geoserverService.DeleteCoverageStore("init", "init", internal.RecurseOption(true)))
			})

			//TODO: purge does not seem to delete the file from geoserver directory (this may be a geoserver bug)
			t.Run("WITH RECURSE AND PURGE", func(t *testing.T) {
				assert.NilError(t, geoserverService.DeleteCoverageStore("init", "init_with_option", internal.RecurseOption(true), internal.PurgeOption("all")))
			})

			t.Run("NON-EXISTENT", func(t *testing.T) {
				err := geoserverService.DeleteCoverageStore("init", "none")
				assert.ErrorType(t, err, &customerrors.NotFoundError{})
				assert.Error(t, err, "coverage store none does not exist")
			})
		})
	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", internal.RecurseOption(true)))
}
