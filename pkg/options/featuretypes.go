package options

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/misc"
)

var FeatureType FeatureTypeOptionsGenerator

type FeatureTypeOptionsGenerator struct{}

type FeatureTypeOption func(ft *internal.FeatureType)

func (ftog FeatureTypeOptionsGenerator) BBOX(bbox [4]float64, bboxSrs string) FeatureTypeOption {
	return func(ft *internal.FeatureType) {
		ft.NativeBoundingBox = &misc.BoundingBox{
			MinX: bbox[0],
			MaxX: bbox[2],
			MinY: bbox[1],
			MaxY: bbox[3],
			CRS:  bboxSrs,
		}
	}
}
