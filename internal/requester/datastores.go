package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/datastores"
	"github.com/canghel3/go-geoserver/internal"
	"io"
	"net/http"
)

type DataStoreRequester struct {
	info *internal.GeoserverInfo
}

func (dr *DataStoreRequester) Create(content []byte) error {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores", dr.info.Connection.URL, dr.info.Workspace), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(dr.info.Connection.Credentials.Username, dr.info.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := dr.info.Client.Do(request)
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

// TODO: implement
func (dr *DataStoreRequester) GetAll() (*datastores.AllDataStoreRetrievalWrapper, error) {
	return nil, errors.New("not implemented")
}

func (dr *DataStoreRequester) Get(name string) (*datastores.DataStoreRetrieval, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s", dr.info.Connection.URL, dr.info.Workspace, name), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(dr.info.Connection.Credentials.Username, dr.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := dr.info.Client.Do(request)
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
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("datastore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

// TODO: implement
func (dr *DataStoreRequester) Update() error {
	return errors.New("not implemented")
}

func (dr *DataStoreRequester) Delete(name string, recurse bool) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s?recurse=%v", dr.info.Connection.URL, dr.info.Workspace, name, recurse), nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(dr.info.Connection.Credentials.Username, dr.info.Connection.Credentials.Password)

	response, err := dr.info.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("datastore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (dr *DataStoreRequester) Reset(name string) error {
	return errors.New("not implemented")
}
