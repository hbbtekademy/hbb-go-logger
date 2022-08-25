package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type logLevel int

const (
	_ logLevel = iota
	_DEBUG
	_INFO
	_ERROR
	_FATAL
	_PANIC
)

func (ll logLevel) String() string {
	switch ll {
	case _DEBUG:
		return "DEBUG"
	case _INFO:
		return "INFO"
	case _ERROR:
		return "ERROR"
	case _FATAL:
		return "FATAL"
	case _PANIC:
		return "PANIC"
	default:
		return ""
	}
}

var (
	once     sync.Once
	llEnvVar string = "LOG_LEVEL"
	instance logger
)

func Debug(v ...interface{}) {
	getInstance().debug(v...)
}

func Debugf(format string, v ...interface{}) {
	getInstance().debugf(format, v...)
}

func Info(v ...interface{}) {
	getInstance().info(v...)
}

func Infof(format string, v ...interface{}) {
	getInstance().infof(format, v...)
}

func Error(v ...interface{}) {
	getInstance().error(v...)
}

func Errorf(format string, v ...interface{}) {
	getInstance().errorf(format, v...)
}

func Fatal(v ...interface{}) {
	getInstance().fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	getInstance().fatalf(format, v...)
}

func Panic(v ...interface{}) {
	getInstance().panic(v...)
}

func Panicf(format string, v ...interface{}) {
	getInstance().panicf(format, v...)
}

type logger interface {
	infof(format string, v ...interface{})
	info(v ...interface{})
	debug(v ...interface{})
	debugf(format string, v ...interface{})
	error(v ...interface{})
	errorf(format string, v ...interface{})
	fatal(v ...interface{})
	fatalf(format string, v ...interface{})
	panic(v ...interface{})
	panicf(format string, v ...interface{})
}

type stdLogger struct {
	logLevel    logLevel
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
	panicLogger *log.Logger
}

func (l stdLogger) debug(v ...interface{}) {
	if l.logLevel <= _DEBUG {
		l.debugLogger.Output(3, fmt.Sprintln(v...))
	}
}

func (l stdLogger) debugf(format string, v ...interface{}) {
	if l.logLevel <= _DEBUG {
		l.debugLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l stdLogger) info(v ...interface{}) {
	if l.logLevel <= _INFO {
		l.infoLogger.Output(3, fmt.Sprintln(v...))
	}
}

func (l stdLogger) infof(format string, v ...interface{}) {
	if l.logLevel <= _INFO {
		l.infoLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) error(v ...interface{}) {
	if l.logLevel <= _ERROR {
		l.errLogger.Output(3, fmt.Sprintln(v...))
	}
}

func (l stdLogger) errorf(format string, v ...interface{}) {
	if l.logLevel <= _ERROR {
		l.errLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) fatal(v ...interface{}) {
	if l.logLevel <= _FATAL {
		l.fatalLogger.Output(3, fmt.Sprintln(v...))
	}
	os.Exit(1)
}

func (l stdLogger) fatalf(format string, v ...interface{}) {
	if l.logLevel <= _FATAL {
		l.fatalLogger.Output(3, fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}

func (l *stdLogger) panic(v ...interface{}) {
	if l.logLevel <= _PANIC {
		s := fmt.Sprintln(v...)
		l.panicLogger.Output(3, s)
		panic(s)
	}
}

func (l stdLogger) panicf(format string, v ...interface{}) {
	if l.logLevel <= _PANIC {
		s := fmt.Sprintf(format, v...)
		l.panicLogger.Output(3, s)
		panic(s)
	}
}

func getInstance() logger {
	once.Do(func() {
		logLevel := getLogLevel()
		instance = &stdLogger{
			logLevel:    logLevel,
			debugLogger: log.New(os.Stdout, "DEBUG ", log.Lshortfile|log.LstdFlags),
			infoLogger:  log.New(os.Stdout, "INFO  ", log.Lshortfile|log.LstdFlags),
			errLogger:   log.New(os.Stderr, "ERROR ", log.Lshortfile|log.LstdFlags),
			fatalLogger: log.New(os.Stderr, "FATAL ", log.Lshortfile|log.LstdFlags),
			panicLogger: log.New(os.Stderr, "PANIC ", log.Lshortfile|log.LstdFlags),
		}
	})

	return instance
}

func SetLogLevelEnvVar(envVar string) {
	if instance != nil {
		instance.error("Logger already created")
		instance.errorf("Log Level Env Variable %s should be set before creating an instance of the logger", envVar)
		return
	}
	llEnvVar = envVar
}

func getLogLevel() logLevel {
	fmt.Fprintf(os.Stdout, "Getting Log Level from %s env var. Call logger.SetLogLevelEnvVar to override this value\n", llEnvVar)
	level := strings.TrimSpace(os.Getenv(llEnvVar))
	ll := _DEBUG
	if level != "" {
		switch strings.ToLower(level) {
		case "debug":
			ll = _DEBUG
		case "info":
			ll = _INFO
		case "error":
			ll = _ERROR
		case "fatal":
			ll = _FATAL
		case "panic":
			ll = _PANIC
		}
	}
	fmt.Fprintf(os.Stdout, "Log Level set to %s\n", ll)
	return ll
}
