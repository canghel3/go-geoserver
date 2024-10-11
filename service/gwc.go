package service

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models/gwc"
	"github.com/canghel3/go-geoserver/utils"
	"io"
	"net/http"
)

func (gs *GeoserverService) GetTileCaching(layer string, options ...utils.Option) (*gwc.GeoServerLayer, error) {
	params := utils.ProcessOptions(options)

	var target string
	if wksp, set := params["workspace"]; set {
		_, err := gs.GetWorkspace(wksp.(string))
		if err != nil {
			return nil, err
		}

		target = fmt.Sprintf("%s/geoserver/gwc/rest/layers/%s:%s", gs.url, wksp, layer)
	} else {
		target = fmt.Sprintf("%s/geoserver/gwc/rest/layers/%s", gs.url, layer)
	}

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/xml")
	request.SetBasicAuth(gs.username, gs.password)

	response, err := gs.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var tiledTab gwc.GeoServerLayer
		err = xml.NewDecoder(response.Body).Decode(&tiledTab)
		if err != nil {
			return nil, err
		}

		return &tiledTab, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (gs *GeoserverService) UpdateTileCaching(layer string, updateData gwc.TileTabUpdateData, options ...utils.Option) error {
	cachedLayer, err := gs.GetTileCaching(layer, options...)
	if err != nil {
		return err
	}

	if len(updateData.Blobstore) > 0 {
		cachedLayer.BlobStore = updateData.Blobstore
	}

	if len(updateData.GridSubsets) > 0 {
		cachedLayer.GridSubsets = updateData.GridSubsets
	}

	if len(updateData.MimeFormats) > 0 {
		cachedLayer.MimeFormats = updateData.MimeFormats
	} else {
		cachedLayer.MimeFormats = []string{"image/png8"}
	}

	if len(updateData.ParameterFilters) > 0 {
		cachedLayer.ParameterFilters = updateData.ParameterFilters
	}

	content, err := xml.Marshal(cachedLayer)
	if err != nil {
		return err
	}

	params := utils.ProcessOptions(options)

	var target string
	if wksp, set := params["workspace"]; set {
		_, err = gs.GetWorkspace(wksp.(string))
		if err != nil {
			return err
		}

		target = fmt.Sprintf("%s/geoserver/gwc/rest/layers/%s:%s", gs.url, wksp, layer)
	} else {
		target = fmt.Sprintf("%s/geoserver/gwc/rest/layers/%s", gs.url, layer)
	}

	request, err := http.NewRequest(http.MethodPut, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/xml")
	request.SetBasicAuth(gs.username, gs.password)

	response, err := gs.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

// Seed gridSetId is of the form "EPSG:4326" or "EPSG:3857"
func (gs *GeoserverService) Seed(workspace, layer, gridSetId, format string, zoomStart, zoomStop, threadCount uint8) error {
	data := gwc.CacheWrapper{
		SeedRequest: gwc.Cache{
			Name:        fmt.Sprintf("%s:%s", workspace, layer),
			GridSetID:   gridSetId,
			ZoomStart:   zoomStart,
			ZoomStop:    zoomStop,
			Type:        "seed",
			ThreadCount: threadCount,
			Format:      format,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	target := fmt.Sprintf("%s/geoserver/gwc/rest/seed/%s.json", gs.url, fmt.Sprintf("%s:%s", workspace, layer))
	request, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gs.username, gs.password)

	response, err := gs.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
