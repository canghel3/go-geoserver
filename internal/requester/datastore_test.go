//go:build mocks

package requester

import (
	"bytes"
	"github.com/canghel3/go-geoserver/internal/mocks"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestDataStoreRequester_Create(t *testing.T) {
	t.Run("201 CREATED", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusCreated,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		dataStoreRequester := &DataStoreRequester{data: testdata.GeoserverInfo(mockClient)}

		err := dataStoreRequester.Create(nil)
		assert.NoError(t, err)
	})
}
