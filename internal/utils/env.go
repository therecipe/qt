package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func QT_DIR() string {
	if dir := os.Getenv("QT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	if runtime.GOOS == "windows" {
		return "C:\\Qt\\Qt5.7.0"
	}
	return filepath.Join(os.Getenv("HOME"), "Qt5.7.0")
}

func QT_STUB() bool {
	return strings.ToLower(os.Getenv("QT_STUB")) == "true"
}
