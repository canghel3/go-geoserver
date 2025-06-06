package validator

import (
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoverageStoreValidator_ArcGrid(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid ArcGrid URL with .asc extension",
			url:     "/path/to/file.asc",
			wantErr: false,
		},
		{
			name:    "Valid ArcGrid URL with .grd extension",
			url:     "/path/to/file.grd",
			wantErr: false,
		},
		{
			name:    "Valid ArcGrid URL with .afd extension",
			url:     "/path/to/file.afd",
			wantErr: false,
		},
		{
			name:         "Empty ArcGrid URL",
			url:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid ArcGrid extension",
			url:          "/path/to/file.txt",
			wantErr:      true,
			errorMessage: "arc grid file extension must be .asc .grd or .afd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.ArcGrid(tt.url)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_GeoPackage(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid GeoPackage URL",
			url:     "/path/to/file.gpkg",
			wantErr: false,
		},
		{
			name:         "Empty GeoPackage URL",
			url:          "",
			wantErr:      true,
			errorMessage: "empty geopackage url",
		},
		{
			name:         "Invalid GeoPackage extension",
			url:          "/path/to/file.txt",
			wantErr:      true,
			errorMessage: "geopackage extension must be .gpkg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.GeoPackage(tt.url)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_GeoTIFF(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid GeoTIFF URL with .tif extension",
			url:     "/path/to/file.tif",
			wantErr: false,
		},
		{
			name:    "Valid GeoTIFF URL with .tiff extension",
			url:     "/path/to/file.tiff",
			wantErr: false,
		},
		{
			name:         "Empty GeoTIFF URL",
			url:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid GeoTIFF extension",
			url:          "/path/to/file.txt",
			wantErr:      true,
			errorMessage: "geotiff file extension must be .tif or .tiff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.GeoTIFF(tt.url)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_WorldImage(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid WorldImage URL with .png extension",
			url:     "/path/to/file.png",
			wantErr: false,
		},
		{
			name:    "Valid WorldImage URL with .jpg extension",
			url:     "/path/to/file.jpg",
			wantErr: false,
		},
		{
			name:    "Valid WorldImage URL with .jpeg extension",
			url:     "/path/to/file.jpeg",
			wantErr: false,
		},
		{
			name:    "Valid WorldImage URL with .gif extension",
			url:     "/path/to/file.gif",
			wantErr: false,
		},
		{
			name:    "Valid WorldImage URL with .tif extension",
			url:     "/path/to/file.tif",
			wantErr: false,
		},
		{
			name:    "Valid WorldImage URL with .tiff extension",
			url:     "/path/to/file.tiff",
			wantErr: false,
		},
		{
			name:    "Valid WorldImage URL with .bmp extension",
			url:     "/path/to/file.bmp",
			wantErr: false,
		},
		{
			name:         "Empty WorldImage URL",
			url:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid WorldImage extension",
			url:          "/path/to/file.txt",
			wantErr:      true,
			errorMessage: "worldimage file must have a valid image extension (.png, .jpg, .jpeg, .gif, .tif, .tiff, .bmp)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.WorldImage(tt.url)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_ImageMosaic(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid ImageMosaic directory",
			dir:     "/path/to/directory",
			wantErr: false,
		},
		{
			name:         "Empty ImageMosaic directory",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty directory path for image mosaic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.ImageMosaic(tt.dir)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_ImagePyramid(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid ImagePyramid directory",
			dir:     "/path/to/directory",
			wantErr: false,
		},
		{
			name:         "Empty ImagePyramid directory",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty directory path for image pyramid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.ImagePyramid(tt.dir)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)

				var inputError *customerrors.InputError
				assert.ErrorAs(t, err, &inputError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Tests for the not implemented methods
func TestCoverageStoreValidator_NotImplementedMethods(t *testing.T) {
	csv := CoverageStoreValidator{}

	t.Run("AIG", func(t *testing.T) {
		err := csv.AIG("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("DTED", func(t *testing.T) {
		err := csv.DTED("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("EHdr", func(t *testing.T) {
		err := csv.EHdr("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("ENVIHdr", func(t *testing.T) {
		err := csv.ENVIHdr("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("ERDASImg", func(t *testing.T) {
		err := csv.ERDASImg("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("NITF", func(t *testing.T) {
		err := csv.NITF("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("RPFTOC", func(t *testing.T) {
		err := csv.RPFTOC("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("RST", func(t *testing.T) {
		err := csv.RST("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("SRP", func(t *testing.T) {
		err := csv.SRP("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("VRT", func(t *testing.T) {
		err := csv.VRT("test")
		assert.Error(t, err)
		assert.EqualError(t, err, "not implemented")
	})
}
