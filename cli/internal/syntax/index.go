package syntax

import (
	"errors"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cli/internal/interpreter"
)

func systemCommand(input string) bool {

	value := strings.ToLower(input)

	if value == "exit" {
		log.Info("")
		log.Info("GOODBYE <3")
		log.Info("")
		os.Exit(0)
		return true
	}

	if value == "help" {
		log.Info("")
		log.Info("GOODBYE <3")
		log.Info("")
		return true
	}

	return false

}

func computeCommandArgs(args []string) []string {
	if len(args) <= 1 {
		return []string{}
	}
	return args[1:]
}

func DoProcess(input string) error {

	if systemCommand(input) {
		return nil
	}

	data := strings.Split(input, " ")
	size := len(data)
	if size > 0 {
		cmdType := interpreter.GetCommandType(strings.ToLower(data[0]))
		if cmdType == interpreter.UNDEFINED {
			return errors.New("unknown command")
		}

		cmdArgs := computeCommandArgs(data)
		executor := interpreter.GetExecutor(cmdType)

		executor(cmdArgs)
		return nil
	}

	return errors.New("please type command!")
}
