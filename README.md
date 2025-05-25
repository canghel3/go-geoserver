# Geoserver Go Client

Making interactions with [GeoServer](https://geoserver.org/) frustration-less in Go.

This client provides a simple and idiomatic way to talk to GeoServer’s REST API from your Go applications.

## ✨ Features

- Support for managing:
    - Workspaces
    - Vector Data Sources
    - Raster Data Sources
    - Feature Types
    - Coverages

Tested Vector Data Sources:
- CSV
- GeoPackage
- PostGIS
- Shapefile

Tested Raster Data Sources:
- GeoTIFF

## 🛠️ Work In Progress - in order of priority

Support:

- WMS, WFS, WCS, WMTS 
- Styles
- Layer Groups
- Caching

## 📦 Installation

```go
go get github.com/canghel3/go-geoserver
```

## 🧪 Examples

### Client Initialization

[`examples/client.go`](./pkg/client/client_test.go)

### Workspaces

[`examples/workspaces.go`](./pkg/client/workspace_test.go)

### DataStores

[`examples/datastores.go`](./pkg/actions/datastore_test.go)

### CoverageStores and Coverages

[`examples/coverages.go`](./pkg/client/coverage_test.go)

### Feature Types

[`examples/featuretypes.go`](./pkg/client/featuretype_test.go)
