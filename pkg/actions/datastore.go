package actions

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/datastores"
	"github.com/canghel3/go-geoserver/pkg/datastores/postgis"
	"github.com/canghel3/go-geoserver/pkg/options"
	"strings"
)

type DataStoreType string

const (
	PostGIS    DataStoreType = "postgis"
	Shapefile  DataStoreType = "shapefile"
	GeoPackage DataStoreType = "geopkg"
	CSV        DataStoreType = "csv"
)

func newDataStoresActions(info *internal.GeoserverData) *DataStores {
	r := requester.NewRequester(info)
	return &DataStores{
		info:      info,
		requester: r,
	}
}

type DataStoreList struct {
	options   *models.DataStoreOptions
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

// Create sets general store options and returns a list of available data stores to create.
func (ds *DataStores) Create(options ...options.DataStoreOption) DataStoreList {
	dsl := DataStoreList{
		requester: ds.requester,
		options:   &models.DataStoreOptions{},
	}

	for _, option := range options {
		option(dsl.options)
	}

	return dsl
}

func (ds *DataStores) Get(name string) (*datastores.DataStore, error) {
	return ds.requester.DataStores().Get(name)
}

func (ds *DataStores) Delete(name string, recurse bool) error {
	return ds.requester.DataStores().Delete(name, recurse)
}

func (dsl DataStoreList) PostGIS(name string, connectionParams postgis.ConnectionParams, options ...options.PostGISOption) error {
	err := validator.DataStore.PostGIS(name)
	if err != nil {
		return err
	}

	cp := datastores.ConnectionParams{
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
	err := validator.DataStore.GeoPackage(filepath)
	if err != nil {
		return err
	}

	var url string
	if strings.HasPrefix(filepath, "file:") {
		url = filepath
	} else {
		url = fmt.Sprintf("file:%s", filepath)
	}

	cp := datastores.ConnectionParams{
		"database": url,
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

func (dsl DataStoreList) Shapefile(name string, filepath string, options ...options.ShapefileOption) error {
	err := validator.DataStore.Shapefile(filepath)
	if err != nil {
		return err
	}

	var url string
	if strings.HasPrefix(filepath, "file:") {
		url = filepath
	} else {
		url = fmt.Sprintf("file:%s", filepath)
	}

	cp := datastores.ConnectionParams{
		"url":      url,
		"filetype": string(Shapefile),
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

func (dsl DataStoreList) Shapefiles(name string, dir string, options ...options.ShapefileOption) error {
	err := validator.DataStore.ShapefileDirectory(dir)
	if err != nil {
		return err
	}

	cp := datastores.ConnectionParams{
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
	err := validator.DataStore.CSV(filepath)
	if err != nil {
		return err
	}

	cp := datastores.ConnectionParams{
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
	cp := datastores.ConnectionParams{
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
