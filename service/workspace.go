package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg"
	"io"
	"net/http"
)

type WorkspaceService struct{}

/*
CreateWorkspace creates a workspace in Geoserver using the provided name.

- Empty names are not allowed.

- The name can only include alphanumerical characters. Any other characters will result in an InputError with a message indicating the restriction.

- If a workspace with the same name already exists, it returns a ConflictError with a message indicating the existing workspace.

- If the workspace name does not exist but any other error occurs while using GetWorkspace, the respective Go error is returned.

- If the request to create a workspace is successful (status code: http.StatusCreated), it returns a nil error.

- If Geoserver responds with a non-successful status code, it returns a GeoserverError providing the status code and server response.

- In case of issues like JSON marshalling, creating the request, or network issues, it returns the respective Go error.
*/
func (gs *GeoserverService) CreateWorkspace(name string) error {
	var enf *customerrors.NotFoundError

	_, err := gs.GetWorkspace(name)
	if err == nil {
		return customerrors.WrapConflictError(fmt.Errorf("workspace %s already exists", name))
	} else if err != nil && !errors.As(err, &enf) {
		return err
	}

	data := pkg.WorkspaceCreationWrapper{
		Workspace: pkg.WorkspaceCreation{
			Name: name,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/workspaces", gs.data.connection.URL), bytes.NewBuffer(content))
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

/*
GetWorkspaces retrieves information about all existing workspaces in the Geoserver.

- In a normal scenario, it returns a pointer to the MultiWorkspaceRetrievalWrapper structure and a nil error.
- If the Geoserver responds with a non-successful status code, it returns a nil pointer and a GeoserverError with a message providing the status code and server response.
- For other issues (like network or JSON parsing problems), it returns the respective Go error and a nil pointer.
*/
func (gs *GeoserverService) GetWorkspaces() (*pkg.MultiWorkspaceRetrievalWrapper, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces", gs.data.connection.URL)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(gs.data.connection.Credentials.Username, gs.data.connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := gs.data.client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var workspace pkg.MultiWorkspaceRetrievalWrapper
		err = json.NewDecoder(response.Body).Decode(&workspace)
		if err != nil {
			return nil, err
		}

		return &workspace, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

/*
GetWorkspace retrieves a single workspace from the Geoserver using the given workspace name.

- If workspace exists and the request is successful, it returns a pointer to the SingleWorkspaceRetrievalWrapper structure and a nil error.

- In case no such workspace exists, a NotFoundError is returned with the message indicating the missing workspace.

- If Geoserver returns a non-successful status code, it returns nil and a GeoserverError with a message containing status code and server response.

- In case of network issues or JSON parsing problems, it returns the respective Go error and a nil pointer.
*/
func (gs *GeoserverService) GetWorkspace(name string) (*pkg.SingleWorkspaceRetrievalWrapper, error) {
	err := internal.ValidateWorkspace(name)
	if err != nil {
		return nil, err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s", gs.data.connection.URL, name)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(gs.data.connection.Credentials.Username, gs.data.connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := gs.data.client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var workspace pkg.SingleWorkspaceRetrievalWrapper
		err = json.NewDecoder(response.Body).Decode(&workspace)
		if err != nil {
			return nil, err
		}

		return &workspace, nil
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

// UpdateWorkspace TODO: implement
func (gs *GeoserverService) UpdateWorkspace(name string) error {
	return nil
}

func (gs *GeoserverService) DeleteWorkspace(name string, options ...internal.Option) error {
	_, err := gs.GetWorkspace(name)
	if err != nil {
		return err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s", gs.data.connection.URL, name)
	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	params := internal.ProcessOptions(options)
	if recurse, set := params["recurse"]; set {
		q := request.URL.Query()
		q.Add("recurse", fmt.Sprintf("%v", recurse.(bool)))
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
