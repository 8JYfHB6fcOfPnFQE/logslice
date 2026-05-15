package reader_test

import (
	"os"
	"testing"

	"github.com/yourorg/logslice/internal/reader"
)

func TestStatFile_ValidFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "logslice-*.log")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	_, _ = f.WriteString("hello log\n")
	f.Close()

	info, err := reader.StatFile(f.Name())
	if err != nil {
		t.Fatalf("StatFile: unexpected error: %v", err)
	}
	if info.Path != f.Name() {
		t.Errorf("Path = %q, want %q", info.Path, f.Name())
	}
	if info.Size != 10 {
		t.Errorf("Size = %d, want 10", info.Size)
	}
	if info.ModTime.IsZero() {
		t.Error("ModTime should not be zero")
	}
}

func TestStatFile_NotFound(t *testing.T) {
	_, err := reader.StatFile("/nonexistent/path/to/file.log")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestStatFile_Directory(t *testing.T) {
	dir := t.TempDir()
	_, err := reader.StatFile(dir)
	if err == nil {
		t.Fatal("expected error for directory, got nil")
	}
}

func TestFileInfo_IsReadable(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "logslice-*.log")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	f.Close()

	info, err := reader.StatFile(f.Name())
	if err != nil {
		t.Fatalf("StatFile: %v", err)
	}
	if !info.IsReadable() {
		t.Error("IsReadable() = false, want true for a normal temp file")
	}
}

func TestFileInfo_String(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "logslice-*.log")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	f.Close()

	info, err := reader.StatFile(f.Name())
	if err != nil {
		t.Fatalf("StatFile: %v", err)
	}
	s := info.String()
	if s == "" {
		t.Error("String() returned empty string")
	}
}
