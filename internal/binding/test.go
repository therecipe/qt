package main

import (
	"fmt"
	"os/exec"
	"path"
	"runtime"
	"time"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	fmt.Println("_________________________Test__________________________")

	var out, err = exec.Command("go", "build", "-o", utils.GetQtPkgPath("internal", "deploy", "deploy"), utils.GetQtPkgPath("internal", "deploy", fmt.Sprintf("deploy_%v.go", runtime.GOOS))).CombinedOutput()
	if err != nil {
		fmt.Println("test", string(out), err)
	} else {
		for _, example := range []string{"widgets", path.Join("quick", "calc"), path.Join("quick", "view"), path.Join("qml", "application"), path.Join("qml", "prop")} {
			var before = time.Now()

			fmt.Print(example)

			exec.Command(utils.GetQtPkgPath("internal", "deploy", "deploy"), path.Join(utils.GetQtPkgPath("internal", "examples"), example)).Run()

			var sep = "\t"
			if len(example) < 8 {
				sep += "\t\t"
			}
			if len(example) >= 8 && len(example) < 17 {
				sep += "\t"
			}

			fmt.Printf("%v\t\t%v\n", sep, time.Since(before).String())
		}
	}
}
