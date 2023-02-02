package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/bluszcz/cutego/internal/utils"
)

var (
	std      []string
	stdMutex = new(sync.Mutex)

	imported      = make(map[string]string)
	importedMutex = new(sync.Mutex)

	importedStd       = make(map[string]struct{})
	importedStdMutex  = new(sync.Mutex)
	importsQmlOrQuick string
	importsInterop    string
	importsFlutter    string
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

//TODO:
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
	} else {
		defer func() {
			importedStdMutex.Lock()
			if importsQmlOrQuick == "" {
				stdImport := make([]string, len(importedStd)+1)
				var c int
				for k := range importedStd {
					stdImport[c] = k
					c++
				}
				stdImport[len(stdImport)-1] = "github.com/bluszcz/cutego/internal/binding/runtime"

				cmd := utils.GoList(fmt.Sprintf("{{if not .Standard}}{{if eq .ImportPath \"%v\"}}{{else}}{{range .Imports}}{{if eq . \"github.com/bluszcz/cutego/qml\" \"github.com/bluszcz/cutego/quick\"}}{{.}}{{end}}{{end}}{{end}}{{end}}", strings.Join(stdImport, "\" \"")), "-deps", utils.BuildTags(tags))
				if !utils.UseGOMOD(path) || (utils.UseGOMOD(path) && !strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/")) {
					cmd.Dir = path
				} else if utils.UseGOMOD(path) && strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/") {
					cmd.Dir = filepath.Dir(utils.GOMOD(path))
					vl := strings.Split(strings.Replace(path, "\\", "/", -1), "/vendor/")
					cmd.Args = append(cmd.Args, vl[len(vl)-1])
				}
				for k, v := range env {
					cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
				}

				importsQmlOrQuick = fmt.Sprint(strings.TrimSpace(utils.RunCmd(cmd, "go list imports qml or quick")) != "")

				utils.Log.WithField("path", path).Debug("project depends on qml or quick: " + importsQmlOrQuick)
			}

			if importsInterop == "" {
				stdImport := make([]string, len(importedStd)+1)
				var c int
				for k := range importedStd {
					stdImport[c] = k
					c++
				}
				stdImport[len(stdImport)-1] = "github.com/bluszcz/cutego/internal/binding/runtime"

				cmd := utils.GoList(fmt.Sprintf("{{if not .Standard}}{{if eq .ImportPath \"%v\"}}{{else}}{{range .Imports}}{{if eq . \"github.com/bluszcz/cutego/interop\"}}{{.}}{{end}}{{end}}{{end}}{{end}}", strings.Join(stdImport, "\" \"")), "-deps", utils.BuildTags(tags))
				if !utils.UseGOMOD(path) || (utils.UseGOMOD(path) && !strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/")) {
					cmd.Dir = path
				} else if utils.UseGOMOD(path) && strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/") {
					cmd.Dir = filepath.Dir(utils.GOMOD(path))
					vl := strings.Split(strings.Replace(path, "\\", "/", -1), "/vendor/")
					cmd.Args = append(cmd.Args, vl[len(vl)-1])
				}
				for k, v := range env {
					cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
				}

				importsInterop = fmt.Sprint(strings.TrimSpace(utils.RunCmd(cmd, "go list imports interop")) != "")

				utils.Log.WithField("path", path).Debug("project depends on interop: " + importsInterop)
			}

			if importsFlutter == "" {
				stdImport := make([]string, len(importedStd)+1)
				var c int
				for k := range importedStd {
					stdImport[c] = k
					c++
				}
				stdImport[len(stdImport)-1] = "github.com/bluszcz/cutego/internal/binding/runtime"

				cmd := utils.GoList(fmt.Sprintf("{{if not .Standard}}{{if eq .ImportPath \"%v\"}}{{else}}{{range .Imports}}{{if eq . \"github.com/bluszcz/cutego/flutter\"}}{{.}}{{end}}{{end}}{{end}}{{end}}", strings.Join(stdImport, "\" \"")), "-deps", utils.BuildTags(tags))
				if !utils.UseGOMOD(path) || (utils.UseGOMOD(path) && !strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/")) {
					cmd.Dir = path
				} else if utils.UseGOMOD(path) && strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/") {
					cmd.Dir = filepath.Dir(utils.GOMOD(path))
					vl := strings.Split(strings.Replace(path, "\\", "/", -1), "/vendor/")
					cmd.Args = append(cmd.Args, vl[len(vl)-1])
				}
				for k, v := range env {
					cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
				}

				importsFlutter = fmt.Sprint(strings.TrimSpace(utils.RunCmd(cmd, "go list imports flutter")) != "")

				utils.Log.WithField("path", path).Debug("project depends on flutter: " + importsFlutter)
			}
			importedStdMutex.Unlock()
		}()
	}

	//TODO: cache
	cmd := utils.GoList("{{join .TestImports \"|\"}}|{{join .XTestImports \"|\"}}|{{join ."+imp+" \"|\"}}", utils.BuildTags(tags))
	if !utils.UseGOMOD(path) || (utils.UseGOMOD(path) && !strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/")) {
		cmd.Dir = path
	} else if utils.UseGOMOD(path) && strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/") {
		cmd.Dir = filepath.Dir(utils.GOMOD(path))
		vl := strings.Split(strings.Replace(path, "\\", "/", -1), "/vendor/")
		cmd.Args = append(cmd.Args, vl[len(vl)-1])
	}

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

			if strings.Contains(l, "github.com/bluszcz/cutego") && !(strings.Contains(l, "qt/internal") || strings.Contains(l, "qt/flutter") || strings.Contains(l, "qt/interop")) {
				if strings.Contains(l, "github.com/bluszcz/cutego/") {
					importedStdMutex.Lock()
					importedStd[l] = struct{}{}
					importedStdMutex.Unlock()
				}
				return
			}

			cmd := utils.GoList("{{.Dir}}", "-find", utils.BuildTags(tags), l)
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
	importedStdMutex.Lock()
	for k := range importedStd {
		o = append(o, strings.TrimPrefix(k, "github.com/bluszcz/cutego/"))
	}
	importedStdMutex.Unlock()
	return
}

func ImportsQmlOrQuick() (o bool) {
	importedStdMutex.Lock()
	o = importsQmlOrQuick == "true"
	importedStdMutex.Unlock()
	return
}

func ImportsInterop() (o bool) {
	importedStdMutex.Lock()
	o = importsInterop == "true"
	importedStdMutex.Unlock()
	return
}

func ImportsFlutter() (o bool) {
	importedStdMutex.Lock()
	o = importsFlutter == "true"
	importedStdMutex.Unlock()
	return
}

func CleanupRegisteredImportsForCI() { importsQmlOrQuick = ""; importsInterop = ""; importsFlutter = "" }

func GetGoFiles(path, target, tagsCustom string) []string {
	utils.Log.WithField("path", path).WithField("target", target).WithField("tagsCustom", tagsCustom).Debug("get go files")

	env, tags, _, _ := BuildEnv(target, "", "")
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}

	//TODO: cache
	cmd := utils.GoList("{{join .GoFiles \"|\"}}|{{join .CgoFiles \"|\"}}|{{join .TestGoFiles \"|\"}}|{{join .XTestGoFiles \"|\"}}", "-find", utils.BuildTags(tags))
	if !utils.UseGOMOD(path) || (utils.UseGOMOD(path) && !strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/")) {
		cmd.Dir = path
	} else if utils.UseGOMOD(path) && strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/") {
		cmd.Dir = filepath.Dir(utils.GOMOD(path))
		vl := strings.Split(strings.Replace(path, "\\", "/", -1), "/vendor/")
		cmd.Args = append(cmd.Args, vl[len(vl)-1])
	}

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

//TODO: directly parse go.mod to make it possible to skip "go mod download"
func QtModVersion(path string) string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "github.com/bluszcz/cutego")
	cmd.Dir = path
	version := strings.TrimSpace(utils.RunCmdOptional(cmd, "get qt tooling version"))
	if !strings.HasPrefix(version, "v") {
		return "latest"
	}
	return version
}

func RestartWithPinnedVersion(path string) bool {
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = path
	utils.RunCmd(cmd, "download qt based on the go.mod version")

	if v := QtModVersion(path); strings.Count(v, "-") == 2 {
		if i, err := strconv.ParseInt(strings.Split(v, "-")[1], 10, 64); !(err == nil && i >= 20191110184604) { //6e660afb3df7
			return false
		}
	} else {
		return false
	}

	cmd = exec.Command("go", "install", "-v", "-tags=no_env", "github.com/bluszcz/cutego/cmd/...")
	cmd.Dir = path
	cmd.Env = append(os.Environ(), "GOBIN="+utils.GOBIN())
	utils.RunCmd(cmd, "re-install qt tooling based on the go.mod version")

	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	p, err := os.StartProcess(filepath.Join(utils.GOBIN(), filepath.Base(os.Args[0])), append(os.Args, "non_recursive"), procAttr)
	if err != nil {
		utils.Log.WithError(err).Error("failed to RestartWithPinnedVersion")
	}
	p.Wait()
	return true
}
