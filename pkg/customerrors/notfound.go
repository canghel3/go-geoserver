package customerrors

type NotFoundError struct {
	Err
}

func (ce *NotFoundError) Error() string {
	return ce.message
}

func (ce *NotFoundError) Unwrap() error {
	return ce.wrapped
}

func NewNotFoundError(message string) *NotFoundError {
	c := &NotFoundError{}
	c.msg(message)
	return c
}

func WrapNotFoundError(err error) *NotFoundError {
	e := &NotFoundError{}
	e.wrap(err)
	e.msg(err.Error())
	return e
}
