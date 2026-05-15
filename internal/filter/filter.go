// Package filter provides line-level filtering capabilities for log entries.
package filter

import (
	"regexp"
	"strings"
)

// Filter defines a predicate applied to a raw log line.
type Filter interface {
	Match(line string) bool
}

// GrepFilter retains lines containing a fixed substring.
type GrepFilter struct {
	substring string
	caseSensitive bool
}

// NewGrepFilter creates a GrepFilter. If caseSensitive is false the match is
// performed on lower-cased input.
func NewGrepFilter(substring string, caseSensitive bool) *GrepFilter {
	if !caseSensitive {
		substring = strings.ToLower(substring)
	}
	return &GrepFilter{substring: substring, caseSensitive: caseSensitive}
}

// Match reports whether line contains the filter substring.
func (f *GrepFilter) Match(line string) bool {
	if !f.caseSensitive {
		line = strings.ToLower(line)
	}
	return strings.Contains(line, f.substring)
}

// RegexpFilter retains lines matching a compiled regular expression.
type RegexpFilter struct {
	re *regexp.Regexp
}

// NewRegexpFilter compiles pattern and returns a RegexpFilter.
// Returns an error if the pattern is invalid.
func NewRegexpFilter(pattern string) (*RegexpFilter, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RegexpFilter{re: re}, nil
}

// Match reports whether line matches the regular expression.
func (f *RegexpFilter) Match(line string) bool {
	return f.re.MatchString(line)
}

// Chain combines multiple filters with AND semantics: all must match.
type Chain struct {
	filters []Filter
}

// NewChain creates a Chain from the provided filters.
func NewChain(filters ...Filter) *Chain {
	return &Chain{filters: filters}
}

// Match returns true only when every contained filter matches line.
func (c *Chain) Match(line string) bool {
	for _, f := range c.filters {
		if !f.Match(line) {
			return false
		}
	}
	return true
}
