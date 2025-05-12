package options

import (
	"github.com/canghel3/go-geoserver/internal"
	"strconv"
)

var GeoTIFF GeoTIFFOptionsGenerator

type GeoTIFFOptionsGenerator struct{}

func (gtog GeoTIFFOptionsGenerator) ConnectionTimeout(timeout uint) GeoTIFFOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Connection timeout"] = strconv.Itoa(int(timeout))
	}
}

type GeoTIFFOptions func(params *internal.ConnectionParams)