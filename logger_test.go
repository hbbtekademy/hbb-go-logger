package logger

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/hbbtekademy/hbb-go-logger/loglevel"
)

var (
	buf bytes.Buffer
)

func TestDefaultLogLevel(t *testing.T) {
	if logLevel != loglevel.INFO {
		t.Errorf("Expected default log level to be INFO but got %s", logLevel)
	}
}

func TestRedirectStdout(t *testing.T) {
	RedirectStdout(os.Stderr)
	if outWriter != os.Stderr {
		t.Error("Expected Stdout to be redirect Stderr but is not")
	}
}

func TestRedirectStderr(t *testing.T) {
	RedirectStderr(os.Stdout)
	if errWriter != os.Stdout {
		t.Error("Expected Stderr to be redirect Stdout but is not")
	}
}

func TestInfoLogging(t *testing.T) {
	logLevel = loglevel.INFO

	redirectLogOutputs()
	defer buf.Reset()

	Info("Test Info Log")
	expectedLog := "Test Info Log\n"
	actualLog := buf.String()

	verify(t, actualLog, expectedLog, logLevel.String())
}

func TestInfofLogging(t *testing.T) {
	logLevel = loglevel.INFO

	redirectLogOutputs()
	defer buf.Reset()

	Infof("Test Info Log: %d %s", 1, "str")
	expectedLog := "Test Info Log: 1 str\n"
	actualLog := buf.String()

	verify(t, actualLog, expectedLog, logLevel.String())
}

func TestDebugLogging(t *testing.T) {
	logLevel = loglevel.DEBUG

	redirectLogOutputs()
	defer buf.Reset()

	Debug("Test Debug Log")
	expectedLog := "Test Debug Log\n"
	actualLog := buf.String()

	verify(t, actualLog, expectedLog, logLevel.String())
}

func TestErrorLogging(t *testing.T) {
	logLevel = loglevel.ERROR

	redirectLogOutputs()
	defer buf.Reset()

	Error("Test Error Log")
	expectedLog := "Test Error Log\n"
	actualLog := buf.String()

	verify(t, actualLog, expectedLog, logLevel.String())
}

func TestDebugfLogging(t *testing.T) {
	logLevel = loglevel.DEBUG

	redirectLogOutputs()
	defer buf.Reset()

	Debugf("Test Debug Log: %d %s", 2, "debug")
	expectedLog := "Test Debug Log: 2 debug\n"
	actualLog := buf.String()

	verify(t, actualLog, expectedLog, logLevel.String())
}

func TestErrorfLogging(t *testing.T) {
	logLevel = loglevel.ERROR

	redirectLogOutputs()
	defer buf.Reset()

	Errorf("Test Errorf Log: %d %s", 2, "errorf")
	expectedLog := "Test Errorf Log: 2 errorf\n"
	actualLog := buf.String()

	verify(t, actualLog, expectedLog, logLevel.String())
}

func TestFatalLogging(t *testing.T) {
	envVar := "TEST_FATAL_LOGGING"
	if os.Getenv(envVar) == "1" {
		logLevel = loglevel.FATAL
		Fatal("Test fatal log")
		return
	}

	actualLog := runCmd(t, "TestFatalLogging", envVar)
	expectedLog := "Test fatal log. Exiting...\n"
	verify(t, actualLog, expectedLog, loglevel.FATAL.String())
}

func TestFatalfLogging(t *testing.T) {
	envVar := "TEST_FATAL_LOGGING"
	if os.Getenv(envVar) == "1" {
		logLevel = loglevel.FATAL
		Fatalf("Test fatalf log: %d %s", 3, "fatalf")
		return
	}

	actualLog := runCmd(t, "TestFatalfLogging", envVar)
	expectedLog := "Test fatalf log: 3 fatalf. Exiting...\n"
	verify(t, actualLog, expectedLog, loglevel.FATAL.String())
}

func TestPanicLogging(t *testing.T) {
	redirectLogOutputs()
	defer buf.Reset()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected code to panic but did not")
		}
		expectedLog := "Test Panic Log. Panicing...\n"
		actualLog := buf.String()
		verify(t, actualLog, expectedLog, logLevel.String())
	}()

	logLevel = loglevel.PANIC
	Panic("Test Panic Log")
}

func TestPanicfLogging(t *testing.T) {
	redirectLogOutputs()
	defer buf.Reset()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected code to panic but did not")
		}

		expectedLog := "Test Panic Log: 10 panic. Panicing...\n"
		actualLog := buf.String()
		verify(t, actualLog, expectedLog, logLevel.String())
	}()

	logLevel = loglevel.PANIC
	Panicf("Test Panic Log: %d %s", 10, "panic")
}

func redirectLogOutputs() {
	if instance == nil {
		RedirectStderr(&buf)
		RedirectStdout(&buf)
	}
}

func runCmd(t *testing.T, testName string, envVar string) string {
	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), envVar+"=1")
	cmdOutput, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed starting cmd with error %v", err)
	}

	b, _ := ioutil.ReadAll(cmdOutput)

	err := cmd.Wait()
	if e, ok := err.(*exec.ExitError); !ok || e.Success() {
		t.Fatalf("Process ran with err %v, wanted exit status 1", err)
	}

	return string(b)
}

func verify(t *testing.T, actualLog, expectedLog, expectedLogLevel string) {
	if !strings.HasPrefix(actualLog, expectedLogLevel) {
		t.Errorf("Expected log to start with %s instead got: %s", expectedLogLevel, actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}
}
