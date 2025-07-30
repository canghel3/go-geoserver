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

const (
	coverageFile  = "../testdata/coverages/coverage.json"
	coveragesFile = "../testdata/coverages/coverages.json"
)

func TestCoverageRequester_Create(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Create(testdata.CoverageStoreGeoTiff, nil)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Create(testdata.CoverageStoreGeoTiff, nil)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Create(testdata.CoverageStoreGeoTiff, nil)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Create(testdata.CoverageStoreGeoTiff, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}
		err := coverageRequester.Create(testdata.CoverageStoreGeoTiff, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestCoverageRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(coverageFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		cov, err := coverageRequester.Get(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.NoError(t, err)
		assert.NotNil(t, cov)
		assert.NotNil(t, cov.LatLonBoundingBox)
		assert.Equal(t, 4.472858110631908, cov.LatLonBoundingBox.MaxX)
		assert.Equal(t, "EPSG:4326", cov.LatLonBoundingBox.CRS.Value)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := coverageRequester.Get(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coverage %s not found", testdata.CoverageGeoTiffName))
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := coverageRequester.Get(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := coverageRequester.Get(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("{")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := coverageRequester.Get(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := coverageRequester.Get(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestCoverageRequester_GetAll(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(coveragesFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		cov, err := coverageRequester.GetAll(testdata.CoverageStoreGeoTiff)
		assert.NoError(t, err)
		assert.NotNil(t, cov)
		assert.Len(t, cov.Entries, 1)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := coverageRequester.GetAll(testdata.CoverageStoreGeoTiff)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := coverageRequester.GetAll(testdata.CoverageStoreGeoTiff)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("{")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := coverageRequester.GetAll(testdata.CoverageStoreGeoTiff)
		assert.Error(t, err)
		assert.EqualError(t, err, "unexpected end of JSON input")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := coverageRequester.Get(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestCoverageRequester_Delete(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Delete(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, true)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Delete(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coverage %s not found", testdata.CoverageGeoTiffName))
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Delete(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Delete(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, true)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}
		err := coverageRequester.Delete(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, true)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestCoverageRequester_Update(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Update(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, nil)
		assert.NoError(t, err)
	})

	t.Run("201 Created", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Update(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, nil)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Update(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, nil)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coverage %s not found", testdata.CoverageGeoTiffName))
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Update(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, nil)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}
		err := coverageRequester.Update(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestCoverageRequester_Reset(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Reset(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Reset(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Reset(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("coverage %s not found", testdata.CoverageGeoTiffName))
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

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}

		err := coverageRequester.Reset(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		coverageRequester := &CoverageRequester{data: testdata.GeoserverInfo(mockClient)}
		err := coverageRequester.Reset(testdata.CoverageStoreGeoTiff, testdata.CoverageGeoTiffName)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
