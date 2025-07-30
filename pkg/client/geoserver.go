package client

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/actions"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"net/http"
)

type GeoserverClient struct {
	data internal.GeoserverData
}

func NewGeoserverClient(url, username, password string, options ...options.GeoserverClientOption) GeoserverClient {
	gc := new(GeoserverClient)
	gc.data = internal.GeoserverData{
		Client: http.DefaultClient,
		Connection: internal.GeoserverConnection{
			URL: url,
			Credentials: internal.GeoserverCredentials{
				Username: username,
				Password: password,
			},
		},
	}

	for _, option := range options {
		option(&gc.data)
	}

	return *gc
}

type GeoserverClientOption func(*GeoserverClient)

func (gc GeoserverClient) About() actions.About {
	return actions.NewAboutAction(gc.data.Clone())
}

func (gc GeoserverClient) Fonts() actions.Fonts {
	return actions.NewFonts(gc.data.Clone())
}

// Workspaces displays available actions inside a workspace.
func (gc GeoserverClient) Workspaces() actions.Workspaces {
	return actions.NewWorkspaceActions(gc.data.Clone())
}

// Workspace is shorthand for Workspaces().Use(name)
func (gc GeoserverClient) Workspace(name string) actions.Workspace {
	return actions.NewWorkspaceActions(gc.data.Clone()).Use(name)
}

func (gc GeoserverClient) WMS(version wms.WMSVersion) actions.WMS {
	return actions.NewWMSActions(gc.data.Clone(), version)
}

func (gc GeoserverClient) LayerGroups() actions.LayerGroups {
	return actions.NewLayerGroup(gc.data.Clone())
}

func (gc GeoserverClient) Logging() actions.Logging {
	return actions.NewLoggingActions(gc.data.Clone())
}

func (gc GeoserverClient) GeoWebCache() actions.GeoWebCache {
	return actions.NewGeoWebCache(gc.data.Clone())
}
