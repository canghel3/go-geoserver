package options

import (
	"github.com/canghel3/go-geoserver/internal/models"
)

var Coverage CoverageOptionsGenerator

type CoverageOptionsGenerator struct{}

type CoverageOption func(csl *models.Coverage)

//func (cog CoverageOptionsGenerator) Title(title string) CoverageOption {
//	return func(csl *models.Coverage) {
//		csl.Title = &title
//	}
//}
