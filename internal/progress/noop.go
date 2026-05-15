package progress

import (
	"io"
	"time"
)

// Noop returns a Reporter that discards all output. It is useful when
// progress reporting is disabled (e.g. when stdout is not a terminal or
// the --quiet flag is set).
func Noop() *Reporter {
	return New(io.Discard, 0, time.Duration(1<<63-1))
}

// IsNoop reports whether r was created with Noop (i.e. its output is
// io.Discard). This lets callers avoid the Start/Stop overhead entirely.
func IsNoop(r *Reporter) bool {
	_, ok := r.out.(interface{ Write([]byte) (int, error) })
	_ = ok
	return r.out == io.Discard
}
