package templater

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/akiyosi/qt/internal/binding/parser"
	"github.com/akiyosi/qt/internal/utils"
)

func GenModule(m, target string, mode int) {
	if os.Getenv("QT_DUMP") == "true" {
		defer parser.Dump()
	}

	if !parser.ShouldBuildForTarget(m, target) {
		utils.Log.WithField("module", m).Debug("skip generation")
		return
	}
	utils.Log.WithField("module", m).Debug("generating")

	var suffix string
	switch m {
	case "AndroidExtras":
		suffix = "_android"

	case "Sailfish":
		suffix = "_sailfish"
	}

	if mode == NONE {
		utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m)))
		utils.MkdirAll(utils.GoQtPkgPath(strings.ToLower(m)))
	}

	if mode == MINIMAL { //TODO: dead code ?
		if suffix != "" {
			return
		}

		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.cpp"), CppTemplate(m, mode, target, ""))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.h"), HTemplate(m, mode, ""))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.go"), GoTemplate(m, false, mode, m, target, ""))

		if !(UseStub(false, "Qt"+m, mode) || utils.QT_GEN_GO_WRAPPER()) {
			CgoTemplate(m, "", target, mode, m, "")
		}

		return
	}

	if !(UseStub(false, "Qt"+m, mode) || utils.QT_GEN_GO_WRAPPER()) {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".cpp"), CppTemplate(m, mode, target, ""))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".h"), HTemplate(m, mode, ""))
	}

	//always generate full
	if suffix != "" {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".go"), GoTemplate(m, false, mode, m, target, ""))
	}

	//may generate stub
	utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+".go"), GoTemplate(m, suffix != "", mode, m, target, ""))

	if !(UseStub(false, "Qt"+m, mode) || utils.QT_GEN_GO_WRAPPER()) {
		CgoTemplate(m, "", target, mode, m, "")
	}

	if utils.QT_DOCKER() {
		if idug, ok := os.LookupEnv("IDUG"); ok {
			if path := utils.GoQtPkgPath(strings.ToLower(m)); utils.UseGOMOD(path) {
				utils.RunCmd(exec.Command("chown", "-R", idug, filepath.Dir(utils.GOMOD(path))), "chown files to user")
			} else {
				utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
			}
		}
	}
}
