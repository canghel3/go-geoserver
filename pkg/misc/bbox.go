package misc

type BoundingBox struct {
	MinX float64 `json:"minx"`
	MaxX float64 `json:"maxx"`
	MinY float64 `json:"miny"`
	MaxY float64 `json:"maxy"`
	CRS  any     `json:"crs"`
}

type CRSClass struct {
	Class string `json:"@class"`
	Value string `json:"$"`
}
