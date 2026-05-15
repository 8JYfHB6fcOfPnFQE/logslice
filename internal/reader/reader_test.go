package reader

import (
	"strings"
	"testing"
)

func TestLineReader_BasicIteration(t *testing.T) {
	input := "line one\nline two\nline three\n"
	lr := NewFromReader(strings.NewReader(input), "test")

	var lines []string
	for lr.Next() {
		lines = append(lines, lr.Line())
	}
	if err := lr.Err(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "line one" {
		t.Errorf("expected 'line one', got %q", lines[0])
	}
	if lines[2] != "line three" {
		t.Errorf("expected 'line three', got %q", lines[2])
	}
}

func TestLineReader_EmptyInput(t *testing.T) {
	lr := NewFromReader(strings.NewReader(""), "empty")
	var count int
	for lr.Next() {
		count++
	}
	if count != 0 {
		t.Errorf("expected 0 lines, got %d", count)
	}
	if err := lr.Err(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLineReader_Path(t *testing.T) {
	lr := NewFromReader(strings.NewReader("hello"), "/var/log/app.log")
	if lr.Path() != "/var/log/app.log" {
		t.Errorf("expected path '/var/log/app.log', got %q", lr.Path())
	}
}

func TestLineReader_SingleLine(t *testing.T) {
	lr := NewFromReader(strings.NewReader("only line"), "single")
	if !lr.Next() {
		t.Fatal("expected at least one line")
	}
	if lr.Line() != "only line" {
		t.Errorf("expected 'only line', got %q", lr.Line())
	}
	if lr.Next() {
		t.Error("expected no more lines")
	}
}

func TestLineReader_NewInvalidPath(t *testing.T) {
	_, err := New("/nonexistent/path/to/file.log")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}
