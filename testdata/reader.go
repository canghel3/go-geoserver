package testdata

import "os"

func Read(file string) ([]byte, error) {
	return os.ReadFile(file)
}
