package setup

import (
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/utils"
)

func Generate(target string) {
	utils.Log.Infof("running: 'qtsetup generate %v'", target)

	parser.LoadModules()

	for _, module := range parser.GetLibs() {
		if !parser.ShouldBuildForTarget(module, target) {
			utils.Log.Debug("skipping generation of %v for %v", module, target)
			continue
		}

		mode := "full"
		if utils.QT_STUB() {
			mode = "stub"
		}
		utils.Log.Infof("generating %v qt/%v", mode, strings.ToLower(module))

		if target == runtime.GOOS {
			templater.GenModule(module, target, templater.NONE)
		} else {
			templater.CgoTemplate(module, "", target, templater.MINIMAL, "")
		}
	}
}
