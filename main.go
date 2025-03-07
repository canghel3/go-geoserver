package main

import (
	"github.com/canghel3/go-geoserver/workspace"
)

func main() {
	w := workspace.NewService()
	v := w.Use("sample").Vectors()
	v.Stores().PostGIS().Create()
	v.Store("s")
}
