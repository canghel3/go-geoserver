package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/style"
	"github.com/canghel3/go-geoserver/models/workspace"
	"github.com/canghel3/go-geoserver/utils"
	"io"
	"net/http"
)

// CreateStyle creates a style in Geoserver.
//
// Parameters:
//   - name: the name of the style
//   - format: the format of the style ("css", "sld", or "zip")
//   - content: the content of the style
//   - options: optional parameters for creating the style (WorkspaceOption for creating the style in a specific workspace)
//
// Returns:
//   - error: an error if any occurred
//
// Example:
// err := gs.CreateStyle("myStyle", "css", []byte(cssContent))
//
// Note: "sld" and "zip" formats are currently not supported for creating styles.
func (gs *GeoserverService) CreateStyle(name, format string, content []byte, options ...utils.Option) error {
	var enf *customerrors.NotFoundError

	_, err := gs.GetStyle(name, format, options...)
	if err == nil {
		return customerrors.WrapConflictError(fmt.Errorf("style %s already exists", name))
	} else if err != nil && !errors.As(err, &enf) {
		return err
	}

	params := utils.ProcessOptions(options)

	var target string
	if wksp, set := params["workspace"]; set {
		_, err := gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/styles", gs.url, wksp)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/styles", gs.url)
	}

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gs.username, gs.password)

	switch format {
	case "css":
		request.Header.Add("Content-Type", "application/vnd.geoserver.geocss+css")
	case "sld":
		return customerrors.NewNotImplementedError("sld style creation is not implemented")
	case "zip":
		return customerrors.NewNotImplementedError("zip style creation is not implemented")
	}

	q := request.URL.Query()
	q.Add("name", name)
	request.URL.RawQuery = q.Encode()

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

func (gs *GeoserverService) GetStyle(name, format string, options ...utils.Option) (*style.GetStyleWrapper, error) {
	err := utils.ValidateStyle(name)
	if err != nil {
		return nil, err
	}

	params := utils.ProcessOptions(options)

	var target string
	wksp, set := params["workspace"]
	if set {
		_, err := gs.GetWorkspace(wksp.(string))
		if err != nil {
			return nil, err
		}

		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/styles/%s", gs.url, wksp, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/styles/%s", gs.url, name)
	}

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(gs.username, gs.password)

	switch format {
	case "css":
		request.Header.Add("Accept", "application/vnd.geoserver.geocss+css")
	case "sld 1.1.0":
		request.Header.Add("Accept", "application/vnd.ogc.se+xml")
	case "sld 1.0.0":
		request.Header.Add("Accept", "application/vnd.ogc.sld+xml")
	case "zip":
		return nil, fmt.Errorf("retrieving zip styles is currently not supported")
	default:
		request.Header.Add("Accept", "application/json")
	}

	response, err := gs.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var s = style.GetStyleWrapper{
			Style: style.GetStyleDetails{
				Name: name,
			},
		}

		if set {
			s.Style.Workspace = &workspace.WorkspaceCreation{Name: wksp.(string)}
		}

		switch format {
		case "css":
			s.Style.Format = format
			content, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}

			s.Style.Content = string(content)
			return &s, nil
		case "sld 1.1.0", "sld 1.0.0":
			s.Style.Format = format
			content, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}

			s.Style.Content = string(content)
			return &s, nil
		case "zip":
			return nil, customerrors.NewNotImplementedError("zip style retrieval is not implemented")
		default:
			err = json.NewDecoder(response.Body).Decode(&s)
			if err != nil {
				return nil, err
			}

			return &s, nil
		}
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("style %s does not exist", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
Available options: WorkspaceOption, PurgeOption (only "true" works in this case), RecurseOption
*/
func (gs *GeoserverService) DeleteStyle(name string, options ...utils.Option) error {
	_, err := gs.GetStyle(name, "json", options...)
	if err != nil {
		return err
	}

	params := utils.ProcessOptions(options)
	var target string
	wksp, set := params["workspace"]
	if set {
		_, err := gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/styles/%s", gs.url, wksp, name)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/styles/%s", gs.url, name)
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

	if purge, set := params["purge"]; set {
		q := request.URL.Query()
		q.Add("purge", fmt.Sprintf("%v", purge.(string)))
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
		return customerrors.NewGeoserverError("style is referenced by other layers and recurse option is missing or set to false")
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
Available options: WorkspaceOption
*/
func (gs *GeoserverService) StyleLayer(layer, style_, format string, default_ bool, options ...utils.Option) error {
	params := utils.ProcessOptions(options)

	_, err := gs.GetLayer(layer, options...)
	if err != nil {
		return err
	}

	_, err = gs.GetStyle(style_, format, options...)
	if err != nil {
		return err
	}

	var target string
	wksp, set := params["workspace"]
	if set {
		_, err = gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/rest/layers/%s:%s/styles?default=%v", gs.url, wksp, layer, default_)
	} else {
		target = fmt.Sprintf("%s/geoserver/rest/layers/%s/styles?default=%v", gs.url, layer, default_)
	}

	var styleContent style.StyleWrapper

	if set {
		styleContent = style.StyleWrapper{
			Style: style.StyleDetails{
				Name: fmt.Sprintf("%s:%s", wksp.(string), style_),
			},
		}
	} else {
		styleContent = style.StyleWrapper{
			Style: style.StyleDetails{
				Name: style_,
			},
		}
	}

	content, err := json.Marshal(styleContent)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(content))
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
