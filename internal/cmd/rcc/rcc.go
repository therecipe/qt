package rcc

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/cmd"

	"github.com/therecipe/qt/internal/utils"
)

var done = make(map[string]struct{})

func Rcc(path, target, tags, output_dir string) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Rcc")

	for i, path := range append([]string{path}, cmd.GetImports(path, target, tags, 0, true)...) {
		fileList, err := ioutil.ReadDir(path)
		if err != nil {
			utils.Log.WithError(err).Error("failed to read dir")
			continue
		}

		var hasQml bool
		for _, file := range fileList {
			if file.IsDir() && file.Name() == "qml" {
				hasQml = true
				break
			}
			ext := filepath.Ext(file.Name())
			if ext == ".qrc" || ext == ".qml" {
				hasQml = true
				break
			}
		}
		if !hasQml {
			continue
		}

		if _, ok := done[path]; !ok && i > 0 {
			done[path] = struct{}{}
			Rcc(path, target, tags, path)
		}
	}

	if dir := filepath.Join(path, "qml"); !utils.ExistsDir(dir) {
		utils.MkdirAll(dir)
	}

	rccQrc := filepath.Join(path, "rcc.qrc")

	pkgCmd := exec.Command("go", "list", "-e", "-f", "{{.Name}}")
	pkgCmd.Dir = path
	pkg := strings.TrimSpace(utils.RunCmd(pkgCmd, "run go list"))
	if pkg == "" {
		pkg = filepath.Base(path)
	}

	rccCpp := filepath.Join(path, "rcc.cpp")
	if output_dir != "" {
		rccCpp = filepath.Join(output_dir, "rcc.cpp")
		templater.CgoTemplate(pkg, output_dir, target, templater.RCC, "", "")
	} else {
		templater.CgoTemplate(pkg, path, target, templater.RCC, "", "")
	}

	rcc := exec.Command(utils.ToolPath("rcc", target), "-project", "-o", rccQrc)
	rcc.Dir = filepath.Join(path, "qml")
	utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.qrc on %v for %v", runtime.GOOS, target))

	content := utils.Load(rccQrc)
	content = strings.Replace(content, "<file>./", "<file>qml/", -1)
	if utils.ExistsFile(filepath.Join(path, "qtquickcontrols2.conf")) {
		content = strings.Replace(content, "<qresource>", "<qresource>\n<file>qtquickcontrols2.conf</file>", -1)
	}
	utils.Save(rccQrc, content)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	var fileList []string
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".qrc" {
			//TODO: check for buildTags
			fileList = append(fileList, filepath.Join(path, f.Name()))
		}
	}

	nameCmd := exec.Command("go", "list", "-f", "{{.ImportPath}}")
	nameCmd.Dir = path
	name := strings.TrimSpace(utils.RunCmd(nameCmd, "run go list"))

	rcc = exec.Command(utils.ToolPath("rcc", target), "-name", strings.Replace(name, "/", "_", -1), "-o", rccCpp)
	rcc.Args = append(rcc.Args, fileList...)
	utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.cpp on %v for %v", runtime.GOOS, target))

	if utils.QT_DEBUG_QML() {
		utils.Save("debug.pro", fmt.Sprintf("RESOURCES += %v", strings.Join(fileList, " ")))
	}
}
