package slicer_test

import (
	"strings"
	"testing"
	"time"

	"github.com/yourusername/logslice/internal/slicer"
)

func mustParse(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

const sampleLog = `2024-01-15T10:00:00Z INFO  server started
2024-01-15T10:01:00Z DEBUG request received
2024-01-15T10:02:00Z INFO  processing
2024-01-15T10:03:00Z WARN  slow query detected
2024-01-15T10:04:00Z ERROR connection failed
no timestamp line
2024-01-15T10:05:00Z INFO  recovered
`

func TestSlice_BasicRange(t *testing.T) {
	opts := slicer.Options{
		Start: mustParse("2024-01-15T10:01:00Z"),
		End:   mustParse("2024-01-15T10:03:00Z"),
	}
	s := slicer.New(opts)
	var out strings.Builder
	n, err := s.Slice(strings.NewReader(sampleLog), &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 3 {
		t.Errorf("expected 3 lines, got %d", n)
	}
	if !strings.Contains(out.String(), "DEBUG request received") {
		t.Error("expected DEBUG line in output")
	}
}

func TestSlice_IncludeUnparsed(t *testing.T) {
	opts := slicer.Options{
		Start:           mustParse("2024-01-15T10:00:00Z"),
		End:             mustParse("2024-01-15T10:05:00Z"),
		IncludeUnparsed: true,
	}
	s := slicer.New(opts)
	var out strings.Builder
	n, err := s.Slice(strings.NewReader(sampleLog), &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 7 {
		t.Errorf("expected 7 lines, got %d", n)
	}
}

func TestSlice_EmptyRange(t *testing.T) {
	opts := slicer.Options{
		Start: mustParse("2024-01-15T11:00:00Z"),
		End:   mustParse("2024-01-15T12:00:00Z"),
	}
	s := slicer.New(opts)
	var out strings.Builder
	n, err := s.Slice(strings.NewReader(sampleLog), &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 0 {
		t.Errorf("expected 0 lines, got %d", n)
	}
}
