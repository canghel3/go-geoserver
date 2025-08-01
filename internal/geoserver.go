package internal

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GeoserverData struct {
	Client     HTTPClient
	Connection GeoserverConnection
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

func (gi GeoserverData) Clone() GeoserverData {
	return GeoserverData{
		Client:     gi.Client,
		Connection: gi.Connection,
		Workspace:  gi.Workspace,
	}
}
