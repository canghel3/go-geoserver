package handler

import (
	"encoding/json"
	datastores2 "github.com/canghel3/go-geoserver/datastores"
	"github.com/canghel3/go-geoserver/datastores/postgis"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type storageParams map[string]string

func newDataStores(info *internal.GeoserverInfo) *DataStores {
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

func (ds *DataStores) Use(name string) *FeatureTypes {
	return newFeatureTypes(name, ds.info.Clone())
}

func (ds *DataStores) Create() DataStoreList {
	return DataStoreList{requester: ds.requester}
}

func (ds *DataStores) Get(name string) (*datastores2.DataStoreRetrieval, error) {
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

	data := datastores2.GenericDataStoreCreationWrapper{
		DataStore: datastores2.GenericDataStoreCreationModel{
			Name: name,
			ConnectionParameters: datastores2.ConnectionParameters{
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

func (params storageParams) toDatastoreEntries() []datastores2.Entry {
	entries := make([]datastores2.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastores2.Entry{Key: k, Value: v})
	}

	return entries
}
