package reader

import (
	"bufio"
	"io"
	"os"
)

// LineReader reads lines from a log file source.
type LineReader struct {
	scanner *bufio.Scanner
	path    string
}

// New creates a LineReader from a file path.
func New(path string) (*LineReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewFromReader(f, path), nil
}

// NewFromReader creates a LineReader from any io.Reader.
func NewFromReader(r io.Reader, path string) *LineReader {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	return &LineReader{
		scanner: scanner,
		path:    path,
	}
}

// Next advances to the next line. Returns false when done or on error.
func (lr *LineReader) Next() bool {
	return lr.scanner.Scan()
}

// Line returns the current line text.
func (lr *LineReader) Line() string {
	return lr.scanner.Text()
}

// Err returns any scanning error (excluding io.EOF).
func (lr *LineReader) Err() error {
	return lr.scanner.Err()
}

// Path returns the source path associated with this reader.
func (lr *LineReader) Path() string {
	return lr.path
}
