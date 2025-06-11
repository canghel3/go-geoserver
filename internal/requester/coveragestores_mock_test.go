package requester

import (
	"bytes"
	"errors"
	"fmt"
	mocks "github.com/canghel3/go-geoserver/internal/mock"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

const coverageStoreFile = "../testdata/coveragestores/coveragestore.json"

func TestCoverageStoreRequester_Create(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageStoreRequester.Create(nil)
		assert.NoError(t, err)
	})

	t.Run("201 Created", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusCreated,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageStoreRequester.Create(nil)
		assert.NoError(t, err)
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

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		err := coverageStoreRequester.Create(nil)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		err := coverageStoreRequester.Create(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestCoverageStoreRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(coverageStoreFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		store, err := coverageStoreRequester.Get(testdata.CoverageStoreGeoTiff)
		assert.NoError(t, err)
		assert.NotNil(t, store)
		assert.Equal(t, "GeoTIFF", store.Type)
		assert.Equal(t, false, store.Default)
	})

	t.Run("201 Created", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(coverageStoreFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		store, err := coverageStoreRequester.Get(testdata.CoverageStoreGeoTiff)
		assert.NoError(t, err)
		assert.NotNil(t, store)
		assert.Equal(t, "GeoTIFF", store.Type)
		assert.Equal(t, false, store.Default)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("not found")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.NotFoundError
		_, err := coverageStoreRequester.Get(testdata.CoverageStoreGeoTiff)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, fmt.Sprintf("coveragestore %s not found", testdata.CoverageStoreGeoTiff))
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

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		_, err := coverageStoreRequester.Get(testdata.CoverageStoreGeoTiff)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := coverageStoreRequester.Get(testdata.CoverageStoreGeoTiff)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestCoverageStoreRequester_Delete(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageStoreRequester.Delete(testdata.CoverageStoreGeoTiff, true)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("not found")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.NotFoundError
		err := coverageStoreRequester.Delete(testdata.CoverageStoreGeoTiff, true)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, fmt.Sprintf("coveragestore %s not found", testdata.CoverageStoreGeoTiff))
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

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		err := coverageStoreRequester.Delete(testdata.CoverageStoreGeoTiff, true)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageStoreRequester := &CoverageStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		err := coverageStoreRequester.Delete(testdata.CoverageStoreGeoTiff, true)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
