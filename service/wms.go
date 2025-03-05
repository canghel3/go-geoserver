package service

import (
	"encoding/xml"
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/wms"
	"io"
	"net/http"
)

const (
	VERSION_1_0_0 = "1.0.0"
	VERSION_1_1_0 = "1.1.0"
	VERSION_1_1_1 = "1.1.1"
	VERSION_1_3_0 = "1.3.0"
)

/*
GetCapabilities operation requests metadata about the operations, services, and data (“capabilities”) that are offered by a WMS server.

Versions: 1.0.0, 1.1.0, 1.1.1, 1.3.0 (if version is left empty, 1.3.0 is used by default)

Options: Workspace (namespace) option - limits response to layers in a given workspace (namespace)

TODO implement rootLayer and format Option
*/
func (gs *GeoserverService) GetCapabilities(version string, options ...internal.Option) (wms.Capabilities, error) {
	if len(version) == 0 {
		version = "1.3.0"
	}

	var target = fmt.Sprintf("%s/geoserver/wms?service=wms&version=%s&request=GetCapabilities", gs.data.connection.URL, version)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return wms.Capabilities{}, err
	}

	request.SetBasicAuth(gs.data.connection.Credentials.Username, gs.data.connection.Credentials.Password)

	params := internal.ProcessOptions(options)
	if wksp, set := params["workspace"]; set {
		target = fmt.Sprintf("%s&namespace=%s", target, wksp.(string))
	}

	response, err := gs.data.client.Do(request)
	if err != nil {
		return wms.Capabilities{}, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var capabilities wms.Capabilities
		err = xml.NewDecoder(response.Body).Decode(&capabilities)
		if err != nil {
			return wms.Capabilities{}, err
		}

		return capabilities, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return wms.Capabilities{}, err
		}

		return wms.Capabilities{}, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
