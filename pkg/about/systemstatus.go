package about

type MetricsResponse struct {
	Metrics Metrics `json:"metrics"`
}

type Metrics struct {
	Metrics []Metric `json:"metric"`
}

type Metric struct {
	Available   bool   `json:"available"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Category    string `json:"category"`
	Identifier  string `json:"identifier"`
	Priority    int    `json:"priority"`
	Value       string `json:"value"`
}
