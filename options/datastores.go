package options

import "github.com/canghel3/go-geoserver/internal"

var Datastore DatastoreOptionGenerator

type DatastoreOptionGenerator struct{}

type DatastoreOptionFunc func(dsl *internal.DatastoreOptions)

func (ds DatastoreOptionGenerator) Description(description string) DatastoreOptionFunc {
	return func(dsl *internal.DatastoreOptions) {
		dsl.Description = description
	}
}

func (ds DatastoreOptionGenerator) DisableConnectionOnFailure(disable bool) DatastoreOptionFunc {
	return func(dsl *internal.DatastoreOptions) {
		dsl.DisableOnConnectionFailure = disable
	}
}
