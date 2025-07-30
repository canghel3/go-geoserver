package featuretypes

import (
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"time"
)

func New(name, nativeName string, options ...options.FeatureTypeOption) models.FeatureType {
	cft := new(models.FeatureType)
	cft.Name = name
	cft.NativeName = nativeName

	for _, option := range options {
		option(cft)
	}

	return *cft
}

type FeatureTypesWrapper struct {
	FeatureTypes FeatureTypes `json:"featureTypes"`
}

type FeatureTypes struct {
	Entries []struct {
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"featureType"`
}

type FeatureTypeWrapper struct {
	FeatureType FeatureType `json:"featureType"`
}

type FeatureType struct {
	Name       string           `json:"name"`
	NativeName string           `json:"nativeName"`
	Namespace  Namespace        `json:"namespace"`
	Title      string           `json:"title"`
	Abstract   string           `json:"abstract"`
	Keywords   *shared.Keywords `json:"keywords,omitempty"`
	//TODO: because GeoServer responds with either a Links structure or just an empty fucking string,
	// we avoid using this until its either fixed or im in the mood to implement it.
	//MetadataLinks          *Links             `json:"metadataLinks,omitempty"`
	//DataLinks              *Links             `json:"dataLinks,omitempty"`
	NativeCRS         *shared.CRSClass    `json:"nativeCRS,omitempty"`
	Srs               string              `json:"srs"`
	NativeBoundingBox *shared.BoundingBox `json:"nativeBoundingBox,omitempty"`
	LatLonBoundingBox *shared.BoundingBox `json:"latLonBoundingBox,omitempty"`
	ProjectionPolicy  string              `json:"projectionPolicy"`
	Enabled           bool                `json:"enabled"`
	Metadata          *Metadata           `json:"metadata,omitempty"`
	Store             Store               `json:"store"`
	CqlFilter         string              `json:"cqlFilter"`
	MaxFeatures       int                 `json:"maxFeatures"`
	NumDecimals       int                 `json:"numDecimals"`
	//TODO: GeoServer responds wit either a struct or a slice of ints (:
	//ResponseSRS            SRS                `json:"responseSRS"`
	OverridingServiceSRS   bool       `json:"overridingServiceSRS"`
	SkipNumberMatched      bool       `json:"skipNumberMatched"`
	CircularArcPresent     bool       `json:"circularArcPresent"`
	LinearizationTolerance any        `json:"linearizationTolerance"`
	Attributes             Attributes `json:"attributes"`
}

type Links struct {
	MetadataLink []MetadataLink `json:"metadataLink"`
}

type MetadataLink struct {
	Type         string `json:"type"`
	MetadataType string `json:"metadataType"`
	Content      string `json:"content"`
}

type Metadata struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	Key           string         `json:"@key"`
	Value         string         `json:"$"`
	DimensionInfo *DimensionInfo `json:"dimensionInfo"`
}

type DimensionInfo struct {
	Enabled      bool      `json:"enabled"`
	Attribute    string    `json:"attribute"`
	Presentation string    `json:"presentation"`
	StartValue   time.Time `json:"startValue"`
	EndValue     time.Time `json:"endValue"`
}

type SRS struct {
	String []int `json:"string"`
}

type Attributes struct {
	Attribute []Attribute `json:"attribute"`
}

type Attribute struct {
	Name      string `json:"name"`
	MinOccurs int    `json:"minOccurs"`
	MaxOccurs int    `json:"maxOccurs"`
	Nillable  bool   `json:"nillable"`
	Binding   string `json:"binding"`
	Length    *int   `json:"length,omitempty"`
}

// Namespace holds workspace configuration details when creating a layer in GeoServer.
type Namespace struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

// Store holds data store configuration details when creating a layer in GeoServer.
type Store struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
	Href  string `json:"href"`
}
