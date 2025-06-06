package actions

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/about"
)

type About struct {
	requester *requester.Requester
}

func NewAboutAction(data *internal.GeoserverData) *About {
	return &About{
		requester: requester.NewRequester(data),
	}
}

func (a *About) Manifest() (*about.Manifest, error) {
	return a.requester.About().Manifest()
}

func (a *About) Version() (*about.Version, error) {
	return a.requester.About().Version()
}

func (a *About) Status() (*about.Status, error) {
	return a.requester.About().Status()
}

func (a *About) SystemStatus() (*about.Metrics, error) {
	return a.requester.About().SystemStatus()
}
