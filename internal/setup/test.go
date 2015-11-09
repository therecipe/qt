package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"time"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var buildTarget = "desktop"
	if len(os.Args) > 1 {
		buildTarget = os.Args[1]
	}

	fmt.Printf("_________________________Test-%v__________________________\n", buildTarget)

	runCmd(exec.Command("go", "build", "-o", path.Join(runtime.GOROOT(), "bin", "qtdeploy"), utils.GetQtPkgPath("internal", "deploy", "deploy.go")), "qtdeploy")

	for _, example := range []string{"widgets", path.Join("quick", "calc"), path.Join("quick", "dialog"), path.Join("quick", "view"), path.Join("qml", "application"), path.Join("qml", "prop")} {
		var before = time.Now()

		fmt.Print(example)

		runCmd(exec.Command("qtdeploy", "test", buildTarget, path.Join(utils.GetQtPkgPath("internal", "examples"), example)), fmt.Sprintf("test.%v", example))

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
		fmt.Printf("=>%v<=\noutput:%s\nerror:%s\n\n", n, out, err)
		os.Exit(1)
	}
	return string(out)
}
