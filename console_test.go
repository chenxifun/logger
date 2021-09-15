package logger

import (
	"testing"
)

// Try each log level in decreasing order of priority.
func testConsoleCalls(bl *LocalLogger) {
	bl.Emer(nil, "emergency")
	bl.Alert(nil, "alert")
	bl.Crit(nil, "critical")
	bl.Error(nil, "error")
	bl.Warn(nil, "warning")
	bl.Debug(nil, "notice")
	bl.Info(nil, "informational")
	bl.Trace(nil, "trace")
}

func TestConsole(t *testing.T) {
	log1 := NewLogger()
	log1.SetLogger("console", "")
	testConsoleCalls(log1)

	log2 := NewLogger()
	log2.SetLogger("console", `{"level":"EROR"}`)
	testConsoleCalls(log2)
}

// Test console without color
func TestNoColorConsole(t *testing.T) {
	log := NewLogger()
	log.SetLogger("console", `{"color":false}`)
	testConsoleCalls(log)
}
