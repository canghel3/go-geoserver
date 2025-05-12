package options

import (
	"github.com/canghel3/go-geoserver/internal"
	"strconv"
)

var Shapefile ShapefileOptionsGenerator

type ShapefileOptionsGenerator struct{}

type ShapefileOptions func(params *internal.ConnectionParams)

// Charset sets the character set for the shapefile
func (sog ShapefileOptionsGenerator) Charset(charset string) ShapefileOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["charset"] = charset
	}
}

// CreateSpatialIndex enables or disables the creation of a spatial index
func (sog ShapefileOptionsGenerator) CreateSpatialIndex(create bool) ShapefileOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["create spatial index"] = strconv.FormatBool(create)
	}
}

// Memory enables or disables loading the shapefile into memory
func (sog ShapefileOptionsGenerator) Memory(inMemory bool) ShapefileOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["memory mapped buffer"] = strconv.FormatBool(inMemory)
	}
}

// CacheAndReuse enables or disables caching and reusing the shapefile
func (sog ShapefileOptionsGenerator) CacheAndReuse(cache bool) ShapefileOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["cache and reuse memory maps"] = strconv.FormatBool(cache)
	}
}

// Directory specifies additional options for a directory of shapefiles
func (sog ShapefileOptionsGenerator) Directory(options map[string]string) ShapefileOptions {
	return func(params *internal.ConnectionParams) {
		for key, value := range options {
			(*params)[key] = value
		}
	}
}

// FilePattern sets the pattern for matching shapefile names in a directory
func (sog ShapefileOptionsGenerator) FilePattern(pattern string) ShapefileOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["file_pattern"] = pattern
	}
}

// Namespace sets the namespace for the shapefiles in a directory
func (sog ShapefileOptionsGenerator) Namespace(namespace string) ShapefileOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["namespace"] = namespace
	}
}
