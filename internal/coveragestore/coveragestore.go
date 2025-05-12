package coveragestore

import (
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type CreateCoverageStoreWrapper struct {
	CoverageStore CreateCoverageStore `json:"coverageStore"`
}

type CreateCoverageStore struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Type        string                   `json:"type"`
	Workspace   workspace.MultiWorkspace `json:"workspace"`
	Enabled     bool                     `json:"enabled"`
	URL         string                   `json:"url"`
}

type GetCoverageStore struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Type        string                   `json:"type"`
	Enabled     bool                     `json:"enabled"`
	Workspace   workspace.MultiWorkspace `json:"workspace"`
	URL         string                   `json:"url"`
	Coverages   string                   `json:"coverages"`
}

type GetCoverageStoreWrapper struct {
	CoverageStore GetCoverageStore `json:"coverageStore"`
}
