package about

type StatusResponse struct {
	Status Status `json:"statuss"` //this is a typo in GeoServer v2.22.2
}

type Status struct {
	Modules []ModuleStatus `json:"status"`
}

// ModuleStatus represents each module/component status entry
type ModuleStatus struct {
	Module      string `json:"module"`
	Name        string `json:"name"`
	Component   string `json:"component,omitempty"`
	Version     string `json:"version,omitempty"`
	Message     string `json:"message,omitempty"`
	IsEnabled   bool   `json:"isEnabled"`
	IsAvailable bool   `json:"isAvailable"`
}
