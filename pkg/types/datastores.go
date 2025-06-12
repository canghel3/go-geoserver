package types

type DataStoreType string

const (
	PostGIS    DataStoreType = "postgis"
	Shapefile  DataStoreType = "shapefile"
	GeoPackage DataStoreType = "geopkg"
	CSV        DataStoreType = "csv"
)
