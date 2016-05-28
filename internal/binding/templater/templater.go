package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

var Minimal bool

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
		if !Minimal {
			utils.RemoveAll(utils.GetQtPkgPath(pkgName))
			utils.RemoveAll(utils.GetQtPkgPath("internal", "binding", "dump", name))
		}

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
		if Minimal {
			if name != "AndroidExtras" {
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+"-minimal"+suffix+".cpp"), CppTemplate("Qt"+name))
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+"-minimal"+suffix+".h"), HTemplate("Qt"+name))
			}
		} else {
			utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".cpp"), CppTemplate("Qt"+name))
			utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".h"), HTemplate("Qt"+name))
		}

		if name == "AndroidExtras" {
			if !Minimal {
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".go"), GoTemplate("Qt"+name, false))
			}
		}

		if Minimal {
			if name != "AndroidExtras" {
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+"-minimal.go"), GoTemplate("Qt"+name, name == "AndroidExtras"))
			}
		} else {
			utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+".go"), GoTemplate("Qt"+name, name == "AndroidExtras"))
		}
	}
}
