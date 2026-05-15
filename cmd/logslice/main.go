// Command logslice extracts time-range segments from large log files.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/yourorg/logslice/internal/output"
	"github.com/yourorg/logslice/internal/reader"
	"github.com/yourorg/logslice/internal/slicer"
)

const timeLayout = "2006-01-02T15:04:05"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "logslice: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		start          = flag.String("start", "", "start timestamp (RFC3339 or 2006-01-02T15:04:05)")
		end            = flag.String("end", "", "end timestamp (RFC3339 or 2006-01-02T15:04:05)")
		outPath        = flag.String("out", "", "output file path (default: stdout)")
		includeUnparsed = flag.Bool("include-unparsed", false, "include lines with no parseable timestamp")
	)
	flag.Usage = usage
	flag.Parse()

	if *start == "" || *end == "" {
		flag.Usage()
		return fmt.Errorf("--start and --end are required")
	}

	if flag.NArg() < 1 {
		flag.Usage()
		return fmt.Errorf("at least one input file is required")
	}

	startTime, err := parseTime(*start)
	if err != nil {
		return fmt.Errorf("invalid --start: %w", err)
	}
	endTime, err := parseTime(*end)
	if err != nil {
		return fmt.Errorf("invalid --end: %w", err)
	}

	tr, err := slicer.NewTimeRange(startTime, endTime)
	if err != nil {
		return fmt.Errorf("invalid time range: %w", err)
	}

	w, err := output.New(*outPath)
	if err != nil {
		return fmt.Errorf("opening output: %w", err)
	}
	defer w.Close()

	for _, path := range flag.Args() {
		r, err := reader.New(path)
		if err != nil {
			return fmt.Errorf("opening %s: %w", path, err)
		}
		s := slicer.New(r, w, tr)
		s.IncludeUnparsed = *includeUnparsed
		if err := s.Slice(); err != nil {
			r.Close()
			return fmt.Errorf("slicing %s: %w", path, err)
		}
		r.Close()
	}
	return nil
}

func parseTime(s string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}
	return time.ParseInLocation(timeLayout, s, time.Local)
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: logslice [flags] <file> [file ...]\n\nFlags:\n`)
	flag.PrintDefaults()
}
