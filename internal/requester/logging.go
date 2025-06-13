package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/customerrors"
	"github.com/canghel3/go-geoserver/pkg/logging"
	"io"
	"net/http"
)

// LoggingRequester handles requests related to logging
type LoggingRequester struct {
	data internal.GeoserverData
}

// Get retrieves logs from the server
func (lr *LoggingRequester) Get() (*logging.Log, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/logging", lr.data.Connection.URL), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(lr.data.Connection.Credentials.Username, lr.data.Connection.Credentials.Password)
	request.Header.Add("Accept", "application/json")

	response, err := lr.data.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var log logging.LogResponse
		err = json.NewDecoder(response.Body).Decode(&log)
		if err != nil {
			return nil, err
		}

		return &log.Log, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

func (lr *LoggingRequester) Put(content []byte) error {
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/geoserver/rest/logging", lr.data.Connection.URL), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	request.SetBasicAuth(lr.data.Connection.Credentials.Username, lr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")

	response, err := lr.data.Client.Do(request)
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

		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
