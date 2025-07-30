package client

import (
	"crypto/tls"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func ExampleNewGeoserverClient() {
	// Basic initialization
	client := NewGeoserverClient(
		"http://localhost:8080",
		"admin",
		"geoserver")

	// Initialize with HttpClient option
	client = NewGeoserverClient(
		"http://localhost:8080",
		"admin",
		"geoserver",
		options.Client.HttpClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}),
	)

	_ = client
}

func TestNewGeoserverClient(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		client := NewGeoserverClient(
			"http://localhost:8080",
			"admin",
			"geoserver")

		assert.NotNil(t, client)
		assert.Equal(t, client.data.Connection.URL, "http://localhost:8080")
	})

	t.Run("Options", func(t *testing.T) {
		t.Run("HttpClient", func(t *testing.T) {
			httpClient := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}

			// Initialize with HttpClient option
			client := NewGeoserverClient(
				"http://localhost:8080",
				"admin",
				"geoserver",
				options.Client.HttpClient(httpClient),
			)

			assert.NotNil(t, client)
			assert.Equal(t, httpClient, client.data.Client)
		})
	})
}
