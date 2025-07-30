package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/layers"
	"io"
	"net/http"
)

type LayerGroupRequester struct {
	data internal.GeoserverData
}

func NewLayerGroupRequester(data internal.GeoserverData) LayerGroupRequester {
	return LayerGroupRequester{data: data}
}

func (lgr LayerGroupRequester) Get(name string) (*layers.Group, error) {
	var target string
	if validator.Empty(lgr.data.Workspace) {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups/%s", lgr.data.Connection.URL, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups/%s", lgr.data.Connection.URL, lgr.data.Workspace, name)
	}

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(lgr.data.Connection.Credentials.Username, lgr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := lgr.data.Client.Do(request)
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
		var layerGroupWrapper layers.GroupWrapper
		err := json.Unmarshal(body, &layerGroupWrapper)
		if err != nil {
			//TODO: handle case for single keyword in response and transform it
			//var interm models.GroupWrapperForSingleKeyword
			//intermErr := json.Unmarshal(body, &interm)
			//if intermErr == nil {
			//	return &layers.Group{
			//		Name:         interm.Group.Name,
			//		Mode:         interm.Group.Mode,
			//		Title:        interm.Group.Title,
			//		Workspace:    interm.Group.Workspace,
			//		Publishables: layers.Publishables{},
			//		Bounds:       shared.BoundingBoxCRSClass{},
			//		Keywords:     &shared.Keywords{Keywords: []string{interm.Group.Keywords.String}},
			//		Styles:       layers.GroupStyles{},
			//		DateCreated:  "",
			//		DateModified: "",
			//	}, nil
			//}

			return nil, err
		}

		return &layerGroupWrapper.Group, nil
	case http.StatusNotFound:
		return nil, customerrors.NewNotFoundError(fmt.Sprintf("layer group %s not found", name))
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (lgr LayerGroupRequester) Create(content []byte) error {
	var target string
	if validator.Empty(lgr.data.Workspace) {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups", lgr.data.Connection.URL)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups", lgr.data.Connection.URL, lgr.data.Workspace)
	}

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(lgr.data.Connection.Credentials.Username, lgr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := lgr.data.Client.Do(request)
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

func (lgr LayerGroupRequester) Update(name string, content []byte) error {
	var target string
	if validator.Empty(lgr.data.Workspace) {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups/%s", lgr.data.Connection.URL, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups/%s", lgr.data.Connection.URL, lgr.data.Workspace, name)
	}

	request, err := http.NewRequest(http.MethodPut, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(lgr.data.Connection.Credentials.Username, lgr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := lgr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.NewNotFoundError(fmt.Sprintf("layer group %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (lgr LayerGroupRequester) Delete(name string) error {
	var target string
	if validator.Empty(lgr.data.Workspace) {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups/%s", lgr.data.Connection.URL, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups/%s", lgr.data.Connection.URL, lgr.data.Workspace, name)
	}

	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(lgr.data.Connection.Credentials.Username, lgr.data.Connection.Credentials.Password)

	response, err := lgr.data.Client.Do(request)
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

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
