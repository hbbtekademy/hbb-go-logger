# Simple Go Logger
This is a simple wrapper over the standard library [log](https://pkg.go.dev/log) package with ability to log at different log levels and set the required log level.

## Installation
```bash
go get -u github.com/hbbtekademy/hbb-go-logger
```

## Log Levels in order of hierarchy
- DEBUG
- INFO
- ERROR
- FATAL
- PANIC

Default Log Level is **INFO**.

**DEBUG** and **INFO** write to **_stdout_**, all other log levels write to **_stderr_** by default.

**FATAL** log executes os.Exit(1) after logging.

**PANIC** log execute panic() after logging.

Setting Log Level and redirecting log output can be done as follows. This should be done before writing any logs.

```go
package main

import (
	"os"

	logger "github.com/hbbtekademy/hbb-go-logger"
	"github.com/hbbtekademy/hbb-go-logger/loglevel"
)

func main() {
	logger.SetLogLevel(loglevel.DEBUG)
	f, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}
	logger.RedirectStdout(f)
	logger.RedirectStderr(f)

	logger.Info("Writting some info logs")
	logger.Infof("Writting some formatted info logs with int: %d, string: %s", 1, "infof")
	logger.Debug("Writting some debug logs")
	logger.Debugf("Writting some formatted debug logs with int: %d, string: %s", 2, "debugf")
	logger.Error("Writting some error logs")
	logger.Errorf("Writting some formatted error logs with int: %d, string: %s", 3, "errorf")
}
```
<pre>
INFO  2022/08/26 13:26:45 main.go:19: Writting some info logs
INFO  2022/08/26 13:26:45 main.go:20: Writting some formatted info logs with int: 1, string: infof
DEBUG 2022/08/26 13:26:45 main.go:21: Writting some debug logs
DEBUG 2022/08/26 13:26:45 main.go:22: Writting some formatted debug logs with int: 2, string: debugf
ERROR 2022/08/26 13:26:45 main.go:23: Writting some error logs
ERROR 2022/08/26 13:26:45 main.go:24: Writting some formatted error logs with int: 3, string: errorf

</pre>
