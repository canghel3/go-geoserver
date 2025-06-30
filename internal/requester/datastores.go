package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/datastores"
	"io"
	"net/http"
)

type DataStoreRequester struct {
	data internal.GeoserverData
}

func NewDataStoreRequester(data internal.GeoserverData) DataStoreRequester {
	return DataStoreRequester{
		data: data,
	}
}

func (dr DataStoreRequester) Create(content []byte) error {
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
	case http.StatusOK, http.StatusCreated:
		return nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (dr DataStoreRequester) GetAll() (*datastores.DataStores, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores", dr.data.Connection.URL, dr.data.Workspace), nil)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err

	}

	switch response.StatusCode {
	case http.StatusOK:
		var dts *datastores.DataStoresWrapper
		err = json.Unmarshal(body, &dts)
		if err != nil {
			//try to unmarshal into empty string because geoserver has a funny way of responding
			type noDataStoreExists struct {
				DataStores string `json:"dataStores"`
			}
			var noDataStoreExistsResponse noDataStoreExists
			noDataStoreExistsError := json.Unmarshal(body, &noDataStoreExistsResponse)
			if noDataStoreExistsError == nil {
				return &datastores.DataStores{Entries: nil}, nil
			}

			return nil, err
		}

		return &dts.DataStores, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (dr DataStoreRequester) Get(name string) (*datastores.DataStore, error) {
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
		var dts *datastores.DataStoreWrapper
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

func (dr DataStoreRequester) Update(name string, content []byte) error {
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s", dr.data.Connection.URL, dr.data.Workspace, name), bytes.NewReader(content))
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
	case http.StatusOK, http.StatusCreated:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("datastore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (dr DataStoreRequester) Delete(name string, recurse bool) error {
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
		return customerrors.WrapNotFoundError(fmt.Errorf("datastore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (dr DataStoreRequester) Reset(name string) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/reset", dr.data.Connection.URL, dr.data.Workspace, name)

	request, err := http.NewRequest(http.MethodPut, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(dr.data.Connection.Credentials.Username, dr.data.Connection.Credentials.Password)

	response, err := dr.data.Client.Do(request)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("datastore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
