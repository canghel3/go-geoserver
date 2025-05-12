package handler

import (
	"encoding/json"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/coveragestores"
	"github.com/canghel3/go-geoserver/pkg/options"
)

func NewCoverageStoreHandler(info *internal.GeoserverData) *CoverageStores {
	r := requester.NewRequester(info)
	return &CoverageStores{
		info:      info,
		requester: r,
	}
}

type CoverageStoreList struct {
	options   *internal.CoveragestoreOptions
	requester *requester.Requester
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

func (csl CoverageStoreList) GeoTIFF(name string, filepath string, options ...options.GeoTIFFOptions) error {
	err := internal.ValidateGeoTIFF(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        "GeoTIFF",
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

func (csl CoverageStoreList) WorldImage(name string, filepath string, options ...options.WorldImageOptions) error {
	err := internal.ValidateWorldImage(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        "WorldImage",
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
			Type:        "ImageMosaic",
			Default:     csl.options.Default,
			Enabled:     csl.options.Enabled,
			URL:         dirpath,
		},
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return csl.requester.CoverageStores().Create(content)
}

func (csl CoverageStoreList) ArcGrid(name string, filepath string, options ...options.ArcGridOptions) error {
	err := internal.ValidateArcGrid(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        "ArcGrid",
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

func (csl CoverageStoreList) GeoPackageRaster(name string, filepath string, options ...options.GeoPackageRasterOptions) error {
	err := internal.ValidateGeoPackage(filepath)
	if err != nil {
		return err
	}

	data := coveragestores.GenericCoverageStoreCreationWrapper{
		CoverageStore: coveragestores.GenericCoverageStoreCreationModel{
			Name:        name,
			Description: csl.options.Description,
			Type:        "GeoPackage (mosaic)",
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
