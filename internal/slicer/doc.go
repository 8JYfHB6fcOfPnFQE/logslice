// Package slicer implements the core log slicing engine for logslice.
//
// The slicer reads log lines from an io.Reader, parses timestamps using the
// parser package, and writes lines that fall within the requested time range
// to an io.Writer.
//
// Basic usage:
//
//	s := slicer.New(slicer.Options{
//		Start: start,
//		End:   end,
//	})
//	n, err := s.Slice(inputFile, outputFile)
//
// Lines that cannot be parsed for a timestamp are skipped by default.
// Set IncludeUnparsed to true to pass them through unchanged.
package slicer
