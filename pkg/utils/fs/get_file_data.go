package fs

import (
	"errors"
	"os"
)

func LoadFile(file string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, errors.New("cannot get current working directory")
	}
	data, err := os.ReadFile(dir + "/" + file)
	if err != nil {
		return nil, errors.New("cannot access file")
	}
	return data, nil
}
