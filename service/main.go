package service

import (
	"fmt"
	"github.com/canghel3/go-geoserver/internal/datastore/postgis"
	"github.com/rs/zerolog/log"
)

func Foo() {
	geosvc := NewGeoserverService("http://localhost:8080", "admin", "geoserver")
	v := geosvc.Workspace("PLAYGROUND").Vectors()
	s := v.Stores().PostGIS("ok", postgis.ConnectionParams{
		Host:     "localhost",
		Database: "vectors",
		User:     "geoserver",
		Password: "geoserver",
		Port:     "5432",
		SSL:      "disable",
	})

	err := v.Stores().Create(s)
	if err != nil {
		log.Error().Err(err)
	}

	store, err := v.Stores().Get("ok")
	if err != nil {
		log.Error().Err(err)
	}

	fmt.Println(store)

	err = v.Stores().Delete("ok", false)
	if err != nil {
		log.Error().Msg(err.Error())
	}

}
