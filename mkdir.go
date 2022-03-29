package fstorage

import (
	"errors"
	"github.com/szks-repo/fstorage/option/diropt"
	"os"
)

func (s *StorageClient) Mkdir(dirname string, mod os.FileMode, opt *diropt.MkdirOption) error {
	return s.mkdir(os.Mkdir, dirname, mod, opt)
}

func (s *StorageClient) MkdirAll(dirname string, mod os.FileMode, opt *diropt.MkdirOption) error {
	return s.mkdir(os.MkdirAll, dirname, mod, opt)
}

func (s *StorageClient) mkdir(fn func(string, os.FileMode) error, dirname string, mod os.FileMode, opt *diropt.MkdirOption) error {
	if opt == nil {
		opt = diropt.DefaultMkdirOption()
	}

	dirname = s.abs(dirname)
	if _, err := s.Stat(dirname); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	} else if err == nil {
		switch opt.OnConflict {
		case diropt.NoAction:
			return nil
		case diropt.ReturnErr:
			return os.ErrExist
		case diropt.Remove:
			if err := s.Remove(dirname); err != nil {
				return err
			}
			break
		case diropt.RemoveAll:
			if err := s.RemoveAll(dirname); err != nil {
				return err
			}
		}
	}

	return fn(dirname, mod)
}
