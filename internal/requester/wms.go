package requester

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/options"
	"github.com/canghel3/go-geoserver/pkg/shared"
	"github.com/canghel3/go-geoserver/pkg/wms"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type WMSRequester struct {
	data internal.GeoserverData
}

func (wmsR *WMSRequester) GetCapabilities(version wms.WMSVersion) ([]byte, error) {
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return body, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (wmsR *WMSRequester) GetMap(width, height uint16, layers []string, bbox shared.BBOX, version wms.WMSVersion, format wms.WMSFormat, options ...options.GetMapOption) ([]byte, error) {
	u, err := url.Parse(fmt.Sprintf("%s/geoserver/wms", wmsR.data.Connection.URL))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("version", string(version))
	q.Add("service", "WMS")
	q.Add("request", "GetMap")
	q.Add("width", strconv.FormatUint(uint64(width), 10))
	q.Add("height", strconv.FormatUint(uint64(height), 10))
	q.Add("layers", strings.Join(layers, ","))
	q.Add("styles", "")
	q.Add("bbox", bbox.ToString())
	q.Add("format", string(format))
	switch version {
	case wms.Version130:
		q.Add("crs", bbox.SRS)
	case wms.Version111:
		q.Add("srs", bbox.SRS)
	}

	for _, option := range options {
		option(&q)
	}

	u.RawQuery = q.Encode()

	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(wmsR.data.Connection.Credentials.Username, wmsR.data.Connection.Credentials.Password)

	response, err := wmsR.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		return body, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
