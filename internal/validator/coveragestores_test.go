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
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_EHdr(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid EHdr filepath",
			dir:     "/path/to/directory/file.bil",
			wantErr: false,
		},
		{
			name:         "Empty EHdr filepath",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid EHdr extension",
			dir:          "/path/to/directory/file.csv",
			wantErr:      true,
			errorMessage: "EHdr file extension must be .bil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.EHdr(tt.dir)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_ENVIHdr(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid ENVIHdr filepath",
			dir:     "/path/to/directory/file.dat",
			wantErr: false,
		},
		{
			name:         "Empty ENVIHdr filepath",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid ENVIHdr extension",
			dir:          "/path/to/directory/file.csv",
			wantErr:      true,
			errorMessage: "ENVIHdr file extension must be .dat",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.ENVIHdr(tt.dir)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_ERDASImg(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid ERDASImg filepath",
			dir:     "/path/to/directory/file.img",
			wantErr: false,
		},
		{
			name:         "Empty ERDASImg filepath",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid ERDASImg extension",
			dir:          "/path/to/directory/file.csv",
			wantErr:      true,
			errorMessage: "ERDASImg file extension must be .img",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.ERDASImg(tt.dir)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
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
				assert.IsType(t, err, &customerrors.InputError{})
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
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_NITF(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid NITF filepath",
			dir:     "/path/to/directory/file.ntf",
			wantErr: false,
		},
		{
			name:         "Empty NITF filepath",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid NITF extension",
			dir:          "/path/to/directory/file.csv",
			wantErr:      true,
			errorMessage: "NITF file extension must be .ntf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.NITF(tt.dir)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_RST(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid RST filepath",
			dir:     "/path/to/directory/file.rst",
			wantErr: false,
		},
		{
			name:         "Empty RST filepath",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid RST extension",
			dir:          "/path/to/directory/file.csv",
			wantErr:      true,
			errorMessage: "RST file extension must be .rst",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.RST(tt.dir)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCoverageStoreValidator_VRT(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid VRT filepath",
			dir:     "/path/to/directory/file.vrt",
			wantErr: false,
		},
		{
			name:         "Empty VRT filepath",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty url",
		},
		{
			name:         "Invalid VRT extension",
			dir:          "/path/to/directory/file.csv",
			wantErr:      true,
			errorMessage: "VRT file extension must be .vrt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := CoverageStoreValidator{}
			err := csv.VRT(tt.dir)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errorMessage)
				assert.IsType(t, err, &customerrors.InputError{})
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
