package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/coveragestores"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"io"
	"net/http"
)

type CoverageStoreRequester struct {
	data internal.GeoserverData
}

func (cr *CoverageStoreRequester) Create(content []byte) error {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores", cr.data.Connection.URL, cr.data.Workspace), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := cr.data.Client.Do(request)
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

// TODO: implement
func (cr *CoverageStoreRequester) GetAll() ([]coveragestores.CoverageStore, error) {
	return nil, errors.New("not implemented")
}

func (cr *CoverageStoreRequester) Get(name string) (*coveragestores.CoverageStore, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s", cr.data.Connection.URL, cr.data.Workspace, name), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var cts *coveragestores.CoverageStoreRetrievalWrapper
		err = json.NewDecoder(response.Body).Decode(&cts)
		if err != nil {
			return nil, err
		}

		return &cts.CoverageStore, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("coveragestore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

// TODO: implement
func (cr *CoverageStoreRequester) Update() error {
	return errors.New("not implemented")
}

func (cr *CoverageStoreRequester) Delete(name string, recurse bool) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s?recurse=%v", cr.data.Connection.URL, cr.data.Workspace, name, recurse), nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("coveragestore %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageStoreRequester) Reset(name string) error {
	return errors.New("not implemented")
}
