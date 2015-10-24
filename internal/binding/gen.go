package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/binding/templater"
)

func main() {
	fmt.Println("________________________Generate________________________")

	for _, module := range templater.GetLibs() {

		var before = time.Now()
		templater.GenModule(module)
		fmt.Printf("\t%v\n", time.Since(before).String())

		var out, err = exec.Command("go", "fmt", fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module))).CombinedOutput()
		if err != nil {
			fmt.Println(string(out), err)
			return
		}

	}

}
