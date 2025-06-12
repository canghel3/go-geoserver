package client

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/types"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWMS_GetMap(t *testing.T) {
	err := addTestWorkspace()
	assert.NoError(t, err)

	err = addTestDataStore(types.GeoPackage)
	assert.NoError(t, err)

	err = addTestVectorLayer(types.GeoPackage)
	assert.NoError(t, err)

	t.Run("Version", func(t *testing.T) {
		t.Run("1.1.1", func(t *testing.T) {
			t.Run("PNG", func(t *testing.T) {
				img, err := geoclient.WMS(wms.Version111).GetMap(759, 768, []string{fmt.Sprintf("%s:%s", testdata.Workspace, testdata.FeatureTypeGeoPackage)}, shared.BBOX{
					MinX: 264970.869,
					MaxX: 840102.83,
					MinY: 270013.039,
					MaxY: 845199.87,
					SRS:  "EPSG:27700",
				}).Png()
				assert.NoError(t, err)
				assert.NotNil(t, img)
				assert.NotNil(t, img.Bounds())

				//var buf bytes.Buffer
				//png.Encode(&buf, img)
				//fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
			})
		})

		t.Run("1.3.0", func(t *testing.T) {

		})

	})

}
