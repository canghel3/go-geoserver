package client

import (
	"crypto/tls"
	"github.com/canghel3/go-geoserver/pkg/options"
	"net/http"
)

func ExampleNewGeoserverClient() {
	// Basic initialization
	client := NewGeoserverClient(
		"http://localhost:8080",
		"admin",
		"geoserver")

	// Initialize with DataDir option
	client = NewGeoserverClient(
		"http://localhost:8080",
		"admin",
		"geoserver",
		options.Client.DataDir("/data"),
	)

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
