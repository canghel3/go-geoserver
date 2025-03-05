package service

import (
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/datastore/postgis"
	"github.com/canghel3/go-geoserver/internal/workspace"
	"gotest.tools/v3/assert"
	"testing"
)

func TestStyle(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)
	assert.NilError(t, geoserverService.CreateWorkspace("init"))
	styleContent := []byte("*{raster-channels: auto; raster-color-map-type:intervals; raster-color-map: color-map-entry('#febb81', 1189.02001953125, 1, '1189.02001953125') color-map-entry('#f8765c', 310000.0, 1, '310000.0') color-map-entry('#d3436e', 660000.0, 1, '660000.0') color-map-entry('#982d80', 1403927.675, 1, '1403927.675') color-map-entry('#5f187f', 3581199.250000002, 1, '3581199.250000002') color-map-entry('#221150', 2671396865.0, 1, '2671396865.0')}")

	t.Run("CREATE", func(t *testing.T) {
		t.Run("CSS STYLE", func(t *testing.T) {
			assert.NilError(t, geoserverService.CreateStyle("init_css_no_workspace", "css", styleContent))
		})

		t.Run("CSS STYLE IN WORKSPACE", func(t *testing.T) {
			assert.NilError(t, geoserverService.CreateStyle("init_css", "css", styleContent, internal.WorkspaceOption("init")))
		})

		t.Run("CSS STYLE IN NON-EXISTENT WORKSPACE", func(t *testing.T) {
			assert.Error(t, geoserverService.CreateStyle("init_css", "css", styleContent, internal.WorkspaceOption("_")), "workspace _ does not exist")
		})

		t.Run("DUPLICATE STYLE", func(t *testing.T) {
			assert.Error(t, geoserverService.CreateStyle("init_css_no_workspace", "css", styleContent), "style init_css_no_workspace already exists")
		})
	})

	t.Run("GET", func(t *testing.T) {
		t.Run("SIMPLE", func(t *testing.T) {
			s, err := geoserverService.GetStyle("init_css_no_workspace", "css")
			assert.NilError(t, err)

			assert.Equal(t, s.Style.Content, string(styleContent))
			assert.Equal(t, s.Style.Format, "css")
			assert.Equal(t, s.Style.Name, "init_css_no_workspace")
			var w *workspace.WorkspaceCreation
			assert.Equal(t, s.Style.Workspace, w)
		})

		t.Run("WITH WORKSPACE OPTION", func(t *testing.T) {
			s, err := geoserverService.GetStyle("init_css", "css", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			assert.Equal(t, s.Style.Content, string(styleContent))
			assert.Equal(t, s.Style.Format, "css")
			assert.Equal(t, s.Style.Name, "init_css")
			assert.Equal(t, s.Style.Workspace.Name, "init")
		})

		t.Run("DEFAULT FORMAT", func(t *testing.T) {
			s, err := geoserverService.GetStyle("init_css_no_workspace", "json")
			assert.NilError(t, err)

			assert.Equal(t, s.Style.Content, "")
			assert.Equal(t, s.Style.Format, "css")
			assert.Equal(t, s.Style.Name, "init_css_no_workspace")

			var w *workspace.WorkspaceCreation
			assert.Equal(t, s.Style.Workspace, w)
		})

		t.Run("NON-EXISTENT", func(t *testing.T) {
			_, err := geoserverService.GetStyle("none", "json")
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "style none does not exist")
		})
	})

	t.Run("APPLY TO LAYER", func(t *testing.T) {
		bbox := [4]float64{}
		connectionParams := postgis.ConnectionParams{
			Host:     host,
			Database: databaseName,
			User:     databaseUser,
			Password: databasePassword,
			Port:     databasePort,
			SSL:      "disable",
		}

		assert.NilError(t, geoserverService.CreatePostGISDataStore("init", "init", connectionParams))
		assert.NilError(t, geoserverService.CreateFeatureType("init", "init", "init", "init", "EPSG:4326", bbox, internal.KeywordsOption([]string{"test", "marian"}), internal.TitleOption("titlu misto"), internal.ProjectionPolicyOption("FORCE_DECLARED")))

		t.Run("WITHOUT WORKSPACE OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.StyleLayer("init", "init_css_no_workspace", "css", true))

			s, err := geoserverService.GetLayer("init")
			assert.NilError(t, err)

			assert.Equal(t, s.Layer.Name, "init")
			assert.Equal(t, s.Layer.Type, "VECTOR")
			assert.Equal(t, s.Layer.Resource.Class, "featureType")
			assert.Equal(t, s.Layer.Resource.Name, "init:init")
			assert.Equal(t, s.Layer.DefaultStyle.Name, "init_css_no_workspace")
		})

		t.Run("WITH WORKSPACE OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.StyleLayer("init", "init_css", "css", true, internal.WorkspaceOption("init")))

			s, err := geoserverService.GetLayer("init", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			assert.Equal(t, s.Layer.Name, "init")
			assert.Equal(t, s.Layer.Type, "VECTOR")
			assert.Equal(t, s.Layer.Resource.Class, "featureType")
			assert.Equal(t, s.Layer.Resource.Name, "init:init")
			assert.Equal(t, s.Layer.DefaultStyle.Name, "init:init_css")
		})

		t.Run("DO NOT MAKE DEFAULT WITH WORKSPACE OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.StyleLayer("init", "init_css", "css", false, internal.WorkspaceOption("init")))

			s, err := geoserverService.GetLayer("init", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			expected := []any{
				map[string]any{
					"name": "init_css_no_workspace",
					"href": fmt.Sprintf("%s/geoserver/rest/styles/init_css_no_workspace.json", geoserverService.url),
				},
				map[string]any{
					"name":      "init:init_css",
					"href":      fmt.Sprintf("%s/geoserver/rest/workspaces/init/styles/init_css.json", geoserverService.url),
					"workspace": "init",
				},
			}

			style, ok := s.Layer.Styles.Style.([]any)
			assert.Equal(t, ok, true)
			assert.Equal(t, len(style), 2)
			assert.DeepEqual(t, style, expected)
			assert.Equal(t, s.Layer.Name, "init")
			assert.Equal(t, s.Layer.Type, "VECTOR")
			assert.Equal(t, s.Layer.Resource.Class, "featureType")
			assert.Equal(t, s.Layer.Resource.Name, "init:init")
			assert.Equal(t, s.Layer.DefaultStyle.Name, "init:init_css")
		})

		t.Run("DO NOT MAKE DEFAULT WITHOUT WORKSPACE OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.StyleLayer("init", "init_css_no_workspace", "css", false))

			s, err := geoserverService.GetLayer("init", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			expected := []any{
				map[string]any{
					"name": "init_css_no_workspace",
					"href": fmt.Sprintf("%s/geoserver/rest/styles/init_css_no_workspace.json", geoserverService.url),
				},
				map[string]any{
					"name":      "init:init_css",
					"href":      fmt.Sprintf("%s/geoserver/rest/workspaces/init/styles/init_css.json", geoserverService.url),
					"workspace": "init",
				},
			}

			styles, ok := s.Layer.Styles.Style.([]any)
			assert.Equal(t, ok, true)
			assert.Equal(t, len(styles), 2)
			assert.DeepEqual(t, styles, expected)
			assert.Equal(t, s.Layer.Name, "init")
			assert.Equal(t, s.Layer.Type, "VECTOR")
			assert.Equal(t, s.Layer.Resource.Class, "featureType")
			assert.Equal(t, s.Layer.Resource.Name, "init:init")
			assert.Equal(t, s.Layer.DefaultStyle.Name, "init:init_css")
		})

		t.Run("NON-EXISTENT STYLE", func(t *testing.T) {
			err := geoserverService.StyleLayer("init", "none", "css", true)
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "style none does not exist")
		})

		t.Run("NON-EXISTENT LAYER", func(t *testing.T) {
			err := geoserverService.StyleLayer("none", "init_css", "css", true)
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "layer none does not exist")
		})
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Run("WITHOUT RECURSE", func(t *testing.T) {
			err := geoserverService.DeleteStyle("init_css", internal.WorkspaceOption("init"))
			assert.ErrorType(t, err, &customerrors.GeoserverError{})
			assert.Error(t, err, "style is referenced by other layers and recurse option is missing or set to false")
		})

		t.Run("WITH RECURSE AND PURGE", func(t *testing.T) {
			assert.NilError(t, geoserverService.DeleteStyle("init_css", internal.WorkspaceOption("init"), internal.RecurseOption(true), internal.PurgeOption("true")))
		})

		t.Run("NON-EXISTENT STYLE", func(t *testing.T) {
			err := geoserverService.DeleteStyle("none")
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "style none does not exist")
		})

		t.Run("NON-EXISTENT WORKSPACE", func(t *testing.T) {
			err := geoserverService.DeleteStyle("init_css_no_workspace", internal.WorkspaceOption("none"))
			assert.ErrorType(t, err, &customerrors.NotFoundError{})
			assert.Error(t, err, "workspace none does not exist")
		})
	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", internal.RecurseOption(true)))
}
