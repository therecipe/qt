package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const packageName = "github.com/therecipe/qt"

var goPath *string

func MustGoBin() string {
	if dir := os.Getenv("GOBIN"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(MustGoPath(), "bin")
}

// MustGoPath is same as GoPath but exits if any error ocurres
// it also caches the result
func MustGoPath() string {
	if goPath != nil {
		return *goPath
	}
	var out, err = gopath()
	if err != nil {
		Log.WithError(err).Errorf("failed to get GOPATH that holds %v", packageName)
	}
	goPath = &out
	return out
}

// GoPath return the GOPATH that holds this package
func gopath() (goPath string, err error) {
	var goPaths = GOPATH()
	if len(goPaths) == 0 {
		err = errors.New("GOPATH environment variable is unset")
		return
	}

	//TODO: fix issue
	var tries = make([]string, 0)
	for _, path := range strings.Split(goPaths, string(os.PathListSeparator)) {
		if _, err = ioutil.ReadDir(path); err != nil {
			err = fmt.Errorf("GOPATH %q point to a non-existing directory", path)
			continue
		}

		if strings.HasPrefix(fmt.Sprintf("%v%v", path, string(os.PathSeparator)), fmt.Sprintf("%v%v", runtime.GOROOT(), string(os.PathSeparator))) {
			err = fmt.Errorf("GOPATH %q is or contains GOROOT", path)
			continue
		}

		if strings.HasPrefix(fmt.Sprintf("%v%v", runtime.GOROOT(), string(os.PathSeparator)), fmt.Sprintf("%v%v", path, string(os.PathSeparator))) {
			err = fmt.Errorf("GOROOT %q is or contains GOPATH", path)
			continue
		}

		var packageDir = filepath.Join(path, "src", packageName)
		if _, err = ioutil.ReadDir(packageDir); err == nil {
			if len(goPath) > 0 {
				err = fmt.Errorf("multiple copies of %s are in GOPATH: in %s and %s", packageName, goPath, path)
				return
			}
			goPath = path
			err = nil
			break
		} else {
			tries = append(tries, packageDir)
		}
	}

	if len(goPath) == 0 {
		err = fmt.Errorf("failed to find %s in all %d GOPATH, tried %s", packageName, len(tries), strings.Join(tries, " "))
	}
	return
}

func GOPATH() string {
	if dir, ok := os.LookupEnv("GOPATH"); ok {
		return dir
	}

	home := "HOME"
	if runtime.GOOS == "windows" {
		home = "USERPROFILE"
	}
	if dir, ok := os.LookupEnv(home); ok {
		return filepath.Join(dir, "go")
	}

	return ""
}
