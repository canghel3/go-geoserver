package wms

// WMSFormat represents the output format for WMS requests
type WMSFormat string

const (
	PNG               WMSFormat = "image/png"
	PNG8              WMSFormat = "image/png8"
	JPEG              WMSFormat = "image/jpeg"
	JPEG_PNG          WMSFormat = "image/vnd.jpeg-png"
	JPEG_PNG8         WMSFormat = "image/vnd.jpeg-png8"
	GIF               WMSFormat = "image/gif"
	TIFF              WMSFormat = "image/tiff"
	TIFF8             WMSFormat = "image/tiff8"
	GeoTIFF           WMSFormat = "image/geotiff"
	GeoTIFF8          WMSFormat = "image/geotiff8"
	SVG               WMSFormat = "image/svg"
	PDF               WMSFormat = "application/pdf"
	GeoRSS            WMSFormat = "rss"
	KML               WMSFormat = "kml"
	KMZ               WMSFormat = "kmz"
	MapML             WMSFormat = "text/mapml"
	MapML_HTML_Viewer WMSFormat = "text/html;subtype=mapml"
	OpenLayers        WMSFormat = "application/openlayers"
	UTFGrid           WMSFormat = "application/json;type=utfgrid"
)
