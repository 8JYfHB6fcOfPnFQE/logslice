// Package parser provides utilities for detecting and extracting timestamps
// from unstructured and semi-structured log lines.
//
// It supports a wide range of common log timestamp formats including:
//   - RFC 3339 / ISO 8601 (with and without timezone)
//   - Datetime with space separator and optional milliseconds
//   - Nginx / Apache combined log format
//   - Syslog (BSD) format
//
// Usage:
//
//	t, offset, err := parser.ExtractTimestamp(line)
//	if errors.Is(err, parser.ErrNoTimestamp) {
//		// line has no detectable timestamp — skip or treat as continuation
//	}
//
// The parser is intentionally lenient: it attempts every known format and
// returns the first successful parse. Callers that need strict format
// enforcement should validate the returned time against their expected range.
package parser
