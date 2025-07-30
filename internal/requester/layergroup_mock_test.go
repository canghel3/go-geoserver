package requester

import (
	"bytes"
	"errors"
	"fmt"
	mocks "github.com/canghel3/go-geoserver/internal/mock"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/layers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

const (
	getLayerGroupResponse = "../testdata/layers/getgroup.json"
)

func TestLayerGroupRequester_Create(t *testing.T) {
	t.Run("201 Created", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusCreated,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Create(nil)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Create(nil)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Create(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Create(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestLayerGroupRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(getLayerGroupResponse)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		group, err := lgr.Get(testdata.LayerGroupName)
		assert.NoError(t, err)
		assert.NotNil(t, group)
		assert.Equal(t, layers.ModeSingle, group.Mode)
		assert.Equal(t, 2, len(group.Publishables.Entries))
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		group, err := lgr.Get(testdata.LayerGroupName)
		assert.Error(t, err)
		assert.Nil(t, group)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("layer group %s not found", testdata.LayerGroupName))
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		group, err := lgr.Get(testdata.LayerGroupName)
		assert.Error(t, err)
		assert.Nil(t, group)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		group, err := lgr.Get(testdata.LayerGroupName)
		assert.Error(t, err)
		assert.Nil(t, group)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		group, err := lgr.Get(testdata.LayerGroupName)
		assert.Error(t, err)
		assert.Nil(t, group)
		assert.EqualError(t, err, "unexpected end of JSON input")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		group, err := lgr.Get(testdata.LayerGroupName)
		assert.Error(t, err)
		assert.Nil(t, group)
		assert.EqualError(t, err, "client error")
	})
}

func TestLayerGroupRequester_Delete(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Delete(testdata.LayerGroupName)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Delete(testdata.LayerGroupName)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Delete(testdata.LayerGroupName)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Delete(testdata.LayerGroupName)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestLayerGroupRequester_Update(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Update(testdata.LayerGroupName, nil)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Update(testdata.LayerGroupName, nil)
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

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Update(testdata.LayerGroupName, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		lgr := &LayerGroupRequester{data: testdata.GeoserverInfo(mockClient)}

		err := lgr.Update(testdata.LayerGroupName, nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
