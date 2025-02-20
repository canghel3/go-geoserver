package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/datastore"
	"github.com/canghel3/go-geoserver/models/datastore/postgis"
	"github.com/canghel3/go-geoserver/utils"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func (gs *GeoserverService) CreatePostGISDataStore(workspace, store string, connectionParams postgis.ConnectionParams) error {
	v := reflect.ValueOf(&connectionParams).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			return customerrors.WrapInputError(fmt.Errorf("%v cannot be empty", strings.ToLower(v.Type().Field(i).Name)))
		}
	}

	cp := map[string]string{
		"host":     connectionParams.Host,
		"database": connectionParams.Database,
		"user":     connectionParams.User,
		"passwd":   connectionParams.Password,
		"port":     connectionParams.Port,
		"dbtype":   "postgis",
	}

	return gs.createDataStore(workspace, store, cp)
}

/*
GetDataStore retrieves a datastore from a Geoserver workspace.

- If the datastore exists and the request is successful, it returns a pointer to the DataStoreRetrieval structure and a nil error.

- In case no such datastore exists within the provided workspace, a NotFoundError is returned indicating the missing datastore.

- If Geoserver returns a non-successful status code, it returns nil and a GeoserverError with a message containing status code and server response.

- In case of network issues or JSON parsing problems, it returns the respective Go error and a nil pointer.
*/
func (gs *GeoserverService) GetDataStore(workspace, store string) (*datastore.DataStoreRetrieval, error) {
	err := utils.ValidateStore(store)
	if err != nil {
		return nil, err
	}

	_, err = gs.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s", gs.data.Connection.URL, workspace, store), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(gs.data.Connection.Credentials.Username, gs.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := gs.data.Client.Do(request)
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
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("datastore %s does not exist", store))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (gs *GeoserverService) UpdatePostGISDataStore(workspace, store string, connectionParams postgis.ConnectionParams) error {
	return fmt.Errorf("not implemented")
}

/*
DeleteDataStore deletes a datastore from geoserver.

Available options: RecurseOption
*/
func (gs *GeoserverService) DeleteDataStore(workspace string, store string, opts ...utils.Option) error {
	_, err := gs.GetWorkspace(workspace)
	var nfe *customerrors.NotFoundError
	if err != nil && errors.Is(err, nfe) {
		return err
	}

	_, err = gs.GetDataStore(workspace, store)
	if err != nil && errors.Is(err, nfe) {
		return err
	}

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s", gs.data.Connection.URL, workspace, store), nil)
	if err != nil {
		return err
	}

	params := utils.ProcessOptions(opts)
	if recurse, set := params["recurse"]; set {
		q := request.URL.Query()
		q.Add("recurse", fmt.Sprintf("%v", recurse.(bool)))
		request.URL.RawQuery = q.Encode()
	}

	request.SetBasicAuth(gs.data.Connection.Credentials.Username, gs.data.Connection.Credentials.Password)

	response, err := gs.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.NewNotFoundError(fmt.Sprintf("datastore %s does not exist", store))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (gs *GeoserverService) createDataStore(workspace, store string, connectionParams map[string]string) error {
	_, err := gs.GetWorkspace(workspace)
	var nfe *customerrors.NotFoundError
	if err != nil && errors.Is(err, nfe) {
		return err
	}

	var enf *customerrors.NotFoundError
	_, err = gs.GetDataStore(workspace, store)
	if err == nil {
		return customerrors.WrapConflictError(fmt.Errorf("datastore %s already exists", store))
	} else if err != nil && !errors.As(err, &enf) {
		return err
	}

	data := datastore.GenericDataStoreCreationWrapper{
		DataStore: datastore.GenericDataStoreCreationModel{
			Name: store,
			ConnectionParameters: datastore.ConnectionParameters{
				Entry: gs.connectionParamsToEntries(connectionParams),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores", gs.data.Connection.URL, workspace), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gs.data.Connection.Credentials.Username, gs.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := gs.data.Client.Do(request)
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

func (gs *GeoserverService) connectionParamsToEntries(params map[string]string) []datastore.Entry {
	entries := make([]datastore.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastore.Entry{Key: k, Value: v})
	}

	return entries
}
