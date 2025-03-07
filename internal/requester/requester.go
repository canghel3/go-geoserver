package requester

import "github.com/canghel3/go-geoserver/internal"

type Requester struct {
	workspaces *WorkspaceRequester
	datastores *DataStoreRequester
}

func NewRequester(info *internal.GeoserverInfo) *Requester {
	return &Requester{
		workspaces: &WorkspaceRequester{
			info: info,
		},
		datastores: &DataStoreRequester{
			info: info,
		},
	}
}

func (r *Requester) Workspaces() *WorkspaceRequester {
	return r.workspaces
}

func (r *Requester) DataStores() *DataStoreRequester {
	return r.datastores
}
