package featuretypes

import (
	"github.com/canghel3/go-geoserver/internal/misc"
)

type FeatureTypeOption func(c *FeatureType)

type featureTypesOptions struct{}

func (fto featureTypesOptions) BBOX(bbox [4]float64, bboxSrs string) FeatureTypeOption {
	return func(f *FeatureType) {
		f.NativeBoundingBox = &misc.BoundingBox{
			MinX: bbox[0],
			MaxX: bbox[2],
			MinY: bbox[1],
			MaxY: bbox[3],
			CRS:  bboxSrs,
		}
	}
}

var Options featureTypesOptions

func New(name, nativeName string, options ...FeatureTypeOption) FeatureType {
	cft := new(FeatureType)
	cft.Name = name
	cft.NativeName = nativeName

	for _, option := range options {
		option(cft)
	}

	return *cft
}

type CreateFeatureTypeWrapper struct {
	FeatureType FeatureType `json:"featureType"`
}

type FeatureType struct {
	Name string `json:"name"`
	//The native Name of the resource. This Name corresponds to the physical resource that feature type is derived from -- a shapefile Name, a database table, etc...
	NativeName        string            `json:"nativeName"`
	Namespace         Namespace         `json:"namespace"`
	Srs               *string           `json:"srs,omitempty"`
	NativeBoundingBox *misc.BoundingBox `json:"nativeBoundingBox,omitempty"`
	ProjectionPolicy  *string           `json:"projectionPolicy,omitempty"`
	Keywords          *misc.Keywords    `json:"keywords,omitempty"`
	Title             *string           `json:"title,omitempty"`
	Store             Store             `json:"store"`
}
