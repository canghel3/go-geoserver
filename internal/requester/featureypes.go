package requester

import (
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/pkg/featuretypes"
)

type FeatureTypeRequester struct {
	info *internal.GeoserverInfo
}

func (ftr *FeatureTypeRequester) Create(content []byte) error {
	return nil
}

func (ftr *FeatureTypeRequester) Delete(name string) error {
	return nil
}

func (ftr *FeatureTypeRequester) Get(name string) (*featuretypes.FeatureType, error) {
	return nil, nil
}
