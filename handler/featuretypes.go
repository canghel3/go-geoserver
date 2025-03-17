package handler

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/featuretypes"
	"github.com/canghel3/go-geoserver/internal"
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
	content, err := json.Marshal(featureType)
	if err != nil {
		return err
	}

	return ft.requester.FeatureTypes().Create(ft.store, content)
}

func (ft *FeatureTypes) GetFeature(name string) (*featuretypes.GetFeatureTypeWrapper, error) {
	return ft.requester.FeatureTypes().Get(ft.store, name)
}

func (ft *FeatureTypes) DeleteFeature(name string, recurse bool) error {
	return ft.requester.FeatureTypes().Delete(ft.store, name, recurse)
}
