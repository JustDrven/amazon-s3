package ping

import (
	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/internal/interpreter"
	"justdrven.dev/storage/shared/src/configuration"
)

func PingCommand(args interpreter.CommandArgs) {

	host := "127.0.0.1"
	port := configuration.GetConfig().Port

	log.Info("Scanning network..", "HOST", host, "PORT", port)

}
