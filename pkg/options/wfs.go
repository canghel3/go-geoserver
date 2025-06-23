package options

import (
	"github.com/canghel3/go-geoserver/pkg/datastores"
	"strconv"
)

var WFS WFSOptionGenerator

type WFSOptionGenerator struct{}

type WFSOptions func(params *datastores.ConnectionParams)

/*
{
                    "@key": "WFSDataStoreFactory:TRY_GZIP",
                    "$": "true"
                },
                {
                    "@key": "usedefaultsrs",
                    "$": "false"
                },
                {
                    "@key": "WFSDataStoreFactory:AXIS_ORDER",
                    "$": "Compliant"
                },
                {
                    "@key": "WFSDataStoreFactory:PROTOCOL",
                    "$": "false"
                },
                {
                    "@key": "WFSDataStoreFactory:GET_CAPABILITIES_URL",
                    "$": "http://geoserver:8080/geoserver/wfs?service=wfs"
                },
                {
                    "@key": "WFSDataStoreFactory:PASSWORD",
                    "$": "crypt2:5yUM0gXhU2vRZOWg+pcdLqz8QmG6Tvi1ly8SBkdgwdE="
                },
                {
                    "@key": "WFSDataStoreFactory:MAXFEATURES",
                    "$": "0"
                },
                {
                    "@key": "WFSDataStoreFactory:FILTER_COMPLIANCE",
                    "$": "0"
                },
                {
                    "@key": "WFSDataStoreFactory:TIMEOUT",
                    "$": "3000"
                },
                {
                    "@key": "WFSDataStoreFactory:WFS_STRATEGY",
                    "$": "auto"
                },
                {
                    "@key": "WFSDataStoreFactory:AXIS_ORDER_FILTER",
                    "$": "Compliant"
                },
                {
                    "@key": "WFSDataStoreFactory:ENCODING",
                    "$": "UTF-8"
                },
                {
                    "@key": "WFSDataStoreFactory:GML_COMPATIBLE_TYPENAMES",
                    "$": "false"
                },
                {
                    "@key": "WFSDataStoreFactory:LENIENT",
                    "$": "false"
                },
                {
                    "@key": "WFSDataStoreFactory:USERNAME",
                    "$": "admin"
                },
                {
                    "@key": "WFSDataStoreFactory:GML_COMPLIANCE_LEVEL",
                    "$": "0"
                },
                {
                    "@key": "WFSDataStoreFactory:MAX_CONNECTION_POOL_SIZE",
                    "$": "6"
                },
                {
                    "@key": "WFSDataStoreFactory:BUFFER_SIZE",
                    "$": "10"
                },
                {
                    "@key": "namespace",
                    "$": "http://PLAYGROUND"
                },
                {
                    "@key": "WFSDataStoreFactory:USE_HTTP_CONNECTION_POOLING",
                    "$": "true"
                }
*/

// Timeout sets the connection timeout in seconds
func (wog WFSOptionGenerator) Timeout(timeout uint) WFSOptions {
	return func(params *datastores.ConnectionParams) {
		(*params)["Time-out"] = strconv.FormatUint(uint64(timeout), 10)
	}
}
