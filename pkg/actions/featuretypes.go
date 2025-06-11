package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
)

type FeatureTypes struct {
	store     string
	info      internal.GeoserverData
	requester *requester.Requester
}

func newFeatureTypes(store string, info internal.GeoserverData) *FeatureTypes {
	return &FeatureTypes{
		store:     store,
		info:      info,
		requester: requester.NewRequester(info),
	}
}

func (ft *FeatureTypes) Publish(featureType models.FeatureType) error {
	featureType.Namespace = models.Namespace{
		Name: ft.info.Workspace,
		Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", ft.info.Connection.URL, ft.info.Workspace),
	}

	featureType.Store = models.Store{
		Class: "dataStore",
		Name:  fmt.Sprintf("%s:%s", ft.info.Workspace, ft.store),
		Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s.json", ft.info.Connection.URL, ft.info.Workspace, ft.store),
	}

	content, err := json.Marshal(models.CreateFeatureTypeWrapper{FeatureType: featureType})
	if err != nil {
		return err
	}

	return ft.requester.FeatureTypes().Create(ft.store, content)
}

func (ft *FeatureTypes) Get(name string) (*featuretypes.GetFeatureType, error) {
	return ft.requester.FeatureTypes().Get(ft.store, name)
}

func (ft *FeatureTypes) GetAll() ([]featuretypes.GetFeatureType, error) {
	return nil, errors.New("not implemented")
}

func (ft *FeatureTypes) Update(featureType models.FeatureType) error {
	return errors.New("not implemented")
}

func (ft *FeatureTypes) Delete(name string, recurse bool) error {
	return ft.requester.FeatureTypes().Delete(ft.store, name, recurse)
}

// Reset the cache of the specified feature type.
func (ft *FeatureTypes) Reset(name string) error {
	return errors.New("not implemented")
}
