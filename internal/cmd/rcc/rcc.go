package rcc

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/cmd"

	"github.com/therecipe/qt/internal/utils"
)

func Rcc(path, target, tagsCustom, output_dir string) {
	rcc(path, target, tagsCustom, output_dir, true)
}

func rcc(path, target, tagsCustom, output_dir string, root bool) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Rcc")

	if root {
		wg := new(sync.WaitGroup)
		defer wg.Wait()
		allImports := cmd.GetImports(path, target, tagsCustom, 0, false)
		wg.Add(len(allImports))
		for _, path := range allImports {
			go func(path string) {
				rcc(path, target, tagsCustom, path, false)
				wg.Done()
			}(path)
		}
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Error("failed to read dir")
		return
	}
	var hasQMLFiles bool
	for _, file := range files {
		if file.IsDir() && file.Name() == "qml" {
			hasQMLFiles = true
			break
		}
		ext := filepath.Ext(file.Name())
		if ext == ".qrc" || ext == ".qml" {
			hasQMLFiles = true
			break
		}
	}
	if !hasQMLFiles {
		return
	}

	rccQrc := filepath.Join(path, "rcc.qrc")

	env, tags, _, _ := cmd.BuildEnv(target, "", "")
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}

	pkgCmd := exec.Command("go", "list", "-e", "-f", "{{.Name}}", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	pkgCmd.Dir = path
	for k, v := range env {
		pkgCmd.Env = append(pkgCmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

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

	if dir := filepath.Join(path, "qml"); utils.ExistsDir(dir) {
		rcc := exec.Command(utils.ToolPath("rcc", target), "-project", "-o", rccQrc)
		rcc.Dir = dir
		utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.qrc on %v for %v", runtime.GOOS, target))

		content := utils.Load(rccQrc)
		content = strings.Replace(content, "<file>./", "<file>qml/", -1)
		if utils.ExistsFile(filepath.Join(path, "qtquickcontrols2.conf")) {
			content = strings.Replace(content, "<qresource>", "<qresource>\n<file>qtquickcontrols2.conf</file>", -1)
		}
		utils.Save(rccQrc, content)
	}

	files, err = ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	var fileList []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".qrc" {
			//TODO: check for buildTags
			fileList = append(fileList, filepath.Join(path, file.Name()))
		}
	}

	nameCmd := exec.Command("go", "list", "-e", "-f", "{{.ImportPath}}", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	nameCmd.Dir = path
	for k, v := range env {
		nameCmd.Env = append(nameCmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	name := strings.TrimSpace(utils.RunCmd(nameCmd, "run go list"))

	rcc := exec.Command(utils.ToolPath("rcc", target), "-name", strings.Replace(name, "/", "_", -1), "-o", rccCpp)
	rcc.Args = append(rcc.Args, fileList...)
	utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.cpp on %v for %v", runtime.GOOS, target))

	if utils.QT_DEBUG_QML() {
		utils.Save("debug.pro", fmt.Sprintf("RESOURCES += %v", strings.Join(fileList, " ")))
	}

	if utils.QT_DOCKER() {
		if idug, ok := os.LookupEnv("IDUG"); ok {
			utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
		}
	}
}
