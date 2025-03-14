package services

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/models/workspace"
	"github.com/canghel3/go-geoserver/vector"
)

type Workspaces struct {
	info      *internal.GeoserverInfo
	requester *requester.Requester
}

func newWorkspaceOperations(info *internal.GeoserverInfo) *Workspaces {
	return &Workspaces{
		info:      info,
		requester: requester.NewRequester(info),
	}
}

func (w *Workspaces) Create(name string, _default bool) error {
	return w.requester.Workspaces().Create(name, _default)
}

func (w *Workspaces) Get(name string) (*workspace.SingleWorkspaceRetrievalWrapper, error) {
	return w.requester.Workspaces().Get(name)
}

func (w *Workspaces) GetAll() (*workspace.MultiWorkspaceRetrievalWrapper, error) {
	return w.requester.Workspaces().GetAll()
}

func (w *Workspaces) Delete(name string, recurse bool) error {
	return w.requester.Workspaces().Delete(name, recurse)
}

type WorkspaceServiceSelector struct {
	info *internal.GeoserverInfo
}

func (w *Workspaces) Use(workspace string) *WorkspaceServiceSelector {
	w.info.Workspace = workspace
	return &WorkspaceServiceSelector{
		info: w.info,
	}
}

func (ss *WorkspaceServiceSelector) Vectors() *vector.Service {
	return vector.NewService(ss.info.Clone())
}

//TODO: add raster service and layers service for handling layers without specifying store
