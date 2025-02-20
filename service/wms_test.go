package service

import (
	"encoding/xml"
	"github.com/canghel3/go-geoserver/models/wms"
	"github.com/canghel3/go-geoserver/utils"
	"gotest.tools/v3/assert"
	"testing"
)

func TestGetCapabilities(t *testing.T) {
	geoserverService := NewGeoserverService(geoserverURL, username, password)
	assert.NilError(t, geoserverService.CreateWorkspace("init"))
	assert.NilError(t, geoserverService.CreateCoverageStore("init", "init", "file:/opt/geoserver/data/shipments_2_geocoded.tif", "GeoTIFF"))

	keywords := []string{"adu", "telefonul", "marian"}
	bbox := [4]float64{-13625746.1231970004737377, -5661864.1336409999057651, 19523253.8768029995262623, 8287135.8663590000942349}
	assert.NilError(t, geoserverService.CreateCoverage("init", "init", "init", "EPSG:3857", bbox, utils.ProjectionPolicyOption("FORCE_DECLARED"), utils.KeywordsOption(keywords)))

	capabilities, err := geoserverService.GetCapabilities("1.3.0")
	assert.NilError(t, err)

	expectedCapabilities := wms.Capabilities{
		XMLName: xml.Name{},
		Service: wms.Service{},
		Capability: wms.Capability{
			Layer: wms.Layer{
				Layers: []wms.Layer{
					{
						Queryable: "1",
						Opaque:    "0",
						Name:      "init:init",
						Title:     "init",
						Abstract:  "",
						CRS:       []string{"EPSG:3857", "CRS:84"},
						EXGeographicBoundingBox: wms.EXGeographicBoundingBox{
							WestBoundLongitude: -122.40216000000181,
							EastBoundLongitude: 175.38037353277835,
							NorthBoundLatitude: 59.490360000000344,
							SouthBoundLatitude: -45.255685383756436,
						},
						BoundingBox: []wms.BoundingBox{
							{
								MinX: -122.40216000000181,
								MaxX: 175.38037353277835,
								MinY: -45.255685383756436,
								MaxY: 59.490360000000344,
								CRS:  "CRS:84",
							},
							{
								MinX: -1.3625746123197e+07,
								MaxX: 1.9523253876803e+07,
								MinY: -5.661864133641e+06,
								MaxY: 8.287135866359e+06,
								CRS:  "EPSG:3857",
							},
						},
						Style: []wms.Style{
							{
								Name:     "raster",
								Title:    "Opaque Raster",
								Abstract: "A sample style for rasters",
							},
						},
						KeywordList: []string{"adu", "telefonul", "marian"},
					},
				},
			},
		},
	}

	expectedCapabilities.Capability.Layer.Layers[0].Style[0].LegendURL = capabilities.Capability.Layer.Layers[0].Style[0].LegendURL
	assert.DeepEqual(t, expectedCapabilities.Capability.Layer.Layers[0], capabilities.Capability.Layer.Layers[0])

	//DONT FORGET TO DELETE YOUR WORKSPACE AFTER TESTING
	assert.NilError(t, geoserverService.DeleteWorkspace("init", utils.RecurseOption(true)))
}
