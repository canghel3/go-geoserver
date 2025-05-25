package options

import (
	"github.com/canghel3/go-geoserver/pkg/datastores"
	"strconv"
)

var GeoPackage GeoPackageOptionsGenerator

type GeoPackageOptionsGenerator struct{}

func (gpog GeoPackageOptionsGenerator) ConnectionTimeout(timeout uint) GeoPackageOptions {
	return func(params *datastores.ConnectionParams) {
		(*params)["Connection timeout"] = strconv.Itoa(int(timeout))
	}
}

type GeoPackageOptions func(params *datastores.ConnectionParams)
