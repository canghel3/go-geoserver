package client

import (
	"github.com/canghel3/go-geoserver/pkg/logging"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoggingIntegration_Get(t *testing.T) {
	logs, err := geoclient.Logging().Get()
	assert.NoError(t, err)
	assert.NotNil(t, logs)
}

func TestLoggingIntegration_Update(t *testing.T) {
	err := geoclient.Logging().Put(logging.NewLog("some message", "INFO", "basement"))
	assert.NoError(t, err)
}
