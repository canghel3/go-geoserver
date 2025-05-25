package customerrors

type NotImplementedError struct {
	Err
}

func (ce *NotImplementedError) Error() string {
	return ce.message
}

func (ce *NotImplementedError) Unwrap() error {
	return ce.wrapped
}

func NewNotImplementedError(message string) *NotImplementedError {
	c := &NotImplementedError{}
	c.msg(message)
	return c
}

func WrapNotImplementedError(err error) *NotImplementedError {
	c := &NotImplementedError{}
	c.wrap(err)
	c.msg(err.Error())
	return c
}
