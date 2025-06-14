package coveragestores

import (
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type CoverageStoreRetrievalWrapper struct {
	CoverageStore CoverageStore `json:"coverageStore"`
}

type CoverageStore struct {
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

type ConnectionParameters struct {
	Entry []Entry `json:"entry"`
}

func (cp *ConnectionParameters) Get(key string) (value string, ok bool) {
	for _, entry := range cp.Entry {
		if entry.Key == key {
			return entry.Value, true
		}
	}
	return value, false
}

type Entry struct {
	Key   string `json:"@key"`
	Value string `json:"$"`
}

type AllCoverageStoreRetrievalWrapper struct {
}
