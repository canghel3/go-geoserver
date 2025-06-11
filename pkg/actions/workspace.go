package actions

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type Workspaces struct {
	info      internal.GeoserverData
	requester *requester.Requester
}

func NewWorkspaceActions(info internal.GeoserverData) *Workspaces {
	return &Workspaces{
		info:      info,
		requester: requester.NewRequester(info),
	}
}

func (ws *Workspaces) Create(name string, _default bool) error {
	err := validator.Workspace.Name(name)
	if err != nil {
		return err
	}

	return ws.requester.Workspaces().Create(name, _default)
}

func (ws *Workspaces) Get(name string) (*workspace.WorkspaceRetrieval, error) {
	err := validator.Workspace.Name(name)
	if err != nil {
		return nil, err
	}

	return ws.requester.Workspaces().Get(name)
}

func (ws *Workspaces) GetAll() ([]workspace.MultiWorkspace, error) {
	return ws.requester.Workspaces().GetAll()
}

func (ws *Workspaces) Update(oldName, newName string) error {
	err := validator.Workspace.Name(oldName)
	if err != nil {
		return err
	}

	err = validator.Workspace.Name(newName)
	if err != nil {
		return err
	}

	return ws.requester.Workspaces().Update(oldName, newName)
}

func (ws *Workspaces) Delete(name string, recurse bool) error {
	err := validator.Workspace.Name(name)
	if err != nil {
		return err
	}

	return ws.requester.Workspaces().Delete(name, recurse)
}

type Workspace struct {
	info internal.GeoserverData
}

func (ws *Workspaces) Use(workspace string) *Workspace {
	ws.info.Workspace = workspace
	return &Workspace{
		info: ws.info,
	}
}

func (w *Workspace) DataStores() *DataStores {
	return newDataStoresActions(w.info.Clone())
}

func (w *Workspace) DataStore(name string) *FeatureTypes {
	return newDataStoresActions(w.info.Clone()).Use(name)
}

func (w *Workspace) CoverageStores() *CoverageStores {
	return newCoverageStoreActions(w.info.Clone())
}

func (w *Workspace) CoverageStore(name string) *Coverages {
	return newCoverageStoreActions(w.info.Clone()).Use(name)
}

//TODO: add raster service and layers service for handling layers without specifying store
