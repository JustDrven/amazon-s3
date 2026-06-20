package system

import (
	"os"

	"github.com/charmbracelet/log"
)

func RunHelp() bool {
	log.Info("")
	log.Info("  user | User management")
	log.Info("  ping | Checks if the server is running")
	log.Info("")
	return true
}

func RunExit() bool {
	log.Info("")
	log.Info("GOODBYE <3")
	log.Info("")

	os.Exit(0)
	return true
}
