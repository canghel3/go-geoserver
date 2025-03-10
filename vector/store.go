package vector

import (
	"github.com/canghel3/go-geoserver/datastore"
	"github.com/canghel3/go-geoserver/datastore/postgis"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type Storage interface {
	Create() error
	Update() error
}

type storageParams map[string]string

type StoreList struct {
	requester *requester.Requester
}

func newStoreList(info *internal.GeoserverInfo) StoreList {
	return StoreList{
		requester: requester.NewRequester(info),
	}
}

func (s StoreList) PostGIS(name string, connectionParams postgis.ConnectionParams) Storage {
	return newPostGISStore(name, connectionParams, s.requester)
}

func (params storageParams) toDatastoreEntries() []datastore.Entry {
	entries := make([]datastore.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastore.Entry{Key: k, Value: v})
	}

	return entries
}
