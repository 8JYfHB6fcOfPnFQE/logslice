package stats_test

import (
	"strings"
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/stats"
)

func TestSummary_ContainsFields(t *testing.T) {
	c := stats.NewCollector()

	ts := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	c.RecordTotal()
	c.RecordTotal()
	c.RecordMatched(ts)
	c.RecordSkipped()
	c.RecordUnparsed()
	c.RecordFiltered()

	r := c.Finish()

	var sb strings.Builder
	stats.Summary(&sb, r)
	out := sb.String()

	cases := []string{
		"total lines",
		"matched",
		"skipped",
		"unparsed",
		"filtered",
		"elapsed",
		"2024-06-15",
	}
	for _, want := range cases {
		if !strings.Contains(out, want) {
			t.Errorf("Summary output missing %q\ngot:\n%s", want, out)
		}
	}
}

func TestSummary_NoMatchTimestamps(t *testing.T) {
	c := stats.NewCollector()
	c.RecordTotal()
	c.RecordSkipped()
	r := c.Finish()

	var sb strings.Builder
	stats.Summary(&sb, r)
	out := sb.String()

	if strings.Contains(out, "first match") {
		t.Errorf("expected no first match line when there are no matches")
	}
	if strings.Contains(out, "last match") {
		t.Errorf("expected no last match line when there are no matches")
	}
}

func TestSummary_Header(t *testing.T) {
	c := stats.NewCollector()
	r := c.Finish()

	var sb strings.Builder
	stats.Summary(&sb, r)

	if !strings.Contains(sb.String(), "logslice summary") {
		t.Error("Summary output should contain header")
	}
}
