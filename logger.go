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
		instance.errorf(3, "LogLevel should be set before writing any logs. LogLevel will continue to be %s", logLevel)
		return
	}
	logLevel = ll
}

// Redirect Stdout to required io.Writer. By default LogLevels DEBUG and INFO are written to Stdout.
// Redirection should be done before any logs are written.
func RedirectStdout(out io.Writer) {
	if instance != nil {
		instance.error(3, "Stdout redirection should be done before writing any logs")
		return
	}
	outWriter = out
}

// Redirect Stderr to required io.Writer. By default all LogLevels except DEBUG and INFO are written to Stderr
// Redirection should be done before any logs are written.
func RedirectStderr(out io.Writer) {
	if instance != nil {
		instance.error(3, "Stderr redirection should be done before writing any logs")
		return
	}
	errWriter = out
}

// Debug writes logs at DEBUG level.
func Debug(v ...interface{}) {
	getInstance().debug(3, v...)
}

// CallDepthDebug is same as Debug with additional Call Depth for number of stack frames to ascend.
func CallDepthDebug(depth int, v ...interface{}) {
	getInstance().debug(depth, v...)
}

// Debugf writes logs at DEBUG level formatted according to the format specifier.
func Debugf(format string, v ...interface{}) {
	getInstance().debugf(3, format, v...)
}

// CallDepthDebugf is same as Debugf with additional Call Depth for number of stack frames to ascend.
func CallDepthDebugf(depth int, format string, v ...interface{}) {
	getInstance().debugf(depth, format, v...)
}

// Info writes logs at INFO level
func Info(v ...interface{}) {
	getInstance().info(3, v...)
}

// CallDepthInfo is same as Info with additional Call Depth for number of stack frames to ascend.
func CallDepthInfo(depth int, v ...interface{}) {
	getInstance().info(depth, v...)
}

// Infof writes logs at INFO level formatted according to the format specifier.
func Infof(format string, v ...interface{}) {
	getInstance().infof(3, format, v...)
}

// CallDepthInfof is same as Infof with additional Call Depth for number of stack frames to ascend.
func CallDepthInfof(depth int, format string, v ...interface{}) {
	getInstance().infof(depth, format, v...)
}

// Error writes logs at ERROR level
func Error(v ...interface{}) {
	getInstance().error(3, v...)
}

// CallDepthError is same as Error with additional Call Depth for number of stack frames to ascend.
func CallDepthError(depth int, v ...interface{}) {
	getInstance().error(depth, v...)
}

// Errorf writes logs at ERROR level formatted according to the format specifier.
func Errorf(format string, v ...interface{}) {
	getInstance().errorf(3, format, v...)
}

// CallDepthErrorf is same as Errorf with additional Call Depth for number of stack frames to ascend.
func CallDepthErrorf(depth int, format string, v ...interface{}) {
	getInstance().errorf(depth, format, v...)
}

// Fatal writes logs at FATAL level followed by call to os.Exit(1)
func Fatal(v ...interface{}) {
	getInstance().fatal(3, v...)
}

// CallDepthFatal is same as Fatal with additional Call Depth for number of stack frames to ascend.
func CallDepthFatal(depth int, v ...interface{}) {
	getInstance().fatal(depth, v...)
}

// Fatalf writes logs at FATAL level formatted according to the format specifier followed by call to os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	getInstance().fatalf(3, format, v...)
}

// CallDepthFatalf is same as Fataf with additional Call Depth for number of stack frames to ascend.
func CallDepthFatalf(depth int, format string, v ...interface{}) {
	getInstance().fatalf(depth, format, v...)
}

// Panic writes logs at PANIC level followed by call to panic()
func Panic(v ...interface{}) {
	getInstance().panic(3, v...)
}

// CallDepthPanic is same as Panic with additional Call Depth for number of stack frames to ascend.
func CallDepthPanic(depth int, v ...interface{}) {
	getInstance().panic(depth, v...)
}

// Panic writes logs at PANIC level formatted according to the format specifier followed by call to panic()
func Panicf(format string, v ...interface{}) {
	getInstance().panicf(3, format, v...)
}

// CallDepthPanicf is same as Panicf with additional Call Depth for number of stack frames to ascend.
func CallDepthPanicf(depth int, format string, v ...interface{}) {
	getInstance().panicf(depth, format, v...)
}

type logger interface {
	infof(depth int, format string, v ...interface{})
	info(depth int, v ...interface{})
	debug(depth int, v ...interface{})
	debugf(depth int, format string, v ...interface{})
	error(depth int, v ...interface{})
	errorf(depth int, format string, v ...interface{})
	fatal(depth int, v ...interface{})
	fatalf(depth int, format string, v ...interface{})
	panic(depth int, v ...interface{})
	panicf(depth int, format string, v ...interface{})
}

type stdLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
	panicLogger *log.Logger
}

func (l stdLogger) debug(depth int, v ...interface{}) {
	if logLevel <= loglevel.DEBUG {
		l.debugLogger.Output(depth, fmt.Sprintln(v...))
	}
}

func (l stdLogger) debugf(depth int, format string, v ...interface{}) {
	if logLevel <= loglevel.DEBUG {
		l.debugLogger.Output(depth, fmt.Sprintf(format, v...))
	}
}

func (l stdLogger) info(depth int, v ...interface{}) {
	if logLevel <= loglevel.INFO {
		l.infoLogger.Output(depth, fmt.Sprintln(v...))
	}
}

func (l stdLogger) infof(depth int, format string, v ...interface{}) {
	if logLevel <= loglevel.INFO {
		l.infoLogger.Output(depth, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) error(depth int, v ...interface{}) {
	if logLevel <= loglevel.ERROR {
		l.errLogger.Output(depth, fmt.Sprintln(v...))
	}
}

func (l stdLogger) errorf(depth int, format string, v ...interface{}) {
	if logLevel <= loglevel.ERROR {
		l.errLogger.Output(depth, fmt.Sprintf(format, v...))
	}
}

func (l *stdLogger) fatal(depth int, v ...interface{}) {
	if logLevel <= loglevel.FATAL {
		l.fatalLogger.Output(depth, fmt.Sprintf("%s. Exiting...", v...))
	}
	os.Exit(1)
}

func (l stdLogger) fatalf(depth int, format string, v ...interface{}) {
	if logLevel <= loglevel.FATAL {
		format = format + ". Exiting..."
		l.fatalLogger.Output(3, fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}

func (l *stdLogger) panic(depth int, v ...interface{}) {
	if logLevel <= loglevel.PANIC {
		s := fmt.Sprintf("%s. Panicing...", v...)
		l.panicLogger.Output(3, s)
		panic(s)
	}
}

func (l stdLogger) panicf(depth int, format string, v ...interface{}) {
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
