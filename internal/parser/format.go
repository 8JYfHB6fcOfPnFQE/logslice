package parser

import "time"

// knownLayouts holds common log timestamp formats ordered by specificity.
// More specific (longer) formats are tried first to avoid partial matches.
var knownLayouts = []string{
	// RFC 3339 variants
	time.RFC3339Nano,
	time.RFC3339,

	// Common syslog / journald formats
	"2006-01-02T15:04:05.999999999",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02 15:04:05",

	// Apache / nginx combined log format
	"02/Jan/2006:15:04:05 -0700",

	// Syslog (BSD)
	"Jan  2 15:04:05",
	"Jan 02 15:04:05",

	// Date only
	"2006-01-02",
}

// Format represents a named timestamp layout.
type Format struct {
	Name   string
	Layout string
}

// BuiltinFormats returns all built-in timestamp formats supported by the parser.
func BuiltinFormats() []Format {
	formats := make([]Format, 0, len(knownLayouts))
	for _, l := range knownLayouts {
		formats = append(formats, Format{
			Name:   l,
			Layout: l,
		})
	}
	return formats
}

// ParseWithLayout attempts to parse s using the provided layout.
// It returns the parsed time and true on success, or the zero Time and false
// if s does not match the layout.
func ParseWithLayout(layout, s string) (time.Time, bool) {
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, false
	}
	return t, true
}
