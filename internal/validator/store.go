package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
)

var Store StoreValidator

type StoreValidator struct{}

func (sv StoreValidator) Name(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("empty store name"))
	}

	return validateAlphaNumerical(name)
}
