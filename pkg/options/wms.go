package options

import (
	"net/url"
	"strings"
)

var GetMap GetMapOptionGenerator

type GetMapOptionGenerator struct{}

type GetMapOption func(values *url.Values)

func (wog GetMapOptionGenerator) Styles(styles []string) GetMapOption {
	return func(values *url.Values) {
		values.Set("styles", strings.Join(styles, ","))
	}
}
