package ping

import (
	"fmt"
	"net"
	"time"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/internal/interpreter"
	"justdrven.dev/storage/shared/src/configuration"
)

func PingCommand(args interpreter.CommandArgs) {

	address := fmt.Sprintf("127.0.0.1:%d", configuration.GetConfig().Port)

	log.Info("Scanning network..", "ADDRESS", address)

	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", address, timeout)

	if err == nil {
		log.Info("The server is running!")
	} else {
		log.Warn("The server isn't running!")
	}
}
