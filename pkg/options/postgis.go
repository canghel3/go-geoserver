package options

import "github.com/canghel3/go-geoserver/internal"

var PostGIS PostGISOptionGenerator

type PostGISOptionGenerator struct{}

type PostGISOptionFunc func(params *internal.ConnectionParams)

func (pgo PostGISOptionGenerator) ValidateConnections() PostGISOptionFunc {
	return func(params *internal.ConnectionParams) {
		(*params)["validate connections"] = "true"
	}
}
