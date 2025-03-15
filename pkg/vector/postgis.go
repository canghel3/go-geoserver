package vector

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/models/datastore"
	"github.com/canghel3/go-geoserver/pkg/models/datastore/postgis"
)

type PostGISStore struct {
	name             string
	connectionParams postgis.ConnectionParams
	requester        *requester.Requester
}

func newPostGISStore(name string, connectionParams postgis.ConnectionParams, requester *requester.Requester) *PostGISStore {
	return &PostGISStore{
		name:             name,
		connectionParams: connectionParams,
		requester:        requester,
	}
}

func (pgs *PostGISStore) Create() error {
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

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return pgs.requester.DataStores().Create(content)
}

func (pgs *PostGISStore) Update() error {
	return nil
}
