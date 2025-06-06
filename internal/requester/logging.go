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
	data *internal.GeoserverData
}

// Get retrieves logs from the server
func (lr *LoggingRequester) Get() (*logging.LogResponse, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geoserver/rest/logs", lr.data.Connection.URL), nil)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		logResponse := &logging.LogResponse{}
		if err = json.Unmarshal(body, logResponse); err != nil {
			return nil, err
		}
		return logResponse, nil
	default:
		return nil, customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}

// Put creates a new log entry
func (lr *LoggingRequester) Put(logRequest *logging.LogRequest) error {
	jsonData, err := json.Marshal(logRequest)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/geoserver/rest/logs", lr.data.Connection.URL), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.SetBasicAuth(lr.data.Connection.Credentials.Username, lr.data.Connection.Credentials.Password)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := lr.data.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return nil
	default:
		return customerrors.WrapGeoserverError(fmt.Errorf("received status code %d from geoserver: %s", response.StatusCode, string(body)))
	}
}
