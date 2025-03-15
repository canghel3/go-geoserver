package handler

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/datastore"
	"github.com/canghel3/go-geoserver/pkg/datastore/postgis"
)

type storageParams map[string]string

func NewDataStores(info *internal.GeoserverInfo) *DataStores {
	r := requester.NewRequester(info)
	return &DataStores{
		info:      info,
		requester: r,
	}
}

type DataStoreList struct {
	requester *requester.Requester
}

type DataStores struct {
	info      *internal.GeoserverInfo
	requester *requester.Requester
}

func (ds *DataStores) Create() DataStoreList {
	return DataStoreList{requester: ds.requester}
}

func (ds *DataStores) Get(name string) (*datastore.DataStoreRetrieval, error) {
	return ds.requester.DataStores().Get(name)
}

func (ds *DataStores) Delete(name string, recurse bool) error {
	return ds.requester.DataStores().Delete(name, recurse)
}

func (dsl DataStoreList) PostGIS(name string, connection postgis.ConnectionParams) error {
	cp := storageParams{
		"host":     connection.Host,
		"database": connection.Database,
		"user":     connection.User,
		"passwd":   connection.Password,
		"port":     connection.Port,
		"dbtype":   "postgis",
	}

	data := datastore.GenericDataStoreCreationWrapper{
		DataStore: datastore.GenericDataStoreCreationModel{
			Name: name,
			ConnectionParameters: datastore.ConnectionParameters{
				Entry: cp.toDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.DataStores().Create(content)
}

func (params storageParams) toDatastoreEntries() []datastore.Entry {
	entries := make([]datastore.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastore.Entry{Key: k, Value: v})
	}

	return entries
}
