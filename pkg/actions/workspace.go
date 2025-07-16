package actions

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type Workspaces struct {
	data      internal.GeoserverData
	requester requester.WorkspaceRequester
}

func NewWorkspaceActions(info internal.GeoserverData) Workspaces {
	return Workspaces{
		data:      info,
		requester: requester.NewWorkspaceRequester(info),
	}
}

func (ws Workspaces) Create(name string, _default bool) error {
	err := validator.Name(name)
	if err != nil {
		return err
	}

	data := workspace.WorkspaceCreationWrapper{
		Workspace: workspace.WorkspaceCreation{
			Name: name,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ws.requester.Create(content, _default)
}

func (ws Workspaces) Get(name string) (*workspace.WorkspaceRetrieval, error) {
	err := validator.Name(name)
	if err != nil {
		return nil, err
	}

	return ws.requester.Get(name)
}

func (ws Workspaces) GetAll() ([]workspace.MultiWorkspace, error) {
	return ws.requester.GetAll()
}

func (ws Workspaces) Update(oldName, newName string) error {
	err := validator.Name(oldName)
	if err != nil {
		return err
	}

	err = validator.Name(newName)
	if err != nil {
		return err
	}

	data := workspace.WorkspaceUpdateWrapper{
		Workspace: workspace.WorkspaceUpdate{
			Name: newName,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ws.requester.Update(content, oldName)
}

func (ws Workspaces) Delete(name string, recurse bool) error {
	err := validator.Name(name)
	if err != nil {
		return err
	}

	return ws.requester.Delete(name, recurse)
}

type Workspace struct {
	data internal.GeoserverData
}

func (ws Workspaces) Use(workspace string) Workspace {
	ws.data.Workspace = workspace
	return Workspace{
		data: ws.data,
	}
}

func (w Workspace) DataStores() DataStores {
	return newDataStoresActions(w.data.Clone())
}

func (w Workspace) DataStore(name string) FeatureTypes {
	return newDataStoresActions(w.data.Clone()).Use(name)
}

func (w Workspace) CoverageStores() CoverageStores {
	return newCoverageStoreActions(w.data.Clone())
}

func (w Workspace) CoverageStore(name string) Coverages {
	return newCoverageStoreActions(w.data.Clone()).Use(name)
}

func (w Workspace) GeoWebCache() GeoWebCache {
	return NewGeoWebCache(w.data.Clone())
}

func (w Workspace) WMS(version wms.WMSVersion) WMS {
	return NewWMSActions(w.data.Clone(), version)
}
