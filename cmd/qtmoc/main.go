package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/moc"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	flag.Usage = func() {
		println("Usage: qtmoc [-docker] [target] [path/to/project]\n")

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

	var fast bool
	flag.BoolVar(&fast, "fast", false, "don't run qtmoc for dependencies")

	var tags string
	flag.StringVar(&tags, "tags", "", "a list of build tags to consider satisfied during the build")

	var slow bool
	flag.BoolVar(&slow, "slow", false, "reduce qtmoc's resource usage")

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
		cmd.Docker([]string{"qtmoc", "-debug", "-tags=" + tags}, target, path, false)
	} else if vagrant {
		cmd.Vagrant([]string{"qtmoc", "-debug", "-tags=" + tags}, target, path, false, vagrantsystem)
	} else {
		moc.Moc(path, target, tags, fast, slow)
	}
}
