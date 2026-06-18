package object

import (
	"errors"
	"io"
	"os"
)

func GetContent(filePath string) ([]byte, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return make([]byte, 0), 404, errors.New("couldn't open the file!")
	}

	bytes, readErr := io.ReadAll(file)
	if readErr != nil {
		return make([]byte, 0), 500, errors.New("couldn't read the file!")
	}

	defer file.Close()

	return bytes, 200, nil
}
