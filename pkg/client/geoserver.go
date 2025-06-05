package client

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/actions"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/wms"
)

type GeoserverClient struct {
	info *internal.GeoserverData
}

func NewGeoserverClient(url, username, password string, options ...options.GeoserverClientOption) *GeoserverClient {
	gc := new(GeoserverClient)
	gc.info = new(internal.GeoserverData)
	gc.info.Connection.URL = url
	gc.info.Connection.Credentials.Username = username
	gc.info.Connection.Credentials.Password = password

	for _, option := range options {
		option(gc.info)
	}

	return gc
}

type GeoserverClientOption func(*GeoserverClient)

// Workspaces displays available actions inside a workspace.
func (s *GeoserverClient) Workspaces() *actions.Workspaces {
	return actions.NewWorkspaceActions(s.info.Clone())
}

// Workspace is shorthand for Workspaces().Use(name)
func (s *GeoserverClient) Workspace(name string) *actions.Workspace {
	return actions.NewWorkspaceActions(s.info.Clone()).Use(name)
}

func (s *GeoserverClient) WMS(version wms.WMSVersion) *actions.WMS {
	return actions.NewWMSActions(s.info.Clone(), version)
}

//TODO: implement wfs and others
