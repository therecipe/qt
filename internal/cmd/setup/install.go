package setup

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"

	"github.com/therecipe/qt/internal/cmd"

	"github.com/therecipe/qt/internal/utils"
)

func Install(target string) {
	utils.Log.Infof("running setup/install %v", target)

	env, tagFlags := cmd.EnvAndTagFlags(target)

	if target == "sailfish" {
		env["GOARCH"] = "386"
		delete(env, "GOARM")

		if _, err := ioutil.ReadDir(filepath.Join(runtime.GOROOT(), "bin", "linux_386")); err != nil {
			build := exec.Command("go", "tool", "dist", "test", "-rebuild", "-run=no_tests")
			for key, value := range env {
				build.Env = append(build.Env, fmt.Sprintf("%v=%v", key, value))
			}
			utils.RunCmd(build, "setup linux go tools for sailfish")
		}
	}

	if target != "desktop" || strings.HasSuffix(target, "-docker") {
		utils.Log.Debugf("build target is %v; skipping installation of modules", target)
		return
	}

	var failed []string
	for _, module := range parser.GetLibs() {
		if !parser.ShouldBuildForTarget(module, target) {
			utils.Log.Debug("skipping installation of %v for %v", module, target)
			continue
		}

		mode := "full"
		if utils.QT_STUB() {
			mode = "stub"
		}
		utils.Log.Infof("installing %v qt/%v", mode, strings.ToLower(module))

		cmd := exec.Command("go", "install", "-p", strconv.Itoa(runtime.GOMAXPROCS(0)), "-v")
		if tagFlags != "" {
			cmd.Args = append(cmd.Args, tagFlags)
		}
		if target != "desktop" {
			cmd.Args = append(cmd.Args, []string{"-installsuffix", strings.Replace(target, "-", "_", -1)}...)
			cmd.Args = append(cmd.Args, []string{"-pkgdir", filepath.Join(utils.MustGoPath(), "pkg", fmt.Sprintf("%v_%v_%v", env["GOOS"], env["GOARCH"], strings.Replace(target, "-", "_", -1)))}...)
		}
		cmd.Args = append(cmd.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module)))
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}

		if _, err := utils.RunCmdOptionalError(cmd, fmt.Sprintf("install %v", strings.ToLower(module))); err != nil {
			failed = append(failed, strings.ToLower(module))
		}
	}

	if l := len(failed); l > 0 {
		utils.Log.Errorf("failed to install the following %v package(s):", l)
		for _, f := range failed {
			utils.Log.Error(f)
		}
		utils.Log.Errorf("you may want to run 'qtsetup install %v' if you depend on these package(s)", target)
	}
}
