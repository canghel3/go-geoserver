package style

import (
	"github.com/canghel3/go-geoserver/internal/misc"
	"github.com/canghel3/go-geoserver/pkg/workspace"
)

type StyleWrapper struct {
	Style StyleDetails `json:"style"`
}

type StyleDetails struct {
	Name     string `json:"name"`
	Filename string `json:"filename,omitempty"`
}

type GetStyleWrapper struct {
	Style GetStyleDetails `json:"style"`
}

type GetStyleDetails struct {
	Name            any                          `json:"name"`
	Content         string                       `json:"content"`
	Workspace       *workspace.WorkspaceCreation `json:"workspace,omitempty"`
	Format          string                       `json:"format,omitempty"`
	LanguageVersion *misc.LanguageVersion        `json:"languageVersion,omitempty"`
	Filename        any                          `json:"filename"`
	DateCreate      string                       `json:"dateCreate,omitempty"`
	DateModified    string                       `json:"dateModified,omitempty"`
}
