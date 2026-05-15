package parser

import (
	"testing"
)

func TestAutoDetect_RFC3339Line(t *testing.T) {
	line := "2024-03-15T10:22:01Z INFO server started"
	res, ok := AutoDetect(line)
	if !ok {
		t.Fatal("expected AutoDetect to succeed for RFC3339 line")
	}
	if res.Format == nil {
		t.Fatal("expected non-nil Format")
	}
	if res.Time.IsZero() {
		t.Error("expected non-zero time")
	}
}

func TestAutoDetect_NoMatch(t *testing.T) {
	line := "this line has no recognizable timestamp at all"
	_, ok := AutoDetect(line)
	if ok {
		t.Error("expected AutoDetect to return ok=false for unrecognized line")
	}
}

func TestAutoDetect_ApacheLine(t *testing.T) {
	line := "127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326"
	res, ok := AutoDetect(line)
	if !ok {
		t.Fatalf("expected AutoDetect to succeed for Apache log line")
	}
	if res.Format == nil || res.Format.Name != "apache" {
		t.Errorf("expected format name 'apache', got %v", res.Format)
	}
}

func TestAutoDetectLayout_MajorityWins(t *testing.T) {
	samples := []string{
		"2024-03-15T10:22:01Z INFO msg1",
		"2024-03-15T10:22:02Z INFO msg2",
		"2024-03-15T10:22:03Z INFO msg3",
		"this line has no timestamp",
	}
	name, ok := AutoDetectLayout(samples)
	if !ok {
		t.Fatal("expected AutoDetectLayout to succeed")
	}
	if name == "" {
		t.Error("expected non-empty layout name")
	}
}

func TestAutoDetectLayout_EmptySamples(t *testing.T) {
	_, ok := AutoDetectLayout(nil)
	if ok {
		t.Error("expected ok=false for empty samples")
	}
}

func TestAutoDetectLayout_NoMatchingSamples(t *testing.T) {
	samples := []string{
		"no timestamp here",
		"another plain line",
	}
	_, ok := AutoDetectLayout(samples)
	if ok {
		t.Error("expected ok=false when no samples match any format")
	}
}
