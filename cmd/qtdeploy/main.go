package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/StarAurryon/qt/internal/cmd"
	"github.com/StarAurryon/qt/internal/cmd/deploy"

	"github.com/StarAurryon/qt/internal/utils"
)

func main() {
	var non_recursive bool
	if os.Args[len(os.Args)-1] == "non_recursive" {
		non_recursive = true
		os.Args = os.Args[:len(os.Args)-1]
	}

	flag.Usage = func() {
		println("Usage: qtdeploy [-docker] [mode] [target] [path/to/project]\n")

		println("Flags:\n")
		flag.PrintDefaults()
		println()

		println("Modes:\n")
		for _, m := range []struct{ name, desc string }{
			{"build", "compile and bundle"},
			{"run", "run the binary"},
			{"test", "build and run"},
			{"help", "print help"},
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

	var ldFlags string
	flag.StringVar(&ldFlags, "ldflags", "", "arguments to pass on each go tool link invocation")

	var fast bool
	flag.BoolVar(&fast, "fast", false, "use cached moc, minimal and dependencies (works for: windows, darwin, linux)")

	var tags string
	flag.StringVar(&tags, "tags", "", "a list of build tags to consider satisfied during the build")

	var device string
	flag.StringVar(&device, "device", "", "a device UUID to be used by the iOS simulator")

	var comply bool
	flag.BoolVar(&comply, "comply", false, "dump object code to make it easier to comply with LGPL obligations for proprietary developments")

	var quickcompiler bool
	flag.BoolVar(&quickcompiler, "quickcompiler", false, "use the quickcompiler")

	var uic bool
	flag.BoolVar(&uic, "uic", true, "use the uic")

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

	var vagrant_system string
	if target_splitted := strings.Split(target, "/"); vagrant && len(target_splitted) == 2 {
		vagrant_system = target_splitted[0]
		target = target_splitted[1]
	}

	if target == "desktop" {
		target = runtime.GOOS
	}
	utils.CheckBuildTarget(target, docker)

	if !filepath.IsAbs(path) {
		oPath := path
		path, err = filepath.Abs(path)
		if err != nil || !utils.ExistsDir(path) {
			utils.Log.WithError(err).WithField("path", path).Debug("can't resolve absolute path")
			dirFunc := func() (string, error) {
				out, err := utils.RunCmdOptionalError(utils.GoList("{{.Dir}}", oPath, "-find"), "get pkg dir")
				return strings.TrimSpace(out), err
			}
			if dir, err := dirFunc(); err != nil || len(dir) == 0 {
				utils.RunCmd(exec.Command("go", "get", "-d", "-v", oPath), "go get pkg")
				path, _ = dirFunc()
			} else {
				path = dir
			}
		}
	}

	cmd.InitEnv(target, docker, path)
	if utils.QT_DOCKER() && utils.UseGOMOD(path) && !non_recursive {
		if cmd.RestartWithPinnedVersion(path) {
			return
		}
	}

	if !(target == runtime.GOOS || strings.HasPrefix(target, "js") || strings.HasPrefix(target, "wasm")) {
		fast = false
	}
	if (docker || vagrant) && !(strings.HasPrefix(target, "js") || strings.HasPrefix(target, "wasm")) {
		fast = false
	}

	deploy.Deploy(mode, target, path, docker, ldFlags, tags, fast, device, vagrant, vagrant_system, comply, uic, quickcompiler)
}
