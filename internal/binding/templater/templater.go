package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GenModule(name string) {
	if ShouldBuild(name) {

		var (
			pkgName = strings.ToLower(name)
			suffix  string
		)

		if name == "AndroidExtras" {
			suffix = "_android"
		}

		//cleanup
		utils.RemoveAll(utils.GetQtPkgPath(pkgName))
		utils.RemoveAll(utils.GetQtPkgPath("internal", "binding", "dump", name))

		//prepare
		utils.MakeFolder(utils.GetQtPkgPath(pkgName))
		CopyCgo(name)
		if name == "AndroidExtras" {
			utils.Save(utils.GetQtPkgPath(pkgName, "utils-androidextras_android.go"), utils.Load(utils.GetQtPkgPath("internal", "binding", "files", "utils-androidextras_android.go")))
		}
		manualWeakLink("Qt" + name)
		for _, c := range parser.ClassMap {
			if strings.TrimPrefix(c.Module, "Qt") == name {
				addCallbackNameFunctions(c)
			}
		}

		//dry run
		CppTemplate("Qt" + name)
		GoTemplate("Qt"+name, false)

		//generate
		utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".cpp"), CppTemplate("Qt"+name))
		utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".h"), HTemplate("Qt"+name))

		if name == "AndroidExtras" {
			utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".go"), GoTemplate("Qt"+name, false))
		}
		utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+".go"), GoTemplate("Qt"+name, name == "AndroidExtras"))
	}
}
