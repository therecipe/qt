package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GenModule(m, buildTarget string, mode int) {
	if !parser.ShouldBuildForTarget(m, buildTarget) {
		utils.Log.WithField("0_module", m).Debug("skip generation")
		return
	}
	utils.Log.WithField("0_module", m).Debug("generating")

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

	if mode == NONE {
		utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m)))
		utils.MkdirAll(utils.GoQtPkgPath(strings.ToLower(m)))
	}

	if !UseStub(false, "Qt"+m, mode) {
		CgoTemplate(m, "", buildTarget, mode, m)
	}

	if mode == MINIMAL {
		if suffix != "" {
			return
		}

		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.cpp"), CppTemplate(m, mode))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.h"), HTemplate(m, mode))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.go"), GoTemplate(m, false, mode, m))

		return
	}

	if m == "AndroidExtras" {
		utils.Save(utils.GoQtPkgPath(strings.ToLower(m), "utils-androidextras_android.go"), utils.Load(utils.GoQtPkgPath("internal", "binding", "files", "utils-androidextras_android.go")))
	}

	if !UseStub(false, "Qt"+m, mode) {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".cpp"), CppTemplate(m, mode))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".h"), HTemplate(m, mode))
	}

	//always generate full
	if suffix != "" {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".go"), GoTemplate(m, false, mode, m))
	}

	//may generate stub
	utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+".go"), GoTemplate(m, suffix != "", mode, m))
}
