package coveragestores

import (
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type CoverageStoreRetrievalWrapper struct {
	CoverageStore CoverageStoreRetrieval `json:"coverageStore"`
}

type CoverageStoreRetrieval struct {
	Name                       string                   `json:"name,omitempty"`
	Description                string                   `json:"description,omitempty"`
	Type                       string                   `json:"type,omitempty"`
	Enabled                    bool                     `json:"enabled,omitempty"`
	Workspace                  workspace.MultiWorkspace `json:"workspace,omitempty"`
	URL                        string                   `json:"url,omitempty"`
	DisableConnectionOnFailure bool                     `json:"disableOnConnFailure,omitempty"`
	ConnectionParameters       ConnectionParameters     `json:"connectionParameters,omitempty"`
	Default                    bool                     `json:"_default,omitempty"`
	Coverages                  string                   `json:"coverages,omitempty"`
}

type AllCoverageStoreRetrievalWrapper struct {
}