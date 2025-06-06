package logging

// Log represents a log entry
type Log struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Level     string `json:"level"`
	Timestamp string `json:"timestamp"`
	Source    string `json:"source"`
	Details   string `json:"details,omitempty"`
}

// LogResponse represents a response containing logs
type LogResponse struct {
	Logs []Log `json:"logs"`
}

// LogRequest represents a request to create a log entry
type LogRequest struct {
	Message string `json:"message"`
	Level   string `json:"level"`
	Source  string `json:"source"`
	Details string `json:"details,omitempty"`
}
