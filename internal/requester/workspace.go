package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/models/workspace"
	"io"
	"net/http"
)

type WorkspaceRequester struct {
	info *internal.GeoserverInfo
}

func (wr *WorkspaceRequester) Create(name string, _default bool) error {
	data := workspace.WorkspaceCreationWrapper{
		Workspace: workspace.WorkspaceCreation{
			Name: name,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces?default=%v", wr.info.Connection.URL, _default)
	request, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(wr.info.Connection.Credentials.Username, wr.info.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := wr.info.Client.Do(request)
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

func (wr *WorkspaceRequester) Get(name string) (*workspace.SingleWorkspaceRetrievalWrapper, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s", wr.info.Connection.URL, name)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(wr.info.Connection.Credentials.Username, wr.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := wr.info.Client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var wksp workspace.SingleWorkspaceRetrievalWrapper
		err = json.NewDecoder(response.Body).Decode(&wksp)
		if err != nil {
			return nil, err
		}

		return &wksp, nil
	case http.StatusNotFound:
		return nil, customerrors.WrapNotFoundError(fmt.Errorf("workspace %s does not exist", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wr *WorkspaceRequester) GetAll() (*workspace.MultiWorkspaceRetrievalWrapper, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces", wr.info.Connection.URL)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(wr.info.Connection.Credentials.Username, wr.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := wr.info.Client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var wksp workspace.MultiWorkspaceRetrievalWrapper
		err = json.NewDecoder(response.Body).Decode(&wksp)
		if err != nil {
			var noWorkspacesExistResponse workspace.NoWorkspacesExist
			noWorkspacesExistError := json.NewDecoder(response.Body).Decode(&noWorkspacesExistResponse)
			if noWorkspacesExistError == nil {
				return &workspace.MultiWorkspaceRetrievalWrapper{
					Workspaces: workspace.MultiWorkspaceRetrieval{
						Workspace: nil,
					},
				}, nil
			}

			return nil, err
		}

		return &wksp, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wr *WorkspaceRequester) Delete(name string, recurse bool) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s?recurse=%v", wr.info.Connection.URL, name, recurse)
	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(wr.info.Connection.Credentials.Username, wr.info.Connection.Credentials.Password)

	response, err := wr.info.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return customerrors.WrapNotFoundError(fmt.Errorf("workspace %s does not exist", name))
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
