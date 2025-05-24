package customerrors

type CreateError struct {
	Message string
	Wrapped error
}

func (ce *CreateError) Error() string {
	return ce.Message
}

func (ce *CreateError) Unwrap() error {
	return ce.Wrapped
}

func NewCreateError(message string) *CreateError {
	return &CreateError{Message: message}
}

func WrapCreateError(err error) *CreateError {
	return &CreateError{
		Message: err.Error(),
		Wrapped: err,
	}
}
