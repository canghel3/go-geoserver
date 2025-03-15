package vector

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type Service struct {
	info            *internal.GeoserverInfo
	storeOperations StoreOperations
	storeList       StoreList
}

func NewService(info *internal.GeoserverInfo) *Service {
	return &Service{
		info:      info,
		storeList: StoreList{requester: requester.NewRequester(info)},
	}
}

func (s *Service) Store(name string) StoreOperations {
	return newStoreOperations(name, s.info)
}

func (s *Service) Stores() StoreList {
	return newStoreList(s.info)
}
