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

var (
	ResourceNames      = make(map[string]string)
	ResourceNamesMutex = new(sync.Mutex)
)

func Rcc(path, target, tagsCustom, output_dir string, quickcompiler bool) {
	if utils.UseGOMOD(path) {
		if !utils.ExistsDir(filepath.Join(path, "vendor")) {
			cmd := exec.Command("go", "mod", "vendor")
			cmd.Dir = path
			utils.RunCmd(cmd, "go mod vendor")
		}
	}

	rcc(path, target, tagsCustom, output_dir, quickcompiler, true)
}

func rcc(path, target, tagsCustom, output_dir string, quickcompiler bool, root bool) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Rcc")

	//TODO: cache non go asset (*.qml, ...) hashes in rcc.go files to indentify staled assets in cached go packages
	//TODO: pure go.rcc for wasm/js targets

	if root {
		wg := new(sync.WaitGroup)
		defer wg.Wait()
		allImports := cmd.GetImports(path, target, tagsCustom, 0, false, false)
		wg.Add(len(allImports))
		for _, path := range allImports {
			go func(path string) {
				rcc(path, target, tagsCustom, path, quickcompiler, false)
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
		if !file.IsDir() && file.Name() == "qml" && !file.Mode().IsRegular() {
			hasQMLFiles = true
			break
		}
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

	pkgCmd := utils.GoList("{{.Name}}", "-find", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	pkgCmd.Dir = path
	for k, v := range env {
		pkgCmd.Env = append(pkgCmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	pkg := strings.TrimSpace(utils.RunCmd(pkgCmd, "run go list"))
	if pkg == "" {
		pkg = filepath.Base(path)
	}

	libs := []string{"Core"}
	if quickcompiler {
		libs = append(libs, "Qml")
	}
	rccCpp := filepath.Join(path, "rcc.cpp")
	if output_dir != "" {
		rccCpp = filepath.Join(output_dir, "rcc.cpp")
		templater.CgoTemplateSafe(pkg, output_dir, target, templater.RCC, "", "", libs)
	} else {
		templater.CgoTemplateSafe(pkg, path, target, templater.RCC, "", "", libs)
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

		//TODO: filter out duplicate assets

		utils.Save(rccQrc, content)
	}

	files, err = ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	var fileList []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".qrc" && !strings.HasSuffix(file.Name(), "_qml_cache.qrc") {
			//TODO: check for buildTags
			fileList = append(fileList, filepath.Join(path, file.Name()))
		}
	}

	nameCmd := utils.GoList("{{.ImportPath}}", "-find", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	nameCmd.Dir = path
	for k, v := range env {
		nameCmd.Env = append(nameCmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	name := "rcc_" + strings.TrimSpace(utils.RunCmd(nameCmd, "run go list"))
	for _, s := range []string{"/", ".", "-"} {
		name = strings.Replace(name, s, "_", -1)
	}

	if cachgen := utils.ToolPath("qmlcachegen", target); utils.ExistsFile(cachgen) && quickcompiler {
		utils.RemoveAll(rccCpp)

		var filteredFiles []string
		var possibleMixedFiles []string
		var pureFiles []string

		for _, f := range fileList {
			newName := filepath.Join(filepath.Dir(f), name+"_"+strings.TrimSuffix(filepath.Base(f), ".qrc")+"_qml_cache.qrc")

			utils.RunCmd(exec.Command(cachgen, "--filter-resource-file", "-o", newName, f), fmt.Sprintf("execute qmlcachegen filter on %v for %v", runtime.GOOS, target))

			for _, tBC := range strings.Split(strings.TrimSpace(utils.RunCmd(exec.Command(utils.ToolPath("rcc", target), "-list", f), "execute rcc")), "\n") {
				tBC = strings.TrimSpace(tBC)
				if strings.HasSuffix(tBC, ".qml") || strings.HasSuffix(tBC, ".js") {
					tBCT := strings.Replace(tBC, ".", "_", -1)

					cmd := exec.Command(cachgen)

					/* TODO: re-enable to warn about duplicate assets?
					for _, fl := range fileList {
						cmd.Args = append(cmd.Args, []string{"--resource", fl}...)
					}
					*/
					cmd.Args = append(cmd.Args, []string{"--resource", f}...)

					cmd.Args = append(cmd.Args, []string{"-o", filepath.Join(filepath.Dir(f), "rcc_"+strings.TrimSuffix(filepath.Base(f), ".qrc")+"_"+filepath.Base(filepath.Dir(tBCT))+"_"+filepath.Base(tBCT+"_qml_cache.cpp")), tBC}...)
					utils.RunCmd(cmd, fmt.Sprintf("execute qmlcachegen cache on %v for %v", runtime.GOOS, target))
				}
			}

			tmpName := filepath.Join(filepath.Dir(f), name+"_"+filepath.Base(f))

			if utils.ExistsFile(newName) {
				filteredFiles = append(filteredFiles, newName)

				utils.Save(tmpName, utils.Load(f))
				possibleMixedFiles = append(possibleMixedFiles, tmpName)

				defer func() { utils.RemoveAll(newName) }()
			} else {

				utils.Save(tmpName, utils.Load(f))
				pureFiles = append(pureFiles, tmpName)
			}

			defer func() { utils.RemoveAll(tmpName) }()
		}

		if len(possibleMixedFiles) > 0 || len(pureFiles) > 0 {
			cmd := exec.Command(utils.ToolPath("qmlcachegen", target))
			cmd.Dir = path

			for i, _ := range possibleMixedFiles {
				cmd.Args = append(cmd.Args, fmt.Sprintf("--resource-file-mapping=%v=%v", possibleMixedFiles[i], filteredFiles[i]))
			}
			for _, f := range pureFiles {
				cmd.Args = append(cmd.Args, fmt.Sprintf("--resource-file-mapping=%v", f))
			}

			cmd.Args = append(cmd.Args, []string{"-o", "rcc_qmlcache_loader.cpp"}...)
			cmd.Args = append(cmd.Args, possibleMixedFiles...)
			cmd.Args = append(cmd.Args, pureFiles...)

			var initNameList []string
			for _, f := range append(possibleMixedFiles, pureFiles...) {
				initNameList = append(initNameList, strings.TrimSuffix(filepath.Base(f), ".qrc"))
			}
			ResourceNamesMutex.Lock()
			ResourceNames[filepath.Join(path, "rcc_qmlcache_loader.cpp")] = strings.Join(initNameList, "|")
			ResourceNamesMutex.Unlock()

			utils.RunCmd(cmd, fmt.Sprintf("execute qmlcachegen loader on %v for %v", runtime.GOOS, target))
		}

		for _, f := range filteredFiles {
			rcc := exec.Command(utils.ToolPath("rcc", target), "-name", strings.TrimSuffix(filepath.Base(f), ".qrc"), "-o", strings.Replace(filepath.Base(strings.TrimSuffix(f, ".qrc")), name, "rcc", -1)+".cpp")
			rcc.Dir = path
			rcc.Args = append(rcc.Args, f)
			utils.RunCmd(rcc, fmt.Sprintf("execute per file rcc *.cpp on %v for %v", runtime.GOOS, target))

			ResourceNamesMutex.Lock()
			ResourceNames[f+".cpp"] = strings.TrimSuffix(filepath.Base(f), ".qrc")
			ResourceNamesMutex.Unlock()
		}

	} else {
		ResourceNamesMutex.Lock()
		ResourceNames[rccCpp] = name
		ResourceNamesMutex.Unlock()

		rcc := exec.Command(utils.ToolPath("rcc", target), "-name", name, "-o", rccCpp)
		rcc.Args = append(rcc.Args, fileList...)
		utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.cpp on %v for %v", runtime.GOOS, target))
	}

	if utils.QT_DEBUG_QML() {
		utils.Save("debug.pro", fmt.Sprintf("RESOURCES += %v", strings.Join(fileList, " ")))
	}

	if utils.QT_DOCKER() {
		if idug, ok := os.LookupEnv("IDUG"); ok {
			utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
		}
	}
}
