package gwc

import (
	"github.com/canghel3/go-geoserver/pkg/formats"
	"github.com/canghel3/go-geoserver/pkg/types"
)

type SeedData struct {
	Layer       string              `json:"name"`
	Format      formats.ImageFormat `json:"format"`
	Type        types.Seeding       `json:"type"`
	ZoomStart   uint                `json:"zoomStart"`
	ZoomStop    uint                `json:"zoomStop"`
	ThreadCount uint                `json:"threadCount"`
	GridSetId   *string             `json:"gridSetId,omitempty"`
}

type SeedStatus struct {
	Info [][]int `json:"long-array-array"`
}
