package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/utils"

	"github.com/therecipe/qt/internal/binding/parser"
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
	var moduleT = name

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
					utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v.%v", strings.ToLower(c.Name), e)), f(c))
					CopyCgo(moduleT)
				}

				c.Dump()

				fcount += len(c.Functions)
			}

		}
		fmt.Printf("%v%v\tfunctions:%v", name, terminalSeperator(name), fcount)
	}

}
