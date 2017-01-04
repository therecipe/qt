package main

import (
	"flag"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/setup"
	"github.com/therecipe/qt/internal/utils"
)

func main() {
	cmd.ParseFlags()

	buildMode, buildTarget := "full", "desktop"

	switch flag.NArg() {
	case 1:
		switch flag.Arg(0) {
		case "desktop", "android", "ios", "ios-simulator",
			"sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows", "darwin", "linux",
			"linux-docker", "windows-docker", "android-docker":
			buildTarget = flag.Arg(0)
			utils.CheckBuildTarget(buildTarget)

		case "prep", "check", "generate", "install", "test", "full":
			buildMode = flag.Arg(0)

		default:
			utils.Log.Fatalln("usage:", "qtsetup", "[ prep | check | generate | install | test | full ]", "[ desktop | android | ios | ios-simulator | sailfish | sailfish-emulator | rpi1 | rpi2 | rpi3 | windows | linux-docker | windows-docker | android-docker ]")
		}

	case 2:
		buildMode = flag.Arg(0)
		buildTarget = flag.Arg(1)
	}

	setup.Setup(buildMode, buildTarget)
}
