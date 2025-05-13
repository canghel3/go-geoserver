package handler

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/datastores"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/options"
)

type DataStoreType string

const (
	PostGIS    DataStoreType = "postgis"
	GeoPackage DataStoreType = "geopkg"
	CSV        DataStoreType = "csv"
)

func newDataStoresHandler(info *internal.GeoserverData) *DataStores {
	r := requester.NewRequester(info)
	return &DataStores{
		info:      info,
		requester: r,
	}
}

type DataStoreList struct {
	options   *internal.DatastoreOptions
	requester *requester.Requester
}

type DataStores struct {
	info      *internal.GeoserverData
	requester *requester.Requester
}

// Reset the caches related to the specified datastore.
func (ds *DataStores) Reset(name string) error {
	return ds.requester.DataStores().Reset(name)
}

func (ds *DataStores) Use(name string) *FeatureTypes {
	return newFeatureTypes(name, ds.info.Clone())
}

func (ds *DataStores) Create(options ...options.DatastoreOptionFunc) DataStoreList {
	dsl := DataStoreList{
		requester: ds.requester,
		options:   &internal.DatastoreOptions{},
	}

	for _, option := range options {
		option(dsl.options)
	}

	return dsl
}

func (ds *DataStores) Get(name string) (*datastores.DataStoreRetrieval, error) {
	return ds.requester.DataStores().Get(name)
}

func (ds *DataStores) Delete(name string, recurse bool) error {
	return ds.requester.DataStores().Delete(name, recurse)
}

func (dsl DataStoreList) PostGIS(name string, connectionParams postgis.ConnectionParams, options ...options.PostGISOptionFunc) error {
	cp := internal.ConnectionParams{
		"host":     connectionParams.Host,
		"database": connectionParams.Database,
		"user":     connectionParams.User,
		"passwd":   connectionParams.Password,
		"port":     connectionParams.Port,
		"dbtype":   string(PostGIS),
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.DisableOnConnectionFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.DataStores().Create(content)
}

func (dsl DataStoreList) GeoPackage(name string, filepath string, options ...options.GeoPackageOptions) error {
	err := internal.ValidateGeoPackage(filepath)
	if err != nil {
		return err
	}

	cp := internal.ConnectionParams{
		"database": filepath,
		"dbtype":   string(GeoPackage),
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.DisableOnConnectionFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.DataStores().Create(content)
}

func (dsl DataStoreList) Shapefile(name string, filepath string, options ...options.ShapefileOptions) error {
	err := internal.ValidateShapefile(filepath)
	if err != nil {
		return err
	}

	cp := internal.ConnectionParams{
		"url": filepath,
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.DisableOnConnectionFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.DataStores().Create(content)
}

func (dsl DataStoreList) Shapefiles(name string, dir string, options ...options.ShapefileOptions) error {
	err := internal.ValidateShapefileDirectory(dir)
	if err != nil {
		return err
	}

	cp := internal.ConnectionParams{
		"url": dir,
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.DisableOnConnectionFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.DataStores().Create(content)
}

func (dsl DataStoreList) CSV(name string, filepath string, options ...options.CSVOptions) error {
	err := internal.ValidateCSV(filepath)
	if err != nil {
		return err
	}

	cp := internal.ConnectionParams{
		"url":    filepath,
		"dbtype": string(CSV),
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.DisableOnConnectionFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.DataStores().Create(content)
}

func (dsl DataStoreList) WebFeatureService(name, getCapabilitiesUrl string, options ...options.WFSOptions) error {
	cp := internal.ConnectionParams{
		"GET_CAPABILITIES_URL": getCapabilitiesUrl,
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.DisableOnConnectionFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.DataStores().Create(content)
}
