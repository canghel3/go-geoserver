package handler

import (
	"encoding/json"
	"errors"
	"github.com/canghel3/go-geoserver/featuretypes"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/misc"
	"github.com/canghel3/go-geoserver/internal/requester"
)

type FeatureTypes struct {
	store     string
	requester *requester.Requester
}

func newFeatureTypes(store string, info *internal.GeoserverInfo) *FeatureTypes {
	return &FeatureTypes{
		store:     store,
		requester: requester.NewRequester(info),
	}
}

func (ft *FeatureTypes) PublishFeature(featureType featuretypes.FeatureType) error {
	//TODO: decide how to automatically infer bbox

	completeFeatureType := internal.CreateFeatureType{
		Name:              featureType.Name,
		NativeName:        featureType.NativeName,
		Namespace:         internal.Namespace{},
		Srs:               "",
		NativeBoundingBox: misc.BoundingBox{},
		ProjectionPolicy:  "",
		Keywords:          nil,
		Title:             featureType.Title,
		Store:             internal.Store{},
	}

	content, err := json.Marshal(completeFeatureType)
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

func (ft *FeatureTypes) UpdateFeature(featureType featuretypes.FeatureType) error {
	return errors.New("not implemented")
}

func (ft *FeatureTypes) DeleteFeature(name string, recurse bool) error {
	return ft.requester.FeatureTypes().Delete(ft.store, name, recurse)
}

// Reset the cache of the specified feature type.
func (ft *FeatureTypes) Reset(name string) error {
	return errors.New("not implemented")
}
