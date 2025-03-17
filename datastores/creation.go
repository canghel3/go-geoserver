package datastores

type GenericDataStoreCreationWrapper struct {
	DataStore GenericDataStoreCreationModel `json:"dataStore"`
}

type GenericDataStoreCreationModel struct {
	Name                 string               `json:"name"`
	ConnectionParameters ConnectionParameters `json:"connectionParameters"`
}

type ConnectionParameters struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	Key   string `json:"@key"`
	Value string `json:"$"`
}
