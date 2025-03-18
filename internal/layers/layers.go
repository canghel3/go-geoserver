package layers

type LayerWrapper struct {
	Layer Layer `json:"layer"`
}

type Layer struct {
	Name         string       `json:"name"`
	Path         string       `json:"path"`
	Type         string       `json:"type"`
	DefaultStyle *Style       `json:"defaultStyle,omitempty"`
	Styles       *Styles      `json:"styles,omitempty"`
	Resource     *Resource    `json:"resource,omitempty"`
	Attribution  *Attribution `json:"attribution,omitempty"`
}

type Style struct {
	Name      string `json:"name"`
	Workspace string `json:"workspace,omitempty"`
	Href      string `json:"href"`
}

type Styles struct {
	Class string `json:"@class"`
	Style any    `json:"style"`
}

type Resource struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
	Href  string `json:"href"`
}

type Attribution struct {
	LogoWidth  int `json:"logoWidth"`
	LogoHeight int `json:"logoHeight"`
}
