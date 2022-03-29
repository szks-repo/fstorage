package testutil

import (
	"github.com/google/uuid"
	"path/filepath"
	"strings"
)

func AbsolutePath(s string) (abs string) {
	abs, _ = filepath.Abs(s)
	return
}

func RandFileName(ext string) string {
	if ext != "" && !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return uuid.New().String() + ext
}
