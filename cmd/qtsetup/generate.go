package main

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

func generate(buildTarget string) {
	utils.Log.Info("running setup/generate")

	if testFile := utils.GoQtPkgPath("core", "cgo_desktop_darwin_amd64.go"); utils.Exists(testFile) && strings.Contains(utils.Load(testFile), utils.QT_DIR()) {
		if buildTarget != "desktop" && utils.QT_STUB() &&
			!utils.Exists(utils.GoQtPkgPath("core", "core.h")) &&
			!utils.Exists(utils.GoQtPkgPath("core", "core.cpp")) {

			utils.Log.Debug("stub files are up to date -> don't re-generate")
			return
		}

		if buildTarget != "desktop" && !utils.QT_STUB() &&
			utils.Exists(utils.GoQtPkgPath("core", "core.h")) &&
			utils.Exists(utils.GoQtPkgPath("core", "core.cpp")) {

			utils.Log.Debug("real files are up to date -> don't re-generate")
			return
		}
	}

	for _, module := range templater.GetLibs() {
		utils.Log.Debugf("loading qt/%v", strings.ToLower(module))
		if _, err := parser.GetModule(module); err != nil {
			utils.Log.WithError(err).Errorf("failed to load qt/%v", strings.ToLower(module))
		}
	}

	for _, module := range templater.GetLibs() {
		utils.Log.Infof("generating%v qt/%v", func() string {
			if utils.QT_STUB() {
				return " stub"
			}
			return " full"
		}(), strings.ToLower(module))
		templater.GenModule(module)
	}
}
