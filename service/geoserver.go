package service

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/raster"
	"github.com/canghel3/go-geoserver/vector"
	"net/http"
)

type GeoserverServiceOption func(gs *GeoserverService)

/*
Sets the data directory field of the GeoserverService
*/
func GeoserverServiceDataDirOption(datadir string) GeoserverServiceOption {
	return func(gs *GeoserverService) {
		gs.data.dataDir = datadir
	}
}

type GeoserverService struct {
	data *internal.GeoserverInfo
}

/*
NewGeoserverService is a function that creates a new instance of GeoserverService.

url - url of geoserver without /geoserver (eg: http://localhost:8080)

username - geoserver username

password - geoserver password

options - GeoserverServiceDataDirOption
*/
func NewGeoserverService(url, username, password string, options ...GeoserverServiceOption) *GeoserverService {
	gs := &GeoserverService{
		data: &internal.GeoserverInfo{
			client: &http.Client{},
			connection: internal.GeoserverConnection{
				URL: url,
				Credentials: internal.GeoserverCredentials{
					Username: username,
					Password: password,
				},
			},
		},
	}

	for _, option := range options {
		option(gs)
	}

	return gs
}

func (gs *GeoserverService) Workspace(name string) *GeoserverService {
	gs.data.Workspace = name
	return gs
}

func (gs *GeoserverService) Vectors() *vector.Service {
	return vector.NewService(*gs.data)
}

func (gs *GeoserverService) Rasters() *raster.Service {
	return nil
}

func (gs *GeoserverService) isDataDirectorySet() bool {
	return len(gs.data.dataDir) > 0
}

func (gs *GeoserverService) getDataDirectory() string {
	return gs.data.dataDir
}
