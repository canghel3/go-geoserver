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
	"github.com/canghel3/go-geoserver/pkg/types"
	"strings"
)

func newDataStoresActions(info internal.GeoserverData) *DataStores {
	return &DataStores{
		info:      info,
		requester: requester.NewDataStoreRequester(info),
	}
}

type DataStoreList struct {
	options   *models.GenericStoreOptions
	requester requester.DataStoreRequester
}

type DataStores struct {
	info      internal.GeoserverData
	requester requester.DataStoreRequester
}

// Reset the caches related to the specified datastore.
func (ds *DataStores) Reset(name string) error {
	return ds.requester.Reset(name)
}

func (ds *DataStores) Use(name string) *FeatureTypes {
	return newFeatureTypes(name, ds.info.Clone())
}

// Create sets general store options and returns a list of available data stores to create.
func (ds *DataStores) Create(options ...options.GenericStoreOption) DataStoreList {
	dsl := DataStoreList{
		requester: ds.requester,
		options:   &models.GenericStoreOptions{},
	}

	for _, option := range options {
		option(dsl.options)
	}

	return dsl
}

func (ds *DataStores) Get(name string) (*datastores.DataStore, error) {
	return ds.requester.Get(name)
}

func (ds *DataStores) GetAll() (*datastores.DataStores, error) {
	return ds.requester.GetAll()
}

func (ds *DataStores) Update(name string, store datastores.DataStore) error {
	content, err := json.Marshal(datastores.DataStoreWrapper{DataStore: store})
	if err != nil {
		return err
	}

	return ds.requester.Update(name, content)
}

func (ds *DataStores) Delete(name string, recurse bool) error {
	return ds.requester.Delete(name, recurse)
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
		"dbtype":   string(types.PostGIS),
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.AutoDisableOnConnFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.Create(content)
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
		"dbtype":   string(types.GeoPackage),
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.AutoDisableOnConnFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.Create(content)
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
		"filetype": string(types.Shapefile),
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.AutoDisableOnConnFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.Create(content)
}

func (dsl DataStoreList) Shapefiles(name string, dir string, options ...options.ShapefileOption) error {
	err := validator.DataStore.ShapefileDirectory(dir)
	if err != nil {
		return err
	}

	var url string
	if strings.HasPrefix(dir, "file:") {
		url = dir
	} else {
		url = fmt.Sprintf("file:%s", dir)
	}

	cp := datastores.ConnectionParams{
		"url":    url,
		"fstype": string(types.DirOfShapefiles),
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       name,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.AutoDisableOnConnFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.Create(content)
}

//func (dsl DataStoreList) CSV(name string, filepath string, options ...options.CSVOptions) error {
//	err := validator.DataStore.CSV(filepath)
//	if err != nil {
//		return err
//	}
//
//	cp := datastores.ConnectionParams{
//		"url":    filepath,
//		"dbtype": string(types.CSV),
//	}
//
//	for _, option := range options {
//		option(&cp)
//	}
//
//	data := datastores.GenericDataStoreCreationWrapper{
//		DataStore: datastores.GenericDataStoreCreationModel{
//			Name:                       name,
//			Description:                dsl.options.Description,
//			DisableOnConnectionFailure: dsl.options.DisableOnConnectionFailure,
//			ConnectionParameters: datastores.ConnectionParameters{
//				Entry: cp.ToDatastoreEntries(),
//			},
//		},
//	}
//
//	content, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	return dsl.requester.DataStores().Create(content)
//}

func (dsl DataStoreList) WebFeatureService(storeName, username, password, wfsCapabilitiesUrl string, options ...options.WFSOptions) error {
	err := validator.DataStore.WebFeatureService(wfsCapabilitiesUrl)
	if err != nil {
		return err
	}

	cp := datastores.ConnectionParams{
		"WFSDataStoreFactory:GET_CAPABILITIES_URL": wfsCapabilitiesUrl,
		"WFSDataStoreFactory:USERNAME":             username,
		"WFSDataStoreFactory:PASSWORD":             password,
	}

	for _, option := range options {
		option(&cp)
	}

	data := datastores.GenericDataStoreCreationWrapper{
		DataStore: datastores.GenericDataStoreCreationModel{
			Name:                       storeName,
			Description:                dsl.options.Description,
			DisableOnConnectionFailure: dsl.options.AutoDisableOnConnFailure,
			ConnectionParameters: datastores.ConnectionParameters{
				Entry: cp.ToDatastoreEntries(),
			},
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dsl.requester.Create(content)
}
