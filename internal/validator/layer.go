package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
)

var Layer LayerValidator

type LayerValidator struct{}

func (lv LayerValidator) Name(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("layer store name"))
	}

	return validateAlphaNumerical(name)
}
