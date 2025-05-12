package customerrors

type Err struct {
	message string
	wrapped error
}

func (e *Err) wrap(err error) {
	e.wrapped = err
}

func (e *Err) msg(msg string) {
	e.message = msg
}
