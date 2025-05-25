package internal

import "net/http"

type GeoserverData struct {
	Client     http.Client
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

func (gi *GeoserverData) Clone() *GeoserverData {
	return &GeoserverData{
		Client:     gi.Client,
		Connection: gi.Connection,
		DataDir:    gi.DataDir,
		Workspace:  gi.Workspace,
	}
}
