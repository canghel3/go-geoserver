package client

import (
	"github.com/canghel3/go-geoserver/handler"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"net/http"
)

type GeoserverClient struct {
	info      *internal.GeoserverInfo
	requester *requester.Requester
}

func New(url, username, password, datadir string) *GeoserverClient {
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
func (s *GeoserverClient) Workspace(name string) *handler.WorkspaceServiceSelector {
	return handler.NewWorkspaceHandler(s.info.Clone()).Use(name)
}

func (s *GeoserverClient) WMS() *handler.WMS {
	return handler.NewWMSHandler(s.info.Clone())
}

//TODO: implement wms, wfs and others
