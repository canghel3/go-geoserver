package actions

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/gwc"
	"strings"
)

type GeoWebCache struct {
	data      internal.GeoserverData
	requester requester.GeoWebCacheRequester
}

func NewGeoWebCache(data internal.GeoserverData) GeoWebCache {
	return GeoWebCache{
		data:      data,
		requester: requester.NewGeoWebCacheRequester(data),
	}
}

func (gwc GeoWebCache) Seed() Seed {
	return Seed{
		data:      gwc.data,
		requester: gwc.requester,
	}
}

type Seed struct {
	data      internal.GeoserverData
	requester requester.GeoWebCacheRequester
}

//func (s Seed) Statuses() ([]string, error) {
//	return nil, nil
//}
//
//func (s Seed) Status(layer string) error {
//	return nil
//}

func (s Seed) Run(seedData gwc.SeedData) error {
	err := validator.WorkspaceLayerFormat(s.data.Workspace, seedData.Layer)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(seedData.Layer, s.data.Workspace) {
		seedData.Layer = fmt.Sprintf("%s:%s", s.data.Workspace, seedData.Layer)
	}

	type seedRequest struct {
		SeedRequest gwc.SeedData `json:"seedRequest"`
	}

	content, err := json.Marshal(seedRequest{SeedRequest: seedData})
	if err != nil {
		return err
	}

	return s.requester.Seed(seedData.Layer, content)
}
