package shared

import (
	"encoding/json"
	"strconv"
	"strings"
)

type BoundingBoxCRSClass struct {
	MinX float64  `json:"minx"`
	MaxX float64  `json:"maxx"`
	MinY float64  `json:"miny"`
	MaxY float64  `json:"maxy"`
	CRS  CRSClass `json:"crs"`
}

type BoundingBox struct {
	MinX float64 `json:"minx"`
	MaxX float64 `json:"maxx"`
	MinY float64 `json:"miny"`
	MaxY float64 `json:"maxy"`
	CRS  string  `json:"crs"`
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

func (c *CRSClass) UnmarshalJSON(data []byte) error {
	type alias CRSClass
	var temp alias
	if err := json.Unmarshal(data, &temp); err == nil {
		*c = CRSClass(temp)
		return nil
	}

	if err := json.Unmarshal(data, &c.Value); err == nil {
		return nil
	}

	return nil
}
