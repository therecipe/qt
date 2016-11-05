package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func UseMsys2() bool {
	return strings.ToLower(os.Getenv("QT_MSYS2")) == "true" || IsMsys2QtDir()
}

func QT_MSYS2_DIR() string {
	if dir := os.Getenv("QT_MSYS2_DIR"); dir != "" {
		if os.Getenv("MSYSTEM") == "MINGW64" {
			return filepath.Join(dir, "mingw64")
		}
		return filepath.Join(dir, "mingw32")
	}

	var prefix = "msys32"
	if runtime.GOARCH == "amd64" {
		prefix = "msys64"
	}
	if os.Getenv("MSYSTEM") == "MINGW64" {
		return fmt.Sprintf("C:\\%v\\mingw64", prefix)
	}
	return fmt.Sprintf("C:\\%v\\mingw32", prefix)
}

func IsMsys2QtDir() bool {
	return Exists(filepath.Join(os.Getenv("QT_MSYS2_DIR"), "msys2.exe"))
}
