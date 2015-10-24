package utils

import (
	"os"
	"path"
)

func GetQtPkgPath(s ...string) string {
	return path.Join(os.Getenv("GOPATH"), "src", "github.com", "therecipe", "qt", path.Join(s...))
}
