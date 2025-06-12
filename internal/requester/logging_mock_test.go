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
	"strings"
	"testing"
)

const loggingFile = "../testdata/logging/logging.json"

func TestLoggingRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(loggingFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		loggingRequester := &LoggingRequester{data: testdata.GeoserverInfo(mockClient)}

		log, err := loggingRequester.Get()
		assert.NoError(t, err)
		assert.NotNil(t, log)
		assert.Equal(t, "logs/geoserver.log", log.Location)
		assert.Equal(t, "DEFAULT_LOGGING", log.Level)
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

		loggingRequester := &LoggingRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := loggingRequester.Get()
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		loggingRequester := &LoggingRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := loggingRequester.Get()
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestLoggingRequester_Put(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(nil),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		loggingRequester := &LoggingRequester{data: testdata.GeoserverInfo(mockClient)}

		err := loggingRequester.Put(nil)
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

		loggingRequester := &LoggingRequester{data: testdata.GeoserverInfo(mockClient)}

		err := loggingRequester.Put(nil)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		loggingRequester := &LoggingRequester{data: testdata.GeoserverInfo(mockClient)}
		err := loggingRequester.Put(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
