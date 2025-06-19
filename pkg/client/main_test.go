package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/coverages"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"github.com/canghel3/go-geoserver/pkg/types"
	"os"
	"path/filepath"
	"testing"
)

const (
	VectorsTestDataDir = "../../internal/testdata/vectors"
	RastersTestDataDir = "../../internal/testdata/rasters"
)

var (
	geoclient = NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)
)

func TestMain(m *testing.M) {
	// VECTORS SETUP
	copyFileToGeoserver(testdata.FileShapefile, true)
	copyFileToGeoserver(testdata.FileGeoPackage, true)

	//RASTERS SETUP
	copyFileToGeoserver(testdata.FileGeoTiff, false)

	code := m.Run()
	os.Exit(code)
}

// copies the file parameter to the testdata.GeoserverDataDir. panics on error
func copyFileToGeoserver(file string, isVector bool) {
	var src string
	if isVector {
		src = filepath.Join(VectorsTestDataDir, file)
	} else {
		src = filepath.Join(RastersTestDataDir, file)
	}

	err := testdata.Copy(src, filepath.Join(testdata.GeoserverDataDir, file))
	if err != nil {
		panic(err)
	}
}

func addTestWorkspace(t *testing.T) {
	geoclient.Workspaces().Delete(testdata.Workspace, true)

	if err := geoclient.Workspaces().Create(testdata.Workspace, false); err != nil {
		t.FailNow()
	}
}

func addTestDataStore(t *testing.T, type_ types.DataStoreType) {
	switch type_ {
	case types.PostGIS:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
			Host:     testdata.PostgisHost,
			Database: testdata.PostgisDb,
			User:     testdata.PostgisUsername,
			Password: testdata.PostgisPassword,
			Port:     testdata.PostgisPort,
			SSL:      testdata.PostgisSsl,
		}); err != nil {
			t.FailNow()
		}
		return
	case types.GeoPackage:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().GeoPackage(testdata.DatastoreGeoPackage, testdata.FileGeoPackage); err != nil {
			t.FailNow()
		}
		return
	case types.Shapefile:
		if err := geoclient.Workspace(testdata.Workspace).DataStores().Create().Shapefile(testdata.DatastoreShapefile, testdata.FileShapefile); err != nil {
			t.FailNow()
		}
		return
	}

	t.Fatal(customerrors.NewUnsupportedError("unsupported data store type"))
}

func addTestVectorLayer(t *testing.T, type_ types.DataStoreType) {
	switch type_ {
	case types.PostGIS:
		feature := featuretypes.New(testdata.FeatureTypePostgis, testdata.FeatureTypePostgisNativeName)
		if err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastorePostgis).Publish(feature); err != nil {
			t.FailNow()
		}
		return
	case types.GeoPackage:
		feature := featuretypes.New(testdata.FeatureTypeGeoPackage, testdata.FeatureTypeGeoPackageNativeName)
		if err := geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastoreGeoPackage).Publish(feature); err != nil {
			t.FailNow()
		}
		return
	}

	t.Fatal(customerrors.NewUnsupportedError("unsupported vector layer type"))
}

func addTestCoverageStore(t *testing.T, type_ types.CoverageStoreType) {
	switch type_ {
	case types.GeoTIFF:
		if err := geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, testdata.FileGeoTiff); err != nil {
			t.FailNow()
		}
		return
	}

	t.Fatal(customerrors.NewUnsupportedError("unsupported coverage store type"))
}

func addTestCoverage(t *testing.T, type_ types.CoverageStoreType) {
	switch type_ {
	case types.GeoTIFF:
		coverage := coverages.New(testdata.CoverageGeoTiffName, testdata.CoverageGeoTiffNativeName)
		if err := geoclient.Workspace(testdata.Workspace).CoverageStore(testdata.CoverageStoreGeoTiff).Publish(coverage); err != nil {
			t.FailNow()
		}
		return
	}

	t.Fatal(customerrors.NewUnsupportedError("unsupported coverage store type"))
}
