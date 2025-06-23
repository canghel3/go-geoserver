package options

import (
	"github.com/canghel3/go-geoserver/internal/models"
)

var GenericStore GenericStoreOptionGenerator

type GenericStoreOptionGenerator struct{}

type GenericStoreOption func(csl *models.GenericStoreOptions)

func (cs GenericStoreOptionGenerator) Description(description string) GenericStoreOption {
	return func(csl *models.GenericStoreOptions) {
		csl.Description = description
	}
}

func (cs GenericStoreOptionGenerator) AutoDisableOnConnFailure() GenericStoreOption {
	return func(csl *models.GenericStoreOptions) {
		csl.AutoDisableOnConnFailure = true
	}
}
