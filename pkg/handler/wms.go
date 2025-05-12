package handler

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/wms"
)

type WMS struct {
	requester *requester.Requester
}

func NewWMSHandler(info *internal.GeoserverData) *WMS {
	return &WMS{
		requester: requester.NewRequester(info),
	}
}

// TODO: should return a Capabilities interface to accomodate all response formats based on the requested version
func (wms *WMS) GetCapabilities(version wms.WMSVersion) (*wms.Capabilities1_3_0, error) {
	//TODO: validate version
	return wms.requester.WMS().GetCapabilities(string(version))
}
