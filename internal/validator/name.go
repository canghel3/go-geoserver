package validator

import (
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"regexp"
	"strings"
)

func Name(name string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return customerrors.WrapInputError(errors.New("empty name"))
	}

	return validateAlphaNumerical(name)
}

func Empty(name string) bool {
	return len(name) == 0 || len(strings.TrimSpace(name)) == 0
}

// WorkspaceLayerFormat validates that if the `workspace` string is empty,
// the `layer` must include a workspace prefix formatted as `<workspace>:<layer>`, otherwise it returns an error.
// If the `workspace` is provided, there is no restriction on the format of the `layer`.
func WorkspaceLayerFormat(workspace, layer string) error {
	split := strings.Split(layer, ":")
	if len(workspace) == 0 && (len(split) != 2 || len(split[0]) == 0) {
		return customerrors.NewInputError(fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", strings.Trim(layer, ":")))
	}

	return nil
}

func validateAlphaNumerical(name string) error {
	regex, err := regexp.Compile(`[^a-zA-Z0-9_-]+`)
	if err != nil {
		return err
	}

	if regex.MatchString(name) {
		return customerrors.NewInputError("name can only contain alphanumerical characters")
	}
	return nil
}
