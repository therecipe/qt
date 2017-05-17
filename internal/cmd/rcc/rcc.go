package rcc

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/utils"
)

func Rcc(path, target string, output_dir string) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Rcc")

	if dir := filepath.Join(path, "qml"); !utils.ExistsDir(dir) {
		utils.MkdirAll(dir)
	}

	rccQrc := filepath.Join(path, "rcc.qrc")

	rccCpp := filepath.Join(path, "rcc.cpp")
	if output_dir != "" {
		rccCpp = filepath.Join(output_dir, "rcc.cpp")
		templater.CgoTemplate("main", output_dir, target, templater.RCC, "main", "")
	} else {
		templater.CgoTemplate("main", path, target, templater.RCC, "main", "")
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

	rcc = exec.Command(utils.ToolPath("rcc", target), "-o", rccCpp)
	rcc.Args = append(rcc.Args, fileList...)
	utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.cpp on %v for %v", runtime.GOOS, target))

	if utils.QT_DEBUG_QML() {
		utils.Save("debug.pro", fmt.Sprintf("RESOURCES += %v", strings.Join(fileList, " ")))
	}
}
