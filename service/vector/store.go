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
	MarshalJSON() ([]byte, error)
}

type storageParams map[string]string

type Stores struct {
	info models.GeoserverInfo
}

func newStores(info models.GeoserverInfo) Stores {
	return Stores{
		info: info,
	}
}

func (s Stores) Get(name string) (*datastore.DataStoreRetrieval, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s", s.info.Connection.URL, s.info.Workspace, name), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := s.info.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var dts *datastore.DataStoreRetrievalWrapper
		err = json.NewDecoder(response.Body).Decode(&dts)
		if err != nil {
			return nil, err
		}

		return &dts.DataStore, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (s Stores) Delete(store string, recurse bool) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s?recurse=%v", s.info.Connection.URL, s.info.Workspace, store, recurse), nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)

	response, err := s.info.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (s Stores) Create(store Storage) error {
	content, err := store.MarshalJSON()
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores", s.info.Connection.URL, s.info.Workspace), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	if s.info.Client == nil {
		s.info.Client = &http.Client{}
	}

	response, err := s.info.Client.Do(request)
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

func (s Stores) PostGIS(name string, connectionParams postgis.ConnectionParams) Storage {
	return newPostGISStore(name, connectionParams)
}

func (s Stores) createGenericDataStore(store string, connectionParams storageParams) error {
	data := datastore.GenericDataStoreCreationWrapper{
		DataStore: datastore.GenericDataStoreCreationModel{
			Name: store,
			ConnectionParameters: datastore.ConnectionParameters{
				Entry: connectionParams.toDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores", s.info.Connection.URL, s.info.Workspace), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	if s.info.Client == nil {
		s.info.Client = &http.Client{}
	}

	response, err := s.info.Client.Do(request)
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

func (params storageParams) toDatastoreEntries() []datastore.Entry {
	entries := make([]datastore.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastore.Entry{Key: k, Value: v})
	}

	return entries
}
