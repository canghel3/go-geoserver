package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFontsIntegration_Get(t *testing.T) {
	fonts, err := geoclient.Fonts().Get()
	assert.NoError(t, err)
	assert.NotNil(t, fonts)
}
