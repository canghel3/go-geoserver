package postgis

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal/datastore"
)

type PostGISStore struct {
	name             string
	connectionParams ConnectionParams
}

func NewPostGISStore(name string, connectionParams ConnectionParams) *PostGISStore {
	return &PostGISStore{
		name:             name,
		connectionParams: connectionParams,
	}
}

func (pgs *PostGISStore) MarshalJSON() ([]byte, error) {
	cp := storageParams{
		"host":     pgs.connectionParams.Host,
		"database": pgs.connectionParams.Database,
		"user":     pgs.connectionParams.User,
		"passwd":   pgs.connectionParams.Password,
		"port":     pgs.connectionParams.Port,
		"dbtype":   "postgis",
	}

	data := datastore.GenericDataStoreCreationWrapper{
		DataStore: datastore.GenericDataStoreCreationModel{
			Name: pgs.name,
			ConnectionParameters: datastore.ConnectionParameters{
				Entry: cp.toDatastoreEntries(),
			},
		},
	}

	return json.Marshal(&data)
}

type storageParams map[string]string

func (params storageParams) toDatastoreEntries() []datastore.Entry {
	entries := make([]datastore.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastore.Entry{Key: k, Value: v})
	}

	return entries
}
