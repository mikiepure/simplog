# simplog

Simple logging library for golang.

It has following features:

* Simple logging interfaces with log level
* Log can be filtered by log level
* Log destination can be changed by setting io.Writer
* Log format can be changed by callback function

## Logger

A local logger can be created by New() function.

It has following functions:

* Fatal(), Error(), Warn(), Info(), or Debug() can be used for logging
  * Arguments for logging are same as fmt.Println() or log.Println()
* SetLevel() can be used for filtering by log level
* SetWriter() can be used to change log destination
* SetFormatter() can be used to change log format

Sample Code:

```go
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mikiepure/simplog"
)

func main() {
	log := simplog.New()

	// output log to standard output (default log destination)
	log.Fatal("fatal log message")
	log.Error("error log message")
	log.Warn("warn log message")
	log.Info("info log message")
	log.Debug("debug log message") // level `Debug` is filtered as default

	fmt.Println("---")

	// set log level `Error`, so only `Fatal` and `Error` can be shown
	log.SetLevel(simplog.LogLevelError)
	log.Fatal("fatal log message")
	log.Error("error log message")
	log.Warn("warn log message")   // level `Warn` is filtered by level `Error`
	log.Info("info log message")   // level `Info` is filtered by level `Error`
	log.Debug("debug log message") // level `Debug` is filtered by level `Error`

	log.SetLevel(simplog.LogLevelInfo)
	fmt.Println("---")

	// change log destination to file "log.txt"
	logfile, err := os.Create("log.txt")
	if err != nil {
		log.Error("failed to create log file:", err.Error())
		return
	}
	defer logfile.Close()
	log.SetWriter(logfile)
	log.Info("write log message to file")

	log.SetWriter(os.Stdout)
	fmt.Println("---")

	// change log format
	log.SetFormatter(simplog.FormatFunc(func(logger *simplog.Logger, level simplog.LogLevel, time time.Time, funcname string, filename string, line int, v ...interface{}) string {
		msglist := []string{}
		msglist = append(msglist, "["+level.String()[:1]+"]")
		logmsg := fmt.Sprintln(v...)
		msglist = append(msglist, logmsg[:len(logmsg)-1])
		return strings.Join(msglist, " ")
	}))
	log.Info("customized simple log message")

	log.SetFormatter(simplog.FormatFunc(simplog.FormatDefault))
	fmt.Println("---")
}
```

Standard Output:

```
2018/04/11 22:42:30 FATAL fatal log message (c:/.../sample/main.go:16)
2018/04/11 22:42:30 ERROR error log message (c:/.../sample/main.go:17)
2018/04/11 22:42:30 WARN. warn log message
2018/04/11 22:42:30 INFO. info log message
---
2018/04/11 22:42:30 FATAL fatal log message (c:/.../sample/main.go:26)
2018/04/11 22:42:30 ERROR error log message (c:/.../sample/main.go:27)
---
---
[I] customized simple log message
---
```

File "log.txt":

```
2018/04/11 22:42:30 INFO. write log message to file
```

## Global Logger

The unique global logger is defined in simplog package and it can be used by exported functions, which has a prefix of "G".

A setting of the global logger can be changed by any other packages, so local logger should be used for a library or official application.

Sample Code:

```go
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/mikiepure/simplog"
)

func main() {
	// output log to standard output (default log destination)
	log.GFatal("fatal log message")
	log.GError("error log message")
	log.GWarn("warn log message")
	log.GInfo("info log message")
	log.GDebug("debug log message") // level `Debug` is filtered as default

	fmt.Println("---")

	// filtered by `Error`, so only `Fatal` and `Error` can be shown
	log.GSetLevel(log.LogLevelError)
	log.GFatal("fatal log message")
	log.GError("error log message")
	log.GWarn("warn log message")   // level `Warn` is filtered by level `Error`
	log.GInfo("info log message")   // level `Info` is filtered by level `Error`
	log.GDebug("debug log message") // level `Debug` is filtered by level `Error`

	log.GSetLevel(log.LogLevelInfo)
	fmt.Println("---")

	// change log destination to file "log.txt"
	logfile, err := os.Create("log.txt")
	if err != nil {
		log.GError("failed to create log file:", err.Error())
		return
	}
	defer logfile.Close()
	log.GSetWriter(logfile)
	log.GInfo("write log message to file")

	log.GSetWriter(os.Stdout)
	fmt.Println("---")

	// change log format
	log.GSetFormatter(log.FormatFunc(func(logger *log.Logger, level log.LogLevel, time time.Time, funcname string, filename string, line int, v ...interface{}) string {
		msglist := []string{}
		msglist = append(msglist, "["+level.String()[:1]+"]")
		logmsg := fmt.Sprintln(v...)
		msglist = append(msglist, logmsg[:len(logmsg)-1])
		return strings.Join(msglist, " ")
	}))
	log.GInfo("customized simple log message")

	log.GSetFormatter(log.FormatFunc(log.FormatDefault))
	fmt.Println("---")
}

Standard Output and File "log.txt" are same as Logger.
