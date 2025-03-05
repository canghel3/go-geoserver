package workspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/workspace"
	"github.com/canghel3/go-geoserver/vector"
	"io"
	"net/http"
)

type Service struct {
	client *http.Client
	info   *internal.GeoserverInfo
}

type ServiceSelector struct {
	info *internal.GeoserverInfo
}

func NewService() *Service {
	return &Service{
		client: &http.Client{},
	}
}

func (s *Service) Create(name string, _default bool) error {
	data := workspace.WorkspaceCreationWrapper{
		Workspace: workspace.WorkspaceCreation{
			Name: name,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces?default=%v", s.info.Connection.URL, _default)
	request, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := s.client.Do(request)
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

func (s *Service) Get(name string) (*workspace.SingleWorkspaceRetrievalWrapper, error) {
	err := internal.ValidateWorkspace(name)
	if err != nil {
		return nil, err
	}

	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s", s.info.Connection.URL, name)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := s.client.Do(request)
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

func (s *Service) GetAll() (*workspace.MultiWorkspaceRetrievalWrapper, error) {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces", s.info.Connection.URL)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		var workspace workspace.MultiWorkspaceRetrievalWrapper
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

func (s *Service) Delete(name string, recurse bool) error {
	var target = fmt.Sprintf("%s/geoserver/rest/workspaces/%s?recurse=%v", s.info.Connection.URL, name, recurse)
	request, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(s.info.Connection.Credentials.Username, s.info.Connection.Credentials.Password)

	response, err := s.client.Do(request)
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

func (s *Service) Use(workspace string) *ServiceSelector {
	return &ServiceSelector{
		info: &internal.GeoserverInfo{
			Client:     s.client,
			Connection: internal.GeoserverConnection{},
			DataDir:    "",
			Workspace:  "",
		},
	}
}

func (ss *ServiceSelector) Vectors() *vector.Service {
	return vector.NewService(ss.info)
}
