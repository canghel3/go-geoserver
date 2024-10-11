package service

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/layers"
	"github.com/canghel3/go-geoserver/utils"
	"io"
	"net/http"
)

func (gs *GeoserverService) GetLayer(name string, options ...utils.Option) (*layers.LayerWrapper, error) {
	params := utils.ProcessOptions(options)

	var target string
	if wksp, set := params["workspace"]; set {
		_, err := gs.GetWorkspace(wksp.(string))
		if err != nil {
			return nil, err
		}

		target = fmt.Sprintf("%s/geoserver/rest/layers/%s:%s", gs.url, wksp, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/layers/%s", gs.url, name)
	}

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(gs.username, gs.password)
	request.Header.Add("Accept", "application/json")

	response, err := gs.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var layer layers.LayerWrapper
		err = json.NewDecoder(response.Body).Decode(&layer)
		if err != nil {
			return nil, err
		}

		return &layer, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("layer %s does not exist", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
Deletes a layer from geoserver.

Available options: WorkspaceOption, RecurseOption
*/
func (gs *GeoserverService) DeleteLayer(name string, options ...utils.Option) error {
	_, err := gs.GetLayer(name, options...)
	if err != nil {
		return err
	}

	params := utils.ProcessOptions(options)

	var target string
	if wksp, set := params["workspace"]; set {
		_, err = gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/rest/layers/%s:%s", gs.url, wksp, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/layers/%s", gs.url, name)
	}

	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	if recurse, set := params["recurse"]; set {
		q := request.URL.Query()
		q.Add("recurse", fmt.Sprintf("%v", recurse.(bool)))
		request.URL.RawQuery = q.Encode()
	}

	request.SetBasicAuth(gs.username, gs.password)

	response, err := gs.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("layer %s does not exist", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
