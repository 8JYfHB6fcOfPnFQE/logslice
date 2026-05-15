package slicer

import (
	"errors"
	"time"
)

// TimeRange represents a closed interval [Start, End].
type TimeRange struct {
	Start time.Time
	End   time.Time
}

// NewTimeRange creates a validated TimeRange.
// It returns an error if start is after end or if either value is zero.
func NewTimeRange(start, end time.Time) (TimeRange, error) {
	if start.IsZero() {
		return TimeRange{}, errors.New("start time must not be zero")
	}
	if end.IsZero() {
		return TimeRange{}, errors.New("end time must not be zero")
	}
	if start.After(end) {
		return TimeRange{}, errors.New("start time must not be after end time")
	}
	return TimeRange{Start: start, End: end}, nil
}

// Contains reports whether t falls within the time range (inclusive).
func (tr TimeRange) Contains(t time.Time) bool {
	return (t.Equal(tr.Start) || t.After(tr.Start)) &&
		(t.Equal(tr.End) || t.Before(tr.End))
}

// Duration returns the duration of the time range.
func (tr TimeRange) Duration() time.Duration {
	return tr.End.Sub(tr.Start)
}

// String returns a human-readable representation of the range.
func (tr TimeRange) String() string {
	return tr.Start.Format(time.RFC3339) + " to " + tr.End.Format(time.RFC3339)
}
