package main

import (
	"os"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cmd"
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
	configFile, configErr := os.Open(config)

	if configErr == nil {
		cmd.Main(configFile)
	} else {
		log.Fatal(configErr.Error())
		os.Exit(1)
	}

}
