package vector

import (
	"github.com/canghel3/go-geoserver/internal"
)

type Service struct {
	info *internal.GeoserverInfo

	layers Layers
	stores Stores
}

func NewService(info *internal.GeoserverInfo) *Service {
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
