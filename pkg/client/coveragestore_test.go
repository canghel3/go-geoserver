package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/coveragestores"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageStoreIntegration_Create(t *testing.T) {
	addTestWorkspace(t)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Generic Options", func(t *testing.T) {
			t.Run("Description", func(t *testing.T) {
				var suffix = "_DESC_OPT"
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create(options.GenericStore.Description("my description")).GeoTIFF(testdata.CoverageStoreGeoTiff+suffix, testdata.FileGeoTiff)
				assert.NoError(t, err)

				cvg, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoTiff + suffix)
				assert.NoError(t, err)
				assert.Equal(t, testdata.CoverageStoreGeoTiff+suffix, cvg.Name)
				assert.Equal(t, "my description", cvg.Description)
			})

			t.Run("AutoDisableOnConnFailure", func(t *testing.T) {
				var suffix = "_AUTO_DISABLE_OPT"
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create(options.GenericStore.AutoDisableOnConnFailure()).GeoTIFF(testdata.CoverageStoreGeoTiff+suffix, testdata.FileGeoTiff)
				assert.NoError(t, err)

				cvg, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoTiff + suffix)
				assert.NoError(t, err)
				assert.Equal(t, testdata.CoverageStoreGeoTiff+suffix, cvg.Name)
				assert.Equal(t, true, cvg.DisableConnectionOnFailure)
			})
		})

		t.Run("GeoTiff", func(t *testing.T) {
			addTestCoverageStore(t, formats.GeoTIFF)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoTiff)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("With file: In Filepath", func(t *testing.T) {
				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreGeoTiff, true)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, fmt.Sprintf("file:%s", testdata.FileGeoTiff))
				assert.NoError(t, err)
			})
		})

		t.Run("EHdr", func(t *testing.T) {
			addTestCoverageStore(t, formats.EHdr)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreEHdr)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("With file: In Filepath", func(t *testing.T) {
				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreEHdr, true)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().EHdr(testdata.CoverageStoreEHdr, fmt.Sprintf("file:%s", testdata.FileEHdr))
				assert.NoError(t, err)
			})
		})

		t.Run("ENVIHdr", func(t *testing.T) {
			addTestCoverageStore(t, formats.ENVIHdr)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreENVIHdr)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("With file: In Filepath", func(t *testing.T) {
				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreENVIHdr, true)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ENVIHdr(testdata.CoverageStoreENVIHdr, fmt.Sprintf("file:%s", testdata.FileENVIHdr))
				assert.NoError(t, err)
			})
		})

		t.Run("ERDASImg", func(t *testing.T) {
			addTestCoverageStore(t, formats.ERDASImg)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreERDASImg)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("With file: In Filepath", func(t *testing.T) {
				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreERDASImg, true)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ERDASImg(testdata.CoverageStoreERDASImg, fmt.Sprintf("file:%s", testdata.FileERDASImg))
				assert.NoError(t, err)
			})
		})

		t.Run("GeoPackage", func(t *testing.T) {
			t.Skip("not working yet")
			addTestCoverageStore(t, formats.GeoPackageMosaic)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoPackage)
			assert.NoError(t, err)
			assert.NotNil(t, store)
		})

		t.Run("NITF", func(t *testing.T) {
			addTestCoverageStore(t, formats.NITF)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreNITF)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("With file: In Filepath", func(t *testing.T) {
				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreNITF, true)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().NITF(testdata.CoverageStoreNITF, fmt.Sprintf("file:%s", testdata.FileNITF))
				assert.NoError(t, err)
			})
		})

		t.Run("RST", func(t *testing.T) {
			addTestCoverageStore(t, formats.RST)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreRST)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("With file: In Filepath", func(t *testing.T) {
				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreRST, true)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().RST(testdata.CoverageStoreRST, fmt.Sprintf("file:%s", testdata.FileRST))
				assert.NoError(t, err)
			})
		})

		t.Run("VRT", func(t *testing.T) {
			addTestCoverageStore(t, formats.VRT)

			store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreVRT)
			assert.NoError(t, err)
			assert.NotNil(t, store)

			t.Run("With file: In Filepath", func(t *testing.T) {
				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Delete(testdata.CoverageStoreVRT, true)
				assert.NoError(t, err)

				err = geoclient.Workspace(testdata.Workspace).CoverageStores().Create().VRT(testdata.CoverageStoreVRT, fmt.Sprintf("file:%s", testdata.FileVRT))
				assert.NoError(t, err)
			})
		})
	})

	t.Run("Validation Error", func(t *testing.T) {
		t.Run("EHdr", func(t *testing.T) {
			t.Run("Store name", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().EHdr(testdata.InvalidName, testdata.FileEHdr)
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "name can only contain alphanumerical characters")
			})

			t.Run("File extension", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().EHdr(testdata.CoverageStoreEHdr, "/path/to/file.csv")
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "EHdr file extension must be .bil")
			})
		})

		t.Run("ENVIHdr", func(t *testing.T) {
			t.Run("Store name", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ENVIHdr(testdata.InvalidName, testdata.FileENVIHdr)
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "name can only contain alphanumerical characters")
			})

			t.Run("File extension", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ENVIHdr(testdata.CoverageStoreENVIHdr, "/path/to/file.csv")
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "ENVIHdr file extension must be .dat")
			})
		})

		t.Run("ERDASImg", func(t *testing.T) {
			t.Run("Store name", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ERDASImg(testdata.InvalidName, testdata.FileERDASImg)
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "name can only contain alphanumerical characters")
			})

			t.Run("File extension", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ERDASImg(testdata.CoverageStoreERDASImg, "/path/to/file.csv")
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "ERDASImg file extension must be .img")
			})
		})

		t.Run("GeoTIFF", func(t *testing.T) {
			t.Run("Store name", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.InvalidName, testdata.FileGeoTiff)
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "name can only contain alphanumerical characters")
			})

			t.Run("File extension", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, "/path/to/file.csv")
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "geotiff file extension must be .tif or .tiff")
			})
		})

		t.Run("NITF", func(t *testing.T) {
			t.Run("Store name", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().NITF(testdata.InvalidName, testdata.FileNITF)
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "name can only contain alphanumerical characters")
			})

			t.Run("File extension", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().NITF(testdata.CoverageStoreNITF, "/path/to/file.csv")
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "NITF file extension must be .ntf")
			})
		})

		t.Run("RST", func(t *testing.T) {
			t.Run("Store name", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().RST(testdata.InvalidName, testdata.FileRST)
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "name can only contain alphanumerical characters")
			})

			t.Run("File extension", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().RST(testdata.CoverageStoreRST, "/path/to/file.csv")
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "RST file extension must be .rst")
			})
		})

		t.Run("VRT", func(t *testing.T) {
			t.Run("Store name", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().VRT(testdata.InvalidName, testdata.FileVRT)
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "name can only contain alphanumerical characters")
			})

			t.Run("File extension", func(t *testing.T) {
				err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().VRT(testdata.CoverageStoreVRT, "/path/to/file.csv")
				assert.IsType(t, err, &customerrors.InputError{})
				assert.EqualError(t, err, "VRT file extension must be .vrt")
			})
		})
	})
}

func TestCoverageStoreIntegration_Delete(t *testing.T) {
	addTestWorkspace(t)

	addTestCoverageStore(t, formats.GeoTIFF)

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
	addTestCoverageStore(t, formats.GeoTIFF)

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

func TestCoverageStoreIntegration_GetAll(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Single CoverageStore", func(t *testing.T) {
			stores, err := geoclient.Workspace(testdata.Workspace).CoverageStores().GetAll()
			assert.NoError(t, err)
			assert.Len(t, stores.Entries, 1)
		})

		t.Run("No CoverageStore", func(t *testing.T) {
			addTestWorkspace(t)

			stores, err := geoclient.Workspace(testdata.Workspace).CoverageStores().GetAll()
			assert.NoError(t, err)
			assert.Nil(t, stores.Entries)
		})
	})
}

func TestCoverageStoreIntegration_Reset(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStores().Reset(testdata.CoverageStoreGeoTiff)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStores().Reset("does-not-exist")
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coveragestore %s not found", "does-not-exist"))
	})
}

func TestCoverageStoreIntegration_Update(t *testing.T) {
	addTestWorkspace(t)
	addTestCoverageStore(t, formats.GeoTIFF)

	t.Run("200 Ok", func(t *testing.T) {
		store, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get(testdata.CoverageStoreGeoTiff)
		assert.NoError(t, err)
		assert.NotNil(t, store)

		store.Name = "changed"

		err = geoclient.Workspace(testdata.Workspace).CoverageStores().Update(testdata.CoverageStoreGeoTiff, *store)
		assert.NoError(t, err)

		cvg, err := geoclient.Workspace(testdata.Workspace).CoverageStores().Get("changed")
		assert.NoError(t, err)
		assert.NotNil(t, cvg)
		assert.Equal(t, "changed", cvg.Name)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		err := geoclient.Workspace(testdata.Workspace).CoverageStores().Update("does-not-exist", coveragestores.CoverageStore{})
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coveragestore %s not found", "does-not-exist"))
	})
}
