package service

import (
	"github.com/canghel3/go-geoserver/internal"
	"gotest.tools/v3/assert"
	"testing"
)

func TestLayers(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)
	assert.NilError(t, geoserverService.CreateWorkspace("init"))

	t.Run("GET", func(t *testing.T) {
		t.Skip()
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Skip()
	})

	t.Run("WITH SINGLE KEYWORD", func(t *testing.T) {})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", internal.RecurseOption(true)))
}
