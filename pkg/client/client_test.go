package client

import (
	"crypto/tls"
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
		Options.DataDir("/data"),
	)

	// Initialize with Client option
	client = NewGeoserverClient(
		"http://localhost:8080",
		"admin",
		"geoserver",
		Options.Client(http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}),
	)

	_ = client
}
