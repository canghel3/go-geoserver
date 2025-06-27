package datastores

import (
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type DataStoreWrapper struct {
	DataStore DataStore `json:"dataStore"`
}

type DataStoresWrapper struct {
	DataStores DataStores `json:"dataStores"`
}

type DataStores struct {
	Entries []struct {
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"dataStore"`
}

type DataStore struct {
	Name                       string                   `json:"name,omitempty"`
	Description                string                   `json:"description,omitempty"`
	DisableConnectionOnFailure bool                     `json:"disableOnConnFailure,omitempty"`
	Enabled                    bool                     `json:"enabled,omitempty"`
	Workspace                  workspace.MultiWorkspace `json:"workspace,omitempty"`
	ConnectionParameters       ConnectionParameters     `json:"connectionParameters"`
	Default                    bool                     `json:"_default,omitempty"`
	DateCreated                string                   `json:"dateCreated,omitempty"`
	DateModified               string                   `json:"dateModified,omitempty"`
	Type                       string                   `json:"type,omitempty"`
	FeatureTypes               string                   `json:"featureTypes,omitempty"`
}
