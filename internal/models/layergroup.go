package models

import (
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type GroupWrapper struct {
	Group Group `json:"layerGroup"`
}

// Group TODO: implement custom unmarshaller in case geoserver responds with numbers for name and title
type Group struct {
	Name         string              `json:"name"`
	Mode         string              `json:"mode"`
	Title        *string             `json:"title,omitempty"`
	Workspace    *workspace.Creation `json:"workspace,omitempty"`
	Publishables Publishables        `json:"publishables"`
	Bounds       *shared.BoundingBox `json:"bounds,omitempty"`
	Keywords     *shared.Keywords    `json:"keywords,omitempty"`
	Styles       *GroupStyles        `json:"styles,omitempty"`
}

type GroupStyles struct {
	Style []shared.Style `json:"style"`
}

type Publishables struct {
	Entries []struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"published"`
}
