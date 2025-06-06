package datastores

import (
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type DataStoreRetrievalWrapper struct {
	DataStore DataStore `json:"dataStore"`
}

type DataStore struct {
	Name                       string                   `json:"name,omitempty"`
	Description                string                   `json:"description,omitempty"`
	DisableConnectionOnFailure bool                     `json:"disableOnConnFailure,omitempty"`
	Enabled                    bool                     `json:"enabled,omitempty"`
	Workspace                  workspace.MultiWorkspace `json:"workspace,omitempty"`
	ConnectionParameters       ConnectionParameters     `json:"connectionParameters"`
	Default                    bool                     `json:"_default,omitempty"`
	FeatureTypes               string                   `json:"featureTypes,omitempty"`
}

type AllDataStoreRetrievalWrapper struct {
}
