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
	DEBUG
	INFO
	ERROR
	FATAL
	PANIC
)

var (
	once     sync.Once
	llEnvVar string = "LOG_LEVEL"
	instance logger
)

type logger interface {
	Infof(format string, v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}

type stdLogger struct {
	logLevel    logLevel
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
	panicLogger *log.Logger
}

func (l stdLogger) Debug(v ...interface{}) {
	if l.logLevel <= DEBUG {
		l.debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l stdLogger) Debugf(format string, v ...interface{}) {
	if l.logLevel <= DEBUG {
		l.debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l stdLogger) Info(v ...interface{}) {
	if l.logLevel <= INFO {
		l.infoLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l stdLogger) Infof(format string, v ...interface{}) {
	if l.logLevel <= INFO {
		l.infoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) Error(v ...interface{}) {
	if l.logLevel <= ERROR {
		l.errLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l stdLogger) Errorf(format string, v ...interface{}) {
	if l.logLevel <= ERROR {
		l.errLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) Fatal(v ...interface{}) {
	if l.logLevel <= FATAL {
		l.fatalLogger.Output(2, fmt.Sprintln(v...))
	}
	os.Exit(1)
}

func (l stdLogger) Fatalf(format string, v ...interface{}) {
	if l.logLevel <= FATAL {
		l.fatalLogger.Output(2, fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}

func (l *stdLogger) Panic(v ...interface{}) {
	if l.logLevel <= PANIC {
		s := fmt.Sprintln(v...)
		l.panicLogger.Output(2, s)
		panic(s)
	}
}

func (l stdLogger) Panicf(format string, v ...interface{}) {
	if l.logLevel <= PANIC {
		s := fmt.Sprintf(format, v...)
		l.panicLogger.Output(2, s)
		panic(s)
	}
}

func New() logger {
	once.Do(func() {
		logLevel := getLogLevel()
		instance = &stdLogger{
			logLevel:    logLevel,
			debugLogger: log.New(os.Stdout, "DEBUG: ", log.Lshortfile|log.LstdFlags),
			infoLogger:  log.New(os.Stdout, "INFO: ", log.Lshortfile|log.LstdFlags),
			errLogger:   log.New(os.Stderr, "ERROR: ", log.Lshortfile|log.LstdFlags),
			fatalLogger: log.New(os.Stderr, "FATAL: ", log.Lshortfile|log.LstdFlags),
			panicLogger: log.New(os.Stderr, "PANIC: ", log.Lshortfile|log.LstdFlags),
		}
		instance.Infof("Log Level set to %s", logLevel)
	})

	return instance
}

func SetLogLevelEnvVar(envVar string) {
	if instance == nil {
		fmt.Fprintf(os.Stderr, "Logger already created.\n")
		fmt.Fprintf(os.Stderr, "Log Level Env Variable %s should be set before creating an instance of the logger", envVar)
	}
	llEnvVar = envVar
}

func getLogLevel() logLevel {
	fmt.Fprintf(os.Stdout, "Getting Log Level from %s", llEnvVar)
	level := strings.TrimSpace(os.Getenv(llEnvVar))
	if level != "" {
		switch strings.ToLower(level) {
		case "debug":
			return DEBUG
		case "info":
			return INFO
		case "error":
			return ERROR
		case "fatal":
			return FATAL
		case "panic":
			return PANIC
		}
	}
	return DEBUG
}
