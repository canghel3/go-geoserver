package requester

import (
	"github.com/canghel3/go-geoserver/internal"
)

type Requester struct {
	about          *AboutRequester
	workspaces     *WorkspaceRequester
	datastores     *DataStoreRequester
	coveragestores *CoverageStoreRequester
	coverages      *CoverageRequester
	featuretypes   *FeatureTypeRequester
	wms            *WMSRequester
}

func NewRequester(data *internal.GeoserverData) *Requester {
	return &Requester{
		about: &AboutRequester{
			data: data,
		},
		workspaces: &WorkspaceRequester{
			data: data,
		},
		datastores: &DataStoreRequester{
			data: data,
		},
		coveragestores: &CoverageStoreRequester{
			data: data,
		},
		featuretypes: &FeatureTypeRequester{
			data: data,
		},
		coverages: &CoverageRequester{
			data: data,
		},
		wms: &WMSRequester{
			data: data,
		},
	}
}

func (r *Requester) About() *AboutRequester {
	return r.about
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

func (r *Requester) Coverages() *CoverageRequester {
	return r.coverages
}
