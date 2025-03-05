package vector

import (
	"github.com/canghel3/go-geoserver/internal"
)

type Layers struct {
	info internal.GeoserverInfo
}

func newLayers(info internal.GeoserverInfo) Layers {
	return Layers{
		info: info,
	}
}

func (l Layers) Feature() {

}
