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

func GetImports(path, target, tagsCustom string, level int, onlyDirect bool) []string {
	utils.Log.WithField("path", path).WithField("level", level).Debug("get imports")

	if std == nil {
		std = append(strings.Split(strings.TrimSpace(utils.RunCmd(exec.Command("go", "list", "std"), "go list std")), "\n"), "C")
	}

	env, tags, _, _ := BuildEnv(target, "", "")
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}

	importMap := make(map[string]struct{})

	imp := "Deps"
	if onlyDirect {
		imp = "Imports"
	}

	cmd := exec.Command("go", "list", "-e", "-f", "'{{ join .TestImports \"|\" }}':'{{ join ."+imp+" \"|\" }}'", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	cmd.Dir = path
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	out := strings.TrimSpace(utils.RunCmd(cmd, "go list deps"))
	out = strings.Replace(out, "'", "", -1)
	out = strings.Replace(out, ":", "|", -1)
	libs := strings.Split(out, "|")

	for i := len(libs) - 1; i >= 0; i-- {
		if strings.TrimSpace(libs[i]) == "" {
			libs = append(libs[:i], libs[i+1:]...)
			continue
		}

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

	for _, l := range libs {
		cmd = exec.Command("go", "list", "-e", "-f", "{{.Dir}}", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")), l)
		for k, v := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
		}

		dep := strings.TrimSpace(utils.RunCmd(cmd, "go list dir"))
		if dep == "" {
			continue
		}

		if strings.Contains(dep, filepath.Join("github.com", "therecipe", "qt")) && !strings.Contains(dep, filepath.Join("qt", "internal")) {
			continue
		}

		if strings.Contains(dep, filepath.Join("github.com", "therecipe", "qt", "q")) {
			iparser.LibDeps[iparser.MOC] = append(iparser.LibDeps[iparser.MOC], "Qml")
		}

		importMap[dep] = struct{}{}
		imported[l] = dep
	}

	var imports []string
	for k := range importMap {
		imports = append(imports, k)
	}
	return imports
}

func GetGoFiles(path, target, tagsCustom string) []string {
	utils.Log.WithField("path", path).WithField("target", target).WithField("tagsCustom", tagsCustom).Debug("get go files")

	env, tags, _, _ := BuildEnv(target, "", "")
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}

	cmd := exec.Command("go", "list", "-e", "-f", "'{{ join .GoFiles \"|\" }}':'{{ join .CgoFiles \"|\" }}':'{{ join .TestGoFiles \"|\" }}'", fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))
	cmd.Dir = path
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	out := strings.TrimSpace(utils.RunCmd(cmd, "go list gofiles"))
	out = strings.Replace(out, "'", "", -1)
	out = strings.Replace(out, ":", "|", -1)

	importMap := make(map[string]struct{})
	for _, v := range strings.Split(out, "|") {
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
