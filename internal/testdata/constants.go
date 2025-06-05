package testdata

const (
	GeoserverUrl      = "http://localhost:1112"
	GeoserverUsername = "admin"
	GeoserverPassword = "geoserver"
	GeoserverDataDir  = "/tmp/data"

	Workspace            = "PLAYGROUND"
	InvalidWorkspaceName = "SO_!@#MEINVALIDNAME"

	DatastorePostgis    = "POSTGIS"
	DatastoreShapefile  = "SHAPEFILE"
	DatastoreGeoPackage = "GEOPACKAGE"

	CoverageStoreGeoTiff = "GEOTIFF"

	FeatureTypePostgis              = "init"
	FeatureTypePostgisNativeName    = "init"
	FeatureTypeGeoPackage           = "buildings"
	FeatureTypeGeoPackageNativeName = "bld_fts_buildingpart"
	FeatureTypeTitle                = "sample"

	CoverageName       = "sample"
	CoverageNativeName = "sample"

	PostgisHost     = "postgis"
	PostgisPort     = "5432"
	PostgisUsername = "geoserver"
	PostgisPassword = "geoserver"
	PostgisDb       = "vectors"
	PostgisSsl      = "disable"

	FileShapefile  = "ne_110m_coastline.shp"
	FileGeoPackage = "bld_fts_buildingpart.gpkg"

	FileGeoTiff = "sample.tif"
)
