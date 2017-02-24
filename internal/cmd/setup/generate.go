package setup

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

func generate(buildTarget string) {
	utils.Log.Info("running setup/generate")

	if testFile := utils.GoQtPkgPath("core", "cgo_desktop_darwin_amd64.go"); utils.ExistsFile(testFile) && strings.Contains(utils.Load(testFile), utils.QT_DIR()) {
		if buildTarget != "desktop" && utils.QT_STUB() &&
			!utils.ExistsFile(utils.GoQtPkgPath("core", "core.h")) &&
			!utils.ExistsFile(utils.GoQtPkgPath("core", "core.cpp")) {

			utils.Log.Debug("stub files are up to date -> don't re-generate")
			return
		}

		if buildTarget != "desktop" && !utils.QT_STUB() &&
			utils.ExistsFile(utils.GoQtPkgPath("core", "core.h")) &&
			utils.ExistsFile(utils.GoQtPkgPath("core", "core.cpp")) {

			utils.Log.Debug("full files are up to date -> don't re-generate")
			return
		}
	}

	parser.LoadModules()

	for _, module := range parser.GetLibs() {
		utils.Log.Infof("generating %v qt/%v", func() string {
			if utils.QT_STUB() {
				return "stub"
			}
			return "full"
		}(), strings.ToLower(module))
		templater.GenModule(module, buildTarget)
	}
}
