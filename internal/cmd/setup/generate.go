package setup

import (
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/utils"

	"github.com/therecipe/qt/internal/cmd"
)

func Generate(target string, docker, vagrant bool) {
	utils.Log.Infof("running: 'qtsetup generate %v' [docker=%v] [vagrant=%v]", target, docker, vagrant)
	if docker {
		cmd.Docker([]string{"/home/user/work/bin/qtsetup", "-debug", "generate"}, "linux", "", true)
		return
	}

	parser.LoadModules()

	for _, module := range parser.GetLibs() {
		if !parser.ShouldBuildForTarget(module, target) {
			utils.Log.Debugf("skipping generation of %v for %v", module, target)
			continue
		}

		mode := "full"
		switch {
		case target != runtime.GOOS:
			mode = "cgo"

		case utils.QT_STUB():
			mode = "stub"
		}
		utils.Log.Infof("generating %v qt/%v", mode, strings.ToLower(module))

		if target == runtime.GOOS || utils.QT_FAT() {
			templater.GenModule(module, target, templater.NONE)
		} else {
			templater.CgoTemplate(module, "", target, templater.MINIMAL, "", "") //TODO: collect errors
		}
	}
}
