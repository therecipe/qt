package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		if err != nil {
			return err
		}
		relPath, relErr := filepath.Rel(root, path)
		if relErr != nil {
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
		if err != nil {
			return err
		}
		if info.IsDir() {
			return f(path, info, err)
		}
		return nil
	}
}

// WalkOnlyFile is opposite of WalkFilterDirectory, it only allow file
func WalkOnlyFile(f filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return f(path, info, err)
		}
		return nil
	}
}

// WalkFilterPrefix only process file that do not have specified prefix
func WalkFilterPrefix(f filepath.WalkFunc, prefixes ...string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		for index := range prefixes {
			if strings.HasPrefix(name, prefixes[index]) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		return f(path, info, err)
	}
}

// WalkOnlyExtension only process file (not directory) that their extension is
// specified, without the `.`  such as `go` or `html`.
func WalkOnlyExtension(f filepath.WalkFunc, extensions ...string) filepath.WalkFunc {
	compiled := make([]string, len(extensions))
	for index := range extensions {
		compiled[index] = fmt.Sprintf(".%s", extensions[index])
	}
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return f(path, info, err)
		}
		extension := filepath.Ext(info.Name())
		for index := range compiled {
			if extension == compiled[index] {
				return f(path, info, err)
			}
		}
		return nil
	}
}
