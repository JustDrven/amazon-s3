package interpreter

var commands = map[CommandType]func(args CommandArgs){}

func (c Command) Register() {
	commands[c.Type] = c.Executor
}

func GetCommandType(cmdType string) CommandType {
	switch cmdType {
	case "user":
		return ENDUSER
	case "ping":
		return PING
	default:
		return UNDEFINED
	}
}

func GetExecutor(cmdType CommandType) func(CommandArgs) {
	return commands[cmdType]
}
