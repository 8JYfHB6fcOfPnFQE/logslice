package progress

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestReporter_AddBytesAndLines(t *testing.T) {
	var buf bytes.Buffer
	r := New(&buf, 1000, time.Hour) // long interval so background never fires

	r.AddBytes(200)
	r.AddBytes(300)
	r.AddLine(true)
	r.AddLine(false)
	r.AddLine(true)

	if got := r.read.Load(); got != 500 {
		t.Errorf("read bytes: want 500, got %d", got)
	}
	if got := r.lines.Load(); got != 3 {
		t.Errorf("lines: want 3, got %d", got)
	}
	if got := r.matched.Load(); got != 2 {
		t.Errorf("matched: want 2, got %d", got)
	}
}

func TestReporter_PrintWithTotal(t *testing.T) {
	var buf bytes.Buffer
	r := New(&buf, 1000, time.Hour)
	r.AddBytes(500)
	r.AddLine(true)

	r.print(true)
	out := buf.String()

	for _, want := range []string{"50.0%", "500/1000", "lines=1", "matched=1"} {
		if !strings.Contains(out, want) {
			t.Errorf("output %q missing %q", out, want)
		}
	}
}

func TestReporter_PrintWithoutTotal(t *testing.T) {
	var buf bytes.Buffer
	r := New(&buf, 0, time.Hour)
	r.AddBytes(128)

	r.print(true)
	out := buf.String()

	if !strings.Contains(out, "128 bytes") {
		t.Errorf("output %q missing byte count", out)
	}
	if strings.Contains(out, "%") {
		t.Errorf("output %q should not contain percentage without total", out)
	}
}

func TestReporter_StartStop(t *testing.T) {
	var buf bytes.Buffer
	r := New(&buf, 500, 20*time.Millisecond)
	r.AddBytes(250)
	r.AddLine(true)

	r.Start()
	time.Sleep(55 * time.Millisecond) // allow ~2 ticks
	r.Stop()
	time.Sleep(10 * time.Millisecond) // let goroutine finish

	out := buf.String()
	if !strings.Contains(out, "lines=") {
		t.Errorf("expected progress output, got: %q", out)
	}
	// final call should end with newline
	if !strings.HasSuffix(out, "\n") {
		t.Errorf("expected trailing newline in output")
	}
}
