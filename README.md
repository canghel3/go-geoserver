[![Go Report Card](https://goreportcard.com/badge/github.com/canghel3/go-geoserver)](https://goreportcard.com/report/github.com/canghel3/go-geoserver)
[![codecov](https://codecov.io/gh/canghel3/go-geoserver/graph/badge.svg?token=OTMJR61Q1H)](https://codecov.io/gh/canghel3/go-geoserver)

# GeoServer Go Client

A GoLang library made to simplify interactions with [GeoServer](https://geoserver.org/), providing an idiomatic way to talk to GeoServer’s REST API from your Go applications.

## Features

- Support for managing:
    - Workspaces
    - Vector Data Sources
    - Raster Data Sources
    - Feature Types
    - Coverages

Tested Vector Data Sources:
- GeoPackage
- PostGIS
- Shapefile
- Directory of Shapefiles

Tested Raster Data Sources:
- EHdr
- ENVIHdr
- GeoTIFF
- NITF
- RST
- VRT

## Work In Progress - in order of priority

- Layer Groups
- Caching
- Styles
- WMS, WFS, WCS, WMTS

## Installation

```go
go get github.com/canghel3/go-geoserver
```

## Usage Examples

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
