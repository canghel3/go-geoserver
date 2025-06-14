package models

type CoverageStoreOptions struct {
	Description string
	Default     bool
}

type GenericCoverageStoreCreationWrapper struct {
	CoverageStore GenericCoverageStoreCreationModel `json:"coverageStore"`
}

type GenericCoverageStoreCreationModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Enabled     bool   `json:"enabled"`
	Default     bool   `json:"__default__"`
	Workspace   struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"workspace"`
	URL       string `json:"url"`
	Coverages struct {
		Link string `json:"link"`
	} `json:"coverages"`
}
