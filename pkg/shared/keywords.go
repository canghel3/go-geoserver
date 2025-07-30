package shared

import (
	"encoding/json"
)

type Keywords struct {
	Keywords []string `json:"string,omitempty"`
}

func (k *Keywords) UnmarshalJSON(data []byte) error {
	type alias Keywords
	var temp alias
	//try the actual slice
	if err := json.Unmarshal(data, &temp); err == nil {
		*k = Keywords(temp)
		return nil
	}

	//try a string
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		k.Keywords = []string{s}
		return nil
	}

	//don't return any error if we couldn't figure out the keywords, just ignore it
	return nil
}
