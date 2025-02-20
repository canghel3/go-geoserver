package vector

import "github.com/canghel3/go-geoserver/models"

type Service struct {
	info models.GeoserverInfo

	layers Layers
	stores Stores
}

func NewService(info models.GeoserverInfo) *Service {
	return &Service{
		info: info,

		layers: Layers{},
		stores: newStores(info),
	}
}

func (s *Service) Store(name string) Layers {
	return s.layers
}

func (s *Service) Layers() Layers {
	return s.layers
}

func (s *Service) Stores() Stores {
	return s.stores
}
