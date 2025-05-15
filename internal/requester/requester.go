package requester

import (
	"github.com/canghel3/go-geoserver/internal"
)

type Requester struct {
	workspaces     *WorkspaceRequester
	datastores     *DataStoreRequester
	coveragestores *CoverageStoreRequester
	featuretypes   *FeatureTypeRequester
	wms            *WMSRequester
}

// TODO: having a clone of the requester for each handler is bad design. this must be refactored
func NewRequester(info *internal.GeoserverData) *Requester {
	return &Requester{
		workspaces: &WorkspaceRequester{
			info: info,
		},
		datastores: &DataStoreRequester{
			data: info,
		},
		coveragestores: &CoverageStoreRequester{
			data: info,
		},
		featuretypes: &FeatureTypeRequester{
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

func (r *Requester) FeatureTypes() *FeatureTypeRequester {
	return r.featuretypes
}

func (r *Requester) WMS() *WMSRequester {
	return r.wms
}

func (r *Requester) CoverageStores() *CoverageStoreRequester {
	return r.coveragestores
}
