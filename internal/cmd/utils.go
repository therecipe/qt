package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	iparser "github.com/therecipe/qt/internal/binding/parser"

	"github.com/therecipe/qt/internal/utils"
)

var std []string

var imported = make(map[string]string)

func GetImports(path, target string, level int) []string {
	utils.Log.WithField("path", path).WithField("level", level).Debug("imports")

	if std == nil {
		std = strings.Split(strings.TrimSpace(utils.RunCmd(exec.Command("go", "list", "std"), "go list std")), "\n")
	}

	env, tags, _, _ := BuildEnv(target, "", "")

	importMap := make(map[string]struct{})

	cmd := exec.Command("go", "list", "-f", "'{{ join .TestImports \"|\" }}':'{{ join .Deps \"|\" }}'", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	cmd.Dir = path
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	out := strings.TrimSpace(utils.RunCmd(cmd, "go list deps"))
	out = strings.Replace(out, "'", "", -1)
	out = strings.Replace(out, ":", "|", -1)
	libs := strings.Split(out, "|")

	for i := len(libs) - 1; i >= 0; i-- {
		if dep, ok := imported[libs[i]]; ok {
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

	if len(libs) == 0 {
		var imports []string
		for k := range importMap {
			imports = append(imports, k)
		}
		return imports
	}

	cmd = exec.Command("go", "list", "-e", "-f", "{{.Dir}}", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	cmd.Args = append(cmd.Args, libs...)
	cmd.Dir = path
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	for i, dep := range strings.Split(strings.TrimSpace(utils.RunCmd(cmd, "go list dir")), "\n") {
		if strings.Contains(dep, filepath.Join("github.com", "therecipe", "qt")) && !strings.Contains(dep, filepath.Join("qt", "internal")) {
			continue
		}

		if strings.Contains(dep, filepath.Join("github.com", "therecipe", "qt", "q")) {
			iparser.LibDeps[iparser.MOC] = append(iparser.LibDeps[iparser.MOC], "Qml")
		}

		importMap[dep] = struct{}{}

		imported[libs[i]] = dep
	}

	var imports []string
	for k := range importMap {
		imports = append(imports, k)
	}
	return imports
}
