// Package reader provides line-by-line reading utilities for log files.
//
// It wraps bufio.Scanner with a larger buffer suitable for long log lines
// and exposes a simple Next/Line iteration interface. Readers can be
// constructed from a file path or any io.Reader, making them easy to use
// in tests and production code alike.
//
// Example usage:
//
//	lr, err := reader.New("/var/log/app.log")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for lr.Next() {
//		fmt.Println(lr.Line())
//	}
//	if err := lr.Err(); err != nil {
//		log.Fatal(err)
//	}
package reader
