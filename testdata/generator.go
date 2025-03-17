package testdata

import "github.com/canghel3/go-geoserver/internal"

func GeoserverInfo(client internal.HTTPClient) *internal.GeoserverInfo {
	return &internal.GeoserverInfo{
		Client: client,
		Connection: internal.GeoserverConnection{
			URL: GEOSERVER_URL,
			Credentials: internal.GeoserverCredentials{
				Username: GEOSERVER_USERNAME,
				Password: GEOSERVER_PASSWORD,
			},
		},
		DataDir:   GEOSERVER_DATADIR,
		Workspace: WORKSPACE,
	}
}
