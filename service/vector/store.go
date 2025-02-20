package vector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models"
	"github.com/canghel3/go-geoserver/models/datastore"
	"github.com/canghel3/go-geoserver/models/datastore/postgis"
	"io"
	"net/http"
)

type Storage interface {
	Create() error
	//Delete() error
}

type Stores struct {
	info models.GeoserverInfo
}

func newStores(info models.GeoserverInfo) Stores {
	return Stores{
		info: info,
	}
}

func (s Stores) Delete(name string) error {
	return nil
}

func (s Stores) PostGIS(name string, connectionParams postgis.ConnectionParams) Storage {
	return newPostGISStore(name, connectionParams, s.info)
}

func createGenericDataStore(workspace, store string, connectionParams map[string]string) error {
	data := datastore.GenericDataStoreCreationWrapper{
		DataStore: datastore.GenericDataStoreCreationModel{
			Name: store,
			ConnectionParameters: datastore.ConnectionParameters{
				Entry: connectionParamsToEntries(connectionParams),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores", gs.url, workspace), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gs.username, gs.password)
	request.Header.Add("Content-Type", "application/json")

	response, err := gs.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusCreated:
		return nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func connectionParamsToEntries(params map[string]string) []datastore.Entry {
	entries := make([]datastore.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastore.Entry{Key: k, Value: v})
	}

	return entries
}
