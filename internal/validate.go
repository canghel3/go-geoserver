package internal

import (
	"errors"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"path/filepath"
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

func ValidateGeoPackage(url string) error {
	if len(url) == 0 {
		return customerrors.WrapInputError(errors.New("empty geopackage url"))
	}

	if filepath.Ext(url) != ".gpkg" {
		return customerrors.WrapInputError(errors.New("geopackage extension must be .gpkg"))
	}

	return nil
}

func ValidateShapefile(url string) error {
	if filepath.Ext(url) != ".shp" {
		return customerrors.WrapInputError(errors.New("shapefile extension must be .shp"))
	}

	return nil
}

func ValidateShapefileDirectory(dir string) error {
	if len(dir) == 0 {
		return customerrors.WrapInputError(errors.New("empty directory path"))
	}

	// Note: In a real implementation, you might want to check if the directory exists
	// and contains at least one shapefile, but for simplicity we'll just check if
	// the path is not empty.

	return nil
}

func ValidateCSV(url string) error {
	if filepath.Ext(url) != ".csv" {
		return customerrors.WrapInputError(errors.New("csv file extension must be .csv"))
	}

	return nil
}

func ValidateGeoTIFF(url string) error {
	ext := filepath.Ext(url)
	if ext != ".tif" && ext != ".tiff" {
		return customerrors.WrapInputError(errors.New("geotiff file extension must be .tif or .tiff"))
	}

	return nil
}

func ValidateWorldImage(url string) error {
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

func ValidateImageMosaic(dir string) error {
	if len(dir) == 0 {
		return customerrors.WrapInputError(errors.New("empty directory path for image mosaic"))
	}

	// Note: In a real implementation, you might want to check if the directory exists
	// and contains valid image files, but for simplicity we'll just check if
	// the path is not empty.

	return nil
}

func ValidateArcGrid(url string) error {
	ext := filepath.Ext(url)
	if ext != ".asc" && ext != ".grd" {
		return customerrors.WrapInputError(errors.New("arcgrid file extension must be .asc or .grd"))
	}

	return nil
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
