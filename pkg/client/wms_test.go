package client

import (
	"bytes"
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/types"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/tiff"
	"image/gif"
	"image/jpeg"
	"image/png"
	"path/filepath"
	"testing"
)

func TestWMS_GetMap(t *testing.T) {
	addTestWorkspace(t)
	addTestDataStore(t, types.GeoPackage)
	addTestFeatureType(t, types.GeoPackage)

	t.Run("Png", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
	})

	t.Run("Png8", func(t *testing.T) {

		t.Run("1.1.1", func(t *testing.T) {
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
	})

	t.Run("Jpeg", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageJpeg), buf.Bytes())
		})
	})

	t.Run("JpegPng", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageJpegPng), buf.Bytes())
		})
	})

	t.Run("JpegPng8", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageJpegPng8), buf.Bytes())
		})
	})

	t.Run("Gif", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = gif.Encode(&buf, img, nil)
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageGif), buf.Bytes())
		})
	})

	t.Run("Tiff", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = tiff.Encode(&buf, img, nil)
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageTiff), buf.Bytes())
		})
	})

	t.Run("Tiff8", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = tiff.Encode(&buf, img, nil)
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageTiff8), buf.Bytes())
		})
	})

	t.Run("GeoTiff", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = tiff.Encode(&buf, img, nil)
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageGeoTiff), buf.Bytes())
		})
	})

	t.Run("GeoTiff8", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
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
			err = tiff.Encode(&buf, img, nil)
			assert.NoError(t, err)

			writeFile(t, filepath.Join(imagesTestDataDir, testdata.ImageGeoTiff8), buf.Bytes())
		})
	})
}
