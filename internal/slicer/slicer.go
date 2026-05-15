// Package slicer provides functionality for extracting time-range segments
// from log files based on parsed timestamps.
package slicer

import (
	"bufio"
	"io"
	"time"

	"github.com/yourusername/logslice/internal/parser"
)

// Options configures the behavior of the log slicer.
type Options struct {
	// Start is the beginning of the time range (inclusive).
	Start time.Time
	// End is the end of the time range (inclusive).
	End time.Time
	// IncludeUnparsed controls whether lines without a timestamp are included.
	IncludeUnparsed bool
}

// Slicer reads log lines and writes those within the specified time range.
type Slicer struct {
	opts Options
}

// New creates a new Slicer with the given options.
func New(opts Options) *Slicer {
	return &Slicer{opts: opts}
}

// Slice reads from r and writes matching log lines to w.
// It returns the number of lines written and any error encountered.
func (s *Slicer) Slice(r io.Reader, w io.Writer) (int, error) {
	scanner := bufio.NewScanner(r)
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	written := 0
	for scanner.Scan() {
		line := scanner.Text()
		t, ok := parser.ExtractTimestamp(line)
		if !ok {
			if s.opts.IncludeUnparsed {
				if _, err := bw.WriteString(line + "\n"); err != nil {
					return written, err
				}
				written++
			}
			continue
		}
		if (t.Equal(s.opts.Start) || t.After(s.opts.Start)) &&
			(t.Equal(s.opts.End) || t.Before(s.opts.End)) {
			if _, err := bw.WriteString(line + "\n"); err != nil {
				return written, err
			}
			written++
		}
	}
	return written, scanner.Err()
}
