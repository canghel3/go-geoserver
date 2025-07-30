package shared

import "encoding/json"

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

type LanguageVersion struct {
	Version string `json:"version"`
}
