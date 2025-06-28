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
	getSingleDataStoreResponse = "../testdata/datastores/getsingle.json"
	getAllDataStoresResponse   = "../testdata/datastores/getall.json"
)

func TestDataStoreRequester_Create(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Create(nil)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Create(nil)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Create(nil)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Create(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Create(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestDataStoreRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(getSingleDataStoreResponse)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		ds, err := dataStoreRequester.Get(testdata.DatastorePostgis)
		assert.NoError(t, err)
		assert.NotNil(t, ds)
		assert.Equal(t, "Shapefile", ds.Type)
		assert.Equal(t, "2025-05-15 19:09:10.716 UTC", ds.DateCreated)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := dataStoreRequester.Get(testdata.DatastorePostgis)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DatastorePostgis))
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := dataStoreRequester.Get(testdata.DatastorePostgis)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := dataStoreRequester.Get(testdata.DatastorePostgis)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := dataStoreRequester.Get(testdata.DatastorePostgis)
		assert.Error(t, err)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := dataStoreRequester.Get(testdata.DatastorePostgis)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestDataStoreRequester_GetAll(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		t.Run("Get All - 3 stores", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			content, err := testdata.Read(getAllDataStoresResponse)
			assert.NoError(t, err)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(content)),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

			ds, err := dataStoreRequester.GetAll()
			assert.NoError(t, err)
			assert.NotNil(t, ds)
			assert.Len(t, ds.Entries, 2)
		})

		t.Run("No DataStores", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("{\"dataStores\": \"\"}")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

			ds, err := dataStoreRequester.GetAll()
			assert.NoError(t, err)
			assert.NotNil(t, ds)
			assert.Nil(t, ds.Entries)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := dataStoreRequester.GetAll()
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := dataStoreRequester.GetAll()
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := dataStoreRequester.GetAll()
		assert.Error(t, err)
		assert.EqualError(t, err, "unexpected end of JSON input")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := dataStoreRequester.GetAll()
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestDataStoreRequester_Update(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Update(testdata.DatastorePostgis, nil)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Update(testdata.DatastorePostgis, nil)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Update(testdata.DatastorePostgis, nil)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DatastorePostgis))
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Update(testdata.DatastorePostgis, nil)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Update(testdata.DatastorePostgis, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestDataStoreRequester_Delete(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Delete(testdata.DatastorePostgis, true)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Delete(testdata.DatastorePostgis, true)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DatastorePostgis))
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Delete(testdata.DatastorePostgis, true)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Delete(testdata.DatastorePostgis, true)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Delete(testdata.DatastorePostgis, true)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestDataStoreRequester_Reset(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		err := dataStoreRequester.Reset(testdata.DatastorePostgis)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		err := dataStoreRequester.Reset(testdata.DatastorePostgis)
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		err := dataStoreRequester.Reset(testdata.DatastorePostgis)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("datastore %s not found", testdata.DatastorePostgis))
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

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		err := dataStoreRequester.Reset(testdata.DatastorePostgis)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}
		err := dataStoreRequester.Reset(testdata.DatastorePostgis)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
