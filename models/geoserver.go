package models

import "net/http"

type GeoserverInfo struct {
	Client     *http.Client
	Connection GeoserverConnection
	DataDir    string
	Workspace  string
}

type GeoserverConnection struct {
	URL         string
	Credentials GeoserverCredentials
}

type GeoserverCredentials struct {
	Username string
	Password string
}
