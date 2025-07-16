package formats

type CoverageStoreFormat string

const (
	AIG              CoverageStoreFormat = "AIG"
	ArcGrid          CoverageStoreFormat = "ArcGrid"
	DTED             CoverageStoreFormat = "DTED"
	EHdr             CoverageStoreFormat = "EHdr"
	ENVIHdr          CoverageStoreFormat = "ENVIHdr"
	ERDASImg         CoverageStoreFormat = "ERDASImg"
	GeoPackageMosaic CoverageStoreFormat = "GeoPackage (mosaic)"
	GeoTIFF          CoverageStoreFormat = "GeoTIFF"
	ImageMosaic      CoverageStoreFormat = "ImageMosaic"
	ImagePyramid     CoverageStoreFormat = "ImagePyramid"
	NITF             CoverageStoreFormat = "NITF"
	RPFTOC           CoverageStoreFormat = "RPFTOC"
	RST              CoverageStoreFormat = "RST"
	SRP              CoverageStoreFormat = "SRP"
	VRT              CoverageStoreFormat = "VRT"
	WorldImage       CoverageStoreFormat = "WorldImage"
)
