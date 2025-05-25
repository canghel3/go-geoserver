package validator

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/models/customerrors"
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

func (csv CoverageStoreValidator) AIG(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) DTED(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) EHdr(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) ENVIHdr(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) ERDASImg(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) NITF(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) RPFTOC(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) RST(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) SRP(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) VRT(url string) error {
	return errors.New("not implemented")
}

func (csv CoverageStoreValidator) GeoPackage(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty geopackage url"))
	}

	if filepath.Ext(url) != ".gpkg" {
		return customerrors.WrapInputError(errors.New("geopackage extension must be .gpkg"))
	}

	return nil
}

func (csv CoverageStoreValidator) GeoTIFF(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty url"))
	}

	ext := filepath.Ext(url)
	if ext != ".tif" && ext != ".tiff" {
		return customerrors.WrapInputError(errors.New("geotiff file extension must be .tif or .tiff"))
	}

	return nil
}

func (csv CoverageStoreValidator) WorldImage(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty url"))
	}

	ext := filepath.Ext(url)
	validExts := []string{".png", ".jpg", ".jpeg", ".gif", ".tif", ".tiff", ".bmp"}

	valid := false
	for _, validExt := range validExts {
		if ext == validExt {
			valid = true
			break
		}
	}

	if !valid {
		return customerrors.WrapInputError(errors.New("worldimage file must have a valid image extension (.png, .jpg, .jpeg, .gif, .tif, .tiff, .bmp)"))
	}

	return nil
}

func (csv CoverageStoreValidator) ImageMosaic(dir string) error {
	if len(dir) == 0 {
		return customerrors.WrapInputError(errors.New("empty directory path for image mosaic"))
	}

	// Note: In a real implementation, you might want to check if the directory exists
	// and contains valid image files, but for simplicity we'll just check if
	// the path is not empty.

	return nil
}

func (csv CoverageStoreValidator) ImagePyramid(dir string) error {
	if len(dir) == 0 {
		return customerrors.WrapInputError(errors.New("empty directory path for image pyramid"))
	}

	// Note: In a real implementation, you might want to check if the directory exists
	// and contains valid image files, but for simplicity we'll just check if
	// the path is not empty.

	return nil
}
