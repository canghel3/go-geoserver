package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"os"
	"path/filepath"
	"testing"
)

var (
	geoclient          = NewGeoserverClient(testdata.GeoserverUrl, testdata.GeoserverUsername, testdata.GeoserverPassword)
	VectorsTestdataDir = filepath.Join("..", "..", "internal", "testdata", "vectors")
	RastersTestdataDir = filepath.Join("..", "..", "internal", "testdata", "rasters")
)

func TestMain(m *testing.M) {
	// VECTORS SETUP
	err := testdata.Copy(filepath.Join(VectorsTestdataDir, testdata.FileShapefile), filepath.Join(testdata.GeoserverDataDir, testdata.FileShapefile))
	if err != nil {
		panic(err)
	}

	err = testdata.Copy(filepath.Join(VectorsTestdataDir, testdata.FileGeoPackage), filepath.Join(testdata.GeoserverDataDir, testdata.FileGeoPackage))
	if err != nil {
		panic(err)
	}

	// RASTERS SETUP
	err = testdata.Copy(filepath.Join(RastersTestdataDir, testdata.FileGeoTiff), filepath.Join(testdata.GeoserverDataDir, testdata.FileGeoTiff))
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

func addTestWorkspace() error {
	geoclient.Workspaces().Delete(testdata.Workspace, true)

	return geoclient.Workspaces().Create(testdata.Workspace, false)
}

func addTestDataStore() error {
	return geoclient.Workspace(testdata.Workspace).DataStores().Create().GeoPackage(testdata.DatastoreGeoPackage, testdata.FileGeoPackage)
}

func addTestVectorLayer() error {
	feature := featuretypes.New(testdata.FeatureTypeGeoPackage, testdata.FeatureTypeGeoPackageNativeName)
	return geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastoreGeoPackage).Publish(feature)
}
