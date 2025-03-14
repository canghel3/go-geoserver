package services

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"net/http"
)

type Service struct {
	client    *http.Client
	info      *internal.GeoserverInfo
	requester *requester.Requester
}

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

func (s *Service) Workspaces() *Workspaces {
	return newWorkspaceOperations(s.info.Clone())
}

func (s *Service) WMS() *WMS {
	return newWMS(s.info.Clone())
}

//TODO: implement wms, wfs and others
