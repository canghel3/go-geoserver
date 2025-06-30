package actions

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/logging"
)

type Logging struct {
	requester requester.LoggingRequester
}

func NewLoggingActions(data internal.GeoserverData) Logging {
	return Logging{
		requester: requester.NewLoggingRequester(data),
	}
}

// Get retrieves logs from the server
func (l Logging) Get() (*logging.Log, error) {
	return l.requester.Get()
}

// Put creates a new log entry
func (l Logging) Put(log models.Log) error {
	content, err := json.Marshal(logging.LogRequestWrapper{LogRequest: logging.LogRequest(log)})
	if err != nil {
		return err
	}

	return l.requester.Put(content)
}
