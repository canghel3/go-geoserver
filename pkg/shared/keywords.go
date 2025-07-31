package shared

import (
	"encoding/json"
)

type Keywords struct {
	Keywords []string `json:"string,omitempty"`
}

func (k *Keywords) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	m = make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	var s string
	if err := json.Unmarshal(m["string"], &s); err == nil {
		k.Keywords = []string{s}
		return nil
	}

	type alias Keywords
	var temp alias
	//try the actual slice
	if err := json.Unmarshal(m["string"], &temp.Keywords); err == nil {
		k.Keywords = temp.Keywords
		return nil
	}

	//don't return any error if we couldn't figure out the keywords, just ignore it
	return nil
}
