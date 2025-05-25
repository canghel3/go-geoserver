package internal

import "github.com/canghel3/go-geoserver/internal/misc"

type CoverageWrapper struct {
	Coverage Coverage `json:"coverage"`
}

type Coverage struct {
	Abstract                   *string               `json:"abstract,omitempty"`
	DefaultInterpolationMethod *string               `json:"defaultInterpolationMethod,omitempty"`
	Description                *string               `json:"description,omitempty"`
	Dimensions                 *CoverageDimensions   `json:"dimensions,omitempty"`
	Enabled                    bool                  `json:"enabled"`
	Grid                       *GridDetails          `json:"grid,omitempty"`
	InterpolationMethods       *InterpolationMethods `json:"interpolationMethods,omitempty"`
	ProjectionPolicy           *string               `json:"projectionPolicy,omitempty"`
	Keywords                   *misc.Keywords        `json:"keywords,omitempty"`
	LatLonBoundingBox          *misc.BoundingBox     `json:"latLonBoundingBox,omitempty"`
	Metadata                   *Metadata             `json:"metadata,omitempty"`
	Name                       string                `json:"name"`
	Namespace                  NamespaceDetails      `json:"namespace"`
	NativeBoundingBox          *misc.BoundingBox     `json:"nativeBoundingBox,omitempty"`
	NativeCRS                  *CRS                  `json:"nativeCRS,omitempty"`
	NativeFormat               *string               `json:"nativeFormat,omitempty"`
	NativeName                 string                `json:"nativeName,omitempty"`
	RequestSRS                 *SRS                  `json:"requestSRS,omitempty"`
	ResponseSRS                *SRS                  `json:"responseSRS,omitempty"`
	Srs                        *string               `json:"srs,omitempty"`
	Store                      StoreDetails          `json:"store"`
	SupportedFormats           *SupportedFormats     `json:"supportedFormats,omitempty"`
	Title                      *string               `json:"title,omitempty"`
}

type CoverageDimensions struct {
	CoverageDimension CoverageDimension `json:"coverageDimension"`
}

type CoverageDimension struct {
	Description string     `json:"description"`
	Name        string     `json:"name"`
	DataType    DataType   `json:"dimensionType"`
	NullValues  NullValues `json:"nullValues"`
	Range       Range      `json:"range"`
}

type DataType struct {
	Name string `json:"name"`
}

type NullValues struct {
	Double float64 `json:"double"`
}

type Range struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type GridDetails struct {
	Dimension int    `json:"@dimension"`
	Crs       string `json:"crs"`
	Range     RangeDetails
	Transform Transform `json:"transform"`
}

type RangeDetails struct {
	High string `json:"high"`
	Low  string `json:"low"`
}

type Transform struct {
	ScaleX     float64 `json:"scaleX"`
	ScaleY     float64 `json:"scaleY"`
	ShearX     float64 `json:"shearX"`
	ShearY     float64 `json:"shearY"`
	TranslateX float64 `json:"translateX"`
	TranslateY float64 `json:"translateY"`
}

type InterpolationMethods struct {
	String []string `json:"string"`
}

type LatLonBoundingBox struct {
	Crs  string  `json:"crs"`
	Maxx float64 `json:"maxx"`
	Maxy float64 `json:"maxy"`
	Minx float64 `json:"minx"`
	Miny float64 `json:"miny"`
}

type Metadata struct {
	Entry []MetadataEntry `json:"entry"`
}

type MetadataEntry struct {
	Key           string        `json:"@key"`
	Text          string        `json:"$"`
	DimensionInfo DimensionInfo `json:"dimensionInfo,omitempty"`
}

type DimensionInfo struct {
	Enabled                bool   `json:"enabled"`
	Presentation           string `json:"presentation"`
	Units                  string `json:"units"`
	DefaultValue           string `json:"defaultValue"`
	NearestMatchEnabled    bool   `json:"nearestMatchEnabled"`
	RawNearestMatchEnabled bool   `json:"rawNearestMatchEnabled"`
	StartValue             int    `json:"startValue"`
	EndValue               string `json:"endValue"`
}

type NamespaceDetails struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type NativeBoundingBox struct {
	Crs  CRS     `json:"crs"`
	Maxx float64 `json:"maxx"`
	Maxy float64 `json:"maxy"`
	Minx float64 `json:"minx"`
	Miny float64 `json:"miny"`
}

type CRS struct {
	Text  string `json:"$"`
	Class string `json:"@class"`
}

type SRS struct {
	String []string `json:"string"`
}

type StoreDetails struct {
	Class string `json:"@class"`
	Href  string `json:"href"`
	Name  string `json:"name"`
}

type SupportedFormats struct {
	String []string `json:"string"`
}
