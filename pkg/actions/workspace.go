package actions

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/models/workspace"
)

type Workspaces struct {
	info      *internal.GeoserverData
	requester *requester.Requester
}

func NewWorkspaceActions(info *internal.GeoserverData) *Workspaces {
	return &Workspaces{
		info:      info,
		requester: requester.NewRequester(info),
	}
}

func (w *Workspaces) Create(name string, _default bool) error {
	err := validator.Workspace.Name(name)
	if err != nil {
		return err
	}

	return w.requester.Workspaces().Create(name, _default)
}

func (w *Workspaces) Get(name string) (*workspace.WorkspaceRetrieval, error) {
	err := validator.Workspace.Name(name)
	if err != nil {
		return nil, err
	}

	return w.requester.Workspaces().Get(name)
}

func (w *Workspaces) GetAll() ([]workspace.MultiWorkspace, error) {
	return w.requester.Workspaces().GetAll()
}

func (w *Workspaces) Update(oldName, newName string) error {
	err := validator.Workspace.Name(oldName)
	if err != nil {
		return err
	}

	err = validator.Workspace.Name(newName)
	if err != nil {
		return err
	}

	return w.requester.Workspaces().Update(oldName, newName)
}

func (w *Workspaces) Delete(name string, recurse bool) error {
	err := validator.Workspace.Name(name)
	if err != nil {
		return err
	}

	return w.requester.Workspaces().Delete(name, recurse)
}

type Workspace struct {
	info *internal.GeoserverData
}

func (w *Workspaces) Use(workspace string) *Workspace {
	w.info.Workspace = workspace
	return &Workspace{
		info: w.info,
	}
}

func (ss *Workspace) DataStores() *DataStores {
	return newDataStoresActions(ss.info.Clone())
}

func (ss *Workspace) DataStore(name string) *FeatureTypes {
	return newDataStoresActions(ss.info.Clone()).Use(name)
}

//func (ss *Workspace) FeatureTypes() *FeatureTypes {
//	return newFeatureTypes("", ss.info.Clone())
//}

func (ss *Workspace) CoverageStores() *CoverageStores {
	return newCoverageStoreActions(ss.info.Clone())
}

func (ss *Workspace) CoverageStore(name string) *Coverages {
	return newCoverageStoreActions(ss.info.Clone()).Use(name)
}

//TODO: add raster service and layers service for handling layers without specifying store
