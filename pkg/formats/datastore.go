package formats

type DataStoreFormat string

const (
	PostGIS           DataStoreFormat = "postgis"
	Shapefile         DataStoreFormat = "shapefile"
	GeoPackage        DataStoreFormat = "geopkg"
	DirOfShapefiles   DataStoreFormat = "shape"
	CSV               DataStoreFormat = "csv"
	WebFeatureService DataStoreFormat = "wfs"
)
