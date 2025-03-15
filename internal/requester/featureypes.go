package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
	"io"
	"net/http"
)

type FeatureTypeRequester struct {
	info *internal.GeoserverInfo
}

func (ftr *FeatureTypeRequester) Create(store string, content []byte) error {
	var target string
	if len(store) == 0 {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/featuretypes", ftr.info.Connection.URL, ftr.info.Workspace)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes", ftr.info.Connection.URL, ftr.info.Workspace, store)
	}

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(ftr.info.Connection.Credentials.Username, ftr.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := ftr.info.Client.Do(request)
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

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ftr *FeatureTypeRequester) Delete(store, feature string, recurse bool) error {
	var target string
	if len(store) == 0 {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/featuretypes/%s?recurse=%v", ftr.info.Connection.URL, ftr.info.Workspace, feature, recurse)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes/%s?recurse=%v", ftr.info.Connection.URL, ftr.info.Workspace, store, feature, recurse)
	}

	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(ftr.info.Connection.Credentials.Username, ftr.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := ftr.info.Client.Do(request)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("featuretype %s does not exist", feature))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ftr *FeatureTypeRequester) Get(store, feature string) (*featuretypes.GetFeatureTypeWrapper, error) {
	var target string
	if len(store) == 0 {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/featuretypes/%s.json", ftr.info.Connection.URL, ftr.info.Workspace, feature)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes/%s.json", ftr.info.Connection.URL, ftr.info.Workspace, store, feature)
	}

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(ftr.info.Connection.Credentials.Username, ftr.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := ftr.info.Client.Do(request)
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

		return &featureType, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("featuretype %s does not exist", feature))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
