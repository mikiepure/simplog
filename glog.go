package simplog

import "io"

var glogger = New()

// GLevel ...
func GLevel() LogLevel {
	return glogger.Level()
}

// GSetLevel ...
func GSetLevel(level LogLevel) {
	glogger.SetLevel(level)
}

// GWriter ...
func GWriter() io.Writer {
	return glogger.Writer()
}

// GSetWriter ...
func GSetWriter(writer io.Writer) {
	glogger.SetWriter(writer)
}

// GSetFormatter ...
func GSetFormatter(formatter Formatter) {
	glogger.SetFormatter(formatter)
}

// GFatal outputs a log message with level `Fatal`.
// The global logger defined in simplog package is used.
func GFatal(v ...interface{}) bool {
	return glogger.log(LogLevelFatal, v...)
}

// GError outputs a log message with level `Error`.
// The global logger defined in simplog package is used.
func GError(v ...interface{}) bool {
	return glogger.log(LogLevelError, v...)
}

// GWarn outputs a log message with level `Warn`.
// The global logger defined in simplog package is used.
func GWarn(v ...interface{}) bool {
	return glogger.log(LogLevelWarn, v...)
}

// GInfo outputs a log message with level `Info`.
// The global logger defined in simplog package is used.
func GInfo(v ...interface{}) bool {
	return glogger.log(LogLevelInfo, v...)
}

// GDebug outputs a log message with level `Debug`.
// The global logger defined in simplog package is used.
func GDebug(v ...interface{}) bool {
	return glogger.log(LogLevelDebug, v...)
}

// GLog outputs a log message with level `level`.
// The global logger defined in simplog package is used.
func GLog(level LogLevel, v ...interface{}) bool {
	return glogger.log(level, v...)
}
