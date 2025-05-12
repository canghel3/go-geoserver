package options

import (
	"github.com/canghel3/go-geoserver/internal"
	"strconv"
)

var GeoPackage GeoPackageOptionsGenerator

type GeoPackageOptionsGenerator struct{}

func (gpog GeoPackageOptionsGenerator) ConnectionTimeout(timeout uint) GeoPackageOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Connection timeout"] = strconv.Itoa(int(timeout))
	}
}

type GeoPackageOptions func(params *internal.ConnectionParams)
