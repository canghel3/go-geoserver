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

type GetFeatureTypeWrapper struct {
	FeatureType GetFeatureType `json:"featureType"`
}

type GetFeatureType struct {
	Name                   string             `json:"name"`
	NativeName             string             `json:"nativeName"`
	Namespace              Namespace          `json:"namespace"`
	Title                  string             `json:"title"`
	Abstract               string             `json:"abstract"`
	Keywords               shared.Keywords    `json:"keywords"`
	MetadataLinks          Links              `json:"metadataLinks"`
	DataLinks              Links              `json:"dataLinks"`
	NativeCRS              any                `json:"nativeCRS"`
	Srs                    string             `json:"srs"`
	NativeBoundingBox      shared.BoundingBox `json:"nativeBoundingBox"`
	LatLonBoundingBox      shared.BoundingBox `json:"latLonBoundingBox"`
	ProjectionPolicy       string             `json:"projectionPolicy"`
	Enabled                bool               `json:"enabled"`
	Metadata               Metadata           `json:"metadata"`
	Store                  Store              `json:"store"`
	CqlFilter              string             `json:"cqlFilter"`
	MaxFeatures            int                `json:"maxFeatures"`
	NumDecimals            int                `json:"numDecimals"`
	ResponseSRS            SRS                `json:"responseSRS"`
	OverridingServiceSRS   bool               `json:"overridingServiceSRS"`
	SkipNumberMatched      bool               `json:"skipNumberMatched"`
	CircularArcPresent     bool               `json:"circularArcPresent"`
	LinearizationTolerance int                `json:"linearizationTolerance"`
	Attributes             Attributes         `json:"attributes"`
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
