// Package progress provides a simple progress reporter for tracking
// bytes read and lines processed during log slicing operations.
package progress

import (
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

// Reporter tracks and reports progress of a log slicing operation.
type Reporter struct {
	total    int64 // total bytes expected (0 = unknown)
	read     atomic.Int64
	lines    atomic.Int64
	matched  atomic.Int64
	start    time.Time
	out      io.Writer
	interval time.Duration
	stop     chan struct{}
}

// New creates a Reporter that writes updates to out every interval.
// Set total to 0 if the file size is unknown.
func New(out io.Writer, total int64, interval time.Duration) *Reporter {
	return &Reporter{
		total:    total,
		out:      out,
		interval: interval,
		start:    time.Now(),
		stop:     make(chan struct{}),
	}
}

// AddBytes records that n bytes have been read.
func (r *Reporter) AddBytes(n int64) { r.read.Add(n) }

// AddLine records a processed line; matched indicates it fell within the range.
func (r *Reporter) AddLine(matched bool) {
	r.lines.Add(1)
	if matched {
		r.matched.Add(1)
	}
}

// Start begins background reporting; call Stop when done.
func (r *Reporter) Start() {
	go func() {
		ticker := time.NewTicker(r.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				r.print(false)
			case <-r.stop:
				r.print(true)
				return
			}
		}
	}()
}

// Stop halts background reporting and prints a final summary line.
func (r *Reporter) Stop() { close(r.stop) }

func (r *Reporter) print(final bool) {
	read := r.read.Load()
	lines := r.lines.Load()
	matched := r.matched.Load()
	elapsed := time.Since(r.start).Truncate(time.Millisecond)

	if r.total > 0 {
		pct := float64(read) / float64(r.total) * 100
		fmt.Fprintf(r.out, "\rprogress: %.1f%% (%d/%d bytes) lines=%d matched=%d elapsed=%s",
			pct, read, r.total, lines, matched, elapsed)
	} else {
		fmt.Fprintf(r.out, "\rprogress: %d bytes lines=%d matched=%d elapsed=%s",
			read, lines, matched, elapsed)
	}
	if final {
		fmt.Fprintln(r.out)
	}
}
