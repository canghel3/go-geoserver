package options

import (
	"github.com/canghel3/go-geoserver/internal/models"
)

var DataStore DatastoreOptionGenerator

type DatastoreOptionGenerator struct{}

type DataStoreOption func(dsl *models.DataStoreOptions)

func (ds DatastoreOptionGenerator) Description(description string) DataStoreOption {
	return func(dsl *models.DataStoreOptions) {
		dsl.Description = description
	}
}

func (ds DatastoreOptionGenerator) DisableConnectionOnFailure(disable bool) DataStoreOption {
	return func(dsl *models.DataStoreOptions) {
		dsl.DisableOnConnectionFailure = disable
	}
}
