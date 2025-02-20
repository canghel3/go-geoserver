package vector

import (
	"fmt"
	"github.com/canghel3/go-geoserver/customerrors"
	"github.com/canghel3/go-geoserver/models"
	"github.com/canghel3/go-geoserver/models/datastore/postgis"
	"reflect"
	"strings"
)

type PostGISStore struct {
	name             string
	connectionParams postgis.ConnectionParams
	geoserverInfo    models.GeoserverInfo
}

func newPostGISStore(name string, connectionParams postgis.ConnectionParams, geoserverInfo models.GeoserverInfo) *PostGISStore {
	return &PostGISStore{
		name:             name,
		connectionParams: connectionParams,
		geoserverInfo:    geoserverInfo,
	}
}

func (pgs *PostGISStore) Create() error {
	v := reflect.ValueOf(&pgs.connectionParams).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			return customerrors.WrapInputError(fmt.Errorf("%v cannot be empty", strings.ToLower(v.Type().Field(i).Name)))
		}
	}

	cp := map[string]string{
		"host":     pgs.connectionParams.Host,
		"database": pgs.connectionParams.Database,
		"user":     pgs.connectionParams.User,
		"passwd":   pgs.connectionParams.Password,
		"port":     pgs.connectionParams.Port,
		"dbtype":   "postgis",
	}

	return createGenericDataStore()
}
