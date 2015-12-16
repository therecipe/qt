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
		fmt.Println("________________________Generate________________________")

		for _, module := range templater.GetLibs() {
			parser.GetModule(strings.ToLower(module))
		}

		for _, module := range templater.GetLibs() {

			var before = time.Now()
			templater.GenModule(module)
			fmt.Printf("\t%v\n", time.Since(before)/time.Second*time.Second)

			var out, err = exec.Command("go", "fmt", fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module))).CombinedOutput()
			if err != nil {
				fmt.Println(string(out), err)
				os.Exit(1)
			}
		}
	}
}
