package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/hbbtekademy/hbb-go-logger/loglevel"
)

var (
	once      sync.Once
	logLevel  loglevel.LogLevel = loglevel.INFO
	instance  logger
	outWriter io.Writer = os.Stdout
	errWriter io.Writer = os.Stderr
)

// Set the LogLevel. Should be done before writing any logs.
func SetLogLevel(ll loglevel.LogLevel) {
	if logLevel == ll {
		return
	}

	if instance != nil {
		instance.errorf("LogLevel should be set before writing any logs. LogLevel will continue to be %s", logLevel)
		return
	}
	logLevel = ll
}

// Redirect Stdout to required io.Writer. By default LogLevels DEBUG and INFO are written to Stdout.
// Redirection should be done before any logs are written.
func RedirectStdout(out io.Writer) {
	if instance != nil {
		instance.error("Stdout redirection should be done before writing any logs")
		return
	}
	outWriter = out
}

// Redirect Stderr to required io.Writer. By default all LogLevels except DEBUG and INFO are written to Stderr
// Redirection should be done before any logs are written.
func RedirectStderr(out io.Writer) {
	if instance != nil {
		instance.error("Stderr redirection should be done before writing any logs")
		return
	}
	errWriter = out
}

// Debug writes logs at DEBUG level.
func Debug(v ...interface{}) {
	getInstance().debug(v...)
}

// Debugf writes logs at DEBUG level formatted according to the format specifier.
func Debugf(format string, v ...interface{}) {
	getInstance().debugf(format, v...)
}

// Info writes logs at INFO level
func Info(v ...interface{}) {
	getInstance().info(v...)
}

// Infof writes logs at INFO level formatted according to the format specifier.
func Infof(format string, v ...interface{}) {
	getInstance().infof(format, v...)
}

// Error writes logs at ERROR level
func Error(v ...interface{}) {
	getInstance().error(v...)
}

// Errorf writes logs at ERROR level formatted according to the format specifier.
func Errorf(format string, v ...interface{}) {
	getInstance().errorf(format, v...)
}

// Fatal writes logs at FATAL level followed by call to os.Exit(1)
func Fatal(v ...interface{}) {
	getInstance().fatal(v...)
}

// Fatalf writes logs at FATAL level formatted according to the format specifier followed by call to os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	getInstance().fatalf(format, v...)
}

// Panic writes logs at PANIC level followed by call to panic()
func Panic(v ...interface{}) {
	getInstance().panic(v...)
}

// Panic writes logs at PANIC level formatted according to the format specifier followed by call to panic()
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
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
	panicLogger *log.Logger
}

func (l stdLogger) debug(v ...interface{}) {
	if logLevel <= loglevel.DEBUG {
		l.debugLogger.Output(3, fmt.Sprintln(v...))
	}
}

func (l stdLogger) debugf(format string, v ...interface{}) {
	if logLevel <= loglevel.DEBUG {
		l.debugLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l stdLogger) info(v ...interface{}) {
	if logLevel <= loglevel.INFO {
		l.infoLogger.Output(3, fmt.Sprintln(v...))
	}
}

func (l stdLogger) infof(format string, v ...interface{}) {
	if logLevel <= loglevel.INFO {
		l.infoLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) error(v ...interface{}) {
	if logLevel <= loglevel.ERROR {
		l.errLogger.Output(3, fmt.Sprintln(v...))
	}
}

func (l stdLogger) errorf(format string, v ...interface{}) {
	if logLevel <= loglevel.ERROR {
		l.errLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) fatal(v ...interface{}) {
	if logLevel <= loglevel.FATAL {
		l.fatalLogger.Output(3, fmt.Sprintf("%s. Exiting...", v...))
	}
	os.Exit(1)
}

func (l stdLogger) fatalf(format string, v ...interface{}) {
	if logLevel <= loglevel.FATAL {
		format = format + ". Exiting..."
		l.fatalLogger.Output(3, fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}

func (l *stdLogger) panic(v ...interface{}) {
	if logLevel <= loglevel.PANIC {
		s := fmt.Sprintf("%s. Panicing...", v...)
		l.panicLogger.Output(3, s)
		panic(s)
	}
}

func (l stdLogger) panicf(format string, v ...interface{}) {
	if logLevel <= loglevel.PANIC {
		format = format + ". Panicing..."
		s := fmt.Sprintf(format, v...)
		l.panicLogger.Output(3, s)
		panic(s)
	}
}

func getInstance() logger {
	once.Do(func() {
		instance = &stdLogger{
			debugLogger: log.New(outWriter, "DEBUG ", log.Lshortfile|log.LstdFlags),
			infoLogger:  log.New(outWriter, "INFO  ", log.Lshortfile|log.LstdFlags),
			errLogger:   log.New(errWriter, "ERROR ", log.Lshortfile|log.LstdFlags),
			fatalLogger: log.New(errWriter, "FATAL ", log.Lshortfile|log.LstdFlags),
			panicLogger: log.New(errWriter, "PANIC ", log.Lshortfile|log.LstdFlags),
		}
	})

	return instance
}
