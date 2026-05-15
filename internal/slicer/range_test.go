package slicer_test

import (
	"testing"
	"time"

	"github.com/yourusername/logslice/internal/slicer"
)

func TestNewTimeRange_Valid(t *testing.T) {
	start := mustParse("2024-01-15T10:00:00Z")
	end := mustParse("2024-01-15T11:00:00Z")
	tr, err := slicer.NewTimeRange(start, end)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tr.Duration() != time.Hour {
		t.Errorf("expected 1h duration, got %v", tr.Duration())
	}
}

func TestNewTimeRange_StartAfterEnd(t *testing.T) {
	start := mustParse("2024-01-15T11:00:00Z")
	end := mustParse("2024-01-15T10:00:00Z")
	_, err := slicer.NewTimeRange(start, end)
	if err == nil {
		t.Fatal("expected error for start after end")
	}
}

func TestNewTimeRange_ZeroValues(t *testing.T) {
	_, err := slicer.NewTimeRange(time.Time{}, mustParse("2024-01-15T10:00:00Z"))
	if err == nil {
		t.Fatal("expected error for zero start")
	}
	_, err = slicer.NewTimeRange(mustParse("2024-01-15T10:00:00Z"), time.Time{})
	if err == nil {
		t.Fatal("expected error for zero end")
	}
}

func TestTimeRange_Contains(t *testing.T) {
	tr, _ := slicer.NewTimeRange(
		mustParse("2024-01-15T10:00:00Z"),
		mustParse("2024-01-15T11:00:00Z"),
	)
	cases := []struct {
		input    string
		expected bool
	}{
		{"2024-01-15T10:00:00Z", true},
		{"2024-01-15T10:30:00Z", true},
		{"2024-01-15T11:00:00Z", true},
		{"2024-01-15T09:59:59Z", false},
		{"2024-01-15T11:00:01Z", false},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			got := tr.Contains(mustParse(c.input))
			if got != c.expected {
				t.Errorf("Contains(%s) = %v, want %v", c.input, got, c.expected)
			}
		})
	}
}
