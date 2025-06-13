package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFontsIntegration_Get(t *testing.T) {
	fonts, err := geoclient.Fonts().Get()
	assert.NoError(t, err)
	assert.NotNil(t, fonts)
}
