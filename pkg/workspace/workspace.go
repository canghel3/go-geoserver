package workspace

type MultiWorkspaceRetrievalWrapper struct {
	Workspaces MultiWorkspaceRetrieval `json:"workspaces" xml:"workspaces"`
}

type MultiWorkspaceRetrieval struct {
	Workspace []MultiWorkspace `json:"workspace" xml:"workspace"`
}

type MultiWorkspace struct {
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

// GetSingleWorkspaceWrapper defines the model for retrieving a single workspace.
type GetSingleWorkspaceWrapper struct {
	Workspace WorkspaceRetrieval `json:"workspace" xml:"workspace"`
}

type WorkspaceRetrieval struct {
	Name           string `json:"name,omitempty" xml:"name"`
	DataStores     string `json:"dataStores,omitempty" xml:"dataStores"`
	CoverageStores string `json:"coverageStores,omitempty" xml:"coverageStores"`
	WMSStores      string `json:"wmsStores,omitempty" xml:"wmsStores"`
}

type WorkspaceCreation struct {
	Name string `json:"name" xml:"name"`
}

// WorkspaceCreationWrapper defines the model for creating a workspace in GeoServer.
type WorkspaceCreationWrapper struct {
	Workspace WorkspaceCreation `json:"workspace" xml:"workspace"`
}

type WorkspaceUpdateWrapper struct {
	Workspace WorkspaceUpdate `json:"workspace" xml:"workspace"`
}

type WorkspaceUpdate struct {
	Name string `json:"name" xml:"name"`
}
