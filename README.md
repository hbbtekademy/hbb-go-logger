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

**FATAL** logs executes os.Exit(1) after logging.

**PANIC** logs execute panic() after logging.

Setting Log Level and redirecting log output can be done as follows. This should be done before writing any logs.

```go
func main() {
  logger.SetLogLevel(loglevel.DEBUG)
  f, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    log.Fatalf("error opening file: %v", err)
  }
  logger.RedirectStdout(f)
  logger.RedirectStderr(f)
  
  logger.Info("Writing some info logs")
  logger.Debug("Writing some debug logs")
}
```
<pre>
INFO  2022/08/26 13:13:08 main.go:23: Writing some info logs
DEBUG 2022/08/26 13:13:08 main.go:24: Writing some debug logs
</pre>
