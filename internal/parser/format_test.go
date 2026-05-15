package parser

import (
	"testing"
	"time"
)

func TestBuiltinFormats_NonEmpty(t *testing.T) {
	formats := BuiltinFormats()
	if len(formats) == 0 {
		t.Fatal("expected at least one built-in format, got none")
	}
}

func TestBuiltinFormats_NameMatchesLayout(t *testing.T) {
	for _, f := range BuiltinFormats() {
		if f.Name != f.Layout {
			t.Errorf("expected Name == Layout for built-in format, got Name=%q Layout=%q", f.Name, f.Layout)
		}
	}
}

func TestParseWithLayout_RFC3339(t *testing.T) {
	input := "2024-03-15T08:30:00Z"
	got, ok := ParseWithLayout(time.RFC3339, input)
	if !ok {
		t.Fatalf("expected successful parse for %q", input)
	}
	want := time.Date(2024, 3, 15, 8, 30, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseWithLayout_Invalid(t *testing.T) {
	_, ok := ParseWithLayout(time.RFC3339, "not-a-timestamp")
	if ok {
		t.Error("expected parse failure for invalid input, got success")
	}
}

func TestParseWithLayout_ApacheFormat(t *testing.T) {
	input := "15/Mar/2024:08:30:00 +0000"
	layout := "02/Jan/2006:15:04:05 -0700"
	got, ok := ParseWithLayout(layout, input)
	if !ok {
		t.Fatalf("expected successful parse for Apache log format %q", input)
	}
	if got.Year() != 2024 || got.Month() != time.March || got.Day() != 15 {
		t.Errorf("unexpected parsed date: %v", got)
	}
}

func TestParseWithLayout_DateOnly(t *testing.T) {
	input := "2024-03-15"
	got, ok := ParseWithLayout("2006-01-02", input)
	if !ok {
		t.Fatalf("expected successful parse for date-only input %q", input)
	}
	if got.Year() != 2024 || got.Month() != time.March || got.Day() != 15 {
		t.Errorf("unexpected parsed date: %v", got)
	}
}

func TestBuiltinFormats_ContainsRFC3339(t *testing.T) {
	found := false
	for _, f := range BuiltinFormats() {
		if f.Layout == time.RFC3339 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected RFC3339 to be present in built-in formats")
	}
}
