// Package pfs implements the pCloud API as a fs.FS interface
// At the moment this is a prototype and should not be used by
// other packages! Because the API is not stable yet and the
// implementation is not complete.
package pfs

import (
	"io/fs"
	"time"

	"github.com/as27/pcloud/api"
)

// FS implements the fs.Fs interface for pCloud-files
type FS struct {
	metadata api.Metadata
}

func (f *FS) Open(name string) (fs.File, error) {
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
	}
	if name == "." {
		rootFile := File{
			FileStat: FileStat{
				Name: ".",
			},
			isDir: true,
		}
		for _, m := range f.metadata.Contents {
			rootFile.files = append(rootFile.files, metaToFile(m))
		}
		return rootFile, nil

	}
	for _, m := range f.metadata.Contents {
		if m.Name == name {
			return metaToFile(m), nil
		}
	}
	return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
}

func metaToFile(m api.Metadata) File {
	f := File{
		FileStat: FileStat{
			Name:        m.Name,
			Size:        m.Size,
			Path:        m.Path,
			ContentType: m.ContentType,
			Created:     time.Now(), // TODO: parse time
			Modified:    time.Now(), // TODO: parse time
		},
	}
	if m.IsFolder {
		f.isDir = true
	}
	return f
}
