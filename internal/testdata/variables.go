package testdata

import "errors"

type ErrorReader struct{}

func (*ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("reader error")
}
