// Package progress provides lightweight progress reporting for logslice
// operations.
//
// A Reporter is created with a target byte count (or 0 for unknown size),
// an output writer, and a reporting interval. It runs in the background,
// printing a single updating line to the writer at each tick.
//
// Typical usage:
//
//	rep := progress.New(os.Stderr, fileSize, time.Second)
//	rep.Start()
//	defer rep.Stop()
//
//	// inside read loop:
//	rep.AddBytes(int64(len(line)))
//	rep.AddLine(matched)
//
// The reporter is safe for concurrent use; AddBytes and AddLine use
// atomic operations so they may be called from multiple goroutines.
package progress
