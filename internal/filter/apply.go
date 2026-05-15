package filter

import (
	"bufio"
	"fmt"
	"io"
)

// ApplyOptions controls the behaviour of Apply.
type ApplyOptions struct {
	// InvertMatch causes Apply to retain lines that do NOT match the filter,
	// analogous to grep -v.
	InvertMatch bool
}

// Apply reads lines from r, passes each through f (respecting opts), and
// writes matching lines to w. It returns the number of lines written and any
// error encountered during reading or writing.
func Apply(r io.Reader, w io.Writer, f Filter, opts ApplyOptions) (int, error) {
	scanner := bufio.NewScanner(r)
	written := 0
	for scanner.Scan() {
		line := scanner.Text()
		matched := f.Match(line)
		if opts.InvertMatch {
			matched = !matched
		}
		if !matched {
			continue
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return written, err
		}
		written++
	}
	if err := scanner.Err(); err != nil {
		return written, err
	}
	return written, nil
}
