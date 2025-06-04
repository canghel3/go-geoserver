package options

import (
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/pkg/misc"
)

var FeatureType FeatureTypeOptionsGenerator

type FeatureTypeOptionsGenerator struct{}

type FeatureTypeOption func(ft *models.FeatureType)

func (ftog FeatureTypeOptionsGenerator) BBOX(bbox [4]float64, bboxSrs string) FeatureTypeOption {
	return func(ft *models.FeatureType) {
		ft.NativeBoundingBox = &misc.BoundingBox{
			MinX: bbox[0],
			MaxX: bbox[2],
			MinY: bbox[1],
			MaxY: bbox[3],
			CRS:  bboxSrs,
		}
	}
}
