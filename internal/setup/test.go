package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildTarget = "desktop"
		ending      string
	)
	if len(os.Args) > 1 {
		buildTarget = os.Args[1]
	}
	if runtime.GOOS == "windows" {
		ending = ".exe"
	}

	fmt.Println("------------------------test----------------------------")

	//TODO:
	var gopath = os.Getenv("GOPATH")
	if gopath != "" {

		utils.MakeFolder(filepath.Join(gopath, "bin"))

		utils.RemoveAll(filepath.Join(gopath, "bin", fmt.Sprintf("qtdeploy%v", ending)))
		utils.RemoveAll(filepath.Join(gopath, "bin", fmt.Sprintf("qtmoc%v", ending)))

		runCmd(exec.Command("go", "build", "-o", filepath.Join(gopath, "bin", fmt.Sprintf("qtdeploy%v", ending)), utils.GoQtPkgPath("internal", "deploy", "deploy.go")), "qtdeploy")
		runCmd(exec.Command("go", "build", "-o", filepath.Join(gopath, "bin", fmt.Sprintf("qtmoc%v", ending)), utils.GoQtPkgPath("internal", "moc", "moc.go")), "qtmoc")

		switch runtime.GOOS {
		case "darwin", "linux":
			{
				utils.RemoveAll(filepath.Join("/usr", "local", "bin", "qtdeploy"))
				utils.RemoveAll(filepath.Join("/usr", "local", "bin", "qtmoc"))

				runCmdOptional(exec.Command("ln", "-s", filepath.Join(gopath, "bin", "qtdeploy"), filepath.Join("/usr", "local", "bin", "qtdeploy")), "symlink.qtdeploy")
				runCmdOptional(exec.Command("ln", "-s", filepath.Join(gopath, "bin", "qtmoc"), filepath.Join("/usr", "local", "bin", "qtmoc")), "symlink.qtmoc")
			}

		case "windows":
			{
				utils.RemoveAll(filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("qtdeploy%v", ending)))
				utils.RemoveAll(filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("qtmoc%v", ending)))

				var cmdDeploy = exec.Command("cmd", "/C", "mklink", "/H", fmt.Sprintf("qtdeploy%v", ending), filepath.Join(gopath, "bin", fmt.Sprintf("qtdeploy%v", ending)))
				cmdDeploy.Dir = filepath.Join(runtime.GOROOT(), "bin")
				runCmdOptional(cmdDeploy, "symlink.qtdeploy")

				var cmdMoc = exec.Command("cmd", "/C", "mklink", "/H", fmt.Sprintf("qtmoc%v", ending), filepath.Join(gopath, "bin", fmt.Sprintf("qtmoc%v", ending)))
				cmdMoc.Dir = filepath.Join(runtime.GOROOT(), "bin")
				runCmdOptional(cmdMoc, "symlink.qtmoc")
			}
		}

	} else {
		panic("GOPATH not set")
	}

	//TODO:
	for _, example := range []string{filepath.Join("widgets", "line_edits"), filepath.Join("widgets", "video_player"), filepath.Join("widgets", "graphicsscene"), filepath.Join("widgets", "dropsite"), filepath.Join("widgets", "table"),
		filepath.Join("quick", "bridge"), filepath.Join("quick", "bridge2"), filepath.Join("quick", "calc"), filepath.Join("quick", "dialog"), filepath.Join("quick", "sailfish"), filepath.Join("quick", "translate"), filepath.Join("quick", "view"),
		filepath.Join("qml", "application"), filepath.Join("qml", "prop"),
		filepath.Join("uitools", "calculator")} {

		if (buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && (!strings.Contains(example, "quick") || (example == filepath.Join("quick", "bridge") || example == filepath.Join("quick", "dialog"))) {
		} else if !(buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && example == filepath.Join("quick", "sailfish") {
		} else {

			var before = time.Now()

			fmt.Print(example)

			//TODO:
			runCmd(exec.Command(filepath.Join(gopath, "bin", "qtdeploy"), "test", buildTarget, filepath.Join(utils.GoQtPkgPath("internal", "examples"), example)), fmt.Sprintf("test.%v", example))

			fmt.Println(strings.Repeat(" ", 30-len(example)), time.Since(before)/time.Second*time.Second)
		}
	}
}

func runCmd(cmd *exec.Cmd, name string) {
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", name, out, err)
		os.Exit(1)
	}
}

func runCmdOptional(cmd *exec.Cmd, name string) {
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", name, out, err)
	}
}
