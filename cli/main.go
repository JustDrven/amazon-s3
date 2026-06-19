package main

import (
	"os"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/client"
	"justdrven.dev/storage/cli/internal/command/enduser"
	"justdrven.dev/storage/cli/internal/command/ping"
	"justdrven.dev/storage/cli/internal/interpreter"
	"justdrven.dev/storage/pkg"
	"justdrven.dev/storage/shared/src/configuration"
)

func claimConfigFileName() string {
	args := os.Args
	config := "./config.json"

	if len(args) > 1 {
		config = args[1]
	}

	return config
}

func registerCommands() {

	interpreter.Command{
		Type:     interpreter.PING,
		Executor: ping.PingCommand,
	}.Register()

	interpreter.Command{
		Type:     interpreter.ENDUSER,
		Executor: enduser.AddEndUserCommand,
	}.Register()

	interpreter.Command{
		Type:     interpreter.ENDUSER,
		Executor: enduser.RemoveEndUserCommand,
	}.Register()

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
	registerCommands()

	client.Start()

}
