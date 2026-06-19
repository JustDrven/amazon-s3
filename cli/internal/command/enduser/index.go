package enduser

import (
	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/internal/interpreter"
)

func AddEndUserCommand(args interpreter.CommandArgs) {

	username := args[0]
	password := args[1]

	log.Info("ADDED NEW USER", "NAME", username, "PASS", password)

}

func RemoveEndUserCommand(args interpreter.CommandArgs) {

	username := args[0]

	log.Info("REMOVE USER", "NAME", username)

}
