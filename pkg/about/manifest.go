package about

type ManifestResponse struct {
	Manifest Manifest `json:"about"`
}

type Manifest struct {
	Resources []Resource `json:"resource"`
}

// Resource represents each software package/library entry
type Resource struct {
	Name                               string `json:"@name"`
	ArchiverVersion                    string `json:"Archiver-Version,omitempty"`
	BundleLicense                      string `json:"Bundle-License,omitempty"`
	SpecificationVersion               any    `json:"Specification-Version,omitempty"` // Can be string or int
	BndLastModified                    int64  `json:"Bnd-LastModified,omitempty"`
	BundleName                         string `json:"Bundle-Name,omitempty"`
	BuildJdk                           string `json:"Build-Jdk,omitempty"`
	BundleDescription                  string `json:"Bundle-Description,omitempty"`
	URL                                string `json:"URL,omitempty"`
	BundleSymbolicName                 string `json:"Bundle-SymbolicName,omitempty"`
	BuiltBy                            string `json:"Built-By,omitempty"`
	RequireCapability                  string `json:"Require-Capability,omitempty"`
	Tool                               string `json:"Tool,omitempty"`
	ExtensionName                      string `json:"Extension-Name,omitempty"`
	ImplementationTitle                string `json:"Implementation-Title,omitempty"`
	ImplementationBuildId              string `json:"Implementation-Build-Id,omitempty"`
	ImplementationVersion              string `json:"Implementation-Version,omitempty"`
	ManifestVersion                    any    `json:"Manifest-Version,omitempty"` // Can be string or int
	CreatedBy                          string `json:"Created-By,omitempty"`
	ImplementationVendorId             string `json:"Implementation-Vendor-Id,omitempty"`
	BundleDocURL                       string `json:"Bundle-DocURL,omitempty"`
	BundleVendor                       string `json:"Bundle-Vendor,omitempty"`
	ImplementationVendor               string `json:"Implementation-Vendor,omitempty"`
	BundleManifestVersion              int    `json:"Bundle-ManifestVersion,omitempty"`
	BundleVersion                      string `json:"Bundle-Version,omitempty"`
	ImplementationURL                  string `json:"Implementation-URL,omitempty"`
	SpecificationTitle                 string `json:"Specification-Title,omitempty"`
	BuildTime                          string `json:"Build-Time,omitempty"`
	BuildJdkSpec                       int    `json:"Build-Jdk-Spec,omitempty"`
	GitCommitId                        string `json:"Git-Commit-Id,omitempty"`
	XCompileSourceJDK                  int    `json:"X-Compile-Source-JDK,omitempty"`
	XCompileTargetJDK                  int    `json:"X-Compile-Target-JDK,omitempty"`
	ProvideCapability                  string `json:"Provide-Capability,omitempty"`
	SpecificationVendor                string `json:"Specification-Vendor,omitempty"`
	AntVersion                         string `json:"Ant-Version,omitempty"`
	BundleRequiredExecutionEnvironment string `json:"Bundle-RequiredExecutionEnvironment,omitempty"`
	ModuleRequires                     string `json:"Module-Requires,omitempty"`
}
