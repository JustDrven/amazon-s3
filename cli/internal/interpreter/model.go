package interpreter

type CommandType int
type CommandArgs []string

const (
	ENDUSER CommandType = iota
	PING

	UNDEFINED
)

type Command struct {
	Type     CommandType
	Executor func(args CommandArgs)
}
