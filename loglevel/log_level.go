package loglevel

type LogLevel int

const (
	_ LogLevel = iota
	DEBUG
	INFO
	ERROR
	FATAL
	PANIC
)

func (ll LogLevel) String() string {
	switch ll {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	default:
		return ""
	}
}
