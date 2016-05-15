package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("------------------------generate------------------------")

		for _, module := range templater.GetLibs() {
			parser.GetModule(strings.ToLower(module))
		}

		for _, module := range templater.GetLibs() {

			var before = time.Now()

			fmt.Print(strings.ToLower(module))

			templater.GenModule(module)

			fmt.Println(strings.Repeat(" ", 30-len(module)), time.Since(before)/time.Second*time.Second)

			//TODO: replace with go/format.Source
			var out, err = exec.Command("go", "fmt", fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module))).CombinedOutput()
			if err != nil {
				fmt.Println(string(out), err)
				os.Exit(1)
			}
		}
	}
}
