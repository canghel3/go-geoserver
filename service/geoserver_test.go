package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeoserverService(t *testing.T) {
	t.Run("CREATE SIMPLE", func(t *testing.T) {
		geoserverService := NewGeoserverService(target, username, password)
		assert.False(t, geoserverService.isDataDirectorySet())
	})

	t.Run("CREATE WITH NORMAL DATADIR", func(t *testing.T) {
		geoserverService := NewGeoserverService(target, username, password, GeoserverServiceDataDirOption("/opt/geoserver/data"))
		assert.True(t, geoserverService.isDataDirectorySet())
		assert.Equal(t, geoserverService.getDataDirectory(), "/opt/geoserver/data")
	})

	t.Run("CREATE WITH EMPTY DATADIR", func(t *testing.T) {
		geoserverService := NewGeoserverService(target, username, password, GeoserverServiceDataDirOption(""))
		assert.False(t, geoserverService.isDataDirectorySet())
	})
}
