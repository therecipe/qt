package utils

import (
	"errors"
	"os"
	"os/exec"
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

	return strings.TrimSpace(RunCmd(exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), "list", "-f", "{{.Root}}", "github.com/therecipe/qt"), "get list gopath")), nil
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
