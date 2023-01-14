package pcloud

import (
	"os"
	"time"
)

// File implements the fs.File interface for pCloud-files
type File struct {
	FileStat
}

type FileStat struct {
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	FileID      int64     `json:"fileid"`
	Path        string    `json:"path"`
	Name        string    `json:"name"`
	ContentType string    `json:"contenttype"`
	Size        int64     `json:"size"`
}

// Stat implements os.File
func (f File) Stat() (os.FileInfo, error) {
	return nil, nil
}

// Read implements os.File
func (f File) Read([]byte) (int, error) {
	return 0, nil
}

// Close implements os.File
func (f File) Close() error {
	return nil
}

// Name implements os.FileInfo
func (f File) Name() string {
	return f.FileStat.Name
}

// Size implements os.FileInfo
func (f File) Size() int64 {
	return f.FileStat.Size
}

// Mode implements os.FileInfo
func (f File) Mode() os.FileMode {
	return 0
}

// ModTime implements os.FileInfo
func (f File) ModTime() time.Time {
	return time.Now()
}

// IsDir implements os.FileInfo
func (f File) IsDir() bool {
	return false
}

// Sys implements os.FileInfo
func (f File) Sys() any {
	return nil
}
