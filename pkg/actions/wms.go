package actions

import (
	"errors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"image"
)

type WMS struct {
	requester *requester.Requester
	version   wms.WMSVersion
}

func NewWMSHandler(info *internal.GeoserverData, version wms.WMSVersion) *WMS {
	return &WMS{
		requester: requester.NewRequester(info),
		version:   version,
	}
}

func (wms *WMS) GetCapabilities() (*wms.Capabilities1_3_0, error) {
	//TODO: validate version
	return wms.requester.WMS().GetCapabilities(wms.version)
}

func (wms *WMS) GetMap() *MapFormats {
	//return wms.requester.WMS().GetMap(version)
	return &MapFormats{
		version: wms.version,
	}
}

type MapFormats struct {
	version wms.WMSVersion
}

func (mf *MapFormats) Png() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) Png8() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) Jpeg() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) JpegPng() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) JpegPng8() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) Gif() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) Tiff() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) Tiff8() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) GeoTiff() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) GeoTiff8() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) Svg() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) Pdf() ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) GeoRSS() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) KML() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) KMZ() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) MapML() (image.Image, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) MapMLHTMLViewer() ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) OpenLayers() ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (mf *MapFormats) UTFGrid() ([]byte, error) {
	return nil, errors.New("not implemented")
}
