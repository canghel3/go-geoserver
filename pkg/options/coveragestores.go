package options

import "github.com/canghel3/go-geoserver/internal"

var Coveragestore CoveragestoreOptionGenerator

type CoveragestoreOptionGenerator struct{}

type CoveragestoreOptionFunc func(csl *internal.CoveragestoreOptions)

func (cs CoveragestoreOptionGenerator) Description(description string) CoveragestoreOptionFunc {
	return func(csl *internal.CoveragestoreOptions) {
		csl.Description = description
	}
}

func (cs CoveragestoreOptionGenerator) DisableConnectionOnFailure(disable bool) CoveragestoreOptionFunc {
	return func(csl *internal.CoveragestoreOptions) {
		csl.DisableOnConnectionFailure = disable
	}
}