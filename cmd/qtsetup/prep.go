package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/therecipe/qt/internal/utils"
)

func prep() {
	utils.Log.Info("running setup/prep")

	utils.RunCmd(exec.Command("go", "get", "golang.org/x/crypto/ssh"), "install ssh pkg")
	utils.RunCmd(exec.Command("go", "get", "github.com/emirpasic/gods/lists/arraylist"), "install arraylist pkg")
	utils.RunCmd(exec.Command("go", "get", "github.com/Sirupsen/logrus"), "install logrus pkg")
	utils.RunCmd(exec.Command("go", "install", "github.com/therecipe/qt/cmd/..."), "install cmd pkgs")

	for _, app := range []string{"qtdeploy", "qtmoc", "qtrcc", "qtminimal"} {
		switch runtime.GOOS {
		case "darwin", "linux":
			{
				if os.Geteuid() == 0 || runtime.GOOS == "darwin" {
					utils.RemoveAll(fmt.Sprintf("/usr/local/bin/%v", app))
					var err = os.Symlink(filepath.Join(utils.MustGoPath(), "bin", app), fmt.Sprintf("/usr/local/bin/%v", app))
					if err != nil {
						utils.Log.WithError(err).Warnf("failed to create %v symlink in your PATH (/usr/local/bin) -> use %v instead", app, filepath.Join(utils.MustGoPath(), "bin", app))
					}
				} else {
					utils.Log.Warnf("(not root) failed to create %v symlink in your PATH (/usr/local/bin) -> use %v instead", app, filepath.Join(utils.MustGoPath(), "bin", app))
				}
			}

		case "windows":
			{
				utils.RemoveAll(filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("%v.exe", app)))
				utils.RunCmdOptional(exec.Command("cmd", "/C", "mklink", "/H", filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("%v.exe", app)), filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("%v.exe", app))), fmt.Sprintf("create %v symlink in your PATH (GOROOT/bin) -> use %v instead", app, filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("%v.exe", app))))
			}
		}
	}
}
