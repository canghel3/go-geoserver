package datastores

import (
	"github.com/canghel3/go-geoserver/workspace"
)

type DataStoreRetrievalWrapper struct {
	DataStore DataStoreRetrieval `json:"dataStore"`
}

type DataStoreRetrieval struct {
	Name                 string                   `json:"name,omitempty"`
	Enabled              bool                     `json:"enabled,omitempty"`
	Workspace            workspace.MultiWorkspace `json:"workspace,omitempty"`
	ConnectionParameters ConnectionParameters     `json:"connectionParameters"`
	Default              bool                     `json:"_default,omitempty"`
	FeatureTypes         string                   `json:"featureTypes,omitempty"`
}
