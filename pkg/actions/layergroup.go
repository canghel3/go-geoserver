package actions

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/layers"
	"strings"
)

type LayerGroups struct {
	data      internal.GeoserverData
	requester requester.LayerGroupRequester
}

func NewLayerGroup(data internal.GeoserverData) LayerGroups {
	return LayerGroups{
		data:      data,
		requester: requester.NewLayerGroupRequester(data),
	}
}

func (lg LayerGroups) Get(name string) (*layers.Group, error) {
	return lg.requester.Get(name)
}

func (lg LayerGroups) Publish(group models.Group) error {
	if err := validator.Name(group.Name); err != nil {
		return err
	}

	for i := range group.Publishables.Entries {
		if !strings.HasPrefix(group.Publishables.Entries[i].Name, lg.data.Workspace) {
			group.Publishables.Entries[i].Name = fmt.Sprintf("%s:%s", lg.data.Workspace, group.Publishables.Entries[i].Name)
		}

		err := validator.WorkspaceLayerFormat(lg.data.Workspace, group.Publishables.Entries[i].Name)
		if err != nil {
			return err
		}
	}

	content, err := json.Marshal(models.GroupWrapper{Group: group})
	if err != nil {
		return err
	}

	return lg.requester.Create(content)
}

func (lg LayerGroups) Update(name string, group models.Group) error {
	if err := validator.Name(name); err != nil {
		return err
	}

	if err := validator.Name(group.Name); err != nil {
		return err
	}

	for _, entry := range group.Publishables.Entries {
		err := validator.WorkspaceLayerFormat(lg.data.Workspace, entry.Name)
		if err != nil {
			return err
		}

		if !strings.HasPrefix(entry.Name, lg.data.Workspace) {
			entry.Name = fmt.Sprintf("%s:%s", lg.data.Workspace, entry.Name)
		}
	}

	content, err := json.Marshal(models.GroupWrapper{Group: group})
	if err != nil {
		return err
	}

	return lg.requester.Update(name, content)
}

func (lg LayerGroups) Delete(name string) error {
	return lg.requester.Delete(name)
}
