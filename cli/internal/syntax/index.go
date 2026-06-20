package syntax

import (
	"errors"
	"strings"

	"justdrven.dev/storage/cli/internal/command/system"
	"justdrven.dev/storage/cli/internal/interpreter"
)

func systemCommand(input string) bool {

	value := strings.ToLower(input)

	switch value {
	case "exit":
		return system.RunExit()
	case "help":
		return system.RunHelp()
	default:
		return false
	}

}

func computeCommandArgs(args interpreter.CommandArgs) []string {
	if len(args) <= 1 {
		return interpreter.CommandArgs{}
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
