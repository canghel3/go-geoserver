package options

import (
	"github.com/canghel3/go-geoserver/pkg/datastores"
)

var PostGIS PostGISOptionGenerator

type PostGISOptionGenerator struct{}

type PostGISOption func(params *datastores.ConnectionParams)

func (pgo PostGISOptionGenerator) ValidateConnections() PostGISOption {
	return func(params *datastores.ConnectionParams) {
		(*params)["validate connections"] = "true"
	}
}
