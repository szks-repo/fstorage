package fstorage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type StorageClient struct {
	basePath string
}

func New(storagePath string) (*StorageClient, error) {
	if !filepath.IsAbs(storagePath) {
		return nil, errors.New("absolute path required")
	}
	storagePath = strings.TrimRight(storagePath, "/")
	info, err := os.Stat(storagePath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("is not directory")
	}

	return &StorageClient{basePath: filepath.ToSlash(storagePath)}, nil
}

func (s *StorageClient) BasePath() string {
	return s.basePath
}

func (s *StorageClient) StoragePath(a ...string) string {
	a = append([]string{s.basePath}, a...)
	return filepath.Join(a...)
}

func (s *StorageClient) abs(filename string) string {
	filename = filepath.ToSlash(filename)
	//fmt.Println(filepath.Match(s.basePath+"/*", filename))
	if strings.HasPrefix(filename, s.basePath) {
		return filename
	} else {
		fmt.Println("filename:", filename)
		fmt.Println("s.basePa:", s.basePath)
	}
	return filepath.Join(s.basePath, strings.TrimPrefix(filename, "/"))
}
