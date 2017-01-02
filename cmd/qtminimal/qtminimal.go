package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/therecipe/qt/internal/cmd/minimal"
	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildTarget = "desktop"
		appPath, _  = os.Getwd()
	)

	switch len(os.Args) {
	case 2:
		{
			buildTarget = os.Args[1]
		}

	case 3:
		{
			buildTarget = os.Args[1]
			appPath = os.Args[2]
		}
	}
	if !filepath.IsAbs(appPath) {
		appPath, _ = utils.Abs(appPath)
	}
	if _, err := ioutil.ReadDir(appPath); err != nil {
		utils.Log.Fatalln("usage:", "qtminimal", "[ desktop | android | ... ]", filepath.Join("path", "to", "project"))
	}

	minimal.Minimal(appPath, buildTarget)
}
