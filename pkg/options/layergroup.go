package options

import (
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

var LayerGroup LayerGroupOptionsGenerator

type LayerGroupOptionsGenerator struct{}

type LayerGroupOption func(group *models.Group)

func (lgog LayerGroupOptionsGenerator) Workspace(name string) LayerGroupOption {
	return func(group *models.Group) {
		group.Workspace = &workspace.Creation{Name: name}
	}
}
