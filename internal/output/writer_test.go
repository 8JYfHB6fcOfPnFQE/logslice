package output_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/output"
)

func TestWriter_ToBuffer(t *testing.T) {
	var buf bytes.Buffer
	w := output.NewFromWriter(&buf)

	lines := []string{"line one", "line two", "line three"}
	for _, l := range lines {
		if err := w.WriteLine(l); err != nil {
			t.Fatalf("WriteLine(%q): unexpected error: %v", l, err)
		}
	}

	if err := w.Close(); err != nil {
		t.Fatalf("Close: unexpected error: %v", err)
	}

	got := buf.String()
	for _, l := range lines {
		if !strings.Contains(got, l) {
			t.Errorf("output missing line %q; got:\n%s", l, got)
		}
	}
}

func TestWriter_LinesWritten(t *testing.T) {
	var buf bytes.Buffer
	w := output.NewFromWriter(&buf)

	for i := 0; i < 5; i++ {
		_ = w.WriteLine("entry")
	}
	_ = w.Close()

	if got := w.LinesWritten(); got != 5 {
		t.Errorf("LinesWritten() = %d; want 5", got)
	}
}

func TestWriter_ToFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.log")

	w, err := output.New(path)
	if err != nil {
		t.Fatalf("New(%q): unexpected error: %v", path, err)
	}

	_ = w.WriteLine("hello from file")
	if err := w.Close(); err != nil {
		t.Fatalf("Close: unexpected error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !strings.Contains(string(data), "hello from file") {
		t.Errorf("file content missing expected line; got: %s", data)
	}
}

func TestWriter_InvalidPath(t *testing.T) {
	_, err := output.New("/nonexistent/dir/out.log")
	if err == nil {
		t.Error("expected error for invalid path, got nil")
	}
}

func TestWriter_StdoutAlias(t *testing.T) {
	// "-" and "" should not error — they resolve to stdout
	for _, alias := range []string{"-", ""} {
		w, err := output.New(alias)
		if err != nil {
			t.Errorf("New(%q): unexpected error: %v", alias, err)
			continue
		}
		// Do not close stdout; just verify construction succeeded.
		_ = w
	}
}
