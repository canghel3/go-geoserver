package models

import (
	"github.com/canghel3/go-geoserver/pkg/shared"
)

type CreateFeatureTypeWrapper struct {
	FeatureType FeatureType `json:"featureType"`
}

type FeatureType struct {
	Name string `json:"name"`
	//The native Name of the resource. This Name corresponds to the physical resource that feature type is derived from -- a shapefile Name, a database table, etc...
	NativeName        string              `json:"nativeName"`
	Namespace         Namespace           `json:"namespace"`
	Srs               *string             `json:"srs,omitempty"`
	NativeBoundingBox *shared.BoundingBox `json:"nativeBoundingBox,omitempty"`
	ProjectionPolicy  *string             `json:"projectionPolicy,omitempty"`
	Keywords          *shared.Keywords    `json:"keywords,omitempty"`
	Title             *string             `json:"title,omitempty"`
	Store             Store               `json:"store"`
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
