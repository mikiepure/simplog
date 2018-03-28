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

// FormatFunc ...
type FormatFunc func(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, msg string) string

// Logger ...
type Logger struct {
	name   string
	level  LogLevel
	out    io.Writer
	format FormatFunc
}

// New ...
func New() *Logger {
	return &Logger{out: os.Stdout, level: LogLevelInfo, format: FormatDefault}
}

// Level ...
func (p *Logger) Level() LogLevel {
	return p.level
}

// SetLevel ...
func (p *Logger) SetLevel(level LogLevel) {
	p.level = level
}

// Fatal outputs a log message with level `Fatal`.
func (p *Logger) Fatal(v ...interface{}) bool {
	return p.log(LogLevelFatal, v...)
}

// Error outputs a log message with level `Error`.
func (p *Logger) Error(v ...interface{}) bool {
	return p.log(LogLevelError, v...)
}

// Warn outputs a log message with level `Warn`.
func (p *Logger) Warn(v ...interface{}) bool {
	return p.log(LogLevelWarn, v...)
}

// Info outputs a log message with level `Info`.
func (p *Logger) Info(v ...interface{}) bool {
	return p.log(LogLevelInfo, v...)
}

// Debug outputs a log message with level `Debug`.
func (p *Logger) Debug(v ...interface{}) bool {
	return p.log(LogLevelDebug, v...)
}

// Log outputs a log message with level `level`.
func (p *Logger) Log(level LogLevel, v ...interface{}) bool {
	return p.log(level, v...)
}

// FormatDefault provides function of default log format
func FormatDefault(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, msg string) string {
	logmsg := []string{time.Format("2006/01/02 15:04:05"), level.String() + ":", msg}
	if level < LogLevelWarn {
		logpos := "(" + filename + ":" + strconv.Itoa(line) + ")"
		logmsg = append(logmsg, logpos)
	}
	return strings.Join(logmsg, " ")
}

func (p *Logger) log(level LogLevel, v ...interface{}) bool {
	if p.level < level {
		return false
	}

	time := time.Now()
	pc, filename, line, ok := runtime.Caller(2)
	funcname := ""
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
	} else {
		filename = ""
		line = 0
	}

	msg := fmt.Sprintln(v...)
	s := p.format(p, level, time, funcname, filename, line, msg[:len(msg)-1])

	_, err := io.WriteString(p.out, s+"\n")
	return err == nil
}
