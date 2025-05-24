package customerrors

type InputError struct {
	Message string
	Wrapped error
}

func (ce *InputError) Error() string {
	return ce.Message
}

func (ce *InputError) Unwrap() error {
	return ce.Wrapped
}

func NewInputError(message string) *InputError {
	return &InputError{Message: message}
}

func WrapInputError(err error) *InputError {
	return &InputError{
		Message: err.Error(),
		Wrapped: err,
	}
}
