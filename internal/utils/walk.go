package utils

import (
	"os"

	"path/filepath"
)

var (
	// files that must be ignored
	blacklist = []string{
		"deploy",
		"qml",
		"android",
		"ios",
		"ios-simulator",
		"sailfish",
		"sailfish-emulator",
		"rpi1",
		"rpi2",
		"rpi3",
		"node_modules",
		".git",
	}
)

// WalkFilterBlacklist filter out blacklisted file
func WalkFilterBlacklist(root string, f filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		relPath, relErr := filepath.Rel(root, path)
		if err != nil {
			return relErr
		}
		for _, n := range blacklist {
			if n == relPath {
				// as the directory is blacklisted, just ignore everything under it
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		return f(path, info, err)
	}
}

// WalkFilterDirectory only allow directory
func WalkOnlyDirectory(f filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return f(path, info, err)
		}
		return nil
	}
}
