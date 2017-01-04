package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GenModule(m string) {
	if !parser.ShouldBuild(m) {
		return
	}

	//prepare state
	parser.State.Module = m

	//generate
	//TODO: remove dry run -->
	CppTemplate(m)
	GoTemplate(m, false)
	//<--

	var suffix = func() string {
		switch m {
		case "AndroidExtras":
			{
				return "_android"
			}

		case "Sailfish":
			{
				return "_sailfish"
			}

		default:
			{
				return ""
			}
		}
	}()

	if !parser.State.Minimal {
		utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m)))
		utils.MkdirAll(utils.GoQtPkgPath(strings.ToLower(m)))
	}

	CgoTemplate(m, "")

	if parser.State.Minimal {
		if suffix != "" {
			return
		}

		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.cpp"), CppTemplate(m))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.h"), HTemplate(m))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.go"), GoTemplate(m, false))

		return
	}

	if m == "AndroidExtras" {
		utils.Save(utils.GoQtPkgPath(strings.ToLower(m), "utils-androidextras_android.go"), utils.Load(utils.GoQtPkgPath("internal", "binding", "files", "utils-androidextras_android.go")))
	}

	if !UseStub() {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".cpp"), CppTemplate(m))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".h"), HTemplate(m))
	}

	//always generate full
	if suffix != "" {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".go"), GoTemplate(m, false))
	}

	//may generate stub
	utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+".go"), GoTemplate(m, suffix != ""))
}
