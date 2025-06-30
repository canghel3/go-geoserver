package actions

import (
	"bytes"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"golang.org/x/image/tiff"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

type WMS struct {
	requester requester.WMSRequester
	version   wms.WMSVersion
}

func NewWMSActions(data internal.GeoserverData, version wms.WMSVersion) WMS {
	return WMS{
		requester: requester.NewWMSRequester(data),
		version:   version,
	}
}

//func (wm *WMS) GetCapabilities() (wms.Capabilities, error) {
//	cap, err := wm.requester.WMS().GetCapabilities(wm.version)
//	if err != nil {
//		return nil, err
//	}
//
//	switch wm.version {
//	case wms.Version111:
//		return nil, customerrors.NewNotImplementedError("not implemented")
//		//unmarshal into v1.1.1 struct
//	case wms.Version130:
//		cap130 := wms.Capabilities1_3_0{}
//		err = xml.Unmarshal(cap, &cap130)
//		if err != nil {
//			return nil, err
//		}
//
//		return &cap130, nil
//	default:
//		return nil, errors.New("unsupported version")
//	}
//}

func (wm WMS) GetMap(width, height uint16, layers []string, bbox shared.BBOX) MapFormats {
	return MapFormats{
		width:     width,
		height:    height,
		bbox:      bbox,
		layers:    layers,
		version:   wm.version,
		requester: wm.requester,
	}
}

type MapFormats struct {
	width     uint16
	height    uint16
	layers    []string
	bbox      shared.BBOX
	version   wms.WMSVersion
	requester requester.WMSRequester
}

func (mf MapFormats) Png() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.PNG)
	if err != nil {
		return nil, err
	}

	return png.Decode(bytes.NewReader(content))
}

func (mf MapFormats) Png8() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.PNG8)
	if err != nil {
		return nil, err
	}

	return png.Decode(bytes.NewReader(content))
}

func (mf MapFormats) Jpeg() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.JPEG)
	if err != nil {
		return nil, err
	}

	return jpeg.Decode(bytes.NewReader(content))
}

func (mf MapFormats) JpegPng() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.JPEG_PNG)
	if err != nil {
		return nil, err
	}

	return jpeg.Decode(bytes.NewReader(content))
}

func (mf MapFormats) JpegPng8() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.JPEG_PNG8)
	if err != nil {
		return nil, err
	}

	return jpeg.Decode(bytes.NewReader(content))
}

func (mf MapFormats) Gif() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.GIF)
	if err != nil {
		return nil, err
	}

	return gif.Decode(bytes.NewReader(content))
}

func (mf MapFormats) Tiff() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.TIFF)
	if err != nil {
		return nil, err
	}

	return tiff.Decode(bytes.NewReader(content))
}

func (mf MapFormats) Tiff8() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.TIFF8)
	if err != nil {
		return nil, err
	}

	return tiff.Decode(bytes.NewReader(content))
}

func (mf MapFormats) GeoTiff() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.GeoTIFF)
	if err != nil {
		return nil, err
	}

	return tiff.Decode(bytes.NewReader(content))
}

func (mf MapFormats) GeoTiff8() (image.Image, error) {
	content, err := mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.GeoTIFF8)
	if err != nil {
		return nil, err
	}

	return tiff.Decode(bytes.NewReader(content))
}

func (mf MapFormats) Svg() ([]byte, error) {
	return mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.SVG)
}

func (mf MapFormats) Pdf() ([]byte, error) {
	return mf.requester.GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.PDF)
}

//func (mf MapFormats) GeoRSS() (image.Image, error) {
//	return nil, customerrors.NewNotImplementedError("not implemented")
//}
//
//func (mf MapFormats) KML() (image.Image, error) {
//	return nil, customerrors.NewNotImplementedError("not implemented")
//}
//
//func (mf MapFormats) KMZ() (image.Image, error) {
//	return nil, customerrors.NewNotImplementedError("not implemented")
//}
//
//func (mf MapFormats) MapML() (image.Image, error) {
//	return nil, customerrors.NewNotImplementedError("not implemented")
//}
//
//func (mf MapFormats) MapMLHTMLViewer() ([]byte, error) {
//	return nil, customerrors.NewNotImplementedError("not implemented")
//}

//func (mf MapFormats) OpenLayers() (*wms.OpenLayersTemplate, error) {
//	raw, err := mf.requester.WMS().GetMap(mf.width, mf.height, mf.layers, mf.bbox, mf.version, wms.OpenLayers)
//	if err != nil {
//		return nil, err
//	}
//
//	templ, err := template.New("OpenLayers").Parse(string(raw))
//	if err != nil {
//		return nil, err
//	}
//
//	return &wms.OpenLayersTemplate{
//		Template: templ,
//		RawHTML:  raw,
//	}, nil
//}

//func (mf MapFormats) UTFGrid() ([]byte, error) {
//	return nil, customerrors.NewNotImplementedError("not implemented")
//}
