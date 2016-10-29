package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildMode   = "full"
		buildTarget = "desktop"
	)

	switch len(os.Args) {
	case 2:
		{
			switch os.Args[1] {
			case "desktop", "android", "ios", "ios-simulator",
				"sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows", "darwin", "linux",
				"linux-docker", "windows-docker", "android-docker":
				{
					buildTarget = os.Args[1]
					var buildDocker = strings.HasSuffix(buildTarget, "-docker")
					switch buildTarget {
					case "windows":
						{
							if runtime.GOOS == "windows" && !buildDocker {
							} else if runtime.GOOS == "linux" || buildDocker {
							} else {
								utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
							}
						}

					case "darwin", "ios", "ios-simulator":
						{
							if runtime.GOOS == "darwin" && !buildDocker {
							} else {
								utils.Log.Fatalf("%v is currently not supported as a deploy target on %v (not even with docker)", buildTarget, runtime.GOOS)
							}
						}

					case "linux":
						{
							if runtime.GOOS == "linux" && !buildDocker {
							} else if buildDocker {
							} else {
								utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
							}
						}
					}
				}

			case "prep", "check", "generate", "install", "test", "full":
				{
					buildMode = os.Args[1]
				}

			default:
				{
					utils.Log.Fatalln("usage:", "qtsetup", "[ prep | check | generate | install | test | full ]", "[ desktop | android | ios | ios-simulator | sailfish | sailfish-emulator | rpi1 | rpi2 | rpi3 | windows | linux-docker | windows-docker | android-docker ]")
				}
			}
		}

	case 3:
		{
			buildMode = os.Args[1]
			buildTarget = os.Args[2]
		}
	}

	switch buildMode {
	case "full":
		{
			prep()
			check(buildTarget)
			generate()
			install(buildTarget)
			test(buildTarget)
		}

	case "prep":
		{
			prep()
		}

	case "check":
		{
			check(buildTarget)
		}

	case "generate":
		{
			generate()
		}

	case "install":
		{
			install(buildTarget)
		}

	case "test":
		{
			test(buildTarget)
		}
	}
}
