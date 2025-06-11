package logging

import "github.com/canghel3/go-geoserver/internal/models"

// Log represents a log entry
type Log struct {
	Level         string `json:"level"`
	Location      string `json:"location"`
	StdOutLogging bool   `json:"stdOutLogging"`
}

// LogResponse represents a response containing logs
type LogResponse struct {
	Log Log `json:"logging"`
}

// LogRequest represents a request to create a log entry
type LogRequest struct {
	Message       string `json:"message"`
	Level         string `json:"level"`
	Source        string `json:"source"`
	StdOutLogging bool   `json:"stdOutLogging"`
}

func NewLog(message, level, source string) *models.Log {
	return &models.Log{
		Message:       message,
		Level:         level,
		Source:        source,
		StdOutLogging: false,
	}
}
