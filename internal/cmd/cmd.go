package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/utils"
)

func ParseFlags() {
	var (
		debug      = flag.Bool("debug", false, "print debug logs")
		p          = flag.Int("p", runtime.NumCPU(), "the number of cpu's used")
		qt_dir     = flag.String("qt_dir", "", "define alternative qt dir")
		qt_version = flag.String("qt_version", "", "define alternative qt version")
	)
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
	}

	runtime.GOMAXPROCS(*p)

	if dir := *qt_dir; dir != "" {
		os.Setenv("QT_DIR", dir)
	}

	if version := *qt_version; version != "" {
		os.Setenv("QT_VERSION", version)
	}
}

func Docker(arg []string, buildTarget, appPath string) {

	var dockerImage string

	switch buildTarget {
	case "desktop":
		{
			switch runtime.GOOS {
			case "windows":
				{
					dockerImage = "base_windows"
					if utils.QT_MXE_ARCH() == "amd64" {
						dockerImage = "base_windows_64"
					}
				}

			case "darwin":
				{
					utils.Log.Fatalf("%v is currently not supported as a deploy target by docker", runtime.GOOS)
				}

			case "linux":
				{
					dockerImage = "base"
				}
			}
		}

	case "windows":
		{
			dockerImage = "base_windows"
			if utils.QT_MXE_ARCH() == "amd64" {
				dockerImage = "base_windows_64"
			}
		}

	case "darwin":
		{
			utils.Log.Fatalf("%v is currently not supported as a deploy target by docker", runtime.GOOS)
		}

	case "linux":
		{
			dockerImage = "base"
		}

	case "android":
		{
			dockerImage = "base_android"
		}

	default:
		{
			utils.Log.Fatalf("%v is currently not supported as a deploy target by docker", runtime.GOOS)
		}
	}

	var (
		args = []string{"run", "--rm"}
		path = make([]string, 0)
	)

	for i, gp := range strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator)) {
		args = append(args, []string{"-v", fmt.Sprintf("%v:/media/sf_GOPATH%v", gp, i)}...)
		path = append(path, fmt.Sprintf("/media/sf_GOPATH%v", i))
	}

	args = append(args, []string{"-e", "GOPATH=" + strings.Join(path, ":")}...)

	args = append(args, []string{"-i", fmt.Sprintf("therecipe/qt:%v", dockerImage)}...)

	args = append(args, arg...)

	var found bool
	for i, gp := range strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator)) {
		gp = filepath.Clean(gp)
		if strings.HasPrefix(appPath, gp) {
			args = append(args, strings.Replace(strings.Replace(appPath, gp, fmt.Sprintf("/media/sf_GOPATH%v", i), -1), "\\", "/", -1))
			found = true
			break
		}
	}

	if !found {
		utils.Log.Panicln("Project needs to be inside GOPATH", appPath, os.Getenv("GOPATH"))
	}

	utils.RunCmd(exec.Command("docker", args...), fmt.Sprintf("deploy binary for %v on %v with docker", buildTarget, runtime.GOOS))
}
