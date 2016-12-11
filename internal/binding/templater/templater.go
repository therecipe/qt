package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GenModule(m string) {
	if !ShouldBuild(m) {
		return
	}

	var (
		pkgName = strings.ToLower(m)
		suffix  string
	)

	switch m {
	case "AndroidExtras":
		{
			suffix = "_android"
		}

	case "Sailfish":
		{
			suffix = "_sailfish"
		}
	}

	//cleanup
	if !parser.CurrentState.Minimal {
		utils.RemoveAll(utils.GoQtPkgPath(pkgName))
		utils.MkdirAll(utils.GoQtPkgPath(pkgName))

		if m == "AndroidExtras" {
			//TODO: move into internal/runtime
			utils.Save(utils.GoQtPkgPath(pkgName, "utils-androidextras_android.go"), utils.Load(utils.GoQtPkgPath("internal", "binding", "files", "utils-androidextras_android.go")))
		}
	}

	//prepare
	CopyCgo(m, "")

	parser.CurrentState.CurrentModule = m

	//dry run
	CppTemplate("Qt" + m)
	GoTemplate("Qt"+m, false)

	//generate
	if parser.CurrentState.Minimal {
		if suffix != "" {
			return
		}

		utils.SaveBytes(utils.GoQtPkgPath(pkgName, pkgName+"-minimal.cpp"), CppTemplate("Qt"+m))
		utils.SaveBytes(utils.GoQtPkgPath(pkgName, pkgName+"-minimal.h"), HTemplate("Qt"+m))
		utils.SaveBytes(utils.GoQtPkgPath(pkgName, pkgName+"-minimal.go"), GoTemplate("Qt"+m, false))

		return
	}

	if !UseStub() {
		utils.SaveBytes(utils.GoQtPkgPath(pkgName, pkgName+suffix+".cpp"), CppTemplate("Qt"+m))
		utils.SaveBytes(utils.GoQtPkgPath(pkgName, pkgName+suffix+".h"), HTemplate("Qt"+m))
	}

	//always generate full
	if suffix != "" {
		utils.SaveBytes(utils.GoQtPkgPath(pkgName, pkgName+suffix+".go"), GoTemplate("Qt"+m, false))
	}

	//may generate stub
	utils.SaveBytes(utils.GoQtPkgPath(pkgName, pkgName+".go"), GoTemplate("Qt"+m, suffix != ""))
}
