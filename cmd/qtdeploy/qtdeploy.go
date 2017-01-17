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

var (
	appPath                string
	buildMode, buildTarget string
	buildDocker            bool
)

func main() {
	var ldFlags = flag.String("ldflags", "", "arguments to pass on each go tool link invocation.")

	cmd.ParseFlags()

	args()

	deploy.Deploy(&deploy.State{
		BuildMode:   buildMode,
		BuildTarget: buildTarget,
		AppPath:     appPath,
		BuildDocker: buildDocker,
		LdFlags:     *ldFlags,
	})
}

func args() {
	utils.Log.Debug("parsing arguments")

	switch flag.NArg() {
	case 0:
		{
			buildMode = "test"
			buildTarget = "desktop"
			appPath, _ = os.Getwd()
		}

	case 1:
		{
			buildMode = flag.Arg(0)
			buildTarget = "desktop"
			appPath, _ = os.Getwd()
		}

	case 2:
		{
			buildMode = flag.Arg(0)
			buildTarget = flag.Arg(1)
			appPath, _ = os.Getwd()
		}

	case 3:
		{
			buildMode = flag.Arg(0)
			buildTarget = flag.Arg(1)
			appPath = flag.Arg(2)
		}

	case 4:
		{
			buildMode = flag.Arg(0)
			buildTarget = flag.Arg(1)
			appPath = flag.Arg(2)
			buildDocker = (flag.Arg(3) == "docker")
		}
	}

	if strings.ToLower(os.Getenv("CI")) == "true" {
		buildMode = "build"
	}

	utils.Log.Debugln("mode:", buildMode, "target:", buildTarget, "path:", appPath, "use docker:", buildDocker)

	switch buildMode {
	case "build", "run", "test":
		{
			switch buildTarget {
			case "desktop", "android", "ios", "ios-simulator", "sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows", "darwin", "linux":
				{
					switch buildTarget {
					case "windows":
						{
							if runtime.GOOS == "windows" && !buildDocker {
								buildTarget = "desktop"
							} else if runtime.GOOS == "linux" || buildDocker {
								buildTarget = "windows"
							} else {
								utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
							}
						}

					case "darwin":
						{
							if runtime.GOOS == "darwin" && !buildDocker {
								buildTarget = "desktop"
							} else {
								utils.Log.Fatalf("%v is currently not supported as a deploy target on %v (not even with docker)", buildTarget, runtime.GOOS)
							}
						}

					case "linux":
						{
							if runtime.GOOS == "linux" && !buildDocker {
								buildTarget = "desktop"
							} else if buildDocker {
								buildTarget = "linux"
							} else {
								utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
							}
						}
					}
				}

			default:
				{
					utils.Log.Fatalln("usage:", "qtdeploy", "[ build | run | test ]", "[ desktop | android | ios | ios-simulator | sailfish | sailfish-emulator | rpi1 | rpi2 | rpi3 | windows ]", fmt.Sprintf("[ %v ]", filepath.Join("path", "to", "project")), "[ docker ]")
				}
			}
		}

	default:
		{
			utils.Log.Fatalln("usage:", "qtdeploy", "[ build | run | test ]", "[ desktop | android | ios | ios-simulator | sailfish | sailfish-emulator | rpi1 | rpi2 | rpi3 | windows ]", fmt.Sprintf("[ %v ]", filepath.Join("path", "to", "project")), "[ docker ]")
		}
	}
}
