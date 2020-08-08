package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/StarAurryon/qt/internal/cmd"
	"github.com/StarAurryon/qt/internal/cmd/rcc"

	"github.com/StarAurryon/qt/internal/utils"
)

func main() {
	var non_recursive bool
	if os.Args[len(os.Args)-1] == "non_recursive" {
		non_recursive = true
		os.Args = os.Args[:len(os.Args)-1]
	}

	flag.Usage = func() {
		println("Usage: qtrcc [-docker] [target] [path/to/project]\n")

		println("Flags:\n")
		flag.PrintDefaults()
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

	var output string
	flag.StringVar(&output, "o", os.Getenv("QTRCC_OUTPUT_DIR"), "specify an alternative output dir")

	var tags string
	flag.StringVar(&tags, "tags", "", "a list of build tags to consider satisfied during the build")

	var quickcompiler bool
	flag.BoolVar(&quickcompiler, "quickcompiler", false, "use the quickcompiler")

	var uic bool
	flag.BoolVar(&uic, "uic", true, "use the uic")

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

	if output != "" && !filepath.IsAbs(output) {
		output, err = filepath.Abs(output)
		if err != nil {
			utils.Log.WithError(err).WithField("output", output).Fatal("can't resolve absolute path")
		}
	}

	dockerArgs := []string{"qtrcc", "-debug"}
	if !uic {
		dockerArgs = append(dockerArgs, "-uic=false")
	}
	if quickcompiler {
		dockerArgs = append(dockerArgs, "-quickcompiler")
	}

	switch {
	case docker:
		cmd.Docker(dockerArgs, target, path, false)
	case vagrant:
		cmd.Vagrant(dockerArgs, target, path, false, vagrant_system)
	default:
		rcc.Rcc(path, target, tags, output, uic, quickcompiler, false, false)
	}
}
