package vector

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal/datastore"
	"github.com/canghel3/go-geoserver/internal/datastore/postgis"
)

type PostGISStore struct {
	name             string
	connectionParams postgis.ConnectionParams
}

func newPostGISStore(name string, connectionParams postgis.ConnectionParams) *PostGISStore {
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
