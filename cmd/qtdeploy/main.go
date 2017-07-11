package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/deploy"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	flag.Usage = func() {
		println("Usage: qtdeploy [-docker] [mode] [target] [path/to/project]\n")

		println("Flags:\n")
		flag.PrintDefaults()
		print("\n")

		println("Modes:\n")
		modes := []struct{ name, desc string }{
			{"build", "compile and bundle"},
			{"run", "run the binary"},
			{"test", "build and run"},
			{"help", "print help"},
		}
		for _, mode := range modes {
			fmt.Printf("  %v%v%v\n", mode.name, strings.Repeat(" ", 12-len(mode.name)), mode.desc)
		}
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

	var ldFlags string
	flag.StringVar(&ldFlags, "ldflags", "", "arguments to pass on each go tool link invocation")

	var fast bool
	flag.BoolVar(&fast, "fast", false, "use cached moc, minimal and dependencies (works for: windows, darwin, linux)")

	var tags string
	flag.StringVar(&tags, "tags", "", "a list of build tags to consider satisfied during the build")

	var device string
	flag.StringVar(&device, "device", "", "a device UUID to be used by the iOS simulator")

	if cmd.ParseFlags() {
		flag.Usage()
	}

	mode := "test"
	target := runtime.GOOS
	path, err := os.Getwd()
	if err != nil {
		utils.Log.WithError(err).Debug("failed to get cwd")
	}

	switch flag.NArg() {
	case 0:
	case 1:
		mode = flag.Arg(0)
	case 2:
		mode = flag.Arg(0)
		target = flag.Arg(1)
	case 3:
		mode = flag.Arg(0)
		target = flag.Arg(1)
		path = flag.Arg(2)
	default:
		flag.Usage()
	}

	if mode == "help" {
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

	//TODO: check mode here

	switch target {
	case "android", "android-emulator", "ios", "ios-simulator", "sailfish", "sailfish-emulator":
		fast = false
	}
	deploy.Deploy(mode, target, path, docker, ldFlags, tags, fast && !docker, device, vagrant, vagrantsystem)
}
