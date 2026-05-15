package filter_test

import (
	"testing"

	"github.com/logslice/logslice/internal/filter"
)

func TestGrepFilter_CaseSensitive(t *testing.T) {
	f := filter.NewGrepFilter("ERROR", true)
	if !f.Match("2024-01-01 ERROR something failed") {
		t.Error("expected match")
	}
	if f.Match("2024-01-01 error something failed") {
		t.Error("expected no match for lowercase with case-sensitive filter")
	}
}

func TestGrepFilter_CaseInsensitive(t *testing.T) {
	f := filter.NewGrepFilter("error", false)
	if !f.Match("2024-01-01 ERROR something failed") {
		t.Error("expected case-insensitive match")
	}
	if !f.Match("2024-01-01 error something failed") {
		t.Error("expected match")
	}
}

func TestGrepFilter_NoMatch(t *testing.T) {
	f := filter.NewGrepFilter("WARN", true)
	if f.Match("INFO all good") {
		t.Error("expected no match")
	}
}

func TestRegexpFilter_ValidPattern(t *testing.T) {
	f, err := filter.NewRegexpFilter(`\bERROR\b`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Match("2024-01-01 ERROR disk full") {
		t.Error("expected match")
	}
	if f.Match("2024-01-01 ERRORS disk full") {
		t.Error("expected no whole-word match")
	}
}

func TestRegexpFilter_InvalidPattern(t *testing.T) {
	_, err := filter.NewRegexpFilter(`[invalid`)
	if err == nil {
		t.Error("expected error for invalid regexp")
	}
}

func TestChain_AllMatch(t *testing.T) {
	f1 := filter.NewGrepFilter("ERROR", true)
	f2 := filter.NewGrepFilter("disk", false)
	c := filter.NewChain(f1, f2)
	if !c.Match("ERROR disk full") {
		t.Error("expected chain match")
	}
}

func TestChain_PartialMatch(t *testing.T) {
	f1 := filter.NewGrepFilter("ERROR", true)
	f2 := filter.NewGrepFilter("disk", true)
	c := filter.NewChain(f1, f2)
	if c.Match("ERROR memory full") {
		t.Error("expected chain to reject when second filter fails")
	}
}

func TestChain_Empty(t *testing.T) {
	c := filter.NewChain()
	if !c.Match("anything") {
		t.Error("empty chain should match everything")
	}
}
