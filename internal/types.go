package internal

import (
	"github.com/canghel3/go-geoserver/pkg/models/datastores"
)

type ConnectionParams map[string]string

func (params ConnectionParams) ToDatastoreEntries() []datastores.Entry {
	entries := make([]datastores.Entry, 0)
	for k, v := range params {
		entries = append(entries, datastores.Entry{Key: k, Value: v})
	}

	return entries
}
