package requester

import (
	"errors"
	mocks "github.com/canghel3/go-geoserver/internal/mock"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestWMSRequester_GetCapabilities(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("some content")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

		capabilities, err := wmsRequester.GetCapabilities(wms.Version130)
		assert.NoError(t, err)
		assert.NotNil(t, capabilities)
		assert.Equal(t, "some content", string(capabilities))
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("some error")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := wmsRequester.GetCapabilities(wms.Version130)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := wmsRequester.GetCapabilities(wms.Version130)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestWMSRequester_GetMap(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		t.Run("1.1.0", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("some content")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

			map_, err := wmsRequester.GetMap(0, 0, nil, shared.BBOX{}, wms.Version110, wms.PNG)
			assert.NoError(t, err)
			assert.NotNil(t, map_)
			assert.Equal(t, "some content", string(map_))
		})

		t.Run("1.1.1", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("some content")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

			map_, err := wmsRequester.GetMap(0, 0, nil, shared.BBOX{}, wms.Version111, wms.PNG)
			assert.NoError(t, err)
			assert.NotNil(t, map_)
			assert.Equal(t, "some content", string(map_))
		})

		t.Run("1.3.0", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("some content")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

			map_, err := wmsRequester.GetMap(0, 0, nil, shared.BBOX{}, wms.Version130, wms.PNG)
			assert.NoError(t, err)
			assert.NotNil(t, map_)
			assert.Equal(t, "some content", string(map_))
		})

		t.Run("With Options", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("some content")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

			map_, err := wmsRequester.GetMap(0, 0, nil, shared.BBOX{}, wms.Version130, wms.PNG, options.GetMap.Styles(nil))
			assert.NoError(t, err)
			assert.NotNil(t, map_)
			assert.Equal(t, "some content", string(map_))
		})
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("some error")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := wmsRequester.GetMap(0, 0, nil, shared.BBOX{}, wms.Version130, wms.PNG)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		wmsRequester := &WMSRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := wmsRequester.GetMap(0, 0, nil, shared.BBOX{}, wms.Version130, wms.PNG)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
