package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAboutIntegration(t *testing.T) {
	t.Run("Manifest", func(t *testing.T) {
		manifest, err := geoclient.About().Manifest()
		assert.NoError(t, err)
		assert.NotNil(t, manifest)
	})

	t.Run("Version", func(t *testing.T) {
		version, err := geoclient.About().Version()
		assert.NoError(t, err)
		assert.NotNil(t, version)
	})

	t.Run("Status", func(t *testing.T) {
		status, err := geoclient.About().Status()
		assert.NoError(t, err)
		assert.NotNil(t, status)
	})

	t.Run("System Status", func(t *testing.T) {
		status, err := geoclient.About().SystemStatus()
		assert.NoError(t, err)
		assert.NotNil(t, status)
	})
}
