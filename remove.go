package fstorage

import "os"

func (s *StorageClient) Remove(filename string) error {
	return os.Remove(s.abs(filename))
}

func (s *StorageClient) RemoveAll(filename string) error {
	return os.RemoveAll(s.abs(filename))
}

func (s *StorageClient) RemoveIfExist(filename string) error {
	return s.rmIfExist(os.Remove, filename)
}

func (s *StorageClient) RemoveAllIfExist(filename string) error {
	return s.rmIfExist(os.RemoveAll, filename)
}

func (s *StorageClient) rmIfExist(fn func(string) error, filename string) error {
	filename = s.abs(filename)
	_, err := s.Stat(filename)
	if err != nil && err != os.ErrNotExist {
		return nil
	}
	if err == nil {
		return fn(filename)
	}
	return nil
}

func (s *StorageClient) RemoveIf(filename string, cond func(info os.FileInfo) bool) error {
	return s.rmIf(os.Remove, filename, cond)
}

func (s *StorageClient) RemoveAllIf(filename string, cond func(info os.FileInfo) bool) error {
	return s.rmIf(os.RemoveAll, filename, cond)
}

func (s *StorageClient) rmIf(fn func(string) error, filename string, cond func(info os.FileInfo) bool) error {
	filename = s.abs(filename)
	info, err := s.Stat(filename)
	if err != nil {
		return err
	}
	if cond(info) {
		return fn(filename)
	}
	return nil
}
