package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/therecipe/qt/internal/utils"
)

var (
	std      []string
	stdMutex = new(sync.Mutex)

	imported      = make(map[string]string)
	importedMutex = new(sync.Mutex)

	importedStd      = make(map[string]struct{})
	importedStdMutex = new(sync.Mutex)
)

func IsStdPkg(pkg string) bool {

	stdMutex.Lock()
	if std == nil {
		std = append(strings.Split(strings.TrimSpace(utils.RunCmd(exec.Command("go", "list", "std"), "go list std")), "\n"), "C")
	}
	stdMutex.Unlock()

	for _, spkg := range std {
		if pkg == spkg {
			return true
		}
	}
	return false
}

func GetImports(path, target, tagsCustom string, level int, onlyDirect bool) []string {
	utils.Log.WithField("path", path).WithField("level", level).Debug("get imports")

	env, tags, _, _ := BuildEnv(target, "", "")

	stdMutex.Lock()
	if std == nil {
		std = append(strings.Split(strings.TrimSpace(utils.RunCmd(exec.Command("go", "list", "std"), "go list std")), "\n"), "C")
	}
	stdMutex.Unlock()

	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}

	importMap := make(map[string]struct{})

	imp := "Deps"
	if onlyDirect {
		imp = "Imports"
	}

	//TODO: cache
	cmd := utils.GoList("{{join .TestImports \"|\"}}|{{join .XTestImports \"|\"}}|{{join ."+imp+" \"|\"}}", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	cmd.Dir = path
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	libs := strings.Split(strings.TrimSpace(utils.RunCmd(cmd, "go list deps")), "|")
	for i := len(libs) - 1; i >= 0; i-- {
		if strings.TrimSpace(libs[i]) == "" {
			libs = append(libs[:i], libs[i+1:]...)
			continue
		}

		importedMutex.Lock()
		dep, ok := imported[libs[i]]
		importedMutex.Unlock()
		if ok {
			importMap[dep] = struct{}{}
			libs = append(libs[:i], libs[i+1:]...)
			continue
		}

		for _, k := range std {
			if libs[i] == k {
				libs = append(libs[:i], libs[i+1:]...)
				break
			}
		}
	}

	wg := new(sync.WaitGroup)
	wc := make(chan bool, runtime.NumCPU()*2)
	wg.Add(len(libs))
	for _, l := range libs {
		wc <- true
		go func(l string) {
			defer func() {
				<-wc
				wg.Done()
			}()

			if strings.Contains(l, "github.com/therecipe/qt") && !strings.Contains(l, "qt/internal") {
				if strings.Contains(l, "github.com/therecipe/qt/") {
					importedStdMutex.Lock()
					importedStd[l] = struct{}{}
					importedStdMutex.Unlock()
				}
				return
			}

			cmd := utils.GoList("{{.Dir}}", "-find", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")), l)
			for k, v := range env {
				cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
			}

			dep := strings.TrimSpace(utils.RunCmd(cmd, "go list dir"))
			if dep == "" {
				return
			}

			importedMutex.Lock()
			importMap[dep] = struct{}{}
			imported[l] = dep
			importedMutex.Unlock()
		}(l)
	}
	wg.Wait()

	var imports []string
	for k := range importMap {
		imports = append(imports, k)
	}
	return imports
}

func GetQtStdImports() (o []string) {
	for k := range importedStd {
		o = append(o, strings.TrimPrefix(k, "github.com/therecipe/qt/"))
	}
	return
}

func GetGoFiles(path, target, tagsCustom string) []string {
	utils.Log.WithField("path", path).WithField("target", target).WithField("tagsCustom", tagsCustom).Debug("get go files")

	env, tags, _, _ := BuildEnv(target, "", "")
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}

	//TODO: cache
	cmd := utils.GoList("{{join .GoFiles \"|\"}}|{{join .CgoFiles \"|\"}}|{{join .TestGoFiles \"|\"}}|{{join .XTestGoFiles \"|\"}}", "-find", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	cmd.Dir = path
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	importMap := make(map[string]struct{})
	for _, v := range strings.Split(strings.TrimSpace(utils.RunCmd(cmd, "go list gofiles")), "|") {
		if strings.TrimSpace(v) != "" {
			importMap[v] = struct{}{}
		}
	}

	olibs := make([]string, 0)
	for k := range importMap {
		olibs = append(olibs, filepath.Join(path, k))
	}
	return olibs
}
