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

func QmakeRcc(path, target string, output_dir *string) {
	if dir := filepath.Join(path, "qml"); !utils.ExistsDir(dir) {
		utils.MkdirAll(dir)
	}

	var (
		rccQrc = filepath.Join(path, "rcc.qrc")

		rccCpp = filepath.Join(path, "rcc.cpp")
	)
	if output_dir != nil {
		rccCpp = filepath.Join(*output_dir, "rcc.cpp")
		templater.QmakeCgoTemplate("main", *output_dir, target, templater.RCC)
	} else {
		templater.QmakeCgoTemplate("main", path, target, templater.RCC)
	}

	var rcc = exec.Command(utils.ToolPath("rcc", target), "-project", "-o", rccQrc)
	rcc.Dir = filepath.Join(path, "qml")
	utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.qrc on %v for %v", runtime.GOOS, target))

	var content = utils.Load(rccQrc)
	content = strings.Replace(content, "<file>./", "<file>qml/", -1)
	if utils.ExistsFile(filepath.Join(path, "qtquickcontrols2.conf")) {
		content = strings.Replace(content, "<qresource>", "<qresource>\n<file>qtquickcontrols2.conf</file>", -1)
	}
	utils.Save(rccQrc, content)

	var (
		files, err = ioutil.ReadDir(path)
		fileList   []string
	)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".qrc") {
			fileList = append(fileList, filepath.Join(path, f.Name()))
		}
	}

	rcc = exec.Command(utils.ToolPath("rcc", target), "-o", rccCpp)
	rcc.Args = append(rcc.Args, fileList...)
	utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.cpp on %v for %v", runtime.GOOS, target))
}

func QmakeCleanPath(path string) {
	var files, err = ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasPrefix(f.Name(), "rcc") {
			utils.RemoveAll(filepath.Join(path, f.Name()))
		}
	}
}
