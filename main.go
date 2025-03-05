package main

import (
	"github.com/canghel3/go-geoserver/workspace"
)

func main() {
	w := workspace.NewService()
	vectors := w.
		vectors.Stores().Get("some")
}
