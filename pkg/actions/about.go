package actions

import (
	"errors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/about"
)

type About struct {
	requester *requester.Requester
}

func newAbout(data *internal.GeoserverData) *About {
	return &About{
		requester: requester.NewRequester(data),
	}
}

func (a *About) Manifest() (*about.Manifest, error) {
	return nil, errors.New("not implemented")
}

func (a *About) Status() (*about.Status, error) {
	return nil, errors.New("not implemented")
}
