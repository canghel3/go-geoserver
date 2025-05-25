package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/models/customerrors"
	"regexp"
)

var Workspace WorkspaceValidator

type WorkspaceValidator struct{}

func (wv WorkspaceValidator) Name(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty workspace name"))
	}

	return validateAlphaNumerical(url)
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
