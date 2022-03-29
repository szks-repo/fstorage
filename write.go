package fstorage

import (
	"bytes"
	"errors"
	"github.com/szks-repo/fstorage/option/diropt"
	"github.com/szks-repo/fstorage/option/fileopt"
	"io"
	"os"
	"path/filepath"
)

func (s *StorageClient) Save(filename string, r io.Reader, opt *fileopt.SaveFileOption) error {
	if opt == nil {
		opt = &fileopt.SaveFileOption{OnConflict: fileopt.Overwrite}
	}
	buf := new(bytes.Buffer)
	filename = s.abs(filename)
	_, err := s.Stat(filename)
	var wc io.WriteCloser
	defer func() {
		if wc != nil {
			_ = wc.Close()
		}
	}()

	if err == nil {
		if opt.OnConflict == fileopt.Overwrite || opt.OnConflict == fileopt.Append {
			if opt.OnConflict == fileopt.Append {
				old, err := os.ReadFile(filename)
				if err != nil {
					return err
				}
				buf.Write(old)
			}
			wc, err = os.Create(filename)
			if err != nil {
				return err
			}

		} else if opt.OnConflict == fileopt.ReturnErr {
			return os.ErrExist
		} else if opt.OnConflict == fileopt.NoAction {
			return nil
		}
	} else if errors.Is(err, os.ErrNotExist) {
		wc, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	_, err = io.Copy(buf, r)
	if _, seekable := r.(io.ReadSeeker); seekable {
		_, _ = r.(io.ReadSeeker).Seek(0, 0)
	}

	if err != nil && err != io.EOF {
		return err
	}

	_, _ = io.Copy(wc, buf)

	return nil
}

func (s *StorageClient) SaveIfNotExist(filename string, r io.Reader) error {
	return s.Save(filename, r, &fileopt.SaveFileOption{OnConflict: fileopt.NoAction})
}

func (s *StorageClient) SaveAll(filename string, r io.Reader, opt *fileopt.SaveFileOption) error {
	filename = s.abs(filename)
	dir, _ := filepath.Split(filename)
	if err := s.MkdirAll(dir, os.ModePerm, &diropt.MkdirOption{OnConflict: diropt.NoAction}); err != nil {
		return err
	}
	return s.Save(filename, r, opt)
}

func (s *StorageClient) SaveAllIfNotExist(filename string, r io.Reader) error {
	return s.SaveAll(filename, r, &fileopt.SaveFileOption{OnConflict: fileopt.NoAction})
}
