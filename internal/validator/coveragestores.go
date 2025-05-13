package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"path/filepath"
)

var CoverageStore CoverageStoreValidator

type CoverageStoreValidator struct{}

func (csv CoverageStoreValidator) ArcGrid(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty url"))
	}

	ext := filepath.Ext(url)
	if ext != ".asc" && ext != ".grd" && ext != ".afd" {
		return customerrors.WrapInputError(errors.New("arc grid file extension must be .asc .grd or .afd"))
	}

	return nil
}
