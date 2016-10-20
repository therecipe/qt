package utils

import (
	"os"
	"strings"
	"io/ioutil"
	"fmt"
	"errors"
	"runtime"
	"path/filepath"
)

const packageName = "github.com/therecipe/qt"

// GoPath return the GOPATH where hold this package.
func GoPath() (goPath string , err error) {
	goPaths := os.Getenv("GOPATH")
	if len(goPaths) == 0 {
		err = errors.New("GOPATH environment variable is unset")
		return
	}

	tries := make([]string, 0)
	for _, path := range strings.Split(goPaths, ":") {
		if _, err = ioutil.ReadDir(path); err != nil {
			err = fmt.Errorf("GOPATH %q point to a non-existing directory", path)
			return
		}

		if strings.Contains(path, runtime.GOROOT()) {
			err = fmt.Errorf("GOPATH %q is or contains GOROOT", path)
			return
		}

		if strings.Contains(runtime.GOROOT(), path) {
			err = fmt.Errorf("nGOROOT %q is or contains GOPATH", path)
			return
		}

		packageDir := filepath.Join(path, "src", packageName)
		if _, err = ioutil.ReadDir(packageDir); err == nil {
			if len(goPath) > 0 {
				err = fmt.Errorf("Multiple copies of %s are in GOPATH: in %s and %s", packageName, goPath, path)
				return
			}
			goPath = path
		} else {
			tries = append(tries, packageDir)
		}
	}

	if len(goPath) == 0 {
		err = fmt.Errorf("Couldn't find %s in all %d GOPATH, tried: %s", packageName, len(tries), strings.Join(tries, " "))
	}
	return
}
