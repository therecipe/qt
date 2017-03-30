package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	iparser "github.com/therecipe/qt/internal/binding/parser"

	"github.com/therecipe/qt/internal/utils"
)

func GetImports(path string, level int) []string {
	utils.Log.WithField("path", path).WithField("level", level).Debug("imports")

	importMap := make(map[string]struct{})

	cmd := exec.Command("go", "list", "-f", "'{{ join .TestImports \"|\" }}':'{{ join .Deps \"|\" }}'")
	cmd.Dir = path

	for _, dep := range strings.Split(utils.RunCmd(cmd, "go list deps"), "|") {
		if strings.Contains(dep, "github.com/therecipe/qt") && !strings.Contains(dep, "qt/internal") {
			continue
		}

		if strings.Contains(dep, "github.com/therecipe/qt/q") {
			iparser.LibDeps[iparser.MOC] = append(iparser.LibDeps[iparser.MOC], "Qml")
		}

		for _, gopath := range strings.Split(utils.GOPATH(), string(os.PathListSeparator)) {
			path := filepath.Join(gopath, "src", strings.Replace(dep, "\"", "", -1))
			if utils.ExistsDir(path) {
				importMap[path] = struct{}{}
			}
		}
	}

	var imports []string

	for k := range importMap {
		imports = append(imports, k)
	}

	return imports
}
