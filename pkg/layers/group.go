package layers

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type GroupMode string

var (
	ModeSingle    GroupMode = "SINGLE"
	ModeContainer GroupMode = "CONTAINER"
	ModeNamed     GroupMode = "NAMED"
	ModeEo        GroupMode = "EO"
)

func NewGroup(name string, mode GroupMode, layers []LayerInput) models.Group {
	publishables := models.Publishables{}
	publishables.Entries = make([]struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	}, 0)

	for _, layer := range layers {
		publishables.Entries = append(publishables.Entries, struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		}{Type: string(layer.Type), Name: layer.Name})
	}

	return models.Group{
		Name:         name,
		Mode:         string(mode),
		Publishables: publishables,
	}
}

type LayerType string

var (
	TypeLayer      LayerType = "layer"
	TypeLayerGroup LayerType = "layerGroup"
)

type LayerInput struct {
	Type LayerType `json:"@type"`
	Name string    `json:"name"`
}

type GroupWrapper struct {
	Group Group `json:"layerGroup"`
}

type Group struct {
	//TODO: Although a layer name can be sent as a string formatted number,
	// geoserver parses it to an actual number and returns it,
	// which will cause panics when decoding the name here.
	Name         string                     `json:"name"`
	Mode         GroupMode                  `json:"mode"`
	Title        string                     `json:"title"`
	Workspace    *workspace.Creation        `json:"workspace,omitempty"`
	Publishables Publishables               `json:"publishables"`
	Bounds       shared.BoundingBoxCRSClass `json:"bounds"`
	Keywords     *shared.Keywords           `json:"keywords,omitempty"`
	Styles       GroupStyles                `json:"styles,omitempty"`
	DateCreated  string                     `json:"dateCreated"`
	DateModified string                     `json:"dateModified"`
}

type GroupStyles struct {
	Style []shared.Style `json:"style"`
}

func (gs *GroupStyles) UnmarshalJSON(data []byte) error {
	type alias GroupStyles
	var temp alias
	if err := json.Unmarshal(data, &temp); err == nil {
		*gs = GroupStyles(temp)
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		gs.Style = []shared.Style{
			{
				Name: s,
			},
		}
		return nil
	}

	return nil
}

type Publishables struct {
	Entries []struct {
		Name string `json:"name"`
		Link string `json:"href"`
	} `json:"published"`
}
