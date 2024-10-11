package service

import (
	"net/http"
)

type Service interface {
	Create() error
	Delete() error
	Ping() error
}

type GeoserverServiceOption func(gs *GeoserverService)

/*
Sets the data directory field of the GeoserverService
*/
func GeoserverServiceDataDirOption(datadir string) GeoserverServiceOption {
	return func(gs *GeoserverService) {
		gs.datadir = datadir
	}
}

type GeoserverService struct {
	client   *http.Client
	url      string
	username string
	password string

	datadir string
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
		client:   &http.Client{},
		url:      url,
		username: username,
		password: password,
	}

	for _, option := range options {
		option(gs)
	}

	return gs
}

func (gs *GeoserverService) isDataDirectorySet() bool {
	return len(gs.datadir) > 0
}

func (gs *GeoserverService) getDataDirectory() string {
	return gs.datadir
}
