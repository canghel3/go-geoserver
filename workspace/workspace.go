package workspace

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg"
	"github.com/canghel3/go-geoserver/vector"
	"net/http"
)

type Service struct {
	client    *http.Client
	info      *internal.GeoserverInfo
	requester *requester.Requester
}

type ServiceSelector struct {
	info *internal.GeoserverInfo
}

// TODO: client options
func NewService(url, username, password, datadir string) *Service {
	return &Service{
		requester: requester.NewRequester(&internal.GeoserverInfo{
			Client: &http.Client{},
			Connection: internal.GeoserverConnection{
				URL: url,
				Credentials: internal.GeoserverCredentials{
					Username: username,
					Password: password,
				},
			},
			DataDir: datadir,
		}),
	}
}

func (s *Service) Create(name string, _default bool) error {
	return s.requester.Workspaces().Create(name, _default)
}

func (s *Service) Get(name string) (*pkg.SingleWorkspaceRetrievalWrapper, error) {
	return s.requester.Workspaces().Get(name)
}

func (s *Service) GetAll() (*pkg.MultiWorkspaceRetrievalWrapper, error) {
	return s.requester.Workspaces().GetAll()
}

func (s *Service) Delete(name string, recurse bool) error {
	return s.requester.Workspaces().Delete(name, recurse)
}

func (s *Service) Use(workspace string) *ServiceSelector {
	return &ServiceSelector{
		info: &internal.GeoserverInfo{
			Client:     s.client,
			Connection: s.info.Connection,
			DataDir:    s.info.DataDir,
			Workspace:  workspace,
		},
	}
}

func (ss *ServiceSelector) Vectors() *vector.Service {
	return vector.NewService(ss.info)
}
