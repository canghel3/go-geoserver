package client

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/handler"
	"net/http"
)

type GeoserverClient struct {
	info *internal.GeoserverData
}

func NewGeoserverClient(url, username, password string, options ...GeoserverClientOption) *GeoserverClient {
	gc := new(GeoserverClient)
	gc.info = new(internal.GeoserverData)
	gc.info.Client = &http.Client{}
	gc.info.Connection.URL = url
	gc.info.Connection.Credentials.Username = username
	gc.info.Connection.Credentials.Password = password

	for _, option := range options {
		option(gc)
	}

	return gc
}

type GeoserverClientOption func(*GeoserverClient)

var Options geoserverClientOptions

type geoserverClientOptions struct{}

func (gco geoserverClientOptions) DataDir(datadir string) GeoserverClientOption {
	return func(c *GeoserverClient) {
		c.info.DataDir = datadir
	}
}

func (gco geoserverClientOptions) Client(client internal.HTTPClient) GeoserverClientOption {
	return func(c *GeoserverClient) {
		c.info.Client = client
	}
}

// Workspaces displays available actions inside a workspace.
func (s *GeoserverClient) Workspaces() *handler.Workspaces {
	return handler.NewWorkspaceHandler(s.info.Clone())
}

// Workspace is shorthand for Workspaces().Use(name)
func (s *GeoserverClient) Workspace(name string) *handler.Workspace {
	return handler.NewWorkspaceHandler(s.info.Clone()).Use(name)
}

func (s *GeoserverClient) WMS() *handler.WMS {
	return handler.NewWMSHandler(s.info.Clone())
}

func (s *GeoserverClient) CoverageStores() *handler.CoverageStores {
	return handler.NewCoverageStoreHandler(s.info.Clone())
}

// CoverageStore is shorthand for CoverageStores().Use(name)
func (s *GeoserverClient) CoverageStore(name string) *handler.CoverageStore {
	return handler.NewCoverageStoreHandler(s.info.Clone()).Use(name)
}

//TODO: implement wfs and others
