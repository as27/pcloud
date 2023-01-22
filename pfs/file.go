package pfs

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

// File implements the fs.File interface for pCloud-files
type File struct {
	FileStat
	body  io.ReadCloser
	isDir bool
	files []File
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
	return f, nil
}

// Read implements os.File
func (f File) Read(b []byte) (int, error) {
	fmt.Println("Read", f.FileStat.Path)
	//return f.body.Read(b)
	return 0, io.EOF
}

// Close implements os.File
func (f File) Close() error {
	//return f.body.Close()
	return nil
}

// ReadDir implementation for fs.FS
func (f File) ReadDir(n int) ([]fs.DirEntry, error) {
	de := make([]fs.DirEntry, len(f.files))
	for i, f := range f.files {
		de[i] = f
	}
	return de, nil
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
func (f File) Mode() fs.FileMode {
	return 0
}

func (f File) Type() fs.FileMode {
	if f.isDir {
		return fs.ModeDir
	}
	return fs.ModeType
}

// ModTime implements os.FileInfo
func (f File) ModTime() time.Time {
	return f.Modified
}

// IsDir implements os.FileInfo
func (f File) IsDir() bool {
	return f.isDir
}

func (f File) Info() (fs.FileInfo, error) {
	return f, nil
}

// Sys implements os.FileInfo
func (f File) Sys() any {
	return 0
}
