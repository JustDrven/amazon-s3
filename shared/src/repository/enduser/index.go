package enduser

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"

	"justdrven.dev/storage/internal/security"
	"justdrven.dev/storage/shared/src/configuration"
)

var data = map[string]*User{}

func Map() map[string]*User {
	return data
}

func Fetch() error {
	fileBytes, err := os.ReadFile(configuration.GetConfig().EndUserFilePath)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(fileBytes)
	json.NewDecoder(reader).Decode(&data)

	return nil
}

func Get(name string) (*User, error) {

	user := data[name]
	if user == nil {
		return nil, errors.New("the user not found!")
	}

	return user, nil
}

func Save(name string, password string) bool {
	if _, founded := data[name]; founded {
		return false
	}

	data[name] = &User{
		Name:     name,
		Password: security.Hash(password),
	}

	err := Flush()

	if err != nil {
		panic(err)
	}

	return true

}

func Flush() error {
	file, fileErr := os.Create(configuration.GetConfig().EndUserFilePath)
	if fileErr != nil {
		return fileErr
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	encoder.Encode(data)

	return nil
}

func Delete(name string) bool {
	_, exist := data[name]
	if exist {

		delete(data, name)
		Flush()
		return true

	}

	return false

}
