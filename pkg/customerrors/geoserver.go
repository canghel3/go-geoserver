package customerrors

type GeoserverError struct {
	Message string
	Wrapped error
}

func (ce *GeoserverError) Error() string {
	return ce.Message
}

func (ce *GeoserverError) Unwrap() error {
	return ce.Wrapped
}

func NewGeoserverError(message string) *GeoserverError {
	return &GeoserverError{Message: message}
}

func WrapGeoserverError(err error) *GeoserverError {
	return &GeoserverError{
		Message: err.Error(),
		Wrapped: err,
	}
}
