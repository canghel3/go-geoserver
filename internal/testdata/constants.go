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
	CoverageStoreEHdr    = "EHDR"
	CoverageStoreENVIHdr = "ENVIHDR"

	FeatureTypePostgis              = "init"
	FeatureTypePostgisNativeName    = "init"
	FeatureTypeGeoPackage           = "buildings"
	FeatureTypeGeoPackageNativeName = "bld_fts_buildingpart"

	CoverageGeoTiffName       = "sample"
	CoverageGeoTiffNativeName = "sample"

	PostgisHost     = "postgis"
	PostgisPort     = "5432"
	PostgisUsername = "geoserver"
	PostgisPassword = "geoserver"
	PostgisDb       = "vectors"
	PostgisSsl      = "disable"

	FileShapefile  = "ne_110m_coastline.shp"
	FileGeoPackage = "bld_fts_buildingpart.gpkg"
	FileGeoTiff    = "geotiff/sample.tif"
	FileEHdr       = "ehdr/output.bil"
	FileENVIHdr    = "envihdr/output.dat"

	DirENVIHdr = "envihdr"
	DirEHdr    = "ehdr"
)
