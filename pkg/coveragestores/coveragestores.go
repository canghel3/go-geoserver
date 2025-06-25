package coveragestores

import (
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type CoverageStoreWrapper struct {
	CoverageStore CoverageStore `json:"coverageStore"`
}

type CoverageStoresWrapper struct {
	CoverageStores CoverageStores `json:"coverageStores"`
}
type CoverageStores struct {
	Entries []struct {
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"coverageStore"`
}

type CoverageStore struct {
	Name                       string                   `json:"name,omitempty"`
	Description                string                   `json:"description,omitempty"`
	Type                       string                   `json:"type,omitempty"`
	Enabled                    bool                     `json:"enabled,omitempty"`
	Workspace                  workspace.MultiWorkspace `json:"workspace,omitempty"`
	URL                        string                   `json:"url,omitempty"`
	DisableConnectionOnFailure bool                     `json:"disableOnConnFailure,omitempty"`
	Default                    bool                     `json:"_default,omitempty"`
	Coverages                  string                   `json:"coverages,omitempty"`
}
