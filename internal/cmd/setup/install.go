package setup

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/utils"
)

func Install(target string, docker, vagrant bool) {
	utils.Log.Infof("running: 'qtsetup install %v' [docker=%v] [vagrant=%v]", target, docker, vagrant)

	if target == "sailfish" || target == "sailfish-emulator" {
		if _, err := ioutil.ReadDir(filepath.Join(runtime.GOROOT(), "bin", "linux_386")); err != nil {
			build := exec.Command("go", "tool", "dist", "test", "-rebuild", "-run=no_tests")

			env := map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": utils.GOPATH(),
				"GOROOT": runtime.GOROOT(),

				"GOOS":   "linux",
				"GOARCH": "386",
			}

			if runtime.GOOS == "windows" {
				env["TMP"] = os.Getenv("TMP")
				env["TEMP"] = os.Getenv("TEMP")
			}

			for k, v := range env {
				build.Env = append(build.Env, fmt.Sprintf("%v=%v", k, v))
			}

			utils.RunCmd(build, "setup linux go tools for sailfish")
		}
	}

	if target != runtime.GOOS && !utils.QT_FAT() {
		utils.Log.Debugf("target is %v; skipping installation of modules", target)
		return
	}

	env, tags, _, _ := cmd.BuildEnv(target, "", "")
	var failed []string
	for _, module := range parser.GetLibs() {
		if !parser.ShouldBuildForTarget(module, target) {
			utils.Log.Debug("skipping installation of %v for %v", module, target)
			continue
		}

		mode := "full"
		if utils.QT_STUB() || docker {
			mode = "stub"
		}
		utils.Log.Infof("installing %v qt/%v", mode, strings.ToLower(module))

		cmd := exec.Command("go", "install", "-p", strconv.Itoa(runtime.GOMAXPROCS(0)), "-v")
		if len(tags) > 0 {
			cmd.Args = append(cmd.Args, fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
		}
		if target != runtime.GOOS {
			cmd.Args = append(cmd.Args, []string{"-pkgdir", filepath.Join(utils.MustGoPath(), "pkg", fmt.Sprintf("%v_%v_%v", strings.Replace(target, "-", "_", -1), env["GOOS"], env["GOARCH"]))}...)
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
		utils.Log.Warn("failed to install:")
		for _, f := range failed {
			utils.Log.Warn(f)
		}
	}
}
