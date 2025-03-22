package main

import (
	"github.com/canghel3/go-geoserver/handler"
	"github.com/canghel3/go-geoserver/internal"
	"net/http"
)

type GeoserverClient struct {
	info *internal.GeoserverInfo
}

func NewGeoserverClient(url, username, password, datadir string) *GeoserverClient {
	return &GeoserverClient{
		info: &internal.GeoserverInfo{
			Client: &http.Client{},
			Connection: internal.GeoserverConnection{
				URL: url,
				Credentials: internal.GeoserverCredentials{
					Username: username,
					Password: password,
				},
			},
			DataDir: datadir,
		},
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

//TODO: implement wms, wfs and others
