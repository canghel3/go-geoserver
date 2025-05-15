package testdata

import (
	"io"
	"os"
)

func Copy(src, dst string) error {
	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	ds, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer ds.Close()

	_, err = io.Copy(ds, sr)
	if err != nil {
		return err
	}

	return nil
}
