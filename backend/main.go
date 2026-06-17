package main

import (
	"os"

	"justdrven.dev/storage/command/engine"
)

func main() {
	args := os.Args
	config := "./config.json"

	if len(args) > 1 {
		config = args[1]
	}

	configFile, configErr := os.Open(config)
	if configErr != nil {
		panic(configErr)
	}

	engine.Main(configFile)
}
