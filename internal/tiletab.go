package internal

// GeoServerLayer describes the response for the tiling tab in xml.
type GeoServerLayer struct {
	ID               string            `xml:"id"`
	Enabled          bool              `xml:"enabled"`
	InMemoryCached   bool              `xml:"inMemoryCached"`
	Name             string            `xml:"name"`
	MimeFormats      []string          `xml:"mimeFormats>string"`
	GridSubsets      []GridSubset      `xml:"gridSubsets>gridSubset"`
	MetaWidthHeight  []int             `xml:"metaWidthHeight>int"`
	ExpireCache      int               `xml:"expireCache"`
	ExpireClients    int               `xml:"expireClients"`
	ParameterFilters []ParameterFilter `xml:"parameterFilters>styleParameterFilter"`
	Gutter           int               `xml:"gutter"`
	BlobStore        string            `xml:"blobStoreId,omitempty"`
}

type GridSubset struct {
	GridSetName    string `xml:"gridSetName"`
	MinCachedLevel int    `xml:"minCachedLevel"`
	MaxCachedLevel int    `xml:"maxCachedLevel"`
}

type ParameterFilter struct {
	Key          string `xml:"key"`
	DefaultValue string `xml:"defaultValue"`
}

type TileTabUpdateData struct {
	GridSubsets      []GridSubset
	ParameterFilters []ParameterFilter
	MimeFormats      []string
	Blobstore        string
}
