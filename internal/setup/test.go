package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildTarget = func() string {
			if len(os.Args) > 1 {
				return os.Args[1]
			}
			return "desktop"
		}()

		ending = func() string {
			if runtime.GOOS == "windows" {
				return ".exe"
			}
			return ""
		}()
	)

	utils.Log.Infof("running setup/test.go %v", buildTarget)

	utils.MakeFolder(filepath.Join(utils.MustGoPath(), "bin"))

	utils.RemoveAll(filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("qtdeploy%v", ending)))
	utils.RemoveAll(filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("qtmoc%v", ending)))

	utils.RunCmd(exec.Command("go", "build", "-o", filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("qtdeploy%v", ending)), utils.GoQtPkgPath("internal", "deploy", "deploy.go")), "qtdeploy")
	utils.RunCmd(exec.Command("go", "build", "-o", filepath.Join(utils.MustGoPath(), "bin", fmt.Sprintf("qtmoc%v", ending)), utils.GoQtPkgPath("internal", "moc", "moc.go")), "qtmoc")

	switch runtime.GOOS {
	case "darwin", "linux":
		{
			if os.Geteuid() == 0 || runtime.GOOS == "darwin" {
				utils.RemoveAll("/usr/local/bin/qtdeploy")
				utils.RemoveAll("/usr/local/bin/qtmoc")

				var err = os.Symlink(filepath.Join(utils.MustGoPath(), "bin/qtdeploy"), "/usr/local/bin/qtdeploy")
				if err != nil {
					utils.Log.WithError(err).Errorf("failed to create qtdeploy symlink in your PATH -> use %v instead", filepath.Join(utils.MustGoPath(), "bin/qtdeploy"))
				}

				err = os.Symlink(filepath.Join(utils.MustGoPath(), "bin/qtmoc"), "/usr/local/bin/qtmoc")
				if err != nil {
					utils.Log.WithError(err).Errorf("failed to create qtmoc symlink in your PATH -> use %v instead", filepath.Join(utils.MustGoPath(), "bin/qtmoc"))
				}
			} else {
				utils.Log.Errorf("failed to create qtdeploy and qtmoc symlinks in your PATH (not root) -> use %v and %v instead", filepath.Join(utils.MustGoPath(), "bin/qtdeploy"), filepath.Join(utils.MustGoPath(), "bin/qtmoc"))
			}
		}

	case "windows":
		{
			utils.RemoveAll(filepath.Join(runtime.GOROOT(), "bin\\qtdeploy.exe"))
			utils.RemoveAll(filepath.Join(runtime.GOROOT(), "bin\\qtmoc.exe"))

			utils.RunCmdOptional(exec.Command("cmd", "/C", "mklink", "/H", filepath.Join(runtime.GOROOT(), "bin\\qtdeploy.exe"), filepath.Join(utils.MustGoPath(), "bin\\qtdeploy.exe")), "symlink.qtdeploy")
			utils.RunCmdOptional(exec.Command("cmd", "/C", "mklink", "/H", filepath.Join(runtime.GOROOT(), "bin\\qtmoc.exe"), filepath.Join(utils.MustGoPath(), "bin\\qtmoc.exe")), "symlink.qtmoc")
		}
	}

	utils.Log.Info("starting to test examples (~10min)")

	//TODO: cleanup
	for _, example := range []string{filepath.Join("widgets", "line_edits"), filepath.Join("widgets", "video_player"), filepath.Join("widgets", "graphicsscene"), filepath.Join("widgets", "dropsite"), filepath.Join("widgets", "table"),
		filepath.Join("quick", "bridge"), filepath.Join("quick", "bridge2"), filepath.Join("quick", "calc"), filepath.Join("quick", "dialog"), filepath.Join("quick", "sailfish"), filepath.Join("quick", "translate"), filepath.Join("quick", "view"),
		filepath.Join("qml", "application"), filepath.Join("qml", "prop"),
		filepath.Join("uitools", "calculator")} {

		if (buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && (!strings.Contains(example, "quick") || (example == filepath.Join("quick", "bridge") || example == filepath.Join("quick", "dialog"))) {
		} else if !(buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && example == filepath.Join("quick", "sailfish") {
		} else {

			utils.Log.Infoln("testing", example)

			utils.RunCmd(exec.Command(filepath.Join(utils.MustGoPath(), "bin", "qtdeploy"),

				"test",

				strings.TrimSuffix(buildTarget, "-docker"),

				filepath.Join(utils.GoQtPkgPath("internal", "examples"), example),

				func() string {
					if strings.HasSuffix(buildTarget, "-docker") {
						return "docker"
					}
					return ""
				}()),

				fmt.Sprintf("test.%v", example))
		}
	}
}
