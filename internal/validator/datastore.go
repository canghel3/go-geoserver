package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"net/url"
	"path/filepath"
	"strings"
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

	//TODO: check the directory contains some shapefiles
	return nil
}

func (dsv DataStoreValidator) CSV(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty csv url"))
	}

	if filepath.Ext(url) != ".csv" {
		return customerrors.WrapInputError(errors.New("csv file extension must be .csv"))
	}

	return nil
}

func (dsv DataStoreValidator) WebFeatureService(u string) error {
	if len(strings.TrimSpace(u)) == 0 {
		return customerrors.WrapInputError(errors.New("empty wfs url"))
	}

	_, err := url.Parse(u)
	if err != nil {
		return customerrors.WrapInputError(err)
	}

	//TODO: should we also check if the service=wfs is sent?

	return nil
}
