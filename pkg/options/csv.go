package options

import (
	"github.com/canghel3/go-geoserver/pkg/datastores"
)

var CSV CSVOptionsGenerator

type CSVOptionsGenerator struct{}

type CSVOptions func(params *datastores.ConnectionParams)

//// Charset sets the character set for the CSV file
//func (cog CSVOptionsGenerator) Charset(charset string) CSVOptions {
//	return func(params *datastores.ConnectionParams) {
//		(*params)["charset"] = charset
//	}
//}
//
//// Delimiter sets the delimiter used in the CSV file (default is comma)
//func (cog CSVOptionsGenerator) Delimiter(delimiter string) CSVOptions {
//	return func(params *datastores.ConnectionParams) {
//		(*params)["delimiter"] = delimiter
//	}
//}
//
//// Quote sets the quote character used in the CSV file (default is double quote)
//func (cog CSVOptionsGenerator) Quote(quote string) CSVOptions {
//	return func(params *datastores.ConnectionParams) {
//		(*params)["quote"] = quote
//	}
//}
//
//// Escape sets the escape character used in the CSV file
//func (cog CSVOptionsGenerator) Escape(escape string) CSVOptions {
//	return func(params *datastores.ConnectionParams) {
//		(*params)["escape"] = escape
//	}
//}
//
//// Headers specifies whether the first row contains headers (default is true)
//func (cog CSVOptionsGenerator) Headers(headers bool) CSVOptions {
//	return func(params *datastores.ConnectionParams) {
//		(*params)["headers"] = strconv.FormatBool(headers)
//	}
//}
//
//// LatField specifies the name of the latitude field
//func (cog CSVOptionsGenerator) LatField(field string) CSVOptions {
//	return func(params *datastores.ConnectionParams) {
//		(*params)["lat.field"] = field
//	}
//}
//
//// LngField specifies the name of the longitude field
//func (cog CSVOptionsGenerator) LngField(field string) CSVOptions {
//	return func(params *datastores.ConnectionParams) {
//		(*params)["lng.field"] = field
//	}
//}
