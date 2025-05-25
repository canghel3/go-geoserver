package customerrors

type ConflictError struct {
	Err
}

func (ce *ConflictError) Error() string {
	return ce.message
}

func (ce *ConflictError) Unwrap() error {
	return ce.wrapped
}

func NewConflictError(message string) *ConflictError {
	c := &ConflictError{}
	c.msg(message)
	return c
}

func WrapConflictError(err error) *ConflictError {
	c := &ConflictError{}
	c.wrap(err)
	c.msg(err.Error())
	return c
}
