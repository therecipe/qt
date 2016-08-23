package templater

import (
	"strings"

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
		if name == "Sailfish" {
			suffix = "_sailfish"
		}

		//cleanup
		if !Minimal {
			utils.RemoveAll(utils.GetQtPkgPath(pkgName))
			utils.RemoveAll(utils.GetQtPkgPath("internal", "binding", "dump", name))
		}

		//prepare
		utils.MakeFolder(utils.GetQtPkgPath(pkgName))
		CopyCgo(name)
		if name == "AndroidExtras" && !Minimal {
			utils.Save(utils.GetQtPkgPath(pkgName, "utils-androidextras_android.go"), utils.Load(utils.GetQtPkgPath("internal", "binding", "files", "utils-androidextras_android.go")))
		}
		manualWeakLink("Qt" + name)

		//dry run
		CppTemplate("Qt" + name)
		GoTemplate("Qt"+name, false)

		//generate
		if Minimal {
			if !(name == "AndroidExtras" || name == "Sailfish") {
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+"-minimal"+suffix+".cpp"), CppTemplate("Qt"+name))
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+"-minimal"+suffix+".h"), HTemplate("Qt"+name))
			}
		} else {
			utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".cpp"), CppTemplate("Qt"+name))
			utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".h"), HTemplate("Qt"+name))
		}

		if Minimal {
			if !(name == "AndroidExtras" || name == "Sailfish") {
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+"-minimal.go"), GoTemplate("Qt"+name, false))
			}
		} else {
			if name == "AndroidExtras" || name == "Sailfish" {
				utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+suffix+".go"), GoTemplate("Qt"+name, false))
			}
			utils.SaveBytes(utils.GetQtPkgPath(pkgName, pkgName+".go"), GoTemplate("Qt"+name, name == "AndroidExtras" || name == "Sailfish"))
		}
	}
}
