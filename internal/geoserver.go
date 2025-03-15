package internal

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GeoserverInfo struct {
	Client     HTTPClient
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

func (gi *GeoserverInfo) Clone() *GeoserverInfo {
	return &GeoserverInfo{
		Client:     gi.Client,
		Connection: gi.Connection,
		DataDir:    gi.DataDir,
		Workspace:  gi.Workspace,
	}
}
