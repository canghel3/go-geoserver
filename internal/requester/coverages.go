package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/coverages"
	customerrors2 "github.com/canghel3/go-geoserver/pkg/customerrors"
	"io"
	"net/http"
)

type CoverageRequester struct {
	data *internal.GeoserverData
}

func (cr *CoverageRequester) Create(store string, content []byte) error {
	var target string
	if len(store) == 0 {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coverages", cr.data.Connection.URL, cr.data.Workspace)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages", cr.data.Connection.URL, cr.data.Workspace, store)
	}

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(content))
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
	case http.StatusCreated, http.StatusOK:
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
func (cr *CoverageRequester) GetAll(store string) ([]coverages.Coverage, error) {
	return nil, errors.New("not implemented")
}

func (cr *CoverageRequester) Get(store, coverage string) (*coverages.Coverage, error) {
	var target string
	if len(store) == 0 {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coverages/%s.json", cr.data.Connection.URL, cr.data.Workspace, coverage)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s.json", cr.data.Connection.URL, cr.data.Workspace, store, coverage)
	}

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var coverageWrapper coverages.CoverageWrapper

		err = json.NewDecoder(response.Body).Decode(&coverageWrapper)
		if err != nil {
			return nil, err
		}

		return &coverageWrapper.Coverage, nil
	case http.StatusNotFound:
		return nil, customerrors2.WrapNotFoundError(fmt.Errorf("featuretype %s does not exist", coverage))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors2.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageRequester) Delete(store, coverage string, recurse bool) error {
	var target string
	if len(store) == 0 {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coverages/%s?recurse=%v", cr.data.Connection.URL, cr.data.Workspace, coverage, recurse)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s?recurse=%v", cr.data.Connection.URL, cr.data.Workspace, store, coverage, recurse)
	}

	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors2.WrapNotFoundError(fmt.Errorf("featuretype %s does not exist", coverage))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors2.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageRequester) Update(store, coverage string, content []byte) error {
	var target string
	if len(store) == 0 {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coverages/%s", cr.data.Connection.URL, cr.data.Workspace, coverage)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s", cr.data.Connection.URL, cr.data.Workspace, store, coverage)
	}

	request, err := http.NewRequest(http.MethodPut, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return err
	}

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
