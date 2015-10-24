package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/binding/templater"
)

func main() {
	fmt.Println("________________________Install________________________")

	for _, m := range templater.GetLibs() {
		var before = time.Now()

		fmt.Print(strings.ToLower(m))

		var out, err = exec.Command("go", "install", fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m))).CombinedOutput()
		if err != nil {
			fmt.Println(string(out), err)
			return
		}

		var sep = "\t"
		if len(m) < 8 {
			sep += "\t\t"
		}
		if len(m) >= 8 && len(m) < 17 {
			sep += "\t"
		}

		fmt.Printf("%v\t\t%v\n", sep, time.Since(before).String())
	}
}
