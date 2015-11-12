package templater

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func terminalSeperator(module string) string {
	var sep = ""
	if len(module) < 8 {
		sep += "\t\t"
	}
	if len(module) >= 8 && len(module) < 17 {
		sep += "\t"
	}
	return sep
}

func GenModule(name string) {
	var (
		moduleT = name
		cppTmp  string
	)

	name = strings.ToLower(name)

	parser.GetModule(name)

	if ShouldBuild(moduleT) {

		utils.RemoveAll(utils.GetQtPkgPath(name))
		utils.RemoveAll(utils.GetQtPkgPath("internal", "binding", "dump", moduleT))

		var fcount = 0
		for _, c := range parser.ClassMap {

			if moduleT == strings.TrimPrefix(c.Module, "Qt") {

				//Dry run to catch all unsupported
				for _, f := range map[string]func(*parser.Class) string{"cpp": CppTemplate, "h": HTemplate, "go": GoTemplate} {
					f(c)
				}

				for e, f := range map[string]func(*parser.Class) string{"cpp": CppTemplate, "h": HTemplate, "go": GoTemplate} {
					utils.MakeFolder(utils.GetQtPkgPath(name))
					CopyCgo(moduleT)
					if moduleT == "AndroidExtras" {
						utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v_android.%v", strings.ToLower(c.Name), e)), f(c))
						if e == "go" {
							c.Stub = true
							utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v.%v", strings.ToLower(c.Name), e)), f(c))
							c.Stub = false
						}
					} else {
						if e == "cpp" && runtime.GOOS == "windows" {
							cppTmp += f(c)
						} else {
							utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v.%v", strings.ToLower(c.Name), e)), f(c))
						}
					}
				}

				c.Dump()

				fcount += len(c.Functions)
			}

		}
		utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v.cpp", strings.ToLower(moduleT))), cppTmp)
		fmt.Printf("%v%v\tfunctions:%v", name, terminalSeperator(name), fcount)
	}
}
