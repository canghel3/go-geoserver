package options

import (
	"github.com/canghel3/go-geoserver/internal"
	"net/http"
)

type GeoserverClientOption func(*internal.GeoserverData)

var Client GeoserverClientOptionsGenerator

type GeoserverClientOptionsGenerator struct{}

func (gco GeoserverClientOptionsGenerator) DataDir(datadir string) GeoserverClientOption {
	return func(i *internal.GeoserverData) {
		i.DataDir = datadir
	}
}

func (gco GeoserverClientOptionsGenerator) HttpClient(client http.Client) GeoserverClientOption {
	return func(i *internal.GeoserverData) {
		i.Client = client
	}
}
