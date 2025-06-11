package actions

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/logging"
)

type Logging struct {
	requester *requester.Requester
}

func NewLoggingActions(data internal.GeoserverData) *Logging {
	return &Logging{
		requester: requester.NewRequester(data),
	}
}

// Get retrieves logs from the server
func (l *Logging) Get() (*logging.LogResponse, error) {
	return l.requester.Logging().Get()
}

// Put creates a new log entry
func (l *Logging) Put(message, level, source string, stdOutLogging bool) error {
	logRequest := logging.LogRequest{
		Message:       message,
		Level:         level,
		Source:        source,
		StdOutLogging: stdOutLogging,
	}

	content, err := json.Marshal(logRequest)
	if err != nil {
		return err
	}

	return l.requester.Logging().Put(content)
}
