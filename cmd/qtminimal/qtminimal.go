package main

import (
	"fmt"
	goparser "go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildTarget = "desktop"
		appPath, _  = os.Getwd()
	)

	switch len(os.Args) {
	case 2:
		{
			buildTarget = os.Args[1]
		}

	case 3:
		{
			buildTarget = os.Args[1]
			appPath = os.Args[2]
		}
	}
	if !filepath.IsAbs(appPath) {
		appPath = utils.GetAbsPath(appPath)
	}
	if _, err := ioutil.ReadDir(appPath); err != nil {
		utils.Log.Fatalln("usage:", "qtminimal", "[ desktop | android | ... ]", filepath.Join("path", "to", "project"))
	}

	var (
		imported       []string
		cached         []string
		importedPkgMap = make(map[string]bool)
		goPaths = strings.Split(os.Getenv("GOPATH"), ":")
	)

	var walkFuncImports = func(appPath string, info os.FileInfo, err error) error {
		if err == nil && !strings.HasPrefix(info.Name(), "moc") && !strings.HasPrefix(info.Name(), "rcc") && strings.HasSuffix(info.Name(), ".go") && !info.IsDir() && !isBlacklisted(appPath, filepath.Dir(filepath.Join(appPath, info.Name()))) {
			var pFile, errParse = goparser.ParseFile(token.NewFileSet(), appPath, nil, 0)
			if errParse != nil {
				utils.Log.WithError(errParse).Panicf("failed to parser file %v", appPath)
			} else {
				for _, i := range pFile.Imports {
					if !strings.Contains(i.Path.Value, "github.com/therecipe/qt") {
						for _, goPath := range goPaths {
							var appPath = filepath.Join(goPath, "src", strings.Replace(i.Path.Value, "\"", "", -1))
							if _, err := ioutil.ReadDir(appPath); err == nil {
								if !isImported(imported, appPath) {
									imported = append(imported, appPath)
								}
							}
						}
					} else {
						if iPath := strings.Replace(i.Path.Value, "\"", "", -1); iPath != "github.com/therecipe/qt" {
							var pkg string
							for _, pkgs := range templater.Libs {
								if strings.ToLower(pkgs) == strings.TrimPrefix(iPath, "github.com/therecipe/qt/") {
									pkg = pkgs
								}
							}
							if _, exists := templater.LibDeps[pkg]; exists {
								importedPkgMap[pkg] = true
								for _, dep := range templater.LibDeps[pkg] {
									if _, exists := templater.LibDeps[dep]; exists {
										importedPkgMap[dep] = true
									}
								}
							}
						}
					}
				}
			}
		}
		return nil
	}

	var walkFunc = func(appPath string, info os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(info.Name(), ".go") && !info.IsDir() && !isBlacklisted(appPath, filepath.Dir(filepath.Join(appPath, info.Name()))) {

			if file := utils.Load(appPath); strings.Contains(file, "github.com/therecipe/qt/") &&
				!(strings.Contains(file, "github.com/therecipe/qt/androidextras") && strings.Count(file, "github.com/therecipe/qt/") == 1) {
				cached = append(cached, file)
			}
		}

		return nil
	}

	filepath.Walk(appPath, walkFuncImports)
	for _, imp := range imported {
		filepath.Walk(imp, walkFuncImports)
	}

	filepath.Walk(appPath, walkFunc)
	for _, imp := range imported {
		filepath.Walk(imp, walkFunc)
	}

	var importedPkgs []string
	for _, dep := range templater.Libs {
		if _, exist := importedPkgMap[dep]; exist {
			importedPkgs = append(importedPkgs, dep)
		}
	}

	for _, module := range importedPkgs {
		utils.Log.Debugf("loading qt/%v", strings.ToLower(module))
		if _, err := parser.GetModule(module); err != nil {
			utils.Log.WithError(err).Errorf("failed to load qt/%v", strings.ToLower(module))
		}
	}

	var file = strings.Join(cached, "")

	for className, c := range parser.ClassMap {
		if strings.Contains(file, className) {
			c.Export = true

			if hasPureVirtualFunctions(className) {
				for _, f := range c.Functions {
					if f.Virtual == parser.PURE {
						exportFunction(c, f)
					}
				}
			}

			for _, bc := range c.GetAllBases() {
				parser.ClassMap[bc].Export = true
			}
		}

		for _, f := range c.Functions {
			switch {
			case (f.Virtual == parser.IMPURE || f.Virtual == parser.PURE || f.Meta == parser.SIGNAL || f.Meta == parser.SLOT):
				{
					for _, signalMode := range []string{parser.CONNECT, parser.DISCONNECT, ""} {
						f.SignalMode = signalMode

						if strings.Contains(file, "."+converter.GoHeaderName(f)+"(") {
							exportFunction(c, f)
						}
					}
				}

			default:
				{
					if f.Static {
						if strings.Contains(file, converter.GoHeaderName(f)+"(") {
							exportFunction(c, f)
						}

						var fTmp = *f
						fTmp.Static = false

						if strings.Contains(file, "."+converter.GoHeaderName(&fTmp)+"(") {
							exportFunction(c, f)
						}
					} else {
						if strings.Contains(file, "."+converter.GoHeaderName(f)+"(") {
							exportFunction(c, f)
						}
					}
				}
			}
		}
	}

	if buildTarget == "sailfish" || buildTarget == "sailfish-emulator" {
		parser.ClassMap["QQuickWidget"].Export = false
	}

	if buildTarget == "ios" || buildTarget == "ios-simulator" {
		parser.ClassMap["QProcess"].Export = false
		parser.ClassMap["QProcessEnvironment"].Export = false
	}

	templater.Minimal = true
	for _, module := range importedPkgs {
		utils.Log.Debugf("generating minimal qt/%v", strings.ToLower(module))
		templater.GenModule(strings.Title(module))
	}
}

func exportFunction(class *parser.Class, function *parser.Function) {
	for _, f := range class.Functions {
		if converter.GoHeaderName(function) == converter.GoHeaderName(f) {

			f.Export = true

			class.Export = true

			for _, bc := range class.GetAllBases() {
				parser.ClassMap[bc].Export = true
			}

			for _, p := range f.Parameters {
				if class, exists := parser.ClassMap[converter.CleanValue(p.Value)]; exists {
					class.Export = true

					for _, bc := range class.GetAllBases() {
						parser.ClassMap[bc].Export = true
					}
				}
			}

			if class, exists := parser.ClassMap[converter.CleanValue(f.Output)]; exists {
				class.Export = true

				for _, bc := range class.GetAllBases() {
					parser.ClassMap[bc].Export = true
				}
			}

			if parser.IsPackedList(f.Output) {
				for _, f := range class.Functions {
					if f.Name == fmt.Sprintf("%v_atList", function.Name) {
						exportFunction(class, f)
					}
				}
			}
		}
	}
}

func hasPureVirtualFunctions(className string) bool {
	for _, f := range parser.ClassMap[className].Functions {
		if f.Virtual == parser.PURE {
			return true
		}
	}
	return false
}

func isImported(imported []string, appPath string) bool {
	for _, i := range imported {
		if i == appPath {
			return true
		}
	}
	return false
}

func isBlacklisted(appPath, currentPath string) bool {

	for _, n := range []string{"deploy", "qml", "android", "ios", "ios-simulator", "sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "node_modules", ".git"} {
		if strings.Contains(filepath.Join(currentPath), filepath.Join(appPath, n)) {
			return true
		}
	}

	return false
}
