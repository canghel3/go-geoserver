package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/models/coverages"
)

type Coverages struct {
	store     string
	data      *internal.GeoserverData
	requester *requester.Requester
}

func newCoverages(store string, data *internal.GeoserverData) *Coverages {
	return &Coverages{
		store:     store,
		data:      data,
		requester: requester.NewRequester(data),
	}
}

func (c *Coverages) Publish(coverage coverages.Coverage) error {
	completeCoverage := coverages.Coverage{
		Abstract:                   coverage.Abstract,
		DefaultInterpolationMethod: coverage.DefaultInterpolationMethod,
		Description:                coverage.Description,
		Dimensions:                 coverage.Dimensions,
		Enabled:                    coverage.Enabled,
		Grid:                       coverage.Grid,
		InterpolationMethods:       coverage.InterpolationMethods,
		ProjectionPolicy:           coverage.ProjectionPolicy,
		Keywords:                   coverage.Keywords,
		LatLonBoundingBox:          coverage.LatLonBoundingBox,
		Metadata:                   coverage.Metadata,
		Name:                       coverage.Name,
		Namespace: coverages.NamespaceDetails{
			Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", c.data.Connection.URL, c.data.Workspace),
			Name: c.data.Workspace,
		},
		NativeBoundingBox: coverage.NativeBoundingBox,
		NativeCRS:         coverage.NativeCRS,
		NativeFormat:      coverage.NativeFormat,
		NativeName:        coverage.NativeName,
		RequestSRS:        coverage.RequestSRS,
		ResponseSRS:       coverage.ResponseSRS,
		Srs:               coverage.Srs,
		Store: coverages.StoreDetails{
			Class: "coverageStore",
			Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s.json", c.data.Connection.URL, c.data.Workspace, c.store),
			Name:  fmt.Sprintf("%s:%s", c.data.Workspace, c.store),
		},
		SupportedFormats: coverage.SupportedFormats,
		Title:            coverage.Title,
	}

	content, err := json.Marshal(coverages.CoverageWrapper{Coverage: completeCoverage})
	if err != nil {
		return err
	}

	return c.requester.Coverages().Create(c.store, content)
}

func (c *Coverages) Get(name string) (*coverages.Coverage, error) {
	return c.requester.Coverages().Get(c.store, name)
}

func (c *Coverages) GetAll() ([]coverages.Coverage, error) {
	return c.requester.Coverages().GetAll(c.store)
}

func (c *Coverages) Update() error {
	return errors.New("not implemented")
}

func (c *Coverages) Delete(name string, recurse bool) error {
	return c.requester.Coverages().Delete(c.store, name, recurse)
}
