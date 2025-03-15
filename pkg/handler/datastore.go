package handler

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/datastores"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
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

func (ds *DataStores) Get(name string) (*datastores.DataStoreRetrieval, error) {
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

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name: name,
			ConnectionParameters: datastores.ConnectionParameters{
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

func (params storageParams) toDatastoreEntries() []datastores.Entry {
	entries := make([]datastores.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastores.Entry{Key: k, Value: v})
	}

	return entries
}
