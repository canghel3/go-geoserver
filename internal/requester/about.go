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
	data *internal.GeoserverData
}

func (ar *AboutRequester) Manifest() (*about.Manifest, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/about/manifest", ar.data.Connection.URL), nil)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusCreated:
		manifest := &about.ManifestResponse{}
		if err = json.Unmarshal(body, manifest); err != nil {
			return nil, err
		}
		return &manifest.Manifest, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ar *AboutRequester) Version() (*about.Version, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/about/version", ar.data.Connection.URL), nil)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusCreated:
		version := &about.VersionResponse{}
		if err = json.Unmarshal(body, version); err != nil {
			return nil, err
		}
		return &version.About, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ar *AboutRequester) Status() (*about.Status, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/about/status", ar.data.Connection.URL), nil)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusCreated:
		status := &about.StatusResponse{}
		if err = json.Unmarshal(body, status); err != nil {
			return nil, err
		}
		return &status.About, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (ar *AboutRequester) SystemStatus() (*about.Metrics, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/geoserver/rest/about/status", ar.data.Connection.URL), nil)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusCreated:
		metrics := &about.MetricsResponse{}
		if err = json.Unmarshal(body, metrics); err != nil {
			return nil, err
		}
		return &metrics.Metrics, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
