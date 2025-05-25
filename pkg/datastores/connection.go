package datastores

type ConnectionParams map[string]string

func (params ConnectionParams) ToDatastoreEntries() []Entry {
	entries := make([]Entry, 0)
	for k, v := range params {
		entries = append(entries, Entry{Key: k, Value: v})
	}

	return entries
}
