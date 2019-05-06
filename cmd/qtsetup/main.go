package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/setup"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	flag.Usage = func() {
		println("Usage: qtsetup [-debug] [mode] [target]\n")

		println("Flags:\n")
		flag.PrintDefaults()
		println()

		println("Modes:\n")
		for _, m := range []struct{ name, desc string }{
			{"prep", "symlink tooling into the PATH"},
			{"check", "perform some basic env checks"},
			{"generate", "generate code for all packages"},
			{"install", "go install all packages"},
			{"test", "build and test some examples"},
			{"full", "run all of the above"},
			{"help", "print help"},
			{"update", "update 'cmd' and 'internal/cmd'"},
			{"upgrade", "update everything"},
		} {
			fmt.Printf("  %v%v%v\n", m.name, strings.Repeat(" ", 12-len(m.name)), m.desc)
		}
		println()

		println("Targets:\n")
		//TODO:
		println()

		os.Exit(0)
	}

	var docker bool
	flag.BoolVar(&docker, "docker", false, "run command inside docker container")

	var vagrant bool
	flag.BoolVar(&vagrant, "vagrant", false, "run command inside vagrant vm")

	var dynamic bool
	if runtime.GOOS != "windows" {
		flag.BoolVar(&dynamic, "dynamic", false, "create and use semi-dynamic libraries during the generation and installation process (experimental; no real replacement for dynamic linking)")
	}

	var failfast bool
	flag.BoolVar(&failfast, "failfast", false, "exit the setup upon the first error encountered during the installation step")

	var test bool
	flag.BoolVar(&test, "test", true, "build and run example applications on the end of the setup")

	if cmd.ParseFlags() {
		flag.Usage()
	}

	mode := "full"
	target := runtime.GOOS

	switch flag.NArg() {
	case 0:
	case 1:
		mode = flag.Arg(0)
	case 2:
		mode = flag.Arg(0)
		target = flag.Arg(1)
	default:
		flag.Usage()
	}

	var vagrant_system string
	if target_splitted := strings.Split(target, "/"); vagrant && len(target_splitted) == 2 {
		vagrant_system = target_splitted[0]
		target = target_splitted[1]
	}

	if target == "desktop" {
		target = runtime.GOOS
	}
	utils.CheckBuildTarget(target, docker)
	cmd.InitEnv(target)

	if dynamic && (target == runtime.GOOS || target == "js" || target == "wasm") {
		os.Setenv("QT_DYNAMIC_SETUP", "true")
	}

	switch mode {
	case "prep":
		setup.Prep(target)
	case "check":
		setup.Check(target, docker, vagrant)
	case "generate":
		setup.Generate(target, docker, vagrant)
	case "install":
		setup.Install(target, docker, vagrant, failfast)
	case "test":
		if !test {
			return
		}
		setup.Test(target, docker, vagrant, vagrant_system)
	case "full":
		setup.Prep(target)
		setup.Check(target, docker, vagrant)
		setup.Generate(target, docker, vagrant)
		setup.Install(target, docker, vagrant, failfast)
		if !test {
			return
		}
		setup.Test(target, docker, vagrant, vagrant_system)
	case "update":
		setup.Update()
	case "upgrade":
		setup.Upgrade()
	default:
		flag.Usage()
	}
}
