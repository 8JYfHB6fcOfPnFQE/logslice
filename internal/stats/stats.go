package stats

import "time"

// Result holds statistics collected during a slice operation.
type Result struct {
	// TotalLines is the total number of lines read from the input.
	TotalLines int

	// MatchedLines is the number of lines that fell within the time range.
	MatchedLines int

	// SkippedLines is the number of lines that were outside the time range.
	SkippedLines int

	// UnparsedLines is the number of lines whose timestamps could not be parsed.
	UnparsedLines int

	// FilteredLines is the number of lines removed by an optional filter chain.
	FilteredLines int

	// FirstMatch is the timestamp of the first matched line, if any.
	FirstMatch *time.Time

	// LastMatch is the timestamp of the last matched line, if any.
	LastMatch *time.Time

	// Elapsed is the wall-clock duration of the slice operation.
	Elapsed time.Duration
}

// Collector accumulates statistics during a slice run.
type Collector struct {
	start time.Time
	r     Result
}

// NewCollector creates a Collector and starts the elapsed timer.
func NewCollector() *Collector {
	return &Collector{start: time.Now()}
}

// RecordTotal increments the total line counter.
func (c *Collector) RecordTotal() { c.r.TotalLines++ }

// RecordMatched increments the matched counter and tracks first/last timestamps.
func (c *Collector) RecordMatched(ts time.Time) {
	c.r.MatchedLines++
	if c.r.FirstMatch == nil {
		t := ts
		c.r.FirstMatch = &t
	}
	t := ts
	c.r.LastMatch = &t
}

// RecordSkipped increments the skipped counter.
func (c *Collector) RecordSkipped() { c.r.SkippedLines++ }

// RecordUnparsed increments the unparsed counter.
func (c *Collector) RecordUnparsed() { c.r.UnparsedLines++ }

// RecordFiltered increments the filtered counter.
func (c *Collector) RecordFiltered() { c.r.FilteredLines++ }

// Finish stops the elapsed timer and returns the final Result.
func (c *Collector) Finish() Result {
	c.r.Elapsed = time.Since(c.start)
	return c.r
}
