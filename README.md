# Geoserver Go Client

Making interactions with [GeoServer](https://geoserver.org/) frustration-less in Go.

This client provides a simple and idiomatic way to talk to GeoServer’s REST API from your Go applications.

## ✨ Features

- Support for managing:
    - Workspaces
    - Vector Data Sources
      - CSV
      - Directory of spatial files (shapefiles)
      - GeoPackage
      - PostGIS
      - Shapefile
      - Web Feature Server
    - Raster Data Sources
      - AIG
      - ArcGrid
      - DTED
      - EHdr
      - ENVIHdr
      - ERDASImg
      - GeoPackage (mosaic)
      - GeoTIFF
      - ImageMosaic
      - ImagePyramid
      - NITF
      - RPFTOC
      - RST
      - SRP
      - VRT
      - WorldImage
    - Feature Types

## 🛠️ Work In Progress - in order of priority

Support:

- Coverages
- Layer Groups
- Styles
- Caching
- WMS, WFS

## 📦 Installation

```go
go get github.com/canghel3/go-geoserver
```

## 🧪 Examples

### Workspaces

[`examples/workspace.go`](./examples/workspace.go)
