package fstorage

import (
	"os"
	"time"
)

func (s *StorageClient) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(s.abs(filename))
}

func (s *StorageClient) Lstat(filename string) (os.FileInfo, error) {
	return os.Lstat(s.abs(filename))
}

func (s *StorageClient) Get(filename string) (*File, error) {
	f, err := os.Open(s.abs(filename))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return &File{
		buf:   nil,
		isDir: false,
	}, nil
}

func (s *StorageClient) GetFile(filename string) (*File, error) {
	buf, err := os.ReadFile(s.abs(filename))
	if err != nil {
		return nil, err
	}
	return &File{
		buf: buf,
	}, nil
}

func (s *StorageClient) LastModified(filename string) (*time.Time, error) {
	info, err := s.Stat(filename)
	if err != nil {
		return nil, err
	}
	mod := info.ModTime()
	return &mod, nil
}
