package actions

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/internal/validator"
	"github.com/canghel3/go-geoserver/pkg/coveragestores"
	"github.com/canghel3/go-geoserver/pkg/options"
)

type CoverageStoreType string

const (
	AIG              CoverageStoreType = "AIG"
	ArcGrid          CoverageStoreType = "ArcGrid"
	DTED             CoverageStoreType = "DTED"
	EHdr             CoverageStoreType = "EHdr"
	ENVIHdr          CoverageStoreType = "ENVIHdr"
	ERDASImg         CoverageStoreType = "ERDASImg"
	GeoPackageMosaic CoverageStoreType = "GeoPackage (mosaic)"
	GeoTIFF          CoverageStoreType = "GeoTIFF"
	ImageMosaic      CoverageStoreType = "ImageMosaic"
	ImagePyramid     CoverageStoreType = "ImagePyramid"
	NITF             CoverageStoreType = "NITF"
	RPFTOC           CoverageStoreType = "RPFTOC"
	RST              CoverageStoreType = "RST"
	SRP              CoverageStoreType = "SRP"
	VRT              CoverageStoreType = "VRT"
	WorldImage       CoverageStoreType = "WorldImage"
)

func newCoverageStoreHandler(info *internal.GeoserverData) *CoverageStores {
	r := requester.NewRequester(info)
	return &CoverageStores{
		info:      info,
		requester: r,
	}
}

type CoverageStoreList struct {
	options   *internal.CoveragestoreOptions
	requester *requester.Requester
	data      *internal.GeoserverData
}

type CoverageStores struct {
	info      *internal.GeoserverData
	requester *requester.Requester
}

// Reset the caches related to the specified coveragestore.
func (cs *CoverageStores) Reset(name string) error {
	return cs.requester.CoverageStores().Reset(name)
}

// Use a specific coverage store
func (cs *CoverageStores) Use(name string) *CoverageStore {
	return &CoverageStore{
		name: name,
		info: cs.info.Clone(),
	}
}

type CoverageStore struct {
	name string
	info *internal.GeoserverData
}

func (cs *CoverageStores) Create(options ...options.CoveragestoreOptionFunc) CoverageStoreList {
	csl := CoverageStoreList{
		requester: cs.requester,
		options:   &internal.CoveragestoreOptions{},
		data:      cs.info.Clone(),
	}

	for _, option := range options {
		option(csl.options)
	}

	return csl
}

func (cs *CoverageStores) Get(name string) (*coveragestores.CoverageStoreRetrieval, error) {
	return cs.requester.CoverageStores().Get(name)
}

func (cs *CoverageStores) Delete(name string, recurse bool) error {
	return cs.requester.CoverageStores().Delete(name, recurse)
}

func (csl CoverageStoreList) AIG(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(AIG),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) ArcGrid(name string, filepath string, options ...options.ArcGridOptions) error {
	err := validator.CoverageStore.ArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(ArcGrid),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) DTED(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(DTED),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) EHdr(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(EHdr),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) ENVIHdr(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(ENVIHdr),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) ERDASImg(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(ERDASImg),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) GeoPackage(name string, filepath string, options ...options.GeoPackageRasterOptions) error {
	err := internal.ValidateGeoPackage(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(GeoPackageMosaic),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) GeoTIFF(name string, filepath string, options ...options.GeoTIFFOptions) error {
	err := internal.ValidateGeoTIFF(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(GeoTIFF),
			Default:     csl.options.Default,
			Enabled:     csl.options.Enabled,
			URL:         filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) ImageMosaic(name string, dirpath string, options ...options.ImageMosaicOptions) error {
	err := internal.ValidateImageMosaic(dirpath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(ImageMosaic),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     dirpath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) ImagePyramid(name string, dirpath string, options ...options.ImageMosaicOptions) error {
	err := internal.ValidateImageMosaic(dirpath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(ImagePyramid),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     dirpath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) NITF(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(NITF),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) RPFTOC(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(RPFTOC),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) RST(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(RST),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) SRP(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(SRP),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) VRT(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(VRT),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) WorldImage(name string, filepath string, options ...options.WorldImageOptions) error {
	err := internal.ValidateWorldImage(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        string(WorldImage),
			Workspace: struct {
				Name string `json:"name"`
				Link string `json:"link"`
			}{Name: csl.data.Workspace, Link: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", csl.data.Connection.URL, csl.data.Workspace)},
			Default: csl.options.Default,
			Enabled: csl.options.Enabled,
			URL:     filepath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}
