package simplog

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDefaultLevel(t *testing.T) {
	log := New()

	if !log.Fatal("test", "message") {
		t.Fatal("fatal log shall be sent as default")
	}
	if !log.Log(LogLevelFatal, "test", "message") {
		t.Fatal("fatal log shall be sent as default")
	}

	if !log.Error("test", "message") {
		t.Fatal("error log shall be sent as default")
	}
	if !log.Log(LogLevelError, "test", "message") {
		t.Fatal("error log shall be sent as default")
	}

	if !log.Warn("test", "message") {
		t.Fatal("warn log shall be sent as default")
	}
	if !log.Log(LogLevelWarn, "test", "message") {
		t.Fatal("warn log shall be sent as default")
	}

	if !log.Info("test", "message") {
		t.Fatal("info log shall be sent as default")
	}
	if !log.Log(LogLevelInfo, "test", "message") {
		t.Fatal("info log shall be sent as default")
	}

	if log.Debug("test", "message") {
		t.Fatal("debug log shall not be sent as default")
	}
	if log.Log(LogLevelDebug, "test", "message") {
		t.Fatal("debug log shall not be sent as default")
	}
}

func TestLevel(t *testing.T) {
	log := New()

	if log.Level() != LogLevelInfo {
		t.Fatal("default log level shall be INFO")
	}

	log.SetLevel(LogLevelError)
	if log.Level() != LogLevelError {
		t.Fatal("failed to set log level of global logger")
	}

	if !log.Fatal("test", "message") {
		t.Fatal("fatal log shall be sent by error level")
	}
	if !log.Error("test", "message") {
		t.Fatal("error log shall be sent by error level")
	}
	if log.Warn("test", "message") {
		t.Fatal("warn log shall not be sent by error level")
	}
	if log.Info("test", "message") {
		t.Fatal("info log shall not be sent by error level")
	}
	if log.Debug("test", "message") {
		t.Fatal("debug log shall not be sent by error level")
	}

	log.SetLevel(LogLevelDebug)
	if log.Level() != LogLevelDebug {
		t.Fatal("failed to set log level of global logger")
	}

	if !log.Fatal("test", "message") {
		t.Fatal("fatal log shall be sent by debug level")
	}
	if !log.Error("test", "message") {
		t.Fatal("error log shall be sent by debug level")
	}
	if !log.Warn("test", "message") {
		t.Fatal("warn log shall be sent by debug level")
	}
	if !log.Info("test", "message") {
		t.Fatal("info log shall be sent by debug level")
	}
	if !log.Debug("test", "message") {
		t.Fatal("debug log shall be sent by debug level")
	}

	log.SetLevel(LogLevelInfo)
	if log.Level() != LogLevelInfo {
		t.Fatal("failed to set log level of global logger")
	}
}

func TestUnknownLevel(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("panic is not occurred by undefined log level")
		}
	}()

	log := New()
	log.Log(10, "undefined level message")
}

func TestWriter(t *testing.T) {
	log := New()

	if log.Writer() != os.Stdout {
		t.Fatal("default writer shall be standard output")
	}

	buf := new(bytes.Buffer)
	log.SetWriter(buf)
	if log.Writer() != buf {
		t.Fatal("failed to set writer of global logger")
	}

	log.Info("test", "message")
	if !strings.Contains(buf.String(), "INFO. test message") {
		t.Fatal("failed to output log to changed writer")
	}

	log.SetWriter(os.Stdout)
	if log.Writer() != os.Stdout {
		t.Fatal("failed to set writer of global logger")
	}
}

func TestFormatter(t *testing.T) {
	log := New()

	log.SetFormatter(FormatFunc(func(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, v ...interface{}) string {
		msglist := []string{}
		msglist = append(msglist, "["+level.String()[:1]+"]")
		logmsg := fmt.Sprintln(v...)
		msglist = append(msglist, logmsg[:len(logmsg)-1])
		return strings.Join(msglist, " ")
	}))

	buf := new(bytes.Buffer)
	log.SetWriter(buf)
	if log.Writer() != buf {
		t.Fatal("failed to set writer of global logger")
	}

	log.Info("test", "message")
	if !strings.Contains(buf.String(), "[I] test message") {
		t.Fatal("failed to output log to changed writer")
	}

	log.SetWriter(os.Stdout)
	if log.Writer() != os.Stdout {
		t.Fatal("failed to set writer of global logger")
	}

	log.SetFormatter(FormatFunc(FormatDefault))
}
