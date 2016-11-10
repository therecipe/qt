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

	if utils.UseMsys2() {
		utils.RemoveAll(filepath.Join(os.Getenv("QT_MSYS2_DIR"), "usr", "bin", "go.exe"))
		utils.RunCmdOptional(exec.Command("cmd", "/C", "mklink", "/H", filepath.Join(os.Getenv("QT_MSYS2_DIR"), "usr", "bin", "go.exe"), filepath.Join(runtime.GOROOT(), "bin", "go.exe")), fmt.Sprintf("create go symlink in your QT_MSYS2_DIR/usr/bin (%v)", filepath.Join(utils.QT_MSYS2_DIR(), "..", "usr", "bin")))
	}

	utils.RunCmd(exec.Command("go", "get", "-v", "github.com/emirpasic/gods/lists/arraylist"), "install arraylist pkg") //needed for widgets/table example
	utils.RunCmd(exec.Command("go", "get", "-v", "github.com/fogleman/ln/ln"), "install ln pkg")                        //needed for widgets/renderer example

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
				if utils.UseMsys2() {
					utils.RemoveAll(filepath.Join(utils.QT_MSYS2_DIR(), "..", "usr", "bin", fmt.Sprintf("%v.exe", app)))
					utils.RunCmdOptional(exec.Command("cmd", "/C", "mklink", "/H", filepath.Join(utils.QT_MSYS2_DIR(), "..", "usr", "bin", fmt.Sprintf("%v.exe", app)), filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("%v.exe", app))), fmt.Sprintf("create %v symlink in your PATH (GOROOT/bin) -> use %v instead", app, filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("%v.exe", app))))
				} else {
					utils.RemoveAll(filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("%v.exe", app)))
					utils.RunCmdOptional(exec.Command("cmd", "/C", "mklink", "/H", filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("%v.exe", app)), filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("%v.exe", app))), fmt.Sprintf("create %v symlink in your PATH (GOROOT/bin) -> use %v instead", app, filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("%v.exe", app))))
				}
			}
		}
	}
}
