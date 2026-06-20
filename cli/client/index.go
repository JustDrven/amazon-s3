package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/internal/syntax"
)

func Start() {

	log.Info("")
	log.Info("")
	log.Info("			WELCOME ON AMAZON CLI")
	log.Info("")
	log.Info("	SYSTEM COMMANDS:")
	log.Info("		- help | Prints all commands to use")
	log.Info("		- exit | Leaves the CLI")
	log.Info("")

	for {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		err := syntax.DoProcess(input)
		if err != nil {
			log.Error(err)
		}

	}

}
