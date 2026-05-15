package reader

import (
	"fmt"
	"os"
	"time"
)

// FileInfo holds metadata about a log file.
type FileInfo struct {
	Path    string
	Size    int64
	ModTime time.Time
	Mode    os.FileMode
}

// String returns a human-readable summary of the file info.
func (f FileInfo) String() string {
	return fmt.Sprintf("%s (size=%d bytes, mod=%s, mode=%s)",
		f.Path, f.Size, f.ModTime.Format(time.RFC3339), f.Mode)
}

// IsReadable reports whether the file has read permission for the current user.
func (f FileInfo) IsReadable() bool {
	return f.Mode&0o444 != 0
}

// StatFile returns FileInfo for the given path, or an error if the file
// cannot be accessed.
func StatFile(path string) (FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, fmt.Errorf("stat %q: %w", path, err)
	}
	if info.IsDir() {
		return FileInfo{}, fmt.Errorf("stat %q: path is a directory", path)
	}
	return FileInfo{
		Path:    path,
		Size:    info.Size(),
		ModTime: info.ModTime(),
		Mode:    info.Mode(),
	}, nil
}
