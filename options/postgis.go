package options

import "github.com/canghel3/go-geoserver/internal"

var PostGIS PostGISOption

type PostGISOption struct{}

type PostGISOptionFunc func(params *internal.ConnectionParams)

func (pgo PostGISOption) ValidateConnections() PostGISOptionFunc {
	return func(params *internal.ConnectionParams) {
		(*params)["validate connections"] = "true"
	}
}
