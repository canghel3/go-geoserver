package vector

import (
	"github.com/canghel3/go-geoserver/internal/datastore"
	"github.com/canghel3/go-geoserver/internal/datastore/postgis"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type Storage interface {
	Create() error
	Update() error
}

type storageParams map[string]string

type StoreManager struct {
	requester *requester.Requester
}

func (s StoreManager) Get(name string) (*datastore.DataStoreRetrieval, error) {
	return s.requester.DataStores().Get(name)
}

func (s StoreManager) Delete(store string, recurse bool) error {
	return s.requester.DataStores().Delete(store, recurse)
}

func (s StoreManager) PostGIS(name string, connectionParams postgis.ConnectionParams) Storage {
	return newPostGISStore(name, connectionParams, s.requester)
}

func (params storageParams) toDatastoreEntries() []datastore.Entry {
	entries := make([]datastore.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastore.Entry{Key: k, Value: v})
	}

	return entries
}
