package shared

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type BoundingBox struct {
	MinX float64  `json:"minx"`
	MaxX float64  `json:"maxx"`
	MinY float64  `json:"miny"`
	MaxY float64  `json:"maxy"`
	CRS  CRSClass `json:"crs,omitempty"`
}

func (b *BoundingBox) MarshalJSON() ([]byte, error) {
	output := map[string]interface{}{
		"minx": b.MinX,
		"maxx": b.MaxX,
		"miny": b.MinY,
		"maxy": b.MaxY,
	}

	class := strings.TrimSpace(b.CRS.Class)
	value := strings.TrimSpace(b.CRS.Value)

	switch {
	case class != "" && value != "":
		output["crs"] = map[string]interface{}{
			"@class": class,
			"$":      value,
		}
	case value != "":
		output["crs"] = value
	}

	return json.Marshal(output)
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

// CRSClass represents coordinate reference system information
type CRSClass struct {
	// Class can be empty
	Class string `json:"@class"`
	// Value will never be empty
	Value string `json:"$"`
}

func (c *CRSClass) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		c.Value = str
		c.Class = ""
		return nil
	}

	type alias CRSClass
	var temp alias
	if err := json.Unmarshal(data, &temp); err == nil {
		*c = CRSClass(temp)
		return nil
	}

	return fmt.Errorf("unable to unmarshal CRS: %s", string(data))
}
