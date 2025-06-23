package models

type GenericStoreOptions struct {
	Description              string
	AutoDisableOnConnFailure bool
}

type GenericCoverageStoreCreationWrapper struct {
	CoverageStore GenericCoverageStoreCreationModel `json:"coverageStore"`
}

type GenericCoverageStoreCreationModel struct {
	Name                     string `json:"name"`
	Description              string `json:"description"`
	Type                     string `json:"type"`
	Enabled                  bool   `json:"enabled"`
	AutoDisableOnConnFailure bool   `json:"disableOnConnFailure"`
	Workspace                struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"workspace"`
	URL       string `json:"url"`
	Coverages struct {
		Link string `json:"link"`
	} `json:"coverages"`
}
