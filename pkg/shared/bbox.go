package shared

import (
	"strconv"
	"strings"
)

type BoundingBox struct {
	MinX float64 `json:"minx"`
	MaxX float64 `json:"maxx"`
	MinY float64 `json:"miny"`
	MaxY float64 `json:"maxy"`
	CRS  any     `json:"crs"`
}

type BBOX struct {
	MinX float64
	MaxX float64
	MinY float64
	MaxY float64
	SRS  string
}

func (b BBOX) ToString() string {
	return strings.Join([]string{
		strconv.FormatFloat(b.MinX, 'f', -1, 64),
		strconv.FormatFloat(b.MinY, 'f', -1, 64),
		strconv.FormatFloat(b.MaxX, 'f', -1, 64),
		strconv.FormatFloat(b.MaxY, 'f', -1, 64),
	}, ",")
}

type CRSClass struct {
	Class string `json:"@class"`
	Value string `json:"$"`
}
