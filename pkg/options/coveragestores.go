package options

import "github.com/canghel3/go-geoserver/internal"

var CoverageStore CoverageStoreOptionGenerator

type CoverageStoreOptionGenerator struct{}

type CoveragestoreOptionFunc func(csl *internal.CoveragestoreOptions)

func (cs CoverageStoreOptionGenerator) Description(description string) CoveragestoreOptionFunc {
	return func(csl *internal.CoveragestoreOptions) {
		csl.Description = description
	}
}

func (cs CoverageStoreOptionGenerator) DisableConnectionOnFailure(disable bool) CoveragestoreOptionFunc {
	return func(csl *internal.CoveragestoreOptions) {
		csl.DisableOnConnectionFailure = disable
	}
}
