package utils

import (
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

const PackageName = "github.com/therecipe/qt"

var (
	mustGoPath      string
	mustGoPathMutex = new(sync.Mutex)
)

func GOBIN() string  { return envOr("GOBIN", filepath.Join(MustGoPath(), "bin")) }
func GOPATH() string { return envOr("GOPATH", build.Default.GOPATH) }

// MustGoPath returns the GOPATH that holds this package
func MustGoPath() string {
	mustGoPathMutex.Lock()
	if mustGoPath == "" {
		if _, err := exec.LookPath("go"); err == nil {
			mustGoPath = strings.TrimSpace(RunCmd(GoList("{{.Root}}", PackageName, "-find"), "get list gopath"))
		}
		if mustGoPath == "" {
			mustGoPath = GOPATH()
		}
	}
	mustGoPathMutex.Unlock()
	return mustGoPath
}

func envOr(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}
