package handler

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
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
	return nil
}

func (ft *FeatureTypes) GetFeature(name string) (*featuretypes.FeatureType, error) {
	return nil, nil
}

func (ft *FeatureTypes) DeleteFeature(name string) error {
	return nil
}
