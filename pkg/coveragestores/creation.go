package coveragestores

type GenericCoverageStoreCreationWrapper struct {
	CoverageStore GenericCoverageStoreCreationModel `json:"coverageStore"`
}

type GenericCoverageStoreCreationModel struct {
	Name                       string               `json:"name"`
	Description                string               `json:"description"`
	Type                       string               `json:"type"`
	Enabled                    bool                 `json:"enabled"`
	URL                        string               `json:"url"`
	DisableOnConnectionFailure bool                 `json:"disableOnConnFailure"`
	ConnectionParameters       ConnectionParameters `json:"connectionParameters,omitempty"`
}

type ConnectionParameters struct {
	Entry []Entry `json:"entry"`
}

func (cp *ConnectionParameters) Get(key string) (value string, ok bool) {
	for _, entry := range cp.Entry {
		if entry.Key == key {
			return entry.Value, true
		}
	}
	return value, false
}

type Entry struct {
	Key   string `json:"@key"`
	Value string `json:"$"`
}