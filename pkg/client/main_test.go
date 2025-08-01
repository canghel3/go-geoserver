package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/coverages"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

// GDAL raster drivers: https://gdal.org/en/stable/drivers/raster/index.html

var (
	vectorsTestDataDir = findTestDataPath("internal/testdata/featuretypes")
	rastersTestDataDir = findTestDataPath("internal/testdata/coverages")
	imagesTestDataDir  = findTestDataPath("internal/testdata/images")
	geoclient          = NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)
)

func TestMain(m *testing.M) {
	// VECTORS SETUP
	copyFileToGeoserver(testdata.FileShapefile, true)
	copyFileToGeoserver(testdata.FileGeoPackage, true)
	copyFileToGeoserver(testdata.FileCSVLatLon, true)
	copyFileToGeoserver(testdata.FileCSVWkt, true)
	copyDirToGeoserver(testdata.DirShapefiles, true)

	//RASTERS SETUP
	copyDirToGeoserver(testdata.DirGeoTiff, false)
	copyDirToGeoserver(testdata.DirEHdr, false)
	copyDirToGeoserver(testdata.DirENVIHdr, false)
	copyDirToGeoserver(testdata.DirERDASImg, false)
	copyDirToGeoserver(testdata.DirGeoPackageRaster, false)
	copyDirToGeoserver(testdata.DirNITF, false)
	copyDirToGeoserver(testdata.DirRST, false)
	copyDirToGeoserver(testdata.DirVRT, false)

	code := m.Run()
	os.Exit(code)
}

// copies the file parameter to the testdata.GeoserverDataDir. panics on error
func copyFileToGeoserver(file string, isVector bool) {
	var src string
	if isVector {
		src = filepath.Join(vectorsTestDataDir, file)
	} else {
		src = filepath.Join(rastersTestDataDir, file)
	}

	err := testdata.CopyFile(src, filepath.Join(testdata.GeoserverDataDir, file))
	if err != nil {
		panic(err)
	}
}

func copyDirToGeoserver(dir string, isVector bool) {
	var src string
	if isVector {
		src = filepath.Join(vectorsTestDataDir, dir)
	} else {
		src = filepath.Join(rastersTestDataDir, dir)
	}

	err := testdata.CopyDir(src, filepath.Join(testdata.GeoserverDataDir, dir))
	if err != nil {
		panic(err)
	}
}

func addTestWorkspace(t *testing.T) {
	geoclient.Workspaces().Delete(testdata.Workspace, true)

	if err := geoclient.Workspaces().Create(testdata.Workspace, false); err != nil {
		t.Fatal(err)
	}
}

func addTestDataStore(t *testing.T, type_ formats.DataStoreFormat) {
	switch type_ {
	case formats.PostGIS:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
			Host:     testdata.PostgisHost,
			Database: testdata.PostgisDb,
			User:     testdata.PostgisUsername,
			Password: testdata.PostgisPassword,
			Port:     testdata.PostgisPort,
			SSL:      testdata.PostgisSsl,
		}); err != nil {
			t.Fatal(err)
		}
		return
	case formats.GeoPackage:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().GeoPackage(testdata.DatastoreGeoPackage, testdata.FileGeoPackage); err != nil {
			t.Fatal(err)
		}
		return
	case formats.Shapefile:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().Shapefile(testdata.DatastoreShapefile, testdata.FileShapefile); err != nil {
			t.Fatal(err)
		}
		return
	case formats.DirOfShapefiles:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().Shapefiles(testdata.DatastoreDirOfShapefiles, testdata.DirShapefiles); err != nil {
			t.Fatal(err)
		}
		return
	//case types.CSV:
	//	if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().CSV(testdata.DatastoreDirOfShapefiles, testdata.DirShapefiles); err != nil {
	//		t.Fatal(err)
	//	}
	//	return
	case formats.WebFeatureService:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().WebFeatureService(testdata.DatastoreWebFeatureService, testdata.GeoserverUsername, testdata.GeoserverPassword, testdata.DatastoreWFSUrl); err != nil {
			t.Fatal(err)
		}
		return
	}

	t.Fatal("unsupported data store type")
}

func addTestFeatureType(t *testing.T, type_ formats.DataStoreFormat) {
	switch type_ {
	case formats.PostGIS:
		feature := featuretypes.New(testdata.FeatureTypePostgis, testdata.FeatureTypePostgisNativeName)
		if err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Publish(feature); err != nil {
			t.Fatal(err)
		}
		return
	case formats.GeoPackage:
		feature := featuretypes.New(testdata.FeatureTypeGeoPackage, testdata.FeatureTypeGeoPackageNativeName)
		if err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastoreGeoPackage).Publish(feature); err != nil {
			t.Fatal(err)
		}
		return
	}

	t.Fatal("unsupported vector layer type")
}

func addTestCoverageStore(t *testing.T, type_ formats.CoverageStoreFormat) {
	switch type_ {
	case formats.GeoTIFF:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, testdata.FileGeoTiff); err != nil {
			t.Fatal(err)
		}
		return
	case formats.EHdr:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().EHdr(testdata.CoverageStoreEHdr, testdata.FileEHdr); err != nil {
			t.Fatal(err)
		}
		return
	case formats.ENVIHdr:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ENVIHdr(testdata.CoverageStoreENVIHdr, testdata.FileENVIHdr); err != nil {
			t.Fatal(err)
		}
		return
	case formats.ERDASImg:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().ERDASImg(testdata.CoverageStoreERDASImg, testdata.FileERDASImg); err != nil {
			t.Fatal(err)
		}
		return
	//case types.GeoPackageMosaic:
	//	if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoPackage(testdata.CoverageStoreGeoPackage, testdata.FileGeoPackageRaster); err != nil {
	//		t.Fatal(err)
	//	}
	//	return
	case formats.NITF:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().NITF(testdata.CoverageStoreNITF, testdata.FileNITF); err != nil {
			t.Fatal(err)
		}
		return
	case formats.RST:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().RST(testdata.CoverageStoreRST, testdata.FileRST); err != nil {
			t.Fatal(err)
		}
		return
	case formats.VRT:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().VRT(testdata.CoverageStoreVRT, testdata.FileVRT); err != nil {
			t.Fatal(err)
		}
		return

	}

	t.Fatal("unsupported coverage store type")
}

func addTestCoverage(t *testing.T, type_ formats.CoverageStoreFormat) {
	switch type_ {
	case formats.GeoTIFF:
		coverage := coverages.New(testdata.CoverageGeoTiffName, testdata.CoverageGeoTiffNativeName)
		if err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Publish(coverage); err != nil {
			t.Fatal(err)
		}
		return
	case formats.EHdr:
		coverage := coverages.New(testdata.CoverageEHdrName, testdata.CoverageEHdrNativeName)
		if err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreEHdr).Publish(coverage); err != nil {
			t.Fatal(err)
		}
		return
	}

	t.Fatal("unsupported coverage store type")
}

func findGoModRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			panic(os.ErrNotExist)
		}
		dir = parent
	}
}

func findTestDataPath(path string) string {
	root := findGoModRoot()
	return filepath.Join(root, path)
}

func writeFile(t *testing.T, path string, content []byte) {
	tmpDir := t.TempDir()
	fullPath := filepath.Join(tmpDir, path)

	err := os.MkdirAll(filepath.Dir(fullPath), 0755)
	assert.NoError(t, err)

	err = os.WriteFile(fullPath, content, 0644)
	assert.NoError(t, err)
}
