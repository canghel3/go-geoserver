package service

import (
	"github.com/canghel3/go-geoserver/utils"
	"gotest.tools/v3/assert"
	"testing"
)

func TestLayers(t *testing.T) {
	geoserverService := NewGeoserverService(target, username, password)
	assert.NilError(t, geoserverService.CreateWorkspace("init"))

	t.Run("GET", func(t *testing.T) {
		t.Skip()
	})

	t.Run("DELETE", func(t *testing.T) {
		t.Skip()
	})

	assert.NilError(t, geoserverService.DeleteWorkspace("init", utils.RecurseOption(true)))
}
