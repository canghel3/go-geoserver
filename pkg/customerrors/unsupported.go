package customerrors

type UnsupportedError struct {
	Err
}

func (ce *UnsupportedError) Error() string {
	return ce.message
}

func (ce *UnsupportedError) Unwrap() error {
	return ce.wrapped
}

func NewUnsupportedError(message string) *UnsupportedError {
	c := &UnsupportedError{}
	c.msg(message)
	return c
}

func WrapUnsupportedError(err error) *UnsupportedError {
	c := &UnsupportedError{}
	c.wrap(err)
	c.msg(err.Error())
	return c
}
