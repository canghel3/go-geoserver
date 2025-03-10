package service

import (
	"github.com/canghel3/go-geoserver/datastore/postgis"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/gwc"
	"gotest.tools/v3/assert"
	"testing"
)

func TestGWC(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)
	assert.NilError(t, geoserverService.CreateWorkspace("init"))

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
	assert.NilError(t, geoserverService.CreateFeatureType("init", "init", "init", "init", "EPSG:3857", bbox, internal.KeywordsOption([]string{"test", "marian"}), internal.TitleOption("titlu misto"), internal.ProjectionPolicyOption("FORCE_DECLARED")))

	t.Run("GET", func(t *testing.T) {
		t.Run("GET FROM FEATURE TYPE", func(t *testing.T) {
			tc, err := geoserverService.GetTileCaching("init", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			assert.Equal(t, tc.Name, "init:init")
			assert.Equal(t, tc.Enabled, true)
			assert.Equal(t, tc.MetaWidthHeight[0], 4)
			assert.DeepEqual(t, &tc.GridSubsets[0], &gwc.GridSubset{
				GridSetName:    "EPSG:4326",
				MinCachedLevel: 0,
				MaxCachedLevel: 0,
			})
		})
	})

	t.Run("UPDATE", func(t *testing.T) {
		updateData := gwc.TileTabUpdateData{
			GridSubsets: []gwc.GridSubset{
				{
					GridSetName:    "EPSG:4326",
					MinCachedLevel: 0,
					MaxCachedLevel: 10,
				},
			},
			MimeFormats: []string{"image/png8"},
		}

		t.Run("OK", func(t *testing.T) {
			assert.NilError(t, geoserverService.UpdateTileCaching("init", updateData, internal.WorkspaceOption("init")))

			tc, err := geoserverService.GetTileCaching("init", internal.WorkspaceOption("init"))
			assert.NilError(t, err)

			assert.Equal(t, len(tc.GridSubsets), 1)
			assert.Equal(t, len(tc.MimeFormats), 1)
			assert.Equal(t, tc.MimeFormats[0], "image/png8")
			assert.Equal(t, tc.GridSubsets[0].GridSetName, "EPSG:4326")
			assert.Equal(t, tc.GridSubsets[0].MaxCachedLevel, 10)

		})
	})

	t.Run("SEED", func(t *testing.T) {
		t.Run("START", func(t *testing.T) {
			assert.NilError(t, geoserverService.Seed("init", "init", "EPSG:4326", "image/png8", 0, 10, 2))
		})

		t.Run("GET STATUS", func(t *testing.T) {
			t.Skip()
		})
	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", internal.RecurseOption(true)))
}
