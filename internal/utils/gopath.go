package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const packageName = "github.com/therecipe/qt"

var mustGoPath string

//GOBIN returns the general GOBIN string
func GOBIN() string {
	if dir, ok := os.LookupEnv("GOBIN"); ok {
		return filepath.Clean(dir)
	}
	return filepath.Join(MustGoPath(), "bin")
}

// MustGoPath returns the GOPATH that holds this package
// it exits if any error occurres and also caches the result
func MustGoPath() string {
	if mustGoPath == "" {
		mustGoPath = strings.TrimSpace(RunCmd(exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), "list", "-f", "{{.Root}}", "github.com/therecipe/qt"), "get list gopath"))
	}
	return mustGoPath
}

// GOPATH returns the general GOPATH string
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
