package stats_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/stats"
)

func TestCollector_Counts(t *testing.T) {
	c := stats.NewCollector()

	c.RecordTotal()
	c.RecordTotal()
	c.RecordTotal()

	ts := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	c.RecordMatched(ts)
	c.RecordMatched(ts.Add(time.Hour))

	c.RecordSkipped()
	c.RecordUnparsed()
	c.RecordFiltered()

	r := c.Finish()

	if r.TotalLines != 3 {
		t.Errorf("TotalLines: got %d, want 3", r.TotalLines)
	}
	if r.MatchedLines != 2 {
		t.Errorf("MatchedLines: got %d, want 2", r.MatchedLines)
	}
	if r.SkippedLines != 1 {
		t.Errorf("SkippedLines: got %d, want 1", r.SkippedLines)
	}
	if r.UnparsedLines != 1 {
		t.Errorf("UnparsedLines: got %d, want 1", r.UnparsedLines)
	}
	if r.FilteredLines != 1 {
		t.Errorf("FilteredLines: got %d, want 1", r.FilteredLines)
	}
}

func TestCollector_FirstLastMatch(t *testing.T) {
	c := stats.NewCollector()

	t1 := time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
	t3 := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)

	c.RecordMatched(t1)
	c.RecordMatched(t2)
	c.RecordMatched(t3)

	r := c.Finish()

	if r.FirstMatch == nil || !r.FirstMatch.Equal(t1) {
		t.Errorf("FirstMatch: got %v, want %v", r.FirstMatch, t1)
	}
	if r.LastMatch == nil || !r.LastMatch.Equal(t3) {
		t.Errorf("LastMatch: got %v, want %v", r.LastMatch, t3)
	}
}

func TestCollector_NoMatches(t *testing.T) {
	c := stats.NewCollector()
	c.RecordTotal()
	c.RecordSkipped()
	r := c.Finish()

	if r.FirstMatch != nil {
		t.Errorf("FirstMatch should be nil when no matches")
	}
	if r.LastMatch != nil {
		t.Errorf("LastMatch should be nil when no matches")
	}
}

func TestCollector_ElapsedPositive(t *testing.T) {
	c := stats.NewCollector()
	time.Sleep(time.Millisecond)
	r := c.Finish()

	if r.Elapsed <= 0 {
		t.Errorf("Elapsed should be positive, got %v", r.Elapsed)
	}
}
