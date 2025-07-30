package shared

import (
	"encoding/json"
	"strings"
)

type Style struct {
	Name string `json:"name"`
	Link string `json:"href"`
}

func (s *Style) UnmarshalJSON(data []byte) error {
	type alias Style
	var temp alias
	if err := json.Unmarshal(data, &temp); err == nil {
		*s = Style(temp)
		return nil
	}

	var ss string
	if err := json.Unmarshal(data, &ss); err == nil {
		s.Name = ss
		return nil
	}

	return nil
}

func (s *Style) MarshalJSON() ([]byte, error) {
	name := strings.TrimSpace(s.Name)
	link := strings.TrimSpace(s.Link)

	type alias Style
	var temp alias
	if name != "" && link != "" {
		temp.Name = name
		temp.Link = link
		return json.Marshal(temp)
	}

	return json.Marshal("")
}

type LanguageVersion struct {
	Version string `json:"version"`
}
