package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/featuretypes"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/misc"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type FeatureTypes struct {
	store     string
	info      *internal.GeoserverInfo
	requester *requester.Requester
}

func newFeatureTypes(store string, info *internal.GeoserverInfo) *FeatureTypes {
	return &FeatureTypes{
		store:     store,
		info:      info,
		requester: requester.NewRequester(info),
	}
}

// TODO: can this be so smart that it computes the bbox from data? decide how to automatically infer bbox
func (ft *FeatureTypes) PublishFeature(featureType featuretypes.CreateFeatureType) error {

	completeFeatureType := internal.CreateFeatureType{
		Name:       featureType.Name,
		NativeName: featureType.NativeName,
		Namespace: internal.Namespace{
			Name: ft.info.Workspace,
			Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", ft.info.Connection.URL, ft.info.Workspace),
		},
		Srs: "EPSG:4326",
		NativeBoundingBox: misc.BoundingBox{
			MinX: -180,
			MaxX: 180,
			MinY: -90,
			MaxY: 90,
			CRS:  "EPSG:4326",
		},
		ProjectionPolicy: "FORCE_DECLARED",
		Keywords:         nil,
		Title:            featureType.Title,
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
