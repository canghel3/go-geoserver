[![Go Report Card](https://goreportcard.com/badge/github.com/canghel3/go-geoserver)](https://goreportcard.com/report/github.com/canghel3/go-geoserver)
[![codecov](https://codecov.io/gh/canghel3/go-geoserver/graph/badge.svg?token=OTMJR61Q1H)](https://codecov.io/gh/canghel3/go-geoserver)

# GeoServer Go Client

A GoLang library made to simplify interactions with [GeoServer](https://geoserver.org/), providing an idiomatic way to
talk to GeoServer’s REST API from your Go applications.

## Features

**Manage:**

- Workspaces
- Vector Data Sources (Data Store)
- Raster Data Sources (Coverage Store)
- Feature Types
- Coverages

## Available Data Sources

### Vector

| Format                  | Status |
|-------------------------|--------|
| CSV                     | ❌      |
| Directory of shapefiles | ✅      |
| GeoPackage              | ✅      |
| PostGIS                 | ✅      |
| Shapefile               | ✅      |
| WebFeatureService       | ✅      |

### Raster

| Format              | Status |
|---------------------|--------|
| AIG                 | ❌      |
| ArcGrid             | ❌      |
| DTED                | ❌      |
| EHdr                | ✅      |
| ENVIHdr             | ✅      |
| ERDASImg            | ✅      |
| GeoPackage (mosaic) | ❌      |
| GeoTIFF             | ✅      |
| ImageMosaic         | ❌      |
| ImagePyramid        | ❌      |
| NITF                | ✅      |
| RPFTOC              | ❌      |
| RST                 | ✅      |
| SRP                 | ❌      |
| VRT                 | ✅      |
| WorldImage          | ❌      |

---

## Tested GeoServer Versions

| Version | Status |
|---------|--------|
| 2.27.1  | ✅      |
| 2.22.2  | ✅      |

## Work In Progress

- Caching
- Layer Groups
- Styles
- WMS, WFS, WCS, WMTS

## Installation

```bash
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
