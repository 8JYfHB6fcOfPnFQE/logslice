package reader

import (
	"fmt"
	"os"
	"time"
)

// FileInfo holds metadata about a log file to be sliced.
type FileInfo struct {
	Path    string
	Size    int64
	ModTime time.Time
}

// StatFile returns FileInfo for the given path.
func StatFile(path string) (FileInfo, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, fmt.Errorf("reader: stat %q: %w", path, err)
	}
	if fi.IsDir() {
		return FileInfo{}, fmt.Errorf("reader: %q is a directory, not a file", path)
	}
	return FileInfo{
		Path:    path,
		Size:    fi.Size(),
		ModTime: fi.ModTime(),
	}, nil
}

// String returns a human-readable summary of the FileInfo.
func (f FileInfo) String() string {
	return fmt.Sprintf("%s (size=%d bytes, modified=%s)",
		f.Path, f.Size, f.ModTime.Format(time.RFC3339))
}
