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

const featureTypeFile = "../testdata/featuretypes/featuretype.json"

func TestFeatureTypeRequester_Create(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Create(testdata.DatastorePostgis, nil)
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

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Create(testdata.DatastorePostgis, nil)
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

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		err := featureTypeRequester.Create(testdata.DatastorePostgis, nil)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Create(testdata.DatastorePostgis, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestFeatureTypeRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(featureTypeFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		ft, err := featureTypeRequester.Get(testdata.DatastorePostgis, testdata.FeatureTypePostgis)
		assert.NoError(t, err)
		assert.NotNil(t, ft)
		assert.Equal(t, "EPSG:27700", ft.Srs)
		assert.Equal(t, -4.253489380362922, ft.LatLonBoundingBox.MinX)
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

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.NotFoundError
		_, err := featureTypeRequester.Get(testdata.DatastorePostgis, testdata.FeatureTypePostgis)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, fmt.Sprintf("featuretype %s not found", testdata.FeatureTypePostgis))
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

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		_, err := featureTypeRequester.Get(testdata.DatastorePostgis, testdata.FeatureTypePostgis)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := featureTypeRequester.Get(testdata.DatastorePostgis, testdata.FeatureTypePostgis)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestFeatureTypeRequester_Delete(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Delete(testdata.DatastorePostgis, testdata.FeatureTypePostgis, true)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.NotFoundError
		err := featureTypeRequester.Delete(testdata.DatastorePostgis, testdata.FeatureTypePostgis, true)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, fmt.Sprintf("featuretype %s not found", testdata.FeatureTypePostgis))
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

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		err := featureTypeRequester.Delete(testdata.DatastorePostgis, testdata.FeatureTypePostgis, true)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Delete(testdata.DatastorePostgis, testdata.FeatureTypePostgis, true)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestFeatureTypeRequester_Update(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Update(testdata.DatastorePostgis, testdata.FeatureTypePostgis, nil)
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

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Update(testdata.DatastorePostgis, testdata.FeatureTypePostgis, nil)
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

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		err := featureTypeRequester.Update(testdata.DatastorePostgis, testdata.FeatureTypePostgis, nil)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		featureTypeRequester := &FeatureTypeRequester{data: testdata.GeoserverInfo(mockClient)}

		err := featureTypeRequester.Update(testdata.DatastorePostgis, testdata.FeatureTypePostgis, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
