package utils

import (
	"errors"
	"github.com/canghel3/go-geoserver/customerrors"
	"regexp"
)

func ValidateWorkspace(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("empty workspace name"))
	}

	return validateAlphaNumerical(name)
}

func ValidateStore(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("empty store name"))
	}

	return validateAlphaNumerical(name)
}

func ValidateLayer(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("empty layer name"))
	}

	return validateAlphaNumerical(name)
}

func ValidateStyle(name string) error {
	if len(name) == 0 {
		return customerrors.WrapInputError(errors.New("empty style name"))
	}

	return validateAlphaNumerical(name)
}

func validateAlphaNumerical(name string) error {
	regex, err := regexp.Compile(`[^a-zA-Z0-9_]+`)
	if err != nil {
		return err
	}

	if regex.MatchString(name) {
		return customerrors.WrapInputError(errors.New("name can only contain alphanumerical characters"))
	}
	return nil
}
