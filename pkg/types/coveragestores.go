package types

type CoverageStoreType string

const (
	AIG              CoverageStoreType = "AIG"
	ArcGrid          CoverageStoreType = "ArcGrid"
	DTED             CoverageStoreType = "DTED"
	EHdr             CoverageStoreType = "EHdr"
	ENVIHdr          CoverageStoreType = "ENVIHdr"
	ERDASImg         CoverageStoreType = "ERDASImg"
	GeoPackageMosaic CoverageStoreType = "GeoPackage (mosaic)"
	GeoTIFF          CoverageStoreType = "GeoTIFF"
	ImageMosaic      CoverageStoreType = "ImageMosaic"
	ImagePyramid     CoverageStoreType = "ImagePyramid"
	NITF             CoverageStoreType = "NITF"
	RPFTOC           CoverageStoreType = "RPFTOC"
	RST              CoverageStoreType = "RST"
	SRP              CoverageStoreType = "SRP"
	VRT              CoverageStoreType = "VRT"
	WorldImage       CoverageStoreType = "WorldImage"
)
