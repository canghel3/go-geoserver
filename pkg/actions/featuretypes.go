package actions

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
)

type FeatureTypes struct {
	store     string
	data      internal.GeoserverData
	requester requester.FeatureTypeRequester
}

func newFeatureTypes(store string, info internal.GeoserverData) FeatureTypes {
	return FeatureTypes{
		store:     store,
		data:      info,
		requester: requester.NewFeatureTypeRequester(info),
	}
}

func (ft FeatureTypes) Publish(featureType models.FeatureType) error {
	featureType.Namespace = models.Namespace{
		Name: ft.data.Workspace,
		Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", ft.data.Connection.URL, ft.data.Workspace),
	}

	featureType.Store = models.Store{
		Class: "dataStore",
		Name:  fmt.Sprintf("%s:%s", ft.data.Workspace, ft.store),
		Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s.json", ft.data.Connection.URL, ft.data.Workspace, ft.store),
	}

	content, err := json.Marshal(models.CreateFeatureTypeWrapper{FeatureType: featureType})
	if err != nil {
		return err
	}

	return ft.requester.Create(ft.store, content)
}

func (ft FeatureTypes) Get(name string) (*featuretypes.FeatureType, error) {
	return ft.requester.Get(ft.store, name)
}

func (ft FeatureTypes) GetAll() (*featuretypes.FeatureTypes, error) {
	return ft.requester.GetAll(ft.store)
}

func (ft FeatureTypes) Update(name string, featureType featuretypes.FeatureType) error {
	featureType.Namespace = featuretypes.Namespace{
		Name: ft.data.Workspace,
		Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", ft.data.Connection.URL, ft.data.Workspace),
	}

	featureType.Store = featuretypes.Store{
		Class: "dataStore",
		Name:  fmt.Sprintf("%s:%s", ft.data.Workspace, ft.store),
		Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s.json", ft.data.Connection.URL, ft.data.Workspace, ft.store),
	}

	content, err := json.Marshal(featuretypes.FeatureTypeWrapper{FeatureType: featureType})
	if err != nil {
		return err
	}

	return ft.requester.Update(ft.store, name, content)
}

func (ft FeatureTypes) Delete(name string, recurse bool) error {
	return ft.requester.Delete(ft.store, name, recurse)
}

// Reset the cache of the specified feature type.
func (ft FeatureTypes) Reset(name string) error {
	return ft.requester.Reset(ft.store, name)
}
