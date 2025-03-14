package requester

import (
	"encoding/xml"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/models/wms"
	"io"
	"net/http"
)

type WMSRequester struct {
	info *internal.GeoserverInfo
}

func (wmsR *WMSRequester) GetCapabilities(version string) (*wms.Capabilities, error) {
	var target = fmt.Sprintf("%s/geoserver/wms?service=wms&version=%s&request=GetCapabilities", wmsR.info.Connection.URL, version)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(wmsR.info.Connection.Credentials.Username, wmsR.info.Connection.Credentials.Password)

	response, err := wmsR.info.Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var capabilities wms.Capabilities
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
