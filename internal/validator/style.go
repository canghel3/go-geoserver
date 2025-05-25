package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/models/customerrors"
)

var Style StyleValidator

type StyleValidator struct{}

func (sv StyleValidator) Name(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("empty style name"))
	}

	return validateAlphaNumerical(name)
}
