package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/coveragestore"
	"github.com/canghel3/go-geoserver/internal/workspace"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

func (gs *GeoserverService) GetCoverageStore(workspace, store string) (*coveragestore.GetCoverageStore, error) {
	err := internal.ValidateStore(store)
	if err != nil {
		return nil, err
	}

	_, err = gs.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s", gs.data.connection.URL, workspace, store)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(gs.data.connection.Credentials.Username, gs.data.connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := gs.data.client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var coverageStore coveragestore.GetCoverageStoreWrapper
		err = json.NewDecoder(response.Body).Decode(&coverageStore)
		if err != nil {
			return nil, err
		}

		return &coverageStore.CoverageStore, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("coverage store %s does not exist", store))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
CreateCoverageStore creates a coverage store in Geoserver.
URL is the url of the file (file location inside geoserver)

Available types (case-insensitive): GeoTIFF
*/
func (gs *GeoserverService) CreateCoverageStore(wksp, name, url, type_ string) error {
	switch strings.ToLower(type_) {
	case "geotiff":
		if !strings.HasSuffix(url, ".tif") {
			return customerrors.NewInputError(fmt.Sprintf("file must be of type .tif, got %s", filepath.Ext(url)))
		}
	default:
		return customerrors.NewUnsupportedError(fmt.Sprintf("unsupported coveragestore type %s", type_))
	}

	if !strings.HasPrefix(url, "file:") {
		//no prefix
		dir, f := filepath.Split(url)
		if len(dir) == 0 {
			if gs.isDataDirectorySet() {
				if strings.HasPrefix(gs.getDataDirectory(), "file:") {
					//provided url is only of file name and data dir is set in geoserver
					url = filepath.Join(gs.getDataDirectory(), f)
				} else {
					//provided url is only of file name and data dir is set in geoserver
					url = fmt.Sprintf("file:%s", filepath.Join(gs.getDataDirectory(), f))
				}
			} else {
				url = fmt.Sprintf("file:%s", url)
			}
		} else {
			//provided url has supposedly fully qualified url
			url = fmt.Sprintf("file:%s", url)
		}
	}

	coverageStore := coveragestore.CreateCoverageStoreWrapper{
		CoverageStore: coveragestore.CreateCoverageStore{
			Name: name,
			Type: type_,
			Workspace: workspace.MultiWorkspace{
				Name: wksp,
				Href: fmt.Sprintf("%s/geoserver/rest/namespaces/%s.json", gs.data.connection.URL, wksp),
			},
			Enabled: true,
			URL:     url,
		},
	}

	content, err := json.Marshal(coverageStore)
	if err != nil {
		return err
	}

	return gs.createCoverageStore(wksp, name, content)
}

/*
DeleteCoverageStore deletes a coverage store from Geoserver.

Available options: RecurseOption, PurgeOption
*/
func (gs *GeoserverService) DeleteCoverageStore(wksp, store string, options ...internal.Option) error {
	_, err := gs.GetWorkspace(wksp)
	if err != nil {
		return err
	}

	_, err = gs.GetCoverageStore(wksp, store)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s", gs.data.connection.URL, wksp, store), nil)
	if err != nil {
		return err
	}

	params := internal.ProcessOptions(options)
	if recurse, set := params["recurse"]; set {
		q := request.URL.Query()
		q.Add("recurse", fmt.Sprintf("%v", recurse.(bool)))
		request.URL.RawQuery = q.Encode()
	}

	if purge, set := params["purge"]; set {
		q := request.URL.Query()
		q.Add("purge", purge.(string))
		request.URL.RawQuery = q.Encode()
	}

	request.SetBasicAuth(gs.data.connection.Credentials.Username, gs.data.connection.Credentials.Password)

	response, err := gs.data.client.Do(request)
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

func (gs *GeoserverService) createCoverageStore(wksp, name string, body []byte) error {
	_, err := gs.GetWorkspace(wksp)
	if err != nil {
		return err
	}

	_, err = gs.GetCoverageStore(wksp, name)
	if err == nil {
		return customerrors.WrapConflictError(fmt.Errorf("coveragestore %s already exists", name))
	}

	target := fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores", gs.data.connection.URL, wksp)
	request, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gs.data.connection.Credentials.Username, gs.data.connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := gs.data.client.Do(request)
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
