package enduser

import (
	"strings"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/internal/interpreter"
)

func addEndUserProcess(name string, password string) {
	log.Info("Added an user", "USERNAME", name)
}

func removeEndUserProcess(name string) {
	log.Info("Removed an user", "USERNAME", name)
}

func EndUserCommand(args interpreter.CommandArgs) {

	if len(args) > 1 {
		action := strings.ToLower(args[0])

		if action == "add" {
			if len(args) < 3 {
				log.Error("please type username and password")
				return
			}

			addEndUserProcess(args[1], args[2])
			return
		}

		if action == "remove" {
			if len(args) < 2 {
				log.Error("please type username")
				return
			}

			removeEndUserProcess(args[1])
			return
		}

	}

	log.Error("unknown arguments")

}
