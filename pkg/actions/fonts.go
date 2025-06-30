package actions

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/fonts"
)

type Fonts struct {
	requester requester.FontsRequester
}

func NewFonts(data internal.GeoserverData) Fonts {
	return Fonts{
		requester: requester.NewFontsRequester(data),
	}
}

func (f Fonts) Get() (*fonts.Fonts, error) {
	return f.requester.Get()
}
