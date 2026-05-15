package main

import (
	"testing"
	"time"
)

func TestParseTime_RFC3339(t *testing.T) {
	input := "2024-03-15T10:30:00Z"
	got, err := parseTime(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseTime_CustomLayout(t *testing.T) {
	input := "2024-03-15T10:30:00"
	got, err := parseTime(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Only check date/time components, not location.
	if got.Year() != 2024 || got.Month() != 3 || got.Day() != 15 {
		t.Errorf("unexpected date components: %v", got)
	}
	if got.Hour() != 10 || got.Minute() != 30 || got.Second() != 0 {
		t.Errorf("unexpected time components: %v", got)
	}
}

func TestParseTime_Invalid(t *testing.T) {
	cases := []string{
		"",
		"not-a-time",
		"2024/03/15",
		"15-03-2024T10:30:00",
	}
	for _, tc := range cases {
		_, err := parseTime(tc)
		if err == nil {
			t.Errorf("expected error for input %q, got nil", tc)
		}
	}
}

func TestParseTime_RFC3339WithOffset(t *testing.T) {
	input := "2024-06-01T08:00:00+02:00"
	got, err := parseTime(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, offset := got.Zone()
	if offset != 2*60*60 {
		t.Errorf("expected +02:00 offset, got offset=%d", offset)
	}
}
