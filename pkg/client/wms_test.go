package client

import (
	"bytes"
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"github.com/stretchr/testify/assert"
	"image/png"
	"path/filepath"
	"testing"
)

func TestWMS_GetMap(t *testing.T) {
	addTestWorkspace(t)
	addTestDataStore(t, formats.GeoPackage)
	addTestFeatureType(t, formats.GeoPackage)

	t.Run("Png", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Png()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Png()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).Png()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("Png8", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Png8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Png8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).Png8()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("Jpeg", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Jpeg()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Jpeg()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).Jpeg()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("JpegPng", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).JpegPng()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).JpegPng()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).JpegPng()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("JpegPng8", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).JpegPng8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).JpegPng8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).JpegPng8()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("Gif", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Gif()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Gif()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).Gif()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("Tiff", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Tiff()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Tiff()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).Tiff()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("Tiff8", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Tiff8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Tiff8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).Tiff8()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("GeoTiff", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).GeoTiff()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).GeoTiff()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).GeoTiff()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})

	t.Run("GeoTiff8", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("From Client", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).GeoTiff8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})

			t.Run("From Workspace", func(t *testing.T) {
				img, err := geoclient.Workspace(testdata.Workspace).WMS(wms.Version111).GetMap(500, 500, []string{fmt.Sprintf(testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 270013.039,
					MinY: 840102.83,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).GeoTiff8()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				var buf bytes.Buffer
				err = png.Encode(&buf, img)
				assert.NoError(t, err)

				writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImagePng8), buf.Bytes())
			})
		})

		t.Run("Input Error", func(t *testing.T) {
			img, err := geoclient.WMS(wms.Version111).GetMap(500, 500, []string{testdata.FeatureTypeGeoPackage}, shared.BBOX{
				MinX: 264970.869,
				MaxX: 270013.039,
				MinY: 840102.83,
				MaxY: 845199.87,
				SRS:  "EPSG:27700",
			}).GeoTiff8()

			expectedError := fmt.Sprintf("unspecified workspace in layer name %[1]s. format the layer name as <workspace>:%[1]s", testdata.FeatureTypeGeoPackage)
			assert.Error(t, err)
			assert.Nil(t, img)
			assert.IsType(t, &customerrors.InputError{}, err)
			assert.EqualError(t, err, expectedError)
		})
	})
}
