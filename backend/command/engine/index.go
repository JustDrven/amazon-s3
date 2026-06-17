package engine

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/command/configuration"
)

func Main(file *os.File) {
	log.Info("Starting to prepare service to startup..")

	config, err := configuration.LoadConfig(file)
	if err != nil {
		panic(err)
	}

	StartAPI(*config)
}

func isPortAvaliable(port int) error {
	if port < 0 {
		return errors.New("The API port can't be below zero!")
	}

	if port > 65535 {
		return errors.New("The API port reached the limit!")
	}

	return nil
}

func StartAPI(config configuration.ConfigData) {
	port := config.Port

	log.Info("Starting API", "port", port)
	err := isPortAvaliable(config.Port)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

}
