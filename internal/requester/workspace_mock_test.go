package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/internal/mock"
	"github.com/canghel3/go-geoserver/internal/testdata"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/workspace"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

const (
	singleWorkspaceResponse = "../testdata/workspace/single_workspace_response.json"
	multiWorkspace          = "../testdata/workspace/multi_workspace_response.json"
	noWorkspacesExist       = "../testdata/workspace/does_not_exist_response.json"
)

func TestWorkspaceRequester_Create(t *testing.T) {
	t.Run("201 Created", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusCreated,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Create(nil, false)
		assert.NoError(t, err)
	})

	t.Run("409 Conflict", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusConflict,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Create(nil, false)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.ConflictError{}, err)
		assert.EqualError(t, err, "workspace already exists")
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("some error")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Create(nil, false)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Create(nil, false)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}
		err := workspaceRequester.Create(nil, false)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestWorkspaceRequester_Delete(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Delete(testdata.Workspace, false)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Delete(testdata.Workspace, false)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace))
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("some error")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Delete(testdata.Workspace, false)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Delete(testdata.Workspace, false)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}
		err := workspaceRequester.Delete(testdata.Workspace, false)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestWorkspaceRequester_Get(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(singleWorkspaceResponse)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		wksp, err := workspaceRequester.Get(testdata.Workspace)
		assert.NoError(t, err)
		assert.NotNil(t, wksp)

		var expectedWorkspace workspace.GetSingleWorkspaceWrapper
		err = json.Unmarshal(content, &expectedWorkspace)
		assert.NoError(t, err)

		assert.Equal(t, expectedWorkspace.Workspace, *wksp)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		wksp, err := workspaceRequester.Get(testdata.Workspace)
		assert.Nil(t, wksp)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace))
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("some error")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := workspaceRequester.Get(testdata.Workspace)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := workspaceRequester.Get(testdata.Workspace)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("{")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := workspaceRequester.Get(testdata.Workspace)
		assert.Error(t, err)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := workspaceRequester.Get(testdata.Workspace)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestWorkspaceRequester_GetAll(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(multiWorkspace)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		wksp, err := workspaceRequester.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, wksp)

		var expectedWorkspace workspace.MultiWorkspaceRetrievalWrapper
		err = json.Unmarshal(content, &expectedWorkspace)
		assert.NoError(t, err)

		assert.Equal(t, expectedWorkspace.Workspaces.Workspace, wksp)
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("some error")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := workspaceRequester.GetAll()
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		_, err := workspaceRequester.GetAll()
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("No workspaces exist", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		content, err := testdata.Read(noWorkspacesExist)
		assert.NoError(t, err)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(content)),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		wksp, err := workspaceRequester.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, wksp)

		var expectedWorkspace workspace.MultiWorkspaceRetrievalWrapper
		err = json.Unmarshal(content, &expectedWorkspace)
		assert.NoError(t, err)

		assert.Equal(t, expectedWorkspace.Workspaces.Workspace, wksp)
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}
		_, err := workspaceRequester.GetAll()
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}

func TestWorkspaceRequester_Update(t *testing.T) {
	t.Run("200 Ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Update(nil, testdata.Workspace)
		assert.NoError(t, err)
	})

	t.Run("404 Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Update(nil, testdata.Workspace)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.NotFoundError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("workspace %s not found", testdata.Workspace))
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString("some error")),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Update(nil, testdata.Workspace)
		assert.Error(t, err)
		assert.IsType(t, &customerrors.GeoserverError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
	})

	t.Run("Invalid Body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockResponse := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Header:     make(http.Header),
			Body:       io.NopCloser(&testdata.ErrorReader{}),
		}

		mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}

		err := workspaceRequester.Update(nil, testdata.Workspace)
		assert.Error(t, err)
		assert.EqualError(t, err, "reader error")
	})

	t.Run("Client Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockClient := mocks.NewMockHTTPClient(ctrl)
		mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("client error"))

		workspaceRequester := &WorkspaceRequester{data: testdata.GeoserverInfo(mockClient)}
		err := workspaceRequester.Update(nil, testdata.Workspace)
		assert.Error(t, err)
		assert.EqualError(t, err, "client error")
	})
}
