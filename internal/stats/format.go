package stats

import (
	"fmt"
	"io"
	"strings"
)

// Summary writes a human-readable summary of r to w.
func Summary(w io.Writer, r Result) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "--- logslice summary ---\n")
	fmt.Fprintf(&sb, "  total lines   : %d\n", r.TotalLines)
	fmt.Fprintf(&sb, "  matched       : %d\n", r.MatchedLines)
	fmt.Fprintf(&sb, "  skipped       : %d\n", r.SkippedLines)
	fmt.Fprintf(&sb, "  unparsed      : %d\n", r.UnparsedLines)
	fmt.Fprintf(&sb, "  filtered      : %d\n", r.FilteredLines)

	if r.FirstMatch != nil {
		fmt.Fprintf(&sb, "  first match   : %s\n", r.FirstMatch.Format("2006-01-02T15:04:05Z07:00"))
	}
	if r.LastMatch != nil {
		fmt.Fprintf(&sb, "  last match    : %s\n", r.LastMatch.Format("2006-01-02T15:04:05Z07:00"))
	}

	fmt.Fprintf(&sb, "  elapsed       : %s\n", r.Elapsed.Round(1000).String())

	fmt.Fprint(w, sb.String())
}
