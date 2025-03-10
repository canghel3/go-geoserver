package service

import (
	"github.com/canghel3/go-geoserver/datastore/postgis"
	"github.com/canghel3/go-geoserver/internal"
	"gotest.tools/v3/assert"
	"testing"
)

func TestLayerGroup(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)

	bbox := [4]float64{-180, -90, 180, 90}
	connectionParams := postgis.ConnectionParams{
		Host:     host,
		Database: databaseName,
		User:     databaseUser,
		Password: databasePassword,
		Port:     databasePort,
		SSL:      "disable",
	}

	assert.NilError(t, geoserverService.CreateWorkspace("init"))

	assert.NilError(t, geoserverService.CreatePostGISDataStore("init", "init_data_store", connectionParams))
	assert.NilError(t, geoserverService.CreateCoverageStore("init", "init_coverage_store", "file:/opt/geoserver/data/shipments_2_geocoded.tif", "GeoTIFF"))

	assert.NilError(t, geoserverService.CreateFeatureType("init", "init_data_store", "init_feature", "init", "EPSG:4326", bbox, internal.KeywordsOption([]string{"test", "marian"}), internal.TitleOption("titlu misto"), internal.ProjectionPolicyOption("FORCE_DECLARED")))
	assert.NilError(t, geoserverService.CreateFeatureType("init", "init_data_store", "new1", "init", "EPSG:4326", bbox, internal.KeywordsOption([]string{"test", "marian"}), internal.TitleOption("titlu misto"), internal.ProjectionPolicyOption("FORCE_DECLARED")))

	assert.NilError(t, geoserverService.CreateCoverage("init", "init_coverage_store", "init_coverage", "EPSG:3857", [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}, internal.ProjectionPolicyOption("FORCE_DECLARED")))
	assert.NilError(t, geoserverService.CreateCoverage("init", "init_coverage_store", "new2", "EPSG:3857", [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}, internal.ProjectionPolicyOption("FORCE_DECLARED")))

	t.Run("CREATE", func(t *testing.T) {
		t.Run("GROUP", func(t *testing.T) {
			err := geoserverService.CreateLayerGroup("group", "EPSG:4326", []string{"init_feature", "init_coverage"}, bbox, internal.WorkspaceOption("init"), internal.ModeOption("NAMED"), internal.TitleOption("GROUP"))
			assert.NilError(t, err)

			g, err := geoserverService.GetLayerGroup("group", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			assert.Equal(t, g.Group.Name, "group")
			assert.Equal(t, g.Group.Title, "GROUP")
			assert.Equal(t, g.Group.Mode, "NAMED")
		})
	})

	t.Run("UPDATE", func(t *testing.T) {
		t.Run("GROUP", func(t *testing.T) {
			g, err := geoserverService.GetLayerGroup("group", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			newBBOX := [4]float64{g.Group.Bounds.MinX + 1, g.Group.Bounds.MinY + 1, g.Group.Bounds.MaxX + 2, g.Group.Bounds.MaxY + 2}

			err = geoserverService.UpdateLayerGroup("group", "EPSG:4326", []string{"new1", "new2"}, newBBOX, internal.WorkspaceOption("init"), internal.ModeOption("NAMED"), internal.TitleOption("GROUP"))
			assert.NilError(t, err)

			g, err = geoserverService.GetLayerGroup("group", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			assert.Equal(t, g.Group.Name, "group")
			assert.Equal(t, len(g.Group.Publishables.Published), 4)
			assert.Equal(t, [4]float64{g.Group.Bounds.MinX, g.Group.Bounds.MinY, g.Group.Bounds.MaxX, g.Group.Bounds.MaxY}, newBBOX)
		})
	})

	t.Run("GET", func(t *testing.T) {
		t.Run("SIMPLE", func(t *testing.T) {
			lg, err := geoserverService.GetLayerGroup("group", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			assert.Equal(t, lg.Group.Mode, "NAMED")
			assert.Equal(t, lg.Group.Name, "group")
			assert.Equal(t, lg.Group.Title, "GROUP")
			assert.Equal(t, lg.Group.Publishables.Published[0].Name, "init:init_feature")
			assert.Equal(t, lg.Group.Publishables.Published[1].Name, "init:init_coverage")
		})

		t.Run("WITHOUT WORKSPACE OPTION", func(t *testing.T) {
			//TODO: create 2 workspaces, if the groups have the same name, we get whichever is found first
			t.Skip()
		})
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Run("WITH WORKSPACE OPTION", func(t *testing.T) {
			assert.NilError(t, geoserverService.DeleteLayerGroup("group", internal.WorkspaceOption("init")))
		})
	})

	t.Run("WITH SINGLE KEYWORD", func(t *testing.T) {

	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", internal.RecurseOption(true)))
}
