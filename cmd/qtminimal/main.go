package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/minimal"

	"github.com/therecipe/qt/internal/utils"
)

func qmake_main() {
	flag.Usage = func() {
		println("Usage: qtminimal [-docker] [target] [path/to/project]\n")

		println("Flags:\n")
		flag.PrintDefaults()
		print("\n")

		os.Exit(0)
	}

	var docker bool
	flag.BoolVar(&docker, "docker", false, "run command inside docker container")

	if cmd.ParseFlags() {
		flag.Usage()
	}

	target := "desktop"
	path, err := os.Getwd()
	if err != nil {
		utils.Log.WithError(err).Debug("failed to get cwd")
	}

	switch flag.NArg() {
	case 0:
	case 1:
		target = flag.Arg(0)
	case 2:
		target = flag.Arg(0)
		path = flag.Arg(1)
	default:
		flag.Usage()
	}

	if target == runtime.GOOS {
		target = "desktop"
	}

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			utils.Log.WithError(err).WithField("path", path).Fatal("can't resolve absolute path")
		}
	}

	utils.CheckBuildTarget(target)
	if docker {
		cmd.Docker([]string{"qtminimal", "-debug"}, target, path)
	} else {
		minimal.QmakeMinimal(path, target)
	}
}
