//go:build mocks

package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal/mocks"
	"github.com/canghel3/go-geoserver/pkg/models/workspace"
	"github.com/canghel3/go-geoserver/testdata"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

const (
	SINGLE_WORKSPACE_RESPONSE    = "../../testdata/workspace/get_single_response.json"
	MULTI_WORKSPACE_RESPONSE     = "../../testdata/workspace/get_multi_response.json"
	NO_WORKSPACES_EXIST_RESPONSE = "../../testdata/workspace/no_workspaces_exist_response.json"
)

func TestWorkspaceRequester(t *testing.T) {
	t.Run("CREATE", func(t *testing.T) {
		t.Run("201 CREATED", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusCreated,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			err := workspaceRequester.Create(testdata.WORKSPACE, false)
			assert.NoError(t, err)
		})

		t.Run("409 CONFLICT", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusConflict,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var econflict *customerrors.ConflictError
			err := workspaceRequester.Create(testdata.WORKSPACE, false)
			assert.Error(t, err)
			assert.EqualError(t, err, "workspace already exists")
			assert.ErrorAs(t, err, &econflict)
		})

		t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("some error")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var econflict *customerrors.GeoserverError
			err := workspaceRequester.Create(testdata.WORKSPACE, false)
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
			assert.ErrorAs(t, err, &econflict)
		})
	})

	t.Run("GET", func(t *testing.T) {
		t.Run("200 OK", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			content, err := testdata.Read(SINGLE_WORKSPACE_RESPONSE)
			assert.NoError(t, err)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(content)),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			wksp, err := workspaceRequester.Get(testdata.WORKSPACE)
			assert.NoError(t, err)
			assert.NotNil(t, wksp)

			var expectedWorkspace *workspace.SingleWorkspaceRetrievalWrapper
			err = json.Unmarshal(content, &expectedWorkspace)
			assert.NoError(t, err)

			assert.Equal(t, expectedWorkspace, wksp)
		})

		t.Run("404 NOT FOUND", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var enotfound *customerrors.NotFoundError
			wksp, err := workspaceRequester.Get(testdata.WORKSPACE)
			assert.Nil(t, wksp)
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.WORKSPACE))
			assert.ErrorAs(t, err, &enotfound)
		})

		t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("some error")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var econflict *customerrors.GeoserverError
			_, err := workspaceRequester.Get(testdata.WORKSPACE)
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
			assert.ErrorAs(t, err, &econflict)
		})
	})

	t.Run("GET ALL", func(t *testing.T) {
		t.Run("NO WORKSPACES EXIST", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			content, err := testdata.Read(NO_WORKSPACES_EXIST_RESPONSE)
			assert.NoError(t, err)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(content)),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			wksp, err := workspaceRequester.GetAll()
			assert.NoError(t, err)
			assert.NotNil(t, wksp)

			var expectedWorkspace *workspace.MultiWorkspaceRetrievalWrapper
			err = json.Unmarshal(content, &expectedWorkspace)
			assert.NoError(t, err)

			assert.Equal(t, expectedWorkspace, wksp)
		})

		t.Run("200 OK", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			content, err := testdata.Read(MULTI_WORKSPACE_RESPONSE)
			assert.NoError(t, err)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(content)),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			wksp, err := workspaceRequester.GetAll()
			assert.NoError(t, err)
			assert.NotNil(t, wksp)

			var expectedWorkspace *workspace.MultiWorkspaceRetrievalWrapper
			err = json.Unmarshal(content, &expectedWorkspace)
			assert.NoError(t, err)

			assert.Equal(t, expectedWorkspace, wksp)
		})

		t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("some error")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var econflict *customerrors.GeoserverError
			_, err := workspaceRequester.GetAll()
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
			assert.ErrorAs(t, err, &econflict)
		})
	})

	t.Run("UPDATE", func(t *testing.T) {
		t.Run("200 OK", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			err := workspaceRequester.Update(testdata.WORKSPACE, "newName")
			assert.NoError(t, err)
		})

		t.Run("404 NOT FOUND", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var econflict *customerrors.NotFoundError
			err := workspaceRequester.Update(testdata.WORKSPACE, "newName")
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.WORKSPACE))
			assert.ErrorAs(t, err, &econflict)
		})

		t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("some error")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var econflict *customerrors.GeoserverError
			err := workspaceRequester.Update(testdata.WORKSPACE, "newName")
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
			assert.ErrorAs(t, err, &econflict)
		})
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Run("200 OK", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			err := workspaceRequester.Delete(testdata.WORKSPACE, false)
			assert.NoError(t, err)
		})

		t.Run("404 NOT FOUND", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var enotfound *customerrors.NotFoundError
			err := workspaceRequester.Delete(testdata.WORKSPACE, false)
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("workspace %s does not exist", testdata.WORKSPACE))
			assert.ErrorAs(t, err, &enotfound)
		})

		t.Run("500 INTERNAL SERVER ERROR", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockClient := mocks.NewMockHTTPClient(ctrl)
			mockResponse := &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString("some error")),
			}

			mockClient.EXPECT().Do(gomock.Any()).Return(mockResponse, nil)

			workspaceRequester := &WorkspaceRequester{info: testdata.GeoserverInfo(mockClient)}

			var econflict *customerrors.GeoserverError
			err := workspaceRequester.Delete(testdata.WORKSPACE, false)
			assert.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("received status code %d from geoserver: some error", http.StatusInternalServerError))
			assert.ErrorAs(t, err, &econflict)
		})
	})
}
