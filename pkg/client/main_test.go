package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"github.com/canghel3/go-geoserver/pkg/types"
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

func addTestDataStore(type_ types.DataStoreType) error {
	switch type_ {
	case types.PostGIS:
		return geoclient.Workspace(testdata.Workspace).DataStores().Create().PostGIS(testdata.DatastorePostgis, postgis.ConnectionParams{
			Host:     testdata.PostgisHost,
			Database: testdata.PostgisDb,
			User:     testdata.PostgisUsername,
			Password: testdata.PostgisPassword,
			Port:     testdata.PostgisPort,
			SSL:      testdata.PostgisSsl,
		})
	case types.GeoPackage:
		return geoclient.Workspace(testdata.Workspace).DataStores().Create().GeoPackage(testdata.DatastoreGeoPackage, testdata.FileGeoPackage)
	case types.Shapefile:
		return geoclient.Workspace(testdata.Workspace).DataStores().Create().Shapefile(testdata.DatastoreShapefile, testdata.FileShapefile)
	}

	return customerrors.NewUnsupportedError("unsupported data store type")
}

func addTestVectorLayer() error {
	feature := featuretypes.New(testdata.FeatureTypeGeoPackage, testdata.FeatureTypeGeoPackageNativeName)
	return geoclient.Workspace(testdata.Workspace).DataStore(testdata.DatastoreGeoPackage).Publish(feature)
}

func addTestCoverageStore(type_ types.CoverageStoreType) error {
	switch type_ {
	case types.GeoTIFF:
		return geoclient.Workspace(testdata.Workspace).CoverageStores().Create().GeoTIFF(testdata.CoverageStoreGeoTiff, testdata.FileGeoTiff)

	}

	return customerrors.NewUnsupportedError("unsupported coverage store type")
}
