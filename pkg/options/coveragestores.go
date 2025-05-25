package options

import (
	"github.com/canghel3/go-geoserver/internal"
)

var CoverageStore CoverageStoreOptionGenerator

type CoverageStoreOptionGenerator struct{}

type CoverageStoreOption func(csl *internal.CoverageStoreOptions)

func (cs CoverageStoreOptionGenerator) Description(description string) CoverageStoreOption {
	return func(csl *internal.CoverageStoreOptions) {
		csl.Description = description
	}
}

func (cs CoverageStoreOptionGenerator) Default() CoverageStoreOption {
	return func(csl *internal.CoverageStoreOptions) {
		csl.Default = true
	}
}
