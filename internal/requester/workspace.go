package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/workspace"
	"io"
	"net/http"
)

type WorkspaceRequester struct {
	data internal.GeoserverData
}

func NewWorkspaceRequester(data internal.GeoserverData) WorkspaceRequester {
	return WorkspaceRequester{data: data}
}

func (wr WorkspaceRequester) Create(content []byte, _default bool) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces?default=%v", wr.data.Connection.URL, _default)
	request, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(wr.data.Connection.Credentials.Username, wr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := wr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusConflict:
		return customerrors.WrapConflictError(fmt.Errorf("workspace already exists"))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wr WorkspaceRequester) Get(name string) (*workspace.Workspace, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s", wr.data.Connection.URL, name)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(wr.data.Connection.Credentials.Username, wr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := wr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var wksp workspace.GetSingleWorkspaceWrapper
		err = json.NewDecoder(response.Body).Decode(&wksp)
		if err != nil {
			return nil, err
		}

		return &wksp.Workspace, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("workspace %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wr WorkspaceRequester) GetAll() ([]workspace.MultiWorkspace, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces", wr.data.Connection.URL)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(wr.data.Connection.Credentials.Username, wr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := wr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err

	}

	switch response.StatusCode {
	case http.StatusOK:
		var wksp workspace.MultiWorkspaceRetrievalWrapper
		err = json.Unmarshal(body, &wksp)
		if err != nil {
			type noWorkspaceExists struct {
				Workspaces string `json:"workspaces" xml:"workspaces"`
			}
			var noWorkspacesExistResponse noWorkspaceExists
			noWorkspacesExistError := json.Unmarshal(body, &noWorkspacesExistResponse)
			if noWorkspacesExistError == nil {
				return []workspace.MultiWorkspace{}, nil
			}

			return nil, err
		}

		return wksp.Workspaces.Workspace, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wr WorkspaceRequester) Update(content []byte, oldName string) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s", wr.data.Connection.URL, oldName)
	request, err := http.NewRequest(http.MethodPut, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(wr.data.Connection.Credentials.Username, wr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := wr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("workspace %s not found", oldName))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wr WorkspaceRequester) Delete(name string, recurse bool) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s?recurse=%v", wr.data.Connection.URL, name, recurse)
	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(wr.data.Connection.Credentials.Username, wr.data.Connection.Credentials.Password)

	response, err := wr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("workspace %s not found", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
