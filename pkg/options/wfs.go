package options

import (
	"github.com/canghel3/go-geoserver/internal"
	"strconv"
)

var WFS WFSOptionGenerator

type WFSOptionGenerator struct{}

type WFSOptions func(params *internal.ConnectionParams)

// Timeout sets the connection timeout in seconds
func (wog WFSOptionGenerator) Timeout(timeout uint) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Time-out"] = strconv.FormatUint(uint64(timeout), 10)
	}
}

// Username sets the username for authentication
func (wog WFSOptionGenerator) Username(username string) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Username"] = username
	}
}

// Password sets the password for authentication
func (wog WFSOptionGenerator) Password(password string) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Password"] = password
	}
}

// Protocol sets the protocol version (1.0.0, 1.1.0, 2.0.0)
func (wog WFSOptionGenerator) Protocol(version string) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Protocol"] = version
	}
}

// BufferSize sets the buffer size for WFS requests
func (wog WFSOptionGenerator) BufferSize(size uint) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Buffer Size"] = strconv.FormatUint(uint64(size), 10)
	}
}

// MaxFeatures sets the maximum number of features to retrieve
func (wog WFSOptionGenerator) MaxFeatures(max uint) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Maximum features"] = strconv.FormatUint(uint64(max), 10)
	}
}

// UseDefaultSRS specifies whether to use the default SRS
func (wog WFSOptionGenerator) UseDefaultSRS(use bool) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Use Default SRS"] = strconv.FormatBool(use)
	}
}

// OutputFormat sets the output format for WFS requests
func (wog WFSOptionGenerator) OutputFormat(format string) WFSOptions {
	return func(params *internal.ConnectionParams) {
		(*params)["Outputformat"] = format
	}
}
