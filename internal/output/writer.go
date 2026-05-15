// Package output handles writing sliced log segments to various destinations.
package output

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Writer writes log lines to an output destination.
type Writer struct {
	w       *bufio.Writer
	closer  io.Closer
	count   int
}

// New creates a Writer that writes to the given file path.
// If path is "-" or empty, it writes to stdout.
func New(path string) (*Writer, error) {
	if path == "" || path == "-" {
		return NewFromWriter(os.Stdout), nil
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("output: create file %q: %w", path, err)
	}

	return &Writer{
		w:      bufio.NewWriter(f),
		closer: f,
	}, nil
}

// NewFromWriter creates a Writer that writes to an existing io.Writer.
func NewFromWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// WriteLine writes a single log line followed by a newline character.
func (w *Writer) WriteLine(line string) error {
	if _, err := fmt.Fprintln(w.w, line); err != nil {
		return fmt.Errorf("output: write line: %w", err)
	}
	w.count++
	return nil
}

// LinesWritten returns the number of lines written so far.
func (w *Writer) LinesWritten() int {
	return w.count
}

// Close flushes any buffered data and closes the underlying file if applicable.
func (w *Writer) Close() error {
	if err := w.w.Flush(); err != nil {
		return fmt.Errorf("output: flush: %w", err)
	}
	if w.closer != nil {
		return w.closer.Close()
	}
	return nil
}
