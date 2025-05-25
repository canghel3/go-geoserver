package options

import (
	"github.com/canghel3/go-geoserver/pkg/models/coverages"
)

var Coverage CoverageOptionsGenerator

type CoverageOptionsGenerator struct{}

type CoverageOptionsFunc func(csl *coverages.Coverage)
