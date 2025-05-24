package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	customerrors2 "github.com/canghel3/go-geoserver/pkg/models/customerrors"
	"github.com/canghel3/go-geoserver/pkg/models/datastores"
	"io"
	"net/http"
)

type DataStoreRequester struct {
	data *internal.GeoserverData
}

func (dr *DataStoreRequester) Create(content []byte) error {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores", dr.data.Connection.URL, dr.data.Workspace), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(dr.data.Connection.Credentials.Username, dr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := dr.data.Client.Do(request)
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

		return customerrors2.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

// TODO: implement
func (dr *DataStoreRequester) GetAll() ([]datastores.DataStore, error) {
	return nil, errors.New("not implemented")
}

func (dr *DataStoreRequester) Get(name string) (*datastores.DataStore, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s", dr.data.Connection.URL, dr.data.Workspace, name), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(dr.data.Connection.Credentials.Username, dr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := dr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var dts *datastores.DataStoreRetrievalWrapper
		err = json.NewDecoder(response.Body).Decode(&dts)
		if err != nil {
			return nil, err
		}

		return &dts.DataStore, nil
	case http.StatusNotFound:
		return nil, customerrors2.WrapNotFoundError(fmt.Errorf("datastore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors2.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

// TODO: implement
func (dr *DataStoreRequester) Update() error {
	return errors.New("not implemented")
}

func (dr *DataStoreRequester) Delete(name string, recurse bool) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s?recurse=%v", dr.data.Connection.URL, dr.data.Workspace, name, recurse), nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(dr.data.Connection.Credentials.Username, dr.data.Connection.Credentials.Password)

	response, err := dr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors2.WrapNotFoundError(fmt.Errorf("datastore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors2.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (dr *DataStoreRequester) Reset(name string) error {
	return errors.New("not implemented")
}
