package options

import (
	"github.com/canghel3/go-geoserver/internal"
)

var Coverage CoverageOptionsGenerator

type CoverageOptionsGenerator struct{}

type CoverageOption func(csl *internal.Coverage)

func (cog CoverageOptionsGenerator) Title(title string) CoverageOption {
	return func(csl *internal.Coverage) {
		csl.Title = &title
	}
}
