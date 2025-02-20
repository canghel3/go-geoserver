package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/layergroups"
	"github.com/canghel3/go-geoserver/models/misc"
	"github.com/canghel3/go-geoserver/utils"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func (gs *GeoserverService) GetLayerGroups() error {
	return customerrors.NewNotImplementedError("route not implemented")
}

/*
Available options: WorkspaceOption
*/
func (gs *GeoserverService) GetLayerGroup(name string, opts ...utils.Option) (*layergroups.GetGroupWrapper, error) {
	err := utils.ValidateLayer(name)
	if err != nil {
		return nil, err
	}

	var target string

	params := utils.ProcessOptions(opts)
	if wksp, set := params["workspace"]; set {
		_, err = gs.GetWorkspace(wksp.(string))
		if err != nil {
			return nil, err
		}

		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups/%s", gs.url, wksp, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups/%s", gs.url, name)
	}

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(gs.username, gs.password)

	response, err := gs.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var group layergroups.GetGroupWrapper
		err = json.NewDecoder(response.Body).Decode(&group)
		if err != nil {
			return nil, err
		}

		return &group, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("layergroup %s does not exist", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
Available options: WorkspaceOption
*/
func (gs *GeoserverService) DeleteLayerGroup(name string, opts ...utils.Option) error {
	_, err := gs.GetLayerGroup(name, opts...)
	if err != nil {
		return err
	}

	params := utils.ProcessOptions(opts)

	var target string
	if wksp, set := params["workspace"]; set {
		_, err := gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups/%s", gs.url, wksp, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups/%s", gs.url, name)
	}

	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
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
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (gs *GeoserverService) CreateLayerGroup(name, srs string, children []string, bbox [4]float64, opts ...utils.Option) error {
	var enf *customerrors.NotFoundError
	_, err := gs.GetLayerGroup(name, opts...)
	if err == nil {
		return customerrors.NewConflictError(fmt.Sprintf("layergroup %s already exists", name))
	} else if !errors.As(err, &enf) {
		return err
	}

	groupContent := layergroups.GroupWrapper{
		Group: layergroups.Group{
			Name:         name,
			Mode:         "SINGLE",
			Title:        name,
			Publishables: layergroups.Publishables{},
			Bounds: misc.BoundingBox{
				MinX: bbox[0],
				MinY: bbox[1],
				MaxX: bbox[2],
				MaxY: bbox[3],
				CRS:  srs,
			},
			Styles: layergroups.GroupStyles{},
		},
	}

	var target string

	params := utils.ProcessOptions(opts)
	if wksp, set := params["workspace"]; set {
		_, err = gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups", gs.url, wksp)
		groupContent.Group.Workspace.Name = wksp.(string)

		publishables := make([]layergroups.Layer, 0, len(children))
		for _, child := range children {
			publishables = append(publishables, layergroups.Layer{
				Type: "layer",
				Name: fmt.Sprintf("%s:%s", wksp.(string), child),
				Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layers/%s.json", gs.url, wksp.(string), child),
			})
		}
		groupContent.Group.Publishables = layergroups.Publishables{Published: publishables}
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups", gs.url)
		publishables := make([]layergroups.Layer, 0, len(children))
		for _, child := range children {
			publishables = append(publishables, layergroups.Layer{
				Type: "layer",
				Name: child,
				Link: fmt.Sprintf("%s/geoserver/rest/layers/%s.json", gs.url, child),
			})
		}
		groupContent.Group.Publishables = layergroups.Publishables{Published: publishables}
	}

	groupStyles := make([]misc.Style, 0, len(children))
	for range children {
		groupStyles = append(groupStyles, misc.Style{})
	}
	groupContent.Group.Styles.Style = groupStyles

	ftVal := reflect.ValueOf(&groupContent.Group).Elem()
	for i := 0; i < ftVal.NumField(); i++ {
		field := ftVal.Type().Field(i)
		f := strings.Split(field.Tag.Get("json"), ",")[0]
		//ignore workspace option in this case because we treated it earlier
		if value, ok := params[f]; ok && f != "workspace" {
			if ftVal.Field(i).CanSet() {
				ftVal.Field(i).Set(reflect.ValueOf(value))
			} else {
				return fmt.Errorf("cannot set option value for field %s", field.Name)
			}
		}
	}

	content, err := json.Marshal(groupContent)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(content))
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

func (gs *GeoserverService) UpdateLayerGroup(name, srs string, layersToAdd []string, bbox [4]float64, opts ...utils.Option) error {
	var (
		enf *customerrors.NotFoundError
	)

	lg, err := gs.GetLayerGroup(name, opts...)
	if err != nil {
		if errors.As(err, &enf) {
			return customerrors.WrapNotFoundError(fmt.Errorf("layer group %s not found", name))
		} else if !errors.As(err, &enf) {
			return err
		}
	}

	groupContent := layergroups.GroupWrapper{
		Group: layergroups.Group{
			Name:         name,
			Mode:         "SINGLE",
			Title:        name,
			Publishables: lg.Group.Publishables,
			Bounds: misc.BoundingBox{
				MinX: bbox[0],
				MinY: bbox[1],
				MaxX: bbox[2],
				MaxY: bbox[3],
				CRS:  srs,
			},
			Styles: layergroups.GroupStyles{},
		},
	}

	var target string

	params := utils.ProcessOptions(opts)
	if wksp, set := params["workspace"]; set {
		_, err = gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layergroups/%s", gs.url, wksp, name)
		groupContent.Group.Workspace.Name = wksp.(string)

		publishables := make([]layergroups.Layer, len(lg.Group.Publishables.Published))
		copy(publishables, lg.Group.Publishables.Published)

		for _, child := range layersToAdd {
			publishables = append(publishables, layergroups.Layer{
				Type: "layer",
				Name: fmt.Sprintf("%s:%s", wksp.(string), child),
				Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s/layers/%s.json", gs.url, wksp.(string), child),
			})
		}
		groupContent.Group.Publishables = layergroups.Publishables{Published: publishables}
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/layergroups/%s", gs.url, name)
		publishables := make([]layergroups.Layer, len(lg.Group.Publishables.Published))
		copy(publishables, lg.Group.Publishables.Published)

		for _, child := range layersToAdd {
			publishables = append(publishables, layergroups.Layer{
				Type: "layer",
				Name: child,
				Link: fmt.Sprintf("%s/geoserver/rest/layers/%s.json", gs.url, child),
			})
		}
		groupContent.Group.Publishables = layergroups.Publishables{Published: publishables}
	}

	groupStyles := make([]misc.Style, 0, len(lg.Group.Publishables.Published)+len(layersToAdd))
	for range len(lg.Group.Publishables.Published) + len(layersToAdd) {
		groupStyles = append(groupStyles, misc.Style{})
	}
	groupContent.Group.Styles.Style = groupStyles

	ftVal := reflect.ValueOf(&groupContent.Group).Elem()
	for i := 0; i < ftVal.NumField(); i++ {
		field := ftVal.Type().Field(i)
		f := strings.Split(field.Tag.Get("json"), ",")[0]
		//ignore workspace option in this case because we treated it earlier
		if value, ok := params[f]; ok && f != "workspace" {
			if ftVal.Field(i).CanSet() {
				ftVal.Field(i).Set(reflect.ValueOf(value))
			} else {
				return fmt.Errorf("cannot set option value for field %s", field.Name)
			}
		}
	}

	content, err := json.Marshal(groupContent)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, target, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gs.username, gs.password)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	response, err := gs.client.Do(request)
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
