package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"strings"
)

func Name(name string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return customerrors.WrapInputError(errors.New("empty name"))
	}

	return validateAlphaNumerical(name)
}
