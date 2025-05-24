//go:build mocks

package requester

import (
	"bytes"
	"encoding/xml"
	"github.com/canghel3/go-geoserver/internal/mocks"
	testdata2 "github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/models/wms"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"testing"
)

const (
	CAPABILITIES_1_3_0 = "../../testdata/wms/capabilities_1_3_0.xml"
)

func TestWMSRequester_GetCapabilitiesRequester(t *testing.T) {
	t.Run("1.3.0", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		content, err := os.ReadFile(CAPABILITIES_1_3_0)
		assert.NoError(t, err)

		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(content)),
			Header:     make(http.Header),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		wmsRequester := WMSRequester{data: testdata2.GeoserverInfo(mockClient)}
		capabilities, err := wmsRequester.GetCapabilities(string(wms.VERSION_1_3_0))

		assert.NoError(t, err)
		assert.NotNil(t, capabilities)

		var expectedCapabilities *wms.Capabilities1_3_0
		err = xml.Unmarshal(content, &expectedCapabilities)
		assert.NoError(t, err)

		assert.Equal(t, expectedCapabilities, capabilities)
	})
}
