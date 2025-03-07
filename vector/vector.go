package vector

import "github.com/canghel3/go-geoserver/internal/requester"

type Service struct {
	layers Layers
	stores StoreManager
}

func NewService(requester *requester.Requester) *Service {
	return &Service{
		layers: Layers{},
		stores: StoreManager{requester: requester},
	}
}

func (s *Service) Store(name string) Layers {
	return s.layers
}

func (s *Service) Layers() Layers {
	return s.layers
}

func (s *Service) Stores() StoreManager {
	return s.stores
}
