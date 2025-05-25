package datastores

type GenericDataStoreCreationWrapper struct {
	DataStore GenericDataStoreCreationModel `json:"dataStore"`
}

type GenericDataStoreCreationModel struct {
	Name                       string               `json:"name"`
	Description                string               `json:"description"`
	DisableOnConnectionFailure bool                 `json:"disableOnConnFailure"`
	ConnectionParameters       ConnectionParameters `json:"connectionParameters"`
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
