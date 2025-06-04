package options

import (
	"github.com/canghel3/go-geoserver/internal/models"
)

var CoverageStore CoverageStoreOptionGenerator

type CoverageStoreOptionGenerator struct{}

type CoverageStoreOption func(csl *models.CoverageStoreOptions)

func (cs CoverageStoreOptionGenerator) Description(description string) CoverageStoreOption {
	return func(csl *models.CoverageStoreOptions) {
		csl.Description = description
	}
}

func (cs CoverageStoreOptionGenerator) Default() CoverageStoreOption {
	return func(csl *models.CoverageStoreOptions) {
		csl.Default = true
	}
}
