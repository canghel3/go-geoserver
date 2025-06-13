package requester

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/about"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"io"
	"net/http"
)

type AboutRequester struct {
	data internal.GeoserverData
}

func (ar *AboutRequester) Manifest() (*about.Manifest, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/about/manifest", ar.data.Connection.URL), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(ar.data.Connection.Credentials.Username, ar.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := ar.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var manifest about.ManifestResponse
		err = json.NewDecoder(response.Body).Decode(&manifest)
		if err != nil {
			return nil, err
		}

		return &manifest.Manifest, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ar *AboutRequester) Version() (*about.Version, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/about/version", ar.data.Connection.URL), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(ar.data.Connection.Credentials.Username, ar.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := ar.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var version about.VersionResponse
		err = json.NewDecoder(response.Body).Decode(&version)
		if err != nil {
			return nil, err
		}

		return &version.About, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ar *AboutRequester) Status() (*about.Status, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/about/status", ar.data.Connection.URL), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(ar.data.Connection.Credentials.Username, ar.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := ar.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var status about.StatusResponse
		err = json.NewDecoder(response.Body).Decode(&status)
		if err != nil {
			return nil, err
		}

		return &status.Status, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ar *AboutRequester) SystemStatus() (*about.Metrics, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/about/status", ar.data.Connection.URL), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(ar.data.Connection.Credentials.Username, ar.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := ar.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var metrics about.MetricsResponse
		err = json.NewDecoder(response.Body).Decode(&metrics)
		if err != nil {
			return nil, err
		}

		return &metrics.Metrics, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
