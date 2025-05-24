package options

import "github.com/canghel3/go-geoserver/internal"

var CoverageStore CoverageStoreOptionGenerator

type CoverageStoreOptionGenerator struct{}

type CoverageStoreOptionFunc func(csl *internal.CoverageStoreOptions)

func (cs CoverageStoreOptionGenerator) Description(description string) CoverageStoreOptionFunc {
	return func(csl *internal.CoverageStoreOptions) {
		csl.Description = description
	}
}
