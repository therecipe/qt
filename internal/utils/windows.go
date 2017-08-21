package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func QT_MSYS2() bool {
	return (os.Getenv("QT_MSYS2") == "true" || IsMsys2QtDir() || MSYSTEM() != "") && !MSYS_DOCKER()
}

func QT_MSYS2_DIR() string {
	if dir, ok := os.LookupEnv("QT_MSYS2_DIR"); ok {
		if QT_MSYS2_ARCH() == "amd64" {
			return filepath.Join(dir, "mingw64")
		}
		return filepath.Join(dir, "mingw32")
	}
	prefix := "msys32"
	if runtime.GOARCH == "amd64" {
		prefix = "msys64"
	}
	suffix := "mingw32"
	if QT_MSYS2_ARCH() == "amd64" {
		suffix = "mingw64"
	}
	return fmt.Sprintf("C:\\%v\\%v", prefix, suffix)
}

func IsMsys2QtDir() bool {
	return ExistsFile(filepath.Join(os.Getenv("QT_MSYS2_DIR"), "msys2.exe"))
}

func QT_MSYS2_ARCH() string {
	if arch, ok := os.LookupEnv("QT_MSYS2_ARCH"); ok {
		return arch
	}
	if MSYSTEM() == "MINGW64" {
		return "amd64"
	}
	return "386"
}

func QT_MSYS2_STATIC() bool {
	return os.Getenv("QT_MSYS2_STATIC") == "true"
}

func MSYSTEM() string {
	return os.Getenv("MSYSTEM")
}

func MSYS_DOCKER() bool {
	_, ok := os.LookupEnv("DOCKER_MACHINE_NAME")
	return ok
}
