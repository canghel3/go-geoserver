package testdata

const (
	GeoserverUrl      = "http://localhost:1112"
	GeoserverUsername = "admin"
	GeoserverPassword = "geoserver"
	GeoserverDataDir  = "/opt/geoserver-for-tests/data"

	Workspace            = "PLAYGROUND"
	InvalidWorkspaceName = "SO_!@#MEINVALIDNAME"

	DatastorePostgis    = "POSTGIS"
	DatastoreShapefile  = "SHAPEFILE"
	DatastoreGeoPackage = "GEOPACKAGE"

	FeatureTypeName       = "init"
	FeatureTypeNativeName = "init"
	FeatureTypeTitle      = "sample"

	PostgisHost     = "postgis"
	PostgisPort     = "5432"
	PostgisUsername = "geoserver"
	PostgisPassword = "geoserver"
	PostgisDb       = "vectors"
	PostgisSsl      = "disable"

	Shapefile  = "ne_110m_coastline.shp"
	GeoPackage = "bld_fts_buildingpart.gpkg"
)
