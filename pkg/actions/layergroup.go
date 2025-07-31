package actions

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/layers"
	"github.com/canghel3/go-geoserver/pkg/workspace"
	"strings"
	"time"
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

	if group.Workspace == nil {
		if validator.Empty(lg.data.Workspace) {
			return customerrors.NewInputError("workspace is required. use Workspace from options.LayerGroup")
		} else {
			group.Workspace = &workspace.Creation{
				Name: lg.data.Workspace,
			}
		}
	}

	if err := validator.Name(group.Workspace.Name); err != nil {
		return err
	}

	for i := range group.Publishables.Entries {
		if !strings.HasPrefix(group.Publishables.Entries[i].Name, group.Workspace.Name) {
			group.Publishables.Entries[i].Name = fmt.Sprintf("%s:%s", group.Workspace.Name, group.Publishables.Entries[i].Name)
		}
	}

	if group.Styles != nil {
		for i := range group.Styles.Style {
			if !validator.Empty(group.Styles.Style[i].Name) && !strings.HasPrefix(group.Styles.Style[i].Name, group.Workspace.Name) {
				group.Styles.Style[i].Name = fmt.Sprintf("%s:%s", group.Workspace.Name, group.Styles.Style[i].Name)
			}
		}
	}

	content, err := json.Marshal(models.GroupWrapper{Group: group})
	if err != nil {
		return err
	}

	return lg.requester.Create(content)
}

func (lg LayerGroups) Update(name string, group layers.Group) error {
	if err := validator.Name(name); err != nil {
		return err
	}

	if err := validator.Name(group.Name); err != nil {
		return err
	}

	if group.Workspace == nil {
		if validator.Empty(lg.data.Workspace) {
			return customerrors.NewInputError("group.Workspace is required")
		} else {
			group.Workspace = &workspace.Creation{
				Name: lg.data.Workspace,
			}
		}
	}

	if err := validator.Name(group.Workspace.Name); err != nil {
		return err
	}

	for i := range group.Publishables.Entries {
		if !strings.HasPrefix(group.Publishables.Entries[i].Name, group.Workspace.Name) {
			group.Publishables.Entries[i].Name = fmt.Sprintf("%s:%s", group.Workspace.Name, group.Publishables.Entries[i].Name)
		}
	}

	if group.Styles != nil {
		for i := range group.Styles.Style {
			if !validator.Empty(group.Styles.Style[i].Name) && !strings.HasPrefix(group.Styles.Style[i].Name, group.Workspace.Name) {
				group.Styles.Style[i].Name = fmt.Sprintf("%s:%s", group.Workspace.Name, group.Styles.Style[i].Name)
			}
		}
	}

	if len(group.DateModified) == 0 {
		group.DateModified = time.Now().UTC().Format(time.RFC3339)
	}

	content, err := json.Marshal(layers.GroupWrapper{Group: group})
	if err != nil {
		return err
	}

	return lg.requester.Update(name, content)
}

func (lg LayerGroups) Delete(name string) error {
	return lg.requester.Delete(name)
}
