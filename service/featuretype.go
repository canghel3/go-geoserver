package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/featuretype"
	"github.com/canghel3/go-geoserver/models/misc"
	"github.com/canghel3/go-geoserver/utils"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func (gs *GeoserverService) CreateFeatureType(workspace, store, layerName, sourceTableName, srs string, bbox [4]float64, options ...utils.Option) error {
	_, err := gs.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	_, err = gs.GetDataStore(workspace, store)
	if err != nil {
		return err
	}

	var enf *customerrors.NotFoundError
	_, err = gs.GetFeatureType(workspace, store, layerName)
	if err == nil {
		return customerrors.WrapConflictError(fmt.Errorf("featuretype %s already exists", layerName))
	} else if err != nil && !errors.As(err, &enf) {
		return err
	}

	data := featuretype.FeatureTypeWrapper{
		FeatureType: featuretype.FeatureType{
			Name:       layerName,
			NativeName: sourceTableName,
			Namespace: featuretype.Namespace{
				Name: workspace,
				Href: fmt.Sprintf("%s/geoserver/rest/namespaces/%s.json", gs.url, workspace),
			},
			Srs: srs,
			NativeBoundingBox: misc.BoundingBox{
				MinX: bbox[0],
				MinY: bbox[1],
				MaxX: bbox[2],
				MaxY: bbox[3],
				CRS:  srs,
			},
			ProjectionPolicy: "NONE",
			Title:            layerName,
			Store: featuretype.Store{
				Class: "dataStore",
				Name:  fmt.Sprintf("%s:%s", workspace, store),
				Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s.json", gs.url, workspace, store),
			},
		},
	}

	params := utils.ProcessOptions(options)

	ftVal := reflect.ValueOf(&data.FeatureType).Elem()
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

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes", gs.url, url.PathEscape(workspace), url.PathEscape(store)), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gs.username, gs.password)
	request.Header.Add("Content-Type", "application/json")

	response, err := gs.client.Do(request)
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

func (gs *GeoserverService) GetFeatureType(workspace, store, layer string) (*featuretype.GetFeatureType, error) {
	err := utils.ValidateLayer(layer)
	if err != nil {
		return nil, err
	}

	_, err = gs.GetWorkspace(workspace)
	var nfe *customerrors.NotFoundError
	if err != nil && errors.Is(err, nfe) {
		return nil, err
	}

	_, err = gs.GetDataStore(workspace, store)
	if err != nil && errors.Is(err, nfe) {
		return nil, err
	}

	target := fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes/%s.json", gs.url, workspace, store, layer)
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

	switch response.StatusCode {
	case http.StatusOK:
		var featureType featuretype.GetFeatureType

		err = json.NewDecoder(response.Body).Decode(&featureType)
		if err != nil {
			return nil, err
		}

		return &featureType, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("featuretype %s does not exist", layer))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
DeleteFeatureType available options: RecurseOption
*/
func (gs *GeoserverService) DeleteFeatureType(workspace, store, layer string, options ...utils.Option) error {
	_, err := gs.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	_, err = gs.GetDataStore(workspace, store)
	if err != nil {
		return err
	}

	_, err = gs.GetFeatureType(workspace, store, layer)
	if err != nil {
		return err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s/featuretypes/%s", gs.url, workspace, store, layer)
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

	request.SetBasicAuth(gs.username, gs.password)

	response, err := gs.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusForbidden:
		return customerrors.WrapGeoserverError(errors.New("layer is being referenced by other resources"))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
