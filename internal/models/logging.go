package models

type Log struct {
	Message       string `json:"message"`
	Level         string `json:"level"`
	Source        string `json:"source"`
	StdOutLogging bool   `json:"stdOutLogging"`
}
