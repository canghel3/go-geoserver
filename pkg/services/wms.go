package services

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/models/wms"
)

type WMS struct {
	requester *requester.Requester
}

func newWMS(info *internal.GeoserverInfo) *WMS {
	return &WMS{
		requester: requester.NewRequester(info),
	}
}

func (wms *WMS) GetCapabilities(version wms.WMSVersion) (*wms.Capabilities, error) {
	return wms.requester.WMS().GetCapabilities(string(version))
}
