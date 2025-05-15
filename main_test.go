package client

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"os"
	"path/filepath"
	"testing"
)

var (
	VectorsTestdataDir = filepath.Join("internal", "testdata", "vectors")
)

func TestMain(m *testing.M) {
	err := testdata.Copy(filepath.Join(VectorsTestdataDir, testdata.Shapefile), filepath.Join(testdata.GeoserverDataDir, testdata.Shapefile))
	if err != nil {
		panic(err)
	}

	err = testdata.Copy(filepath.Join(VectorsTestdataDir, testdata.GeoPackage), filepath.Join(testdata.GeoserverDataDir, testdata.GeoPackage))
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}
