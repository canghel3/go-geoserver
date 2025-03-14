package vector

import (
	"github.com/canghel3/go-geoserver/datastore"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type StoreOperations struct {
	store     string
	requester *requester.Requester
}

func newStoreOperations(name string, info *internal.GeoserverInfo) StoreOperations {
	return StoreOperations{
		store:     name,
		requester: requester.NewRequester(info),
	}
}

func (s StoreOperations) Get() (*datastore.DataStoreRetrieval, error) {
	return s.requester.DataStores().Get(s.store)
}

func (s StoreOperations) Delete(recurse bool) error {
	return s.requester.DataStores().Delete(s.store, recurse)
}
