package parser

import (
	"time"
)

// AutoDetectResult holds the result of an auto-detection attempt.
type AutoDetectResult struct {
	// Format is the matched BuiltinFormat, or nil if none matched.
	Format *BuiltinFormat
	// Time is the parsed timestamp.
	Time time.Time
}

// AutoDetect attempts to parse line using each of the BuiltinFormats in order,
// returning the first successful match. If no format matches, ok is false.
func AutoDetect(line string) (result AutoDetectResult, ok bool) {
	for i := range BuiltinFormats {
		f := &BuiltinFormats[i]
		t, err := ParseWithLayout(line, f.Layout)
		if err == nil {
			return AutoDetectResult{Format: f, Time: t}, true
		}
	}
	return AutoDetectResult{}, false
}

// AutoDetectLayout probes the provided sample lines and returns the name of
// the BuiltinFormat that successfully parses the most of them. If no format
// matches any line, it returns an empty string and false.
func AutoDetectLayout(samples []string) (name string, ok bool) {
	if len(samples) == 0 {
		return "", false
	}

	scores := make(map[string]int, len(BuiltinFormats))
	for _, line := range samples {
		for _, f := range BuiltinFormats {
			_, err := ParseWithLayout(line, f.Layout)
			if err == nil {
				scores[f.Name]++
			}
		}
	}

	best := ""
	bestScore := 0
	for n, s := range scores {
		if s > bestScore {
			best = n
			bestScore = s
		}
	}

	if best == "" {
		return "", false
	}
	return best, true
}
