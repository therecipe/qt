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
		ending = func() string {
			if runtime.GOOS == "windows" {
				return ".exe"
			}
			return ""
		}()

		buildTarget = func() string {
			if len(os.Args) > 1 {
				return os.Args[1]
			}
			return "desktop"
		}()
	)

	fmt.Println("------------------------test----------------------------")

	//TODO:
	runCmd(exec.Command("go", "build", "-o", filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("qtdeploy%v", ending)), utils.GetQtPkgPath("internal", "deploy", "deploy.go")), "qtdeploy")
	runCmd(exec.Command("go", "build", "-o", filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("qtmoc%v", ending)), utils.GetQtPkgPath("internal", "moc", "moc.go")), "qtmoc")

	//TODO:
	for _, example := range []string{filepath.Join("widgets", "line_edits"), filepath.Join("widgets", "video_player"), filepath.Join("widgets", "graphicsscene"),
		filepath.Join("quick", "bridge"), filepath.Join("quick", "bridge2"), filepath.Join("quick", "calc"), filepath.Join("quick", "dialog"), filepath.Join("quick", "translate"), filepath.Join("quick", "view"),
		filepath.Join("qml", "application"), filepath.Join("qml", "prop"),
		filepath.Join("uitools", "calculator")} {

		var before = time.Now()

		fmt.Print(example)

		//TODO:
		runCmd(exec.Command(fmt.Sprintf("qtdeploy%v", ending), "test", buildTarget, filepath.Join(utils.GetQtPkgPath("internal", "examples"), example)), fmt.Sprintf("test.%v", example))

		fmt.Println(strings.Repeat(" ", 30-len(example)), time.Since(before)/time.Second*time.Second)
	}

}

func runCmd(cmd *exec.Cmd, name string) {
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", name, out, err)
		os.Exit(1)
	}
}
