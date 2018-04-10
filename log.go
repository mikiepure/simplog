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
	// LogLevelDebug ...
	LogLevelDebug LogLevel = iota
	// LogLevelInfo ...
	LogLevelInfo
	// LogLevelWarn ...
	LogLevelWarn
	// LogLevelError ...
	LogLevelError
	// LogLevelFatal ...
	LogLevelFatal
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

// Logger ...
type Logger struct {
	level     LogLevel
	writer    io.Writer
	formatter Formatter
}

// New ...
func New() *Logger {
	return &Logger{level: LogLevelInfo, writer: os.Stdout, formatter: DefaultFormatter{showTime: true, showLevel: true, showPositionLevel: LogLevelError}}
}

// Level ...
func (p *Logger) Level() LogLevel {
	return p.level
}

// SetLevel ...
func (p *Logger) SetLevel(level LogLevel) {
	p.level = level
}

// Writer ...
func (p *Logger) Writer() io.Writer {
	return p.writer
}

// SetWriter ...
func (p *Logger) SetWriter(writer io.Writer) {
	p.writer = writer
}

// Formatter ...
func (p *Logger) Formatter() Formatter {
	return p.formatter
}

// SetFormatter ...
func (p *Logger) SetFormatter(formatter Formatter) {
	p.formatter = formatter
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

func (p *Logger) log(level LogLevel, v ...interface{}) bool {
	if p.level > level {
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

	msg := p.formatter.Format(p, level, time, funcname, filename, line, v...)

	_, err := io.WriteString(p.writer, msg+"\n")
	return err == nil
}

// Formatter can change log message format as you like
type Formatter interface {
	Format(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, v ...interface{}) string
}

// DefaultFormatter provides default log format of simplog
type DefaultFormatter struct {
	showTime          bool
	showLevel         bool
	showPositionLevel LogLevel
}

// Format provides default log format of simplog
func (p DefaultFormatter) Format(logger *Logger, level LogLevel, time time.Time, funcname string, filename string, line int, v ...interface{}) string {
	msglist := []string{}

	if p.showTime {
		msglist = append(msglist, time.Format("2006/01/02 15:04:05"))
	}

	if p.showLevel {
		msglist = append(msglist, level.String())
	}

	logmsg := fmt.Sprintln(v...)
	msglist = append(msglist, logmsg[:len(logmsg)-1])

	if level >= p.showPositionLevel {
		msglist = append(msglist, "("+filename+":"+strconv.Itoa(line)+")")
	}

	return strings.Join(msglist, " ")
}
