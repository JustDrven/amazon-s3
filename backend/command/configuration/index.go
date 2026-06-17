package configuration

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/charmbracelet/log"
)

var cachedConfig ConfigData

func LoadConfig(config *os.File) (*ConfigData, error) {
	defer config.Close()
	log.Info("Scanning config..", "file", config.Name())

	byteStream, err := io.ReadAll(config)
	if err != nil {
		return nil, errors.New("failed to read config file")
	}

	data := &ConfigData{}
	decodeErr := json.NewDecoder(bytes.NewReader(byteStream)).Decode(data)

	if decodeErr != nil {
		return nil, errors.New("failed to decode json data")
	}

	cachedConfig = *data

	return data, nil
}

func GetConfig() ConfigData {
	return cachedConfig
}
