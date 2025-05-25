package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/models/customerrors"
	"path/filepath"
)

var DataStore DataStoreValidator

type DataStoreValidator struct{}

func (dsv DataStoreValidator) PostGIS(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("empty store name"))
	}

	return validateAlphaNumerical(name)
}

func (dsv DataStoreValidator) GeoPackage(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty geopackage url"))
	}

	if filepath.Ext(url) != ".gpkg" {
		return customerrors.WrapInputError(errors.New("geopackage extension must be .gpkg"))
	}

	return nil
}

func (dsv DataStoreValidator) Shapefile(url string) error {
	if filepath.Ext(url) != ".shp" {
		return customerrors.WrapInputError(errors.New("shapefile extension must be .shp"))
	}

	return nil
}

func (dsv DataStoreValidator) ShapefileDirectory(dir string) error {
	if len(dir) == 0 {
		return customerrors.WrapInputError(errors.New("empty directory path"))
	}

	// Note: In a real implementation, you might want to check if the directory exists
	// and contains at least one shapefile, but for simplicity we'll just check if
	// the path is not empty.

	return nil
}

func (dsv DataStoreValidator) CSV(url string) error {
	if filepath.Ext(url) != ".csv" {
		return customerrors.WrapInputError(errors.New("csv file extension must be .csv"))
	}

	return nil
}
