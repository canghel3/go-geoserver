package customerrors

type ResourceError struct {
	Message string
	Wrapped error
}

func (ce *ResourceError) Error() string {
	return ce.Message
}

func (ce *ResourceError) Unwrap() error {
	return ce.Wrapped
}

func NewResourceError(message string) *ResourceError {
	return &ResourceError{Message: message}
}

func WrapResourceError(err error) *ResourceError {
	return &ResourceError{
		Message: err.Error(),
		Wrapped: err,
	}
}
