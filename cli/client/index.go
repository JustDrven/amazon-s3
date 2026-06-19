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
	log.Info("	If you don't know what to do... type")
	log.Info("		- help")
	log.Info("")
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
