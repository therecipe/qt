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
	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/utils"
)

func Install(target string, docker, vagrant, failfast bool) {
	utils.Log.Infof("running: 'qtsetup install %v' [docker=%v] [vagrant=%v]", target, docker, vagrant)

	if strings.HasPrefix(target, "sailfish") && !utils.QT_SAILFISH() {
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

	if !(target == runtime.GOOS || target == "js" || target == "wasm") && !utils.QT_FAT() {
		utils.Log.Debugf("target is %v; skipping installation of modules", target)
		return
	}

	env, tags, _, _ := cmd.BuildEnv(target, "", "")
	var failed []string
	for _, module := range parser.GetLibs() {
		if !parser.ShouldBuildForTarget(module, target) {
			utils.Log.Debugf("skipping installation of %v for %v", module, target)
			continue
		}

		mode := "full"
		if utils.QT_STUB() || docker {
			mode = "stub"
		}

		var license string
		switch module {
		case "Charts", "DataVisualization", "VirtualKeyboard":
			license = strings.Repeat(" ", 21-len(module)) + "[GPLv3]"
		}
		utils.Log.Infof("installing %v qt/%v%v", mode, strings.ToLower(module), license)

		if utils.QT_DYNAMIC_SETUP() && mode == "full" {
			cc, com := templater.ParseCgo(strings.ToLower(module), target)
			if cc != "" {
				if target == "js" || target == "wasm" {
					com = strings.Replace(com, strings.ToLower(module)+".js_plugin_import.cpp ", "", -1)
				}
				cmd := exec.Command(cc, strings.Split(com, " ")...)
				cmd.Dir = utils.GoQtPkgPath(strings.ToLower(module))
				if target == "js" || target == "wasm" {
					for key, value := range env {
						cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
					}
				}
				utils.RunCmdOptional(cmd, fmt.Sprintf("failed to create dynamic lib for %v (%v) on %v", target, strings.ToLower(module), runtime.GOOS))

				if !(target == "js" || target == "wasm") {
					utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(module), strings.ToLower(module)+".cpp"))
					utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(module), "_obj"))
					templater.ReplaceCgo(strings.ToLower(module), target)
				}
			}
		}

		if target == "js" || target == "wasm" {
			env["CGO_ENABLED"] = "0"
		}

		cmd := exec.Command("go", "install", "-p", strconv.Itoa(runtime.GOMAXPROCS(0)), "-v")
		if len(tags) > 0 {
			cmd.Args = append(cmd.Args, fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
		}
		if target != runtime.GOOS {
			cmd.Args = append(cmd.Args, []string{"-pkgdir", filepath.Join(utils.MustGoPath(), "pkg", fmt.Sprintf("%v_%v_%v", strings.Replace(target, "-", "_", -1), env["GOOS"], env["GOARCH"]))}...)
		}

		if target == "js" {
			cmd = exec.Command(filepath.Join(utils.GOBIN(), "gopherjs"), "install")
		}

		cmd.Args = append(cmd.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module)))

		if target == "js" {
			cmd.Args = append(cmd.Args, "-v")
		} else {
			if target == "linux" {
				env["CGO_LDFLAGS"] = strings.Replace(env["CGO_LDFLAGS"], "-Wl,-rpath,$ORIGIN/lib -Wl,--disable-new-dtags", "", -1)
			}
			for key, value := range env {
				cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
			}
		}

		if msg, err := utils.RunCmdOptionalError(cmd, fmt.Sprintf("install %v", strings.ToLower(module))); err != nil {
			println(msg)
			failed = append(failed, strings.ToLower(module))
			if strings.ToLower(module) == "core" || failfast {
				utils.Log.Errorf("failed to install '%v'; aborting setup", strings.ToLower(module))
				os.Exit(1)
			}
		}
	}

	if l := len(failed); l > 0 {
		utils.Log.Warn("failed to install:")
		for _, f := range failed {
			utils.Log.Warn(f)
		}
	}
}
