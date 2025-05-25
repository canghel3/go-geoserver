package options

import (
	"github.com/canghel3/go-geoserver/pkg/datastores"
	"strconv"
)

var Shapefile ShapefileOptionsGenerator

type ShapefileOptionsGenerator struct{}

type ShapefileOption func(params *datastores.ConnectionParams)

// Charset sets the character set for the shapefile
func (sog ShapefileOptionsGenerator) Charset(charset string) ShapefileOption {
	return func(params *datastores.ConnectionParams) {
		(*params)["charset"] = charset
	}
}

// CreateSpatialIndex enables or disables the creation of a spatial index
func (sog ShapefileOptionsGenerator) CreateSpatialIndex(create bool) ShapefileOption {
	return func(params *datastores.ConnectionParams) {
		(*params)["create spatial index"] = strconv.FormatBool(create)
	}
}

// Memory enables or disables loading the shapefile into memory
func (sog ShapefileOptionsGenerator) Memory(inMemory bool) ShapefileOption {
	return func(params *datastores.ConnectionParams) {
		(*params)["memory mapped buffer"] = strconv.FormatBool(inMemory)
	}
}

// CacheAndReuse enables or disables caching and reusing the shapefile
func (sog ShapefileOptionsGenerator) CacheAndReuse(cache bool) ShapefileOption {
	return func(params *datastores.ConnectionParams) {
		(*params)["cache and reuse memory maps"] = strconv.FormatBool(cache)
	}
}

// Directory specifies additional options for a directory of shapefiles
func (sog ShapefileOptionsGenerator) Directory(options map[string]string) ShapefileOption {
	return func(params *datastores.ConnectionParams) {
		for key, value := range options {
			(*params)[key] = value
		}
	}
}

// FilePattern sets the pattern for matching shapefile names in a directory
func (sog ShapefileOptionsGenerator) FilePattern(pattern string) ShapefileOption {
	return func(params *datastores.ConnectionParams) {
		(*params)["file_pattern"] = pattern
	}
}

// Namespace sets the namespace for the shapefiles in a directory
func (sog ShapefileOptionsGenerator) Namespace(namespace string) ShapefileOption {
	return func(params *datastores.ConnectionParams) {
		(*params)["namespace"] = namespace
	}
}
