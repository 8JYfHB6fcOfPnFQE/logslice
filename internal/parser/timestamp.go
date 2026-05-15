package parser

import (
	"errors"
	"regexp"
	"time"
)

// Common log timestamp formats to attempt parsing
var timestampFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05.000",
	"2006-01-02 15:04:05,000",
	"02/Jan/2006:15:04:05 -0700",
	"Jan 02 15:04:05",
	"Jan  2 15:04:05",
}

// timestampRegex captures a likely timestamp substring from a log line
var timestampRegex = regexp.MustCompile(
	`(\d{4}-\d{2}-\d{2}[T ]\d{2}:\d{2}:\d{2}[.,]?\d*(?:Z|[+-]\d{2}:?\d{2})?|` +
		`\d{2}/\w+/\d{4}:\d{2}:\d{2}:\d{2} [+-]\d{4}|` +
		`\w{3}\s+\d{1,2} \d{2}:\d{2}:\d{2})`,
)

// ErrNoTimestamp is returned when no recognisable timestamp is found in a line.
var ErrNoTimestamp = errors.New("no timestamp found in log line")

// ExtractTimestamp attempts to parse a timestamp from a raw log line.
// It returns the parsed time and the byte offset where the match was found.
func ExtractTimestamp(line string) (time.Time, int, error) {
	loc := timestampRegex.FindStringIndex(line)
	if loc == nil {
		return time.Time{}, -1, ErrNoTimestamp
	}

	candidate := line[loc[0]:loc[1]]

	for _, format := range timestampFormats {
		t, err := time.Parse(format, candidate)
		if err == nil {
			return t, loc[0], nil
		}
	}

	return time.Time{}, -1, ErrNoTimestamp
}
