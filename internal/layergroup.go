package internal

import (
	"github.com/canghel3/go-geoserver/pkg/misc"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type GroupWrapper struct {
	Group Group `json:"layerGroup"`
}

type Group struct {
	Name         string                      `json:"name"`
	Mode         string                      `json:"mode"`
	Title        string                      `json:"title"`
	Workspace    workspace.WorkspaceCreation `json:"workspace"`
	Publishables Publishables                `json:"publishables"`
	Bounds       misc.BoundingBox            `json:"bounds"`
	Keywords     *misc.Keywords              `json:"keywords,omitempty"`
	Styles       GroupStyles                 `json:"styles"`
}

type Publishables struct {
	Published []Published `json:"published"`
}

type Published struct {
	Type string `json:"@type"`
	Name string `json:"name"`
	Link string `json:"href"`
}

type GroupStyles struct {
	Style []misc.Style `json:"style"`
}

type GetGroupWrapper struct {
	Group GetGroup `json:"layerGroup"`
}

type GetGroup struct {
	//even though a layer name can be sent as a string number, geoserver parses it to an actual number and returns it because it is stupids
	Name         string                      `json:"name"`
	Mode         string                      `json:"mode"`
	Title        string                      `json:"title"`
	Workspace    workspace.WorkspaceCreation `json:"workspace"`
	Publishables Publishables                `json:"publishables"`
	Bounds       misc.BoundingBox            `json:"bounds"`
	Keywords     *misc.Keywords              `json:"keywords,omitempty"`
	Styles       any                         `json:"styles,omitempty"`
}
