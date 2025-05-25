package internal

//
//import (
//	"github.com/canghel3/go-geoserver/internal/misc"
//	"strings"
//)
//
//type Option func() map[string]any
//
//func KeywordsOption(keywords []string) Option {
//	return func() map[string]any {
//		return map[string]any{"keywords": &misc.Keywords{Keywords: keywords}}
//	}
//}
//
//func TitleOption(title string) Option {
//	return func() map[string]any {
//		return map[string]any{"title": title}
//	}
//}
//
///*
//Available options: NONE=KEEP_NATIVE, FORCE_DECLARED, REPROJECT_TO_DECLARED
//*/
//func ProjectionPolicyOption(prjPolicy string) Option {
//	return func() map[string]any {
//		return map[string]any{"projectionPolicy": prjPolicy}
//	}
//}
//
//func WorkspaceOption(workspace string) Option {
//	return func() map[string]any {
//		return map[string]any{"workspace": workspace}
//	}
//}
//
///*
//Available options: none, metadata, all
//
//The purge parameter specifies if and how the underlying raster data source is deleted. When set to “none” data and auxiliary files are preserved. When set to “metadata” delete only auxiliary files and metadata. It’s recommended when data files (such as granules) should not be deleted from disk. Finally, when set to “all” both data and auxiliary files are removed.
//*/
//func PurgeOption(purge string) Option {
//	return func() map[string]any {
//		return map[string]any{"purge": purge}
//	}
//}
//
//func RecurseOption(recurse bool) Option {
//	return func() map[string]any {
//		return map[string]any{"recurse": recurse}
//	}
//}
//
//func EnabledOption(enabled bool) Option {
//	return func() map[string]any {
//		return map[string]any{"enabled": enabled}
//	}
//}
//
///*
//Available options: SINGLE, NAMED, CONTAINER, EO
//*/
//func ModeOption(mode string) Option {
//	return func() map[string]any {
//		return map[string]any{"mode": mode}
//	}
//}
//
//func ProcessOptions(options []Option) map[string]any {
//	params := make(map[string]any)
//	for _, option := range options {
//		for key, value := range option() {
//			params[key] = value
//		}
//	}
//
//	return params
//}
//
//func CoverageDimensionsOption(dimension geotiff.CoverageDimension) Option {
//	return func() map[string]any {
//		return map[string]any{
//			"dimensions": &geotiff.CoverageDimensions{CoverageDimension: dimension},
//		}
//	}
//}
//
///*
//MakeCoverageDimension generates a valid coverage dimension to be used in CoverageDimensionsOption.
//
//dataType - in general REAL_32BITS/REAL_64BITS
//nullValue - in general 0, but can be left empty
//*/
//func MakeCoverageDimension(name, dataType string, nullValue, min, max float64, description ...string) geotiff.CoverageDimension {
//	return geotiff.CoverageDimension{
//		Description: strings.Join(description, "\n"),
//		Name:        name,
//		DataType: geotiff.DataType{
//			Name: dataType,
//		},
//		NullValues: geotiff.NullValues{
//			Double: nullValue,
//		},
//		Range: geotiff.Range{
//			Max: max,
//			Min: min,
//		},
//	}
//}
