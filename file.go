package fstorage

import (
	"bytes"
	"io"
)

type File struct {
	buf   []byte
	isDir bool
}

func (f *File) IsDir() bool {
	return f.isDir
}

func (f *File) String() string {
	if f.isDir {
		return "is a directory"
	}
	return string(f.buf)
}

func (f *File) Bytes() []byte {
	return f.buf
}

func (f *File) Size() int {
	return len(f.buf)
}

func (f *File) Reader() io.Reader {
	return bytes.NewReader(f.buf)
}

func (f *File) Close() error {
	return nil
}
