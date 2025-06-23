package validator

import (
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataStoreValidator_PostGIS(t *testing.T) {
	tests := []struct {
		name         string
		storeName    string
		wantErr      bool
		errorMessage string
	}{
		{
			name:      "Valid PostGIS store name",
			storeName: testdata.DatastorePostgis,
			wantErr:   false,
		},
		{
			name:         "Empty PostGIS store name",
			storeName:    "",
			wantErr:      true,
			errorMessage: "empty store name",
		},
		{
			name:         "Invalid PostGIS store name with special characters",
			storeName:    "invalid!@#$%^&*()",
			wantErr:      true,
			errorMessage: "name can only contain alphanumerical characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsv := DataStoreValidator{}
			err := dsv.PostGIS(tt.storeName)

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

func TestDataStoreValidator_GeoPackage(t *testing.T) {
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
			dsv := DataStoreValidator{}
			err := dsv.GeoPackage(tt.url)

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

func TestDataStoreValidator_Shapefile(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid Shapefile URL",
			url:     "/path/to/file.shp",
			wantErr: false,
		},
		{
			name:         "Invalid Shapefile extension",
			url:          "/path/to/file.txt",
			wantErr:      true,
			errorMessage: "shapefile extension must be .shp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsv := DataStoreValidator{}
			err := dsv.Shapefile(tt.url)

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

func TestDataStoreValidator_ShapefileDirectory(t *testing.T) {
	tests := []struct {
		name         string
		dir          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid Shapefile directory",
			dir:     "/path/to/directory",
			wantErr: false,
		},
		{
			name:         "Empty Shapefile directory",
			dir:          "",
			wantErr:      true,
			errorMessage: "empty directory path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsv := DataStoreValidator{}
			err := dsv.ShapefileDirectory(tt.dir)

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

func TestDataStoreValidator_CSV(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid CSV URL",
			url:     "/path/to/file.csv",
			wantErr: false,
		},
		{
			name:         "Invalid CSV extension",
			url:          "/path/to/file.txt",
			wantErr:      true,
			errorMessage: "csv file extension must be .csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsv := DataStoreValidator{}
			err := dsv.CSV(tt.url)

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

func TestDataStoreValidator_WebFeatureService(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "Valid WebFeatureService URL",
			url:     "http://localhost:8080/geoserver",
			wantErr: false,
		},
		{
			name:         "Empty WebFeatureService URL",
			url:          "",
			wantErr:      true,
			errorMessage: "empty wfs url",
		},
		{
			name:         "Blank WebFeatureService URL",
			url:          "   ",
			wantErr:      true,
			errorMessage: "empty wfs url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsv := DataStoreValidator{}
			err := dsv.WebFeatureService(tt.url)

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
