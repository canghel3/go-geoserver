package testdata

const (
	GeoserverUrl      = "http://localhost:1112"
	GeoserverUsername = "admin"
	GeoserverPassword = "geoserver"
	GeoserverDataDir  = "/tmp/data"

	Workspace   = "PLAYGROUND"
	InvalidName = "SO_!@#MEINVALIDNAME"

	DatastorePostgis           = "POSTGIS"
	DatastoreShapefile         = "SHAPEFILE"
	DatastoreGeoPackage        = "GEOPACKAGE"
	DatastoreDirOfShapefiles   = "SHAPEFILES"
	DatastoreCSV               = "CSV"
	DatastoreWebFeatureService = "WEBFEATURESERVICE"

	PostgisHost     = "postgis"
	PostgisPort     = "5432"
	PostgisUsername = "geoserver"
	PostgisPassword = "geoserver"
	PostgisDb       = "vectors"
	PostgisSsl      = "disable"

	DatastoreWFSUrl = "http://geoserver:8080/geoserver/wfs?service=wfs&version=1.3.0&request=GetCapabilities"

	FeatureTypePostgis              = "init"
	FeatureTypePostgisNativeName    = "init"
	FeatureTypeGeoPackage           = "buildings"
	FeatureTypeGeoPackageNativeName = "bld_fts_buildingpart"

	CoverageStoreGeoTiff    = "GEOTIFF"
	CoverageStoreEHdr       = "EHDR"
	CoverageStoreENVIHdr    = "ENVIHDR"
	CoverageStoreERDASImg   = "ERDASIMG"
	CoverageStoreGeoPackage = "GEOPACKAGE"
	CoverageStoreNITF       = "NITF"
	CoverageStoreRST        = "RST"
	CoverageStoreVRT        = "VRT"

	CoverageGeoTiffName       = "sample"
	CoverageGeoTiffNativeName = "sample"

	DirGeoTiff          = "geotiff"
	DirENVIHdr          = "envihdr"
	DirEHdr             = "ehdr"
	DirERDASImg         = "erdasimg"
	DirGeoPackageRaster = "gpkg"
	DirNITF             = "nitf"
	DirRST              = "rst"
	DirVRT              = "vrt"
	DirShapefiles       = "shps"

	FileShapefile        = "shp/ne_110m_coastline.shp"
	FileGeoPackage       = "gpkg/bld_fts_buildingpart.gpkg"
	FileGeoTiff          = "geotiff/sample.tif"
	FileEHdr             = DirEHdr + "/output.bil"
	FileENVIHdr          = "envihdr/output.dat"
	FileERDASImg         = DirERDASImg + "/output.img"
	FileGeoPackageRaster = "gpkg/output.gpkg"
	FileNITF             = "nitf/output.ntf"
	FileRST              = "rst/output.rst"
	FileVRT              = "vrt/output.vrt"
	FileCSVLatLon        = "csv/latlon.csv"
	FileCSVWkt           = "csv/wkt.csv"
)
