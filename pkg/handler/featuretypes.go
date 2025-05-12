package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
)

type FeatureTypes struct {
	store     string
	info      *internal.GeoserverData
	requester *requester.Requester
}

func newFeatureTypes(store string, info *internal.GeoserverData) *FeatureTypes {
	return &FeatureTypes{
		store:     store,
		info:      info,
		requester: requester.NewRequester(info),
	}
}

func (ft *FeatureTypes) PublishFeature(featureType featuretypes.CreateFeatureType) error {
	completeFeatureType := internal.CreateFeatureType{
		Name:       featureType.Name,
		NativeName: featureType.NativeName,
		Namespace: internal.Namespace{
			Name: ft.info.Workspace,
			Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", ft.info.Connection.URL, ft.info.Workspace),
		},
		Srs:               featureType.Srs,
		NativeBoundingBox: featureType.Bbox,
		ProjectionPolicy:  featureType.ProjectionPolicy,
		Keywords:          featureType.Keywords,
		Title:             featureType.Title,
		Store: internal.Store{
			Class: "dataStore",
			Name:  fmt.Sprintf("%s:%s", ft.info.Workspace, ft.store),
			Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/datastores/%s.json", ft.info.Connection.URL, ft.info.Workspace, ft.store),
		},
	}

	content, err := json.Marshal(internal.CreateFeatureTypeWrapper{FeatureType: completeFeatureType})
	if err != nil {
		return err
	}

	return ft.requester.FeatureTypes().Create(ft.store, content)
}

func (ft *FeatureTypes) GetFeature(name string) (*featuretypes.GetFeatureTypeWrapper, error) {
	return ft.requester.FeatureTypes().Get(ft.store, name)
}

func (ft *FeatureTypes) GetFeatureTypes() ([]featuretypes.GetFeatureTypeWrapper, error) {
	return nil, errors.New("not implemented")
}

func (ft *FeatureTypes) UpdateFeature(featureType featuretypes.CreateFeatureType) error {
	return errors.New("not implemented")
}

func (ft *FeatureTypes) DeleteFeature(name string, recurse bool) error {
	return ft.requester.FeatureTypes().Delete(ft.store, name, recurse)
}

// Reset the cache of the specified feature type.
func (ft *FeatureTypes) Reset(name string) error {
	return errors.New("not implemented")
}
