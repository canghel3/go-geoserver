package requester

import (
	"encoding/xml"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"io"
	"net/http"
)

type WMSRequester struct {
	data *internal.GeoserverData
}

func (wmsR *WMSRequester) GetCapabilities(version wms.WMSVersion) (*wms.Capabilities1_3_0, error) {
	var target = fmt.Sprintf("%s/geoserver/wms?service=wms&version=%s&request=GetCapabilities", wmsR.data.Connection.URL, version)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(wmsR.data.Connection.Credentials.Username, wmsR.data.Connection.Credentials.Password)

	response, err := wmsR.data.Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var capabilities wms.Capabilities1_3_0
		err = xml.NewDecoder(response.Body).Decode(&capabilities)
		if err != nil {
			return nil, err
		}

		return &capabilities, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wmsR *WMSRequester) GetMap(width, height uint, layers []string, srs string, bbox [4]float64, version wms.WMSVersion, format wms.WMSFormat) {

}
