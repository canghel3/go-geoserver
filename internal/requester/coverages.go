package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/coverages"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"io"
	"net/http"
)

type CoverageRequester struct {
	data internal.GeoserverData
}

func NewCoverageRequester(data internal.GeoserverData) CoverageRequester {
	return CoverageRequester{
		data: data,
	}
}

func (cr *CoverageRequester) Create(store string, content []byte) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages", cr.data.Connection.URL, cr.data.Workspace, store)

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

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageRequester) GetAll(store string) (*coverages.Coverages, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages", cr.data.Connection.URL, cr.data.Workspace, store)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err

	}

	switch response.StatusCode {
	case http.StatusOK:
		var cvgs *coverages.CoveragesWrapper
		err = json.Unmarshal(body, &cvgs)
		if err != nil {
			//try to unmarshal into empty string because geoserver has a funny way of responding
			type noCoveragesExists struct {
				Coverages string `json:"coverages"`
			}
			var noCoveragesExistsResponse noCoveragesExists
			noCoveragesExistsError := json.Unmarshal(body, &noCoveragesExistsResponse)
			if noCoveragesExistsError == nil {
				return &coverages.Coverages{Entries: nil}, nil
			}

			return nil, err
		}

		return &cvgs.Coverages, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageRequester) Get(store, coverage string) (*coverages.Coverage, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s.json", cr.data.Connection.URL, cr.data.Workspace, store, coverage)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

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
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("coverage %s not found", coverage))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageRequester) Delete(store, coverage string, recurse bool) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s?recurse=%v", cr.data.Connection.URL, cr.data.Workspace, store, coverage, recurse)

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
		return customerrors.WrapNotFoundError(fmt.Errorf("coverage %s not found", coverage))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageRequester) Update(store, coverage string, content []byte) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s", cr.data.Connection.URL, cr.data.Workspace, store, coverage)

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
	case http.StatusOK, http.StatusCreated:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("coverage %s not found", coverage))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (cr *CoverageRequester) Reset(store, coverage string) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s/reset", cr.data.Connection.URL, cr.data.Workspace, store, coverage)

	request, err := http.NewRequest(http.MethodPut, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(cr.data.Connection.Credentials.Username, cr.data.Connection.Credentials.Password)

	response, err := cr.data.Client.Do(request)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("coverage %s not found", coverage))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
