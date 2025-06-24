package actions

import (
	"encoding/json"
	"fmt"
	"github.com/canghel3/go-geoserver/internal"
	"github.com/canghel3/go-geoserver/internal/models"
	"github.com/canghel3/go-geoserver/internal/requester"
	"github.com/canghel3/go-geoserver/pkg/coverages"
)

type Coverages struct {
	store     string
	data      internal.GeoserverData
	requester *requester.Requester
}

func newCoverages(store string, data internal.GeoserverData) *Coverages {
	return &Coverages{
		store:     store,
		data:      data,
		requester: requester.NewRequester(data),
	}
}

func (c *Coverages) Publish(coverage models.Coverage) error {
	coverage.Namespace = models.NamespaceDetails{
		Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", c.data.Connection.URL, c.data.Workspace),
		Name: c.data.Workspace,
	}

	coverage.Store = models.StoreDetails{
		Class: "coverageStore",
		Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s.json", c.data.Connection.URL, c.data.Workspace, c.store),
		Name:  fmt.Sprintf("%s:%s", c.data.Workspace, c.store),
	}

	content, err := json.Marshal(models.CoverageWrapper{Coverage: coverage})
	if err != nil {
		return err
	}

	return c.requester.Coverages().Create(c.store, content)
}

func (c *Coverages) Get(name string) (*coverages.Coverage, error) {
	return c.requester.Coverages().Get(c.store, name)
}

func (c *Coverages) GetAll() (*coverages.Coverages, error) {
	return c.requester.Coverages().GetAll(c.store)
}

func (c *Coverages) Update(name string, coverage models.Coverage) error {
	coverage.Namespace = models.NamespaceDetails{
		Href: fmt.Sprintf("%s/geoserver/rest/workspaces/%s.json", c.data.Connection.URL, c.data.Workspace),
		Name: c.data.Workspace,
	}

	coverage.Store = models.StoreDetails{
		Class: "coverageStore",
		Href:  fmt.Sprintf("%s/geoserver/rest/workspaces/%s/coveragestores/%s.json", c.data.Connection.URL, c.data.Workspace, c.store),
		Name:  fmt.Sprintf("%s:%s", c.data.Workspace, c.store),
	}

	content, err := json.Marshal(models.CoverageWrapper{Coverage: coverage})
	if err != nil {
		return err
	}

	return c.requester.Coverages().Update(c.store, name, content)
}

func (c *Coverages) Delete(name string, recurse bool) error {
	return c.requester.Coverages().Delete(c.store, name, recurse)
}
