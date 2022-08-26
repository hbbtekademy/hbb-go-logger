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
	logLevel = loglevel.DEBUG

	redirectLogOutputs()
	defer buf.Reset()

	Info("Test Info Log")
	expectedLog := "Test Info Log\n"
	actualLog := buf.String()

	if !strings.HasPrefix(actualLog, "INFO") {
		t.Errorf("Expected log to start with INFO instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}
}

func TestInfofLogging(t *testing.T) {
	logLevel = loglevel.INFO
	redirectLogOutputs()
	defer buf.Reset()

	Infof("Test Info Log: %d %s", 1, "str")
	expectedLog := "Test Info Log: 1 str\n"
	actualLog := buf.String()

	if !strings.HasPrefix(actualLog, "INFO") {
		t.Errorf("Expected log to start with INFO instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}
}

func TestDebugLogging(t *testing.T) {
	logLevel = loglevel.DEBUG

	redirectLogOutputs()
	defer buf.Reset()

	Debug("Test Debug Log")
	expectedLog := "Test Debug Log\n"
	actualLog := buf.String()

	if !strings.HasPrefix(actualLog, "DEBUG") {
		t.Errorf("Expected log to start with DEBUG instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}
}

func TestErrorLogging(t *testing.T) {
	logLevel = loglevel.ERROR

	redirectLogOutputs()
	defer buf.Reset()

	Error("Test Error Log")
	expectedLog := "Test Error Log\n"
	actualLog := buf.String()

	if !strings.HasPrefix(actualLog, "ERROR") {
		t.Errorf("Expected log to start with ERROR instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}
}

func TestDebugfLogging(t *testing.T) {
	logLevel = loglevel.DEBUG

	redirectLogOutputs()
	defer buf.Reset()

	Debugf("Test Debug Log: %d %s", 2, "debug")
	expectedLog := "Test Debug Log: 2 debug\n"
	actualLog := buf.String()

	if !strings.HasPrefix(actualLog, "DEBUG") {
		t.Errorf("Expected log to start with DEBUG instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}
}

func TestErrorfLogging(t *testing.T) {
	logLevel = loglevel.ERROR

	redirectLogOutputs()
	defer buf.Reset()

	Errorf("Test Errorf Log: %d %s", 2, "errorf")
	expectedLog := "Test Errorf Log: 2 errorf\n"
	actualLog := buf.String()

	if !strings.HasPrefix(actualLog, "ERROR") {
		t.Errorf("Expected log to start with ERROR instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}
}

func TestFatalLogging(t *testing.T) {
	if os.Getenv("TEST_FATAL_LOGGING") == "1" {
		Fatal("Test fatal log")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalLogging")
	cmd.Env = append(os.Environ(), "TEST_FATAL_LOGGING=1")
	cmdOutput, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	b, _ := ioutil.ReadAll(cmdOutput)
	actualLog := string(b)
	expectedLog := "Test fatal log\n"

	if !strings.HasPrefix(actualLog, "FATAL") {
		t.Errorf("Expected log to start with FATAL instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}

	err := cmd.Wait()
	if e, ok := err.(*exec.ExitError); !ok || e.Success() {
		t.Fatalf("Process ran with err %v, want exit status 1", err)
	}
}

func TestFatalfLogging(t *testing.T) {
	if os.Getenv("TEST_FATAL_LOGGING") == "1" {
		Fatalf("Test fatalf log: %d %s", 3, "fatalf")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalfLogging")
	cmd.Env = append(os.Environ(), "TEST_FATAL_LOGGING=1")
	cmdOutput, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	b, _ := ioutil.ReadAll(cmdOutput)
	actualLog := string(b)
	expectedLog := "Test fatalf log: 3 fatalf\n"

	if !strings.HasPrefix(actualLog, "FATAL") {
		t.Errorf("Expected log to start with FATAL instead got: %s", actualLog)
	}
	if !strings.HasSuffix(actualLog, expectedLog) {
		t.Errorf("Expected log to end with '%s' instead got: %s", expectedLog, actualLog)
	}

	err := cmd.Wait()
	if e, ok := err.(*exec.ExitError); !ok || e.Success() {
		t.Fatalf("Process ran with err %v, want exit status 1", err)
	}
}

func redirectLogOutputs() {
	if instance == nil {
		RedirectStderr(&buf)
		RedirectStdout(&buf)
	}
}
