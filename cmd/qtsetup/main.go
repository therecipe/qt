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

func qmake_main() {
	flag.Usage = func() {
		println("Usage: qtsetup [-debug] [mode] [target]\n")

		println("Flags:\n")
		flag.PrintDefaults()
		print("\n")

		println("Modes:\n")
		modes := []struct{ name, desc string }{
			{"prep", "try to symlink tooling into the PATH"},
			{"check", "do some basic env checks"},
			{"generate", "generate the code for all packages"},
			{"install", "run go install for all packages"},
			{"test", "build some examples"},
			{"full", "run all of the above"},
			{"help", "print help"},
			{"update", "update 'cmd' and 'internal/cmd'"},
			{"upgrade", "update everything"},
		}
		for _, mode := range modes {
			fmt.Printf("  %v%v%v\n", mode.name, strings.Repeat(" ", 12-len(mode.name)), mode.desc)
		}
		print("\n")

		os.Exit(0)
	}

	if cmd.ParseFlags() {
		flag.Usage()
	}

	mode := "full"
	target := "desktop"

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

	if target == runtime.GOOS {
		target = "desktop"
	}

	utils.CheckBuildTarget(target)

	switch mode {
	case "full":
		setup.Prep()
		setup.Check(target)
		setup.Generate(target)
		setup.Install(target)
		setup.Test(target)
	case "prep":
		setup.Prep()
	case "check":
		setup.Check(target)
	case "generate":
		setup.Generate(target)
	case "install":
		setup.Install(target)
	case "test":
		setup.Test(target)
	case "update":
		setup.Update()
	case "upgrade":
		setup.Upgrade()
	default:
		flag.Usage()
	}
}
