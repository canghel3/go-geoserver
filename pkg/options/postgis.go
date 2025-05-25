package options

import "github.com/canghel3/go-geoserver/internal"

var PostGIS PostGISOptionGenerator

type PostGISOptionGenerator struct{}

type PostGISOption func(params *internal.ConnectionParams)

func (pgo PostGISOptionGenerator) ValidateConnections() PostGISOption {
	return func(params *internal.ConnectionParams) {
		(*params)["validate connections"] = "true"
	}
}
