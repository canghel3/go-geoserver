package wms

import "html/template"

type OpenLayersTemplate struct {
	Template *template.Template
	RawHTML  []byte
}
