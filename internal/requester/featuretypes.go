package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"io"
	"net/http"
)

type FeatureTypeRequester struct {
	data internal.GeoserverData
}

func (ftr *FeatureTypeRequester) Create(store string, content []byte) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes", ftr.data.Connection.URL, ftr.data.Workspace, store)

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(ftr.data.Connection.Credentials.Username, ftr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := ftr.data.Client.Do(request)
	if err != nil {
		return err
	}

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

func (ftr *FeatureTypeRequester) Delete(store, feature string, recurse bool) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes/%s?recurse=%v", ftr.data.Connection.URL, ftr.data.Workspace, store, feature, recurse)

	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(ftr.data.Connection.Credentials.Username, ftr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := ftr.data.Client.Do(request)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("featuretype %s not found", feature))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ftr *FeatureTypeRequester) Get(store, feature string) (*featuretypes.FeatureType, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes/%s.json", ftr.data.Connection.URL, ftr.data.Workspace, store, feature)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(ftr.data.Connection.Credentials.Username, ftr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := ftr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var featureType featuretypes.GetFeatureTypeWrapper

		err = json.NewDecoder(response.Body).Decode(&featureType)
		if err != nil {
			return nil, err
		}

		return &featureType.FeatureType, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("featuretype %s not found", feature))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ftr *FeatureTypeRequester) Update(store, feature string, content []byte) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes/%s", ftr.data.Connection.URL, ftr.data.Workspace, store, feature)

	request, err := http.NewRequest(http.MethodPut, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(ftr.data.Connection.Credentials.Username, ftr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := ftr.data.Client.Do(request)
	if err != nil {
		return err
	}

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

func (ftr *FeatureTypeRequester) Reset(name string) error {
	return customerrors.NewNotImplementedError("not implemented")
}
