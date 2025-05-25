package options

import (
	"github.com/canghel3/go-geoserver/internal"
)

var DataStore DatastoreOptionGenerator

type DatastoreOptionGenerator struct{}

type DataStoreOption func(dsl *internal.DataStoreOptions)

func (ds DatastoreOptionGenerator) Description(description string) DataStoreOption {
	return func(dsl *internal.DataStoreOptions) {
		dsl.Description = description
	}
}

func (ds DatastoreOptionGenerator) DisableConnectionOnFailure(disable bool) DataStoreOption {
	return func(dsl *internal.DataStoreOptions) {
		dsl.DisableOnConnectionFailure = disable
	}
}
