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

const SINGLE_WORKSPACE_RESPONSE = "../../testdata/workspace/get_single_response.json"

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
			assert.ErrorContains(t, err, "workspace already exists")
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
			assert.ErrorContains(t, err, fmt.Sprintf("workspace %s does not exist", testdata.WORKSPACE))
			assert.ErrorAs(t, err, &enotfound)
		})
	})
}
