package internal

type Cache struct {
	Name        string `json:"name"`
	GridSetID   string `json:"gridSetId"`
	ZoomStart   uint8  `json:"zoomStart"`
	ZoomStop    uint8  `json:"zoomStop"`
	Type        string `json:"type"`
	ThreadCount uint8  `json:"threadCount"`
	Format      string `json:"format"`
}

type CacheWrapper struct {
	SeedRequest Cache `json:"seedRequest"`
}
