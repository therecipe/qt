package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/minimal"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	flag.Usage = func() {
		println("Usage: qtminimal [-docker] [target] [path/to/project]\n")

		println("Flags:\n")
		flag.PrintDefaults()
		print("\n")

		println("Targets:\n")
		//TODO:
		print("\n")

		os.Exit(0)
	}

	var docker bool
	flag.BoolVar(&docker, "docker", false, "run command inside docker container")

	var vagrant bool
	flag.BoolVar(&vagrant, "vagrant", false, "run command inside vagrant vm")

	var tags string
	flag.StringVar(&tags, "tags", "", "a list of build tags to consider satisfied during the build")

	if cmd.ParseFlags() {
		flag.Usage()
	}

	target := runtime.GOOS
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

	var vagrantsystem string
	if vagrant && strings.Contains(target, "/") {
		vagrantsystem = strings.Split(target, "/")[0]
		target = strings.Split(target, "/")[1]
	}

	if target == "desktop" {
		target = runtime.GOOS
	}

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			utils.Log.WithError(err).WithField("path", path).Fatal("can't resolve absolute path")
		}
	}

	utils.CheckBuildTarget(target)
	if docker {
		cmd.Docker([]string{"qtminimal", "-debug", "-tags=" + tags}, target, path, false)
	} else if vagrant {
		cmd.Vagrant([]string{"qtminimal", "-debug", "-tags=" + tags}, target, path, false, vagrantsystem)
	} else {
		minimal.Minimal(path, target, tags)
	}
}
