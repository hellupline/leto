package leto

import (
	"bytes"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	defaultFileSystem = New(defaultFiles)
	defaultFiles      = Files{}

	Open = defaultFileSystem.Open

	Register = defaultFiles.Register
)

type FileSystem struct {
	Files
}

func New(files Files) *FileSystem {
	return &FileSystem{Files: files}
}

func Default() *FileSystem {
	return defaultFileSystem
}

func (fs *FileSystem) Open(name string) (http.File, error) {
	f, present := fs.Files[path.Clean(name)]
	if !present {
		return os.Open(name)
	}
	return f.File()
}

type Files map[string]*File

func (s Files) Register(name string, data []byte) {
	s[name] = &File{
		modtime: time.Now().Unix(),
		size:    int64(len(data)),
		data:    data,
		name:    name,
	}
}

type File struct {
	data    []byte
	modtime int64
	size    int64
	isDir   bool
	name    string
}

func (f File) File() (http.File, error) {
	httpFile := struct {
		*bytes.Reader
		File
	}{
		Reader: bytes.NewReader(f.data),
		File:   f,
	}
	return &httpFile, nil
}

func (File) Close() error {
	return nil
}

func (File) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *File) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int64 {
	return f.size
}

func (f *File) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *File) Mode() os.FileMode {
	return 0
}

func (f *File) Sys() interface{} {
	return f
}

func (f *File) IsDir() bool {
	return f.isDir
}
