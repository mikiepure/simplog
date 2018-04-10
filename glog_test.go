package simplog

import "testing"

func TestDefaultLevel(t *testing.T) {
	if GFatal("test", "message") == false {
		t.Fatal("fatal log shall be sent as default")
	}
	if GError("test", "message") == false {
		t.Fatal("error log shall be sent as default")
	}
	if GWarn("test", "message") == false {
		t.Fatal("warn log shall be sent as default")
	}
	if GInfo("test", "message") == false {
		t.Fatal("info log shall be sent as default")
	}
	if GDebug("test", "message") == true {
		t.Fatal("debug log shall not be sent as default")
	}
}

func TestLevel(t *testing.T) {
	GSetLevel(LogLevelError)
	if GLevel() != LogLevelError {
		t.Fatal("failed to set log level of global logger")
	}

	if GFatal("test", "message") == false {
		t.Fatal("fatal log shall be sent by error level")
	}
	if GError("test", "message") == false {
		t.Fatal("error log shall be sent by error level")
	}
	if GWarn("test", "message") == true {
		t.Fatal("warn log shall not be sent by error level")
	}
	if GInfo("test", "message") == true {
		t.Fatal("info log shall not be sent by error level")
	}
	if GDebug("test", "message") == true {
		t.Fatal("debug log shall not be sent by error level")
	}

	GSetLevel(LogLevelInfo)
	if GLevel() != LogLevelInfo {
		t.Fatal("failed to set log level of global logger")
	}
}
