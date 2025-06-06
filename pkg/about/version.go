package about

type VersionResponse struct {
	About Version `json:"about"`
}

type Version struct {
	Resources []VersionResource `json:"resource"`
}

type VersionResource struct {
	Name           string `json:"@name"`
	BuildTimestamp string `json:"Build-Timestamp"`
	Version        string `json:"Version"`
	GitRevision    string `json:"Git-Revision"`
}
