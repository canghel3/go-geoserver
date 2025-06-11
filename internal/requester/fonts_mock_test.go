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

const fontsFile = "../testdata/fonts/fonts.json"

func TestFontsRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(fontsFile)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		fontsRequester := &FontsRequester{data: testdata.GeoserverInfo(mockClient)}

		font, err := fontsRequester.Get()
		assert.NoError(t, err)
		assert.NotNil(t, font)
		assert.Equal(t, 13, len(font.Fonts))
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

		fontsRequester := &FontsRequester{data: testdata.GeoserverInfo(mockClient)}

		var geoserverError *customerrors.GeoserverError
		_, err := fontsRequester.Get()
		assert.Error(t, err)
		assert.ErrorAs(t, err, &geoserverError)
		assert.EqualError(t, err, "received status code 500 from geoserver: some error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		fontsRequester := &FontsRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := fontsRequester.Get()
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
