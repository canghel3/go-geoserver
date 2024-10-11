package datastore

import (
	"github.com/canghel3/go-geoserver/models/workspace"
)

//CREATION MODELS

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

//RETRIEVAL MODELS

type DataStoreRetrievalWrapper struct {
	DataStore DataStoreRetrieval `json:"dataStore"`
}

type DataStoreRetrieval struct {
	Name                 string                   `json:"name,omitempty"`
	Enabled              bool                     `json:"enabled,omitempty"`
	Workspace            workspace.MultiWorkspace `json:"workspace,omitempty"`
	ConnectionParameters ConnectionParameters     `json:"connectionParameters"`
	Default              bool                     `json:"_default,omitempty"`
	FeatureTypes         string                   `json:"featureTypes,omitempty"`
}
