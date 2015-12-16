package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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
			fmt.Println(name(module), functions(module), time.Since(before)/time.Second*time.Second)

			var out, err = exec.Command("go", "fmt", fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module))).CombinedOutput()
			if err != nil {
				fmt.Println(string(out), err)
				os.Exit(1)
			}
		}
	}
}

func name(name string) string {
	return fmt.Sprintf("%v%v\t", strings.ToLower(name), strings.Repeat(" ", 17-len(name)))
}

func functions(name string) string {
	var (
		sC, usC = functionsCount(name)
		s       = strconv.Itoa(sC)
		us      = strconv.Itoa(usC)
	)
	return fmt.Sprintf("funcs: %v%v (-%v%v)\t", strings.Repeat(" ", 5-len(s)), s, strings.Repeat(" ", 5-len(us)), us)
}

func functionsCount(name string) (int, int) {
	var (
		supported   int
		unsupported int
	)

	for _, c := range parser.ClassMap {
		if strings.TrimPrefix(c.Module, "Qt") == name {
			for _, f := range c.Functions {
				if f.Access == "public" || f.Access == "protected" {
					supported++
				} else {
					unsupported++
				}
			}
		}
	}

	return supported, unsupported
}
