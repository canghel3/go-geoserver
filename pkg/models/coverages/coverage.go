package coverages

// Coverage contains details about a coverage.
// This structure contains many fields with pointer values because otherwise empty values would not be omitted from the json marshaller.
//type Coverage struct {
//	Abstract                   string                   `json:"abstract,omitempty"`
//	DefaultInterpolationMethod string                   `json:"defaultInterpolationMethod,omitempty"`
//	Description                string                   `json:"description,omitempty"`
//	Dimensions                 *CoverageDimensions      `json:"dimensions,omitempty"`
//	Enabled                    bool                     `json:"enabled"`
//	Grid                       *GridDetails             `json:"grid,omitempty"`
//	InterpolationMethods       *InterpolationMethods    `json:"interpolationMethods,omitempty"`
//	ProjectionPolicy           string                   `json:"projectionPolicy,omitempty"`
//	Keywords                   *misc.Keywords           `json:"keywords,omitempty"`
//	LatLonBoundingBox          *misc.BoundingBox        `json:"latLonBoundingBox,omitempty"`
//	Metadata                   *Metadata                `json:"metadata,omitempty"`
//	Name                       string                   `json:"name"`
//	Namespace                  workspace.MultiWorkspace `json:"namespace"`
//	NativeBoundingBox          misc.BoundingBox         `json:"nativeBoundingBox"`
//	NativeCRS                  *CRS                     `json:"nativeCRS,omitempty"`
//	NativeFormat               string                   `json:"nativeFormat,omitempty"`
//	NativeName                 string                   `json:"nativeName,omitempty"`
//	RequestSRS                 *SRS                     `json:"requestSRS,omitempty"`
//	ResponseSRS                *SRS                     `json:"responseSRS,omitempty"`
//	Srs                        string                   `json:"srs"`
//	Store                      StoreDetails             `json:"store"`
//	SupportedFormats           *SupportedFormats        `json:"supportedFormats,omitempty"`
//	Title                      string                   `json:"title"`
//}
