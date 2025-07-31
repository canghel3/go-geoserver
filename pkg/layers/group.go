package layers

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type GroupMode string

var (
	ModeSingle    GroupMode = "SINGLE"
	ModeContainer GroupMode = "CONTAINER"
	ModeNamed     GroupMode = "NAMED"
	//ModeEo       GroupMode = "EO"	TODO: mode EO requires a root layer which is not supported.
)

func NewGroup(name string, mode GroupMode, layers []LayerInput, options ...options.LayerGroupOption) models.Group {
	publishables := models.Publishables{}
	publishables.Entries = make([]struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	}, 0)
	styles := models.GroupStyles{}
	styles.Style = make([]shared.Style, 0)

	for _, layer := range layers {
		publishables.Entries = append(publishables.Entries, struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		}{Type: string(layer.Type), Name: layer.Name})

		if !validator.Empty(layer.Style) {
			styles.Style = append(styles.Style, shared.Style{
				Name: layer.Style,
			})
		}
	}

	g := models.Group{
		Name:         name,
		Mode:         string(mode),
		Publishables: publishables,
	}

	if len(styles.Style) > 0 {
		g.Styles = &styles
	}

	for _, option := range options {
		option(&g)
	}

	return g
}

type LayerType string

var (
	TypeLayer      LayerType = "layer"
	TypeLayerGroup LayerType = "layerGroup"
)

type LayerInput struct {
	Type  LayerType
	Name  string
	Style string
}

type GroupWrapper struct {
	Group Group `json:"layerGroup"`
}

// TODO: Although a layer name can be sent as a string formatted number,
// geoserver parses it to an actual number and returns it,
// which will cause panics when decoding the name here.
type Group struct {
	Name         string              `json:"name"`
	Mode         GroupMode           `json:"mode"`
	Title        string              `json:"title"`
	Workspace    *workspace.Creation `json:"workspace,omitempty"`
	Publishables *Publishables       `json:"publishables"`
	Bounds       *shared.BoundingBox `json:"bounds,omitempty"`
	Keywords     *shared.Keywords    `json:"keywords,omitempty"`
	Styles       *GroupStyles        `json:"styles,omitempty"`
	DateCreated  string              `json:"dateCreated"`
	DateModified string              `json:"dateModified"`
}

func (g *Group) AddPublishables(layers ...LayerInput) {
	for _, layer := range layers {
		if g.Publishables == nil {
			g.Publishables = &Publishables{}
		}

		g.Publishables.Entries = append(g.Publishables.Entries, struct {
			Type string `json:"@type"`
			Name string `json:"name"`
			Link string `json:"href,omitempty"`
		}{Type: string(layer.Type), Name: layer.Name})

		if g.Styles == nil {
			g.Styles = &GroupStyles{}
		}

		g.Styles.Style = append(g.Styles.Style, shared.Style{
			Name: layer.Style,
		})
	}
}

//func (g *Group) MarshalJSON() ([]byte, error) {
//	type alias Group
//	var temp alias
//
//	temp.Name = g.Name
//	temp.Mode = g.Mode
//	temp.Title = g.Title
//	temp.Workspace = g.Workspace
//	temp.Publishables = g.Publishables
//	temp.Bounds = g.Bounds
//	temp.Keywords = g.Keywords
//	temp.Styles = g.Styles
//
//	return json.Marshal(temp)
//}

type GroupStyles struct {
	Style []shared.Style `json:"style"`
}

func (gs *GroupStyles) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	type alias GroupStyles
	var temp alias
	if err := json.Unmarshal(m["style"], &temp.Style); err == nil {
		*gs = GroupStyles(temp)
		return nil
	}

	//geoserver responds with a single string when the group contains a single layer with the default style
	var s string
	if err := json.Unmarshal(m["style"], &s); err == nil {
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
	Entries []Entries `json:"published"`
}

func (p *Publishables) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	m = make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	type alias Publishables
	var temp alias
	if err := json.Unmarshal(m["published"], &temp.Entries); err == nil {
		p.Entries = temp.Entries
		return nil
	}

	var entry Entries
	if err := json.Unmarshal(m["published"], &entry); err == nil {
		p.Entries = []Entries{entry}
		return nil
	}

	return nil
}

type Entries struct {
	Type string `json:"@type"`
	Name string `json:"name"`
	Link string `json:"href,omitempty"`
}
