package requester

import (
	"bytes"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"io"
	"net/http"
)

type GeoWebCacheRequester struct {
	data internal.GeoserverData
}

func NewGeoWebCacheRequester(data internal.GeoserverData) GeoWebCacheRequester {
	return GeoWebCacheRequester{data: data}
}

//func (gwcr GeoWebCacheRequester) Layers() ([]string, error) {
//	var target = fmt.Sprintf("%s/geoserver/gwc/rest/layers", gwcr.data.Connection.URL)
//
//	request, err := http.NewRequest(http.MethodGet, target, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	request.SetBasicAuth(gwcr.data.Connection.Credentials.Username, gwcr.data.Connection.Credentials.Password)
//
//	response, err := gwcr.data.Client.Do(request)
//	if err != nil {
//		return nil, err
//	}
//	defer response.Body.Close()
//
//	switch response.StatusCode {
//	case http.StatusOK:
//		var layers []string
//
//		err = json.NewDecoder(response.Body).Decode(&layers)
//		if err != nil {
//			return nil, err
//		}
//
//		return layers, nil
//	default:
//		body, err := io.ReadAll(response.Body)
//		if err != nil {
//			return nil, err
//		}
//
//		return nil, customerrors.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
//	}
//}
//
//func (gwcr GeoWebCacheRequester) Layer(name string) (*gwc.Layer, error) {
//	return nil, nil
//}

func (gwcr GeoWebCacheRequester) Seed(name string, content []byte) error {
	var target = fmt.Sprintf("%s/geoserver/gwc/rest/seed/%s.json", gwcr.data.Connection.URL, name)

	request, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(gwcr.data.Connection.Credentials.Username, gwcr.data.Connection.Credentials.Password)

	response, err := gwcr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return customerrors.NewGeoserverError(fmt.Sprintf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
