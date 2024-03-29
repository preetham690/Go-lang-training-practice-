package fileops

import (
	"errors"
	"os"
)

func GetFileCharCount(filename string) (int, error) {
	if filename == nil {
		return 0, errors.New("nil filename")
	}
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	return len(bytes), nil
}
