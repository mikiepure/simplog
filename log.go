package simplog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// LogLevel defines severity and priority of a log message.
type LogLevel uint8

const (
	// LogLevelFatal ...
	LogLevelFatal LogLevel = iota
	// LogLevelError ...
	LogLevelError
	// LogLevelWarn ...
	LogLevelWarn
	// LogLevelInfo ...
	LogLevelInfo
	// LogLevelDebug ...
	LogLevelDebug
)

func (p LogLevel) String() string {
	switch p {
	case LogLevelFatal:
		return "FATAL"
	case LogLevelError:
		return "ERROR"
	case LogLevelWarn:
		return "WARN."
	case LogLevelInfo:
		return "INFO."
	case LogLevelDebug:
		return "DEBUG"
	}

	panic("undefined log level")
}

// FuncFormat ...
type FuncFormat func(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, msg string) string

// Logger ...
type Logger struct {
	name       string
	level      LogLevel
	out        io.Writer
	funcFormat FuncFormat
}

// New ...
func New() *Logger {
	return &Logger{out: os.Stdout, level: LogLevelInfo, funcFormat: funcFormatDefault}
}

// LogLevel ...
func (p *Logger) LogLevel() LogLevel {
	return p.level
}

// SetLogLevel ...
func (p *Logger) SetLogLevel(level LogLevel) {
	p.level = level
}

// Fatal ...
func (p *Logger) Fatal(v ...interface{}) bool {
	return p.logging(LogLevelFatal, v...)
}

// Error ...
func (p *Logger) Error(v ...interface{}) bool {
	return p.logging(LogLevelError, v...)
}

// Warn ...
func (p *Logger) Warn(v ...interface{}) bool {
	return p.logging(LogLevelWarn, v...)
}

// Info ...
func (p *Logger) Info(v ...interface{}) bool {
	return p.logging(LogLevelInfo, v...)
}

// Debug ...
func (p *Logger) Debug(v ...interface{}) bool {
	return p.logging(LogLevelDebug, v...)
}

// Log ...
func (p *Logger) Log(level LogLevel, v ...interface{}) bool {
	return p.logging(level, v...)
}

func (p *Logger) logging(level LogLevel, v ...interface{}) bool {
	// filtering by log level
	if p.level < level {
		return false
	}

	// logging information
	time := time.Now()
	pc, filename, line, ok := runtime.Caller(2)
	funcname := ""
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
	} else {
		filename = ""
		line = 0
	}

	// format log msg
	msg := fmt.Sprintln(v...)
	s := p.funcFormat(p, level, time, funcname, filename, line, msg[:len(msg)-1])

	// output log msg to writer
	_, err := io.WriteString(p.out, s+"\n")
	return err == nil
}

func funcFormatDefault(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, msg string) string {
	logmsg := []string{time.Format("2006/01/02 15:04:05"), level.String() + ":", msg}
	if level < LogLevelWarn {
		logpos := "(" + filename + ":" + strconv.Itoa(line) + ")"
		logmsg = append(logmsg, logpos)
	}
	return strings.Join(logmsg, " ")
}
