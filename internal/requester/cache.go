package requester

import "github.com/canghel3/go-geoserver/internal"

type GeoWebCacheRequester struct {
	data internal.GeoserverData
}

func NewGeoWebCacheRequester(data internal.GeoserverData) GeoWebCacheRequester {
	return GeoWebCacheRequester{data: data}
}
