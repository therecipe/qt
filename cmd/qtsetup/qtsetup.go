package main

import (
	"os"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	buildMode, buildTarget := "full", "desktop"

	switch len(os.Args) {
	case 2:
		switch os.Args[1] {
		case "desktop", "android", "ios", "ios-simulator",
			"sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows", "darwin", "linux",
			"linux-docker", "windows-docker", "android-docker":
			buildTarget = os.Args[1]
			utils.CheckBuildTarget(buildTarget)

		case "prep", "check", "generate", "install", "test", "full":
			buildMode = os.Args[1]

		default:
			utils.Log.Fatalln("usage:", "qtsetup", "[ prep | check | generate | install | test | full ]", "[ desktop | android | ios | ios-simulator | sailfish | sailfish-emulator | rpi1 | rpi2 | rpi3 | windows | linux-docker | windows-docker | android-docker ]")
		}

	case 3:
		buildMode = os.Args[1]
		buildTarget = os.Args[2]
	}

	switch buildMode {
	case "full":
		prep()
		check(buildTarget)
		generate(buildTarget)
		install(buildTarget)
		test(buildTarget)

	case "prep":
		prep()

	case "check":
		check(buildTarget)

	case "generate":
		generate(buildTarget)

	case "install":
		install(buildTarget)

	case "test":
		test(buildTarget)
	}
}
