package testdata

import "github.com/canghel3/go-geoserver/internal"

func GeoserverInfo(client internal.HTTPClient) *internal.GeoserverInfo {
	return &internal.GeoserverInfo{
		Client: client,
		Connection: internal.GeoserverConnection{
			URL: GeoserverUrl,
			Credentials: internal.GeoserverCredentials{
				Username: GeoserverUsername,
				Password: GeoserverPassword,
			},
		},
		DataDir:   GeoserverDatadir,
		Workspace: WORKSPACE,
	}
}
