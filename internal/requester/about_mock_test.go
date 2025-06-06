package requester

import (
	"bytes"
	"errors"
	mocks "github.com/canghel3/go-geoserver/internal/mock"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

const manifestFile = "../testdata/about/manifest.json"

func TestAboutRequester_Manifest(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(manifestFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		aboutRequester := &AboutRequester{data: testdata.GeoserverInfo(mockClient)}

		manifest, err := aboutRequester.Manifest()
		assert.NoError(t, err)
		assert.NotNil(t, manifest)
		assert.NotNil(t, manifest.Resources)
		assert.Len(t, manifest.Resources, 2)
		assert.Equal(t, float64(1), manifest.Resources[0].ManifestVersion)
		assert.Equal(t, float64(1), manifest.Resources[1].ManifestVersion)
		assert.Equal(t, int64(1535553538488), manifest.Resources[0].BndLastModified)
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader([]byte("some error"))),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		aboutRequester := &AboutRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		manifest, err := aboutRequester.Manifest()
		assert.Nil(t, manifest)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		aboutRequester := &AboutRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := aboutRequester.Manifest()
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
