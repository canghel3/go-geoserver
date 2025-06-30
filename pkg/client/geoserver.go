package client

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/actions"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"net/http"
)

type GeoserverClient struct {
	info internal.GeoserverData
}

func NewGeoserverClient(url, username, password string, options ...options.GeoserverClientOption) *GeoserverClient {
	gc := new(GeoserverClient)
	gc.info = internal.GeoserverData{
		Client: &http.Client{},
		Connection: internal.GeoserverConnection{
			URL: url,
			Credentials: internal.GeoserverCredentials{
				Username: username,
				Password: password,
			},
		},
	}

	for _, option := range options {
		option(&gc.info)
	}

	return gc
}

type GeoserverClientOption func(*GeoserverClient)

func (s *GeoserverClient) About() actions.About {
	return actions.NewAboutAction(s.info.Clone())
}

func (s *GeoserverClient) Fonts() actions.Fonts {
	return actions.NewFonts(s.info.Clone())
}

// Workspaces displays available actions inside a workspace.
func (s *GeoserverClient) Workspaces() actions.Workspaces {
	return actions.NewWorkspaceActions(s.info.Clone())
}

// Workspace is shorthand for Workspaces().Use(name)
func (s *GeoserverClient) Workspace(name string) actions.Workspace {
	return actions.NewWorkspaceActions(s.info.Clone()).Use(name)
}

func (s *GeoserverClient) WMS(version wms.WMSVersion) actions.WMS {
	return actions.NewWMSActions(s.info.Clone(), version)
}

func (s *GeoserverClient) Logging() actions.Logging {
	return actions.NewLoggingActions(s.info.Clone())
}
