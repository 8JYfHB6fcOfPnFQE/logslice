// Package filter provides composable predicates for filtering log lines.
//
// Three filter types are available:
//
//	- GrepFilter: matches lines containing a fixed substring, with optional
//	  case-insensitive comparison.
//
//	- RegexpFilter: matches lines against a compiled regular expression.
//
//	- Chain: combines multiple Filter values with AND semantics so that all
//	  contained filters must match for a line to be retained.
//
// Filters are designed to be applied after time-range slicing so that only
// lines already within the desired time window are subject to further
// predicate evaluation.
//
// Example:
//
//	grep := filter.NewGrepFilter("ERROR", false)
//	re, _ := filter.NewRegexpFilter(`disk|memory`)
//	chain := filter.NewChain(grep, re)
//	if chain.Match(line) { /* keep line */ }
package filter
