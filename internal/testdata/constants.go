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

	PostgisHost     = "postgis"
	PostgisPort     = "5432"
	PostgisUsername = "geoserver"
	PostgisPassword = "geoserver"
	PostgisDb       = "vectors"
	PostgisSsl      = "disable"

	FeatureTypePostgis              = "init"
	FeatureTypePostgisNativeName    = "init"
	FeatureTypeGeoPackage           = "buildings"
	FeatureTypeGeoPackageNativeName = "bld_fts_buildingpart"

	CoverageStoreGeoTiff    = "GEOTIFF"
	CoverageStoreEHdr       = "EHDR"
	CoverageStoreENVIHdr    = "ENVIHDR"
	CoverageStoreGeoPackage = "GEOPACKAGE"
	CoverageStoreNITF       = "NITF"
	CoverageStoreRST        = "RST"
	CoverageStoreVRT        = "VRT"

	CoverageGeoTiffName       = "sample"
	CoverageGeoTiffNativeName = "sample"

	FileShapefile        = "ne_110m_coastline.shp"
	FileGeoPackage       = "bld_fts_buildingpart.gpkg"
	FileGeoTiff          = "geotiff/sample.tif"
	FileEHdr             = "ehdr/output.bil"
	FileENVIHdr          = "envihdr/output.dat"
	FileGeoPackageRaster = "gpkg/output.gpkg"
	FileNITF             = "nitf/output.ntf"
	FileRST              = "rst/output.rst"
	FileVRT              = "vrt/output.vrt"

	DirGeoTiff          = "geotiff"
	DirENVIHdr          = "envihdr"
	DirEHdr             = "ehdr"
	DirGeoPackageRaster = "gpkg"
	DirNITF             = "nitf"
	DirRST              = "rst"
	DirVRT              = "vrt"
)
