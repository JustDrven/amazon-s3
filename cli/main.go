package main

import (
	"os"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cmd/configuration"
	"justdrven.dev/storage/pkg"
)

func claimConfigFileName() string {
	args := os.Args
	config := "./config.json"

	if len(args) > 1 {
		config = args[1]
	}

	return config
}

func main() {
	pkg.PrintLogo()

	config := claimConfigFileName()
	configFile, err := os.Open(config)

	if err != nil {
		log.Fatal("we can't find your config file")
		os.Exit(1)
	}

	configuration.LoadConfig(configFile)

}
