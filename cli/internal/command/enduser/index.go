package enduser

import (
	"strings"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/internal/interpreter"
	"justdrven.dev/storage/shared/src/repository/enduser"
)

func addEndUserProcess(name string, password string) {
	if enduser.Save(name, password) {
		log.Info("Added an user", "USERNAME", name)
	} else {
		log.Error("The user with this name already exists!")
	}
}

func removeEndUserProcess(name string) {
	if enduser.Delete(name) {
		log.Info("Removed an user", "USERNAME", name)

	} else {
		log.Error("The user doesn't exist!")
	}
}

func EndUserCommand(args interpreter.CommandArgs) {

	if len(args) > 0 {
		action := strings.ToLower(args[0])

		if action == "list" {
			users := enduser.FindAll()
			if len(users) == 0 {
				log.Warn("users are empty!")
				return
			}
			index := 1

			for i := range users {
				user := users[i]

				log.Info("User's info", "INDEX", index, "USER", user.Name)

				index++
			}

			return
		}

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
