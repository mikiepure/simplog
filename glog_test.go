package simplog

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGDefaultLevel(t *testing.T) {
	if !GFatal("test", "message") {
		t.Fatal("fatal log shall be sent as default")
	}
	if !GLog(LogLevelFatal, "test", "message") {
		t.Fatal("fatal log shall be sent as default")
	}

	if !GError("test", "message") {
		t.Fatal("error log shall be sent as default")
	}
	if !GLog(LogLevelError, "test", "message") {
		t.Fatal("error log shall be sent as default")
	}

	if !GWarn("test", "message") {
		t.Fatal("warn log shall be sent as default")
	}
	if !GLog(LogLevelWarn, "test", "message") {
		t.Fatal("warn log shall be sent as default")
	}

	if !GInfo("test", "message") {
		t.Fatal("info log shall be sent as default")
	}
	if !GLog(LogLevelInfo, "test", "message") {
		t.Fatal("info log shall be sent as default")
	}

	if GDebug("test", "message") {
		t.Fatal("debug log shall not be sent as default")
	}
	if GLog(LogLevelDebug, "test", "message") {
		t.Fatal("debug log shall not be sent as default")
	}
}

func TestGLevel(t *testing.T) {
	if GLevel() != LogLevelInfo {
		t.Fatal("default log level shall be INFO")
	}

	GSetLevel(LogLevelError)
	if GLevel() != LogLevelError {
		t.Fatal("failed to set log level of global logger")
	}

	if !GFatal("test", "message") {
		t.Fatal("fatal log shall be sent by error level")
	}
	if !GError("test", "message") {
		t.Fatal("error log shall be sent by error level")
	}
	if GWarn("test", "message") {
		t.Fatal("warn log shall not be sent by error level")
	}
	if GInfo("test", "message") {
		t.Fatal("info log shall not be sent by error level")
	}
	if GDebug("test", "message") {
		t.Fatal("debug log shall not be sent by error level")
	}

	GSetLevel(LogLevelDebug)
	if GLevel() != LogLevelDebug {
		t.Fatal("failed to set log level of global logger")
	}

	if !GFatal("test", "message") {
		t.Fatal("fatal log shall be sent by debug level")
	}
	if !GError("test", "message") {
		t.Fatal("error log shall be sent by debug level")
	}
	if !GWarn("test", "message") {
		t.Fatal("warn log shall be sent by debug level")
	}
	if !GInfo("test", "message") {
		t.Fatal("info log shall be sent by debug level")
	}
	if !GDebug("test", "message") {
		t.Fatal("debug log shall be sent by debug level")
	}

	GSetLevel(LogLevelInfo)
	if GLevel() != LogLevelInfo {
		t.Fatal("failed to set log level of global logger")
	}
}

func TestGUnknownLevel(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("panic is not occurred by undefined log level")
		}
	}()

	GLog(10, "undefined level message")
}

func TestGWriter(t *testing.T) {
	if GWriter() != os.Stdout {
		t.Fatal("default writer shall be standard output")
	}

	buf := new(bytes.Buffer)
	GSetWriter(buf)
	if GWriter() != buf {
		t.Fatal("failed to set writer of global logger")
	}

	GInfo("test", "message")
	if !strings.Contains(buf.String(), "INFO. test message") {
		t.Fatal("failed to output log to changed writer")
	}

	GSetWriter(os.Stdout)
	if GWriter() != os.Stdout {
		t.Fatal("failed to set writer of global logger")
	}
}

func TestGFormatter(t *testing.T) {
	GSetFormatter(FormatFunc(func(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, v ...interface{}) string {
		msglist := []string{}
		msglist = append(msglist, "["+level.String()[:1]+"]")
		logmsg := fmt.Sprintln(v...)
		msglist = append(msglist, logmsg[:len(logmsg)-1])
		return strings.Join(msglist, " ")
	}))

	buf := new(bytes.Buffer)
	GSetWriter(buf)
	if GWriter() != buf {
		t.Fatal("failed to set writer of global logger")
	}

	GInfo("test", "message")
	if !strings.Contains(buf.String(), "[I] test message") {
		t.Fatal("failed to output log to changed writer")
	}

	GSetWriter(os.Stdout)
	if GWriter() != os.Stdout {
		t.Fatal("failed to set writer of global logger")
	}

	GSetFormatter(FormatFunc(FormatDefault))
}
