package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/coverage/geotiff"
	"github.com/canghel3/go-geoserver/models/misc"
	"github.com/canghel3/go-geoserver/models/workspace"
	"github.com/canghel3/go-geoserver/utils"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func (gs *GeoserverService) GetCoverage(workspace, store, name string) (*geotiff.CoverageWrapper, error) {
	err := utils.ValidateLayer(name)
	if err != nil {
		return nil, err
	}

	_, err = gs.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	_, err = gs.GetCoverageStore(workspace, store)
	if err != nil {
		return nil, err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s", gs.data.Connection.URL, workspace, store, name)

	request, err := http.NewRequest(http.MethodGet, target, nil)
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
		var coverage geotiff.CoverageWrapper
		err = json.NewDecoder(response.Body).Decode(&coverage)
		if err != nil {
			return nil, err
		}

		return &coverage, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("coverage %s does not exist", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
CreateCoverage creates a coverage in geoserver.
BBOX values should adhere to the specified srs.

Available options: TitleOption, ProjectionPolicyOption

TODO include example files with CoverageDimension Option
*/
func (gs *GeoserverService) CreateCoverage(wksp, store, layerName, srs string, bbox [4]float64, options ...utils.Option) error {
	var enf *customerrors.NotFoundError

	_, err := gs.GetWorkspace(wksp)
	if err != nil {
		return err
	}

	_, err = gs.GetCoverageStore(wksp, store)
	if err != nil {
		return err
	}

	_, err = gs.GetCoverage(wksp, store, layerName)
	if err == nil {
		return customerrors.WrapConflictError(fmt.Errorf("coverage %s already exists", layerName))
	} else if !errors.As(err, &enf) {
		return err
	}

	coverage := geotiff.CoverageWrapper{
		Coverage: geotiff.CoverageDetails{
			Enabled: true,
			Name:    layerName,
			Namespace: workspace.MultiWorkspace{
				Name: wksp,
				Href: fmt.Sprintf("%s/geoserver/rest/namespaces/%s.json", gs.data.Connection.URL, wksp),
			},
			NativeBoundingBox: misc.BoundingBox{
				MinX: bbox[0],
				MinY: bbox[1],
				MaxX: bbox[2],
				MaxY: bbox[3],
				CRS:  srs,
			},
			ProjectionPolicy: "NONE",
			Srs:              srs,
			Store: geotiff.StoreDetails{
				Class: "coverageStore",
				Name:  fmt.Sprintf("%s:%s", wksp, store),
				Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s.json", gs.data.Connection.URL, wksp, store),
			},
			Title: layerName,
		},
	}

	params := utils.ProcessOptions(options)
	//TODO: extract as function
	ftVal := reflect.ValueOf(&coverage.Coverage).Elem()
	for i := 0; i < ftVal.NumField(); i++ {
		field := ftVal.Type().Field(i)
		if value, ok := params[strings.Split(field.Tag.Get("json"), ",")[0]]; ok {
			if ftVal.Field(i).CanSet() {
				ftVal.Field(i).Set(reflect.ValueOf(value))
			} else {
				return fmt.Errorf("cannot set option value for field %s", field.Name)
			}
		}
	}

	content, err := json.Marshal(coverage)
	if err != nil {
		return err
	}

	return gs.createCoverage(wksp, store, content)
}

func (gs *GeoserverService) createCoverage(workspace, store string, coverage []byte) error {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages", gs.data.Connection.URL, workspace, store), bytes.NewReader(coverage))
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

// DeleteCoverage deletes a coverage from Geoserver.
// The recurse controls recursive deletion. When set to true all stores containing the resource are also removed.
func (gs *GeoserverService) DeleteCoverage(workspace, store, name string, options ...utils.Option) error {
	_, err := gs.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	_, err = gs.GetCoverageStore(workspace, store)
	if err != nil {
		return err
	}

	_, err = gs.GetCoverage(workspace, store, name)
	if err != nil {
		return err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s/coverages/%s", gs.data.Connection.URL, workspace, store, name)
	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	params := utils.ProcessOptions(options)
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
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
