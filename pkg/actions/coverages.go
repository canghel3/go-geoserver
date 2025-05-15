package actions

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type Coverages struct {
	data      *internal.GeoserverData
	requester *requester.Requester
}

func newCoverages(data *internal.GeoserverData) *Coverages {
	return &Coverages{
		data:      data,
		requester: requester.NewRequester(data),
	}
}
