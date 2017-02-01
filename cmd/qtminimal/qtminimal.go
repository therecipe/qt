package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/minimal"
	"github.com/therecipe/qt/internal/utils"
)

func main() {
	cmd.ParseFlags()

	var (
		buildTarget = "desktop"
		appPath, _  = os.Getwd()
		buildDocker bool
	)

	switch flag.NArg() {
	case 1:
		{
			buildTarget = flag.Arg(0)
		}

	case 2:
		{
			buildTarget = flag.Arg(0)
			appPath = flag.Arg(1)
		}

	case 3:
		{
			buildTarget = flag.Arg(0)
			appPath = flag.Arg(1)
			buildDocker = true
		}
	}
	if !filepath.IsAbs(appPath) {
		appPath, _ = utils.Abs(appPath)
	}
	if _, err := ioutil.ReadDir(appPath); err != nil || strings.ContainsAny(buildTarget, "./\\") {
		utils.Log.Fatalln("usage:", "qtminimal", "[ desktop | android | ... ]", filepath.Join("path", "to", "project"))
	}

	if buildDocker {
		cmd.Docker([]string{"qtminimal", "-debug"}, buildTarget, appPath)
	} else {
		minimal.Minimal(appPath, buildTarget)
	}
}
