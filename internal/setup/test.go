package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		ending      string
		buildTarget = "desktop"
	)
	if len(os.Args) > 1 {
		buildTarget = os.Args[1]
	}
	if runtime.GOOS == "windows" {
		ending = ".exe"
	}

	fmt.Printf("_________________________Test-%v__________________________\n", buildTarget)

	runCmd(exec.Command("go", "build", "-o", filepath.Join(runtime.GOROOT(), "bin", fmt.Sprintf("qtdeploy%v", ending)), utils.GetQtPkgPath("internal", "deploy", "deploy.go")), "qtdeploy")

	for _, example := range []string{filepath.Join("widgets", "line_edits"), filepath.Join("widgets", "video_player"), filepath.Join("widgets", "graphicsscene"), filepath.Join("quick", "calc"), filepath.Join("quick", "dialog"), filepath.Join("quick", "view"), filepath.Join("qml", "application"), filepath.Join("qml", "prop"), filepath.Join("uitools", "calculator")} {
		var before = time.Now()

		fmt.Print(example)

		runCmd(exec.Command(fmt.Sprintf("qtdeploy%v", ending), "test", buildTarget, filepath.Join(utils.GetQtPkgPath("internal", "examples"), example)), fmt.Sprintf("test.%v", example))

		var sep = "\t"
		if len(example) < 8 {
			sep += "\t\t"
		}
		if len(example) >= 8 && len(example) < 17 {
			sep += "\t"
		}

		fmt.Printf("%v\t\t%v\n", sep, time.Since(before)/time.Second*time.Second)
	}

}

func runCmd(cmd *exec.Cmd, n string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", n, out, err)
		os.Exit(1)
	}
	return string(out)
}
