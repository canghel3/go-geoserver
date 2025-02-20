package service

import "github.com/canghel3/go-geoserver/models/datastore/postgis"

func x() {
	geosvc := NewGeoserverService("http://localhost:8080", "admin", "geoserver")
	err := geosvc.Workspace("a").Vectors().Stores().PostGIS("ok", postgis.ConnectionParams{
		Host:     "",
		Database: "",
		User:     "",
		Password: "",
		Port:     "",
		SSL:      "",
	}).Create()
	if err != nil {
		return
	}
}
