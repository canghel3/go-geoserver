package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"os"
	"path/filepath"
	"testing"
)

var (
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
