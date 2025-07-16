package gwc

type LayerWrapper struct {
	Layer Layer `json:"GeoServerLayer"`
}

type Layer struct {
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Enabled         bool      `json:"enabled"`
	BlobStoreId     string    `json:"blobStoreId"`
	MimeFormats     []string  `json:"mimeFormats"`
	InMemoryCached  bool      `json:"inMemoryCached"`
	MetaWidthHeight []int     `json:"metaWidthHeight"`
	GridSets        []GridSet `json:"gridSubsets"`
	ExpireCache     int       `json:"expireCache"`
	Gutter          int       `json:"gutter"`
	ExpireClient    int       `json:"expireClients"`
}

type GridSet struct {
	Name string `json:"gridSetName"`
}
