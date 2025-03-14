package requester

import (
	"github.com/canghel3/go-geoserver/internal"
)

type Requester struct {
	workspaces *WorkspaceRequester
	datastores *DataStoreRequester
	wms        *WMSRequester
}

func NewRequester(info *internal.GeoserverInfo) *Requester {
	return &Requester{
		workspaces: &WorkspaceRequester{
			info: info,
		},
		datastores: &DataStoreRequester{
			info: info,
		},
		wms: &WMSRequester{
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

func (r *Requester) WMS() *WMSRequester {
	return r.wms
}
