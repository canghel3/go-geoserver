### This library is meant to be a wrapper for GeoServer REST API

```go
go get github.com/canghel3/go-geoserver
```

Define a geoserver service instance and use it to manage GeoServer resources (CRUD).

```go
package main

import "github.com/canghel3/go-geoserver/service"

func main() {
	geosvc := service.NewGeoserverService("http://localhost:8080", "admin", "geoserver")
	
	geosvc.CreateWorkspace("my first workspace")
	geosvc.CreateCoverageStore("my first workspace", "coveragestore", "/opt/geoserver/data/myfile.tif", "GeoTIFF")
	geosvc.CreateCoverage("coverage", "coveragestore", "layer", "EPSG:3857", [4]float64{0, 0, 0, 0})
}
```