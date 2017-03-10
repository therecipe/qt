package minimal

import (
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

func Minimal(appPath, buildTarget string) {
	if utils.QT_QMAKE_CGO() {
		QmakeMinimal(appPath, buildTarget)
		return
	}

	var (
		imported       []string
		cached         []string
		importedPkgMap = make(map[string]struct{})
		goPaths        = strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))
	)

	var walkFuncImports = func(appPath string, info os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(info.Name(), ".go") && !info.IsDir() && !isBlacklisted(appPath, filepath.Dir(filepath.Join(appPath, info.Name()))) {
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
							for _, pkgs := range parser.Libs {
								if strings.ToLower(pkgs) == strings.TrimPrefix(iPath, "github.com/therecipe/qt/") {
									pkg = pkgs
								}
							}
							if _, ok := parser.LibDeps[pkg]; ok {
								importedPkgMap[pkg] = struct{}{}
								for _, dep := range parser.LibDeps[pkg] {
									if _, ok := parser.LibDeps[dep]; ok {
										importedPkgMap[dep] = struct{}{}
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
	for _, dep := range parser.Libs {
		if _, ok := importedPkgMap[dep]; ok {
			importedPkgs = append(importedPkgs, dep)
		}
	}

	if len(parser.State.ClassMap) == 0 {
		parser.LoadModules()
	}

	var file = strings.Join(cached, "")

	for className, c := range parser.State.ClassMap {
		if strings.Contains(file, className) {
			c.Export = true

			if hasPureVirtualFunctions(className) {
				for _, f := range c.Functions {
					if f.Virtual == parser.PURE {
						var cdna, _ = f.PossibleDerivationsInAllModules(true)
						for _, cdn := range cdna {
							var c, ok = parser.State.ClassMap[cdn]
							if !ok || !c.HasFunctionWithName(f.Name) {
								continue
							}
							exportFunction(c, c.GetFunction(f.Name))
						}
					}
				}
			}

			for _, bc := range c.GetAllBases() {
				parser.State.ClassMap[bc].Export = true
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
		if _, ok := parser.State.ClassMap["TestCase"]; ok {
			delete(parser.State.ClassMap, "TestCase")
		}

		//TODO: use parseFloat
		for _, c := range parser.State.ClassMap {
			switch c.Since {
			case "5.3", "5.4", "5.5", "5.6", "5.7", "5.8":
				delete(parser.State.ClassMap, c.Name)
			}

			if !parser.IsWhiteListedSailfishLib(strings.TrimPrefix(c.Module, "Qt")) {
				delete(parser.State.ClassMap, c.Name)
			}

			for _, f := range c.Functions {
				switch f.Since {
				case "5.3", "5.4", "5.5", "5.6", "5.7", "5.8":
					f.Export = false
				}
			}
		}
	}

	if buildTarget == "asteroid" {
		if _, ok := parser.State.ClassMap["TestCase"]; ok {
			delete(parser.State.ClassMap, "TestCase")
		}
		if _, ok := parser.State.ClassMap["QQuickWidget"]; ok {
			parser.State.ClassMap["QQuickWidget"].Export = false
		}

		for k, c := range parser.State.ClassMap {
			switch c.Since {
			case "5.7", "5.8":
				delete(parser.State.ClassMap, c.Name)
			}

			for _, f := range c.Functions {
				switch f.Since {
				case "5.7", "5.8":
					f.Export = false
				}
			}

			if strings.HasPrefix(k, "QAccessible") {
				delete(parser.State.ClassMap, k)
			}
		}
	}

	if buildTarget == "ios" || buildTarget == "ios-simulator" {
		parser.State.ClassMap["QProcess"].Export = false
		parser.State.ClassMap["QProcessEnvironment"].Export = false
	}

	parser.State.Minimal = true
	for _, module := range importedPkgs {
		utils.Log.Debugf("generating minimal qt/%v", strings.ToLower(module))
		templater.GenModule(strings.Title(module), buildTarget, templater.MINIMAL)
	}
}

func exportFunction(class *parser.Class, function *parser.Function) {
	for _, f := range class.Functions {
		if converter.GoHeaderName(function) == converter.GoHeaderName(f) {

			f.Export = true

			class.Export = true

			for _, bc := range class.GetAllBases() {
				parser.State.ClassMap[bc].Export = true
			}

			for _, p := range f.Parameters {
				if class, ok := parser.State.ClassMap[parser.CleanValue(p.Value)]; ok {
					class.Export = true

					for _, bc := range class.GetAllBases() {
						parser.State.ClassMap[bc].Export = true
					}
				}
			}

			if class, ok := parser.State.ClassMap[parser.CleanValue(f.Output)]; ok {
				exportClass(class)

				for _, bc := range class.GetAllBases() {
					parser.State.ClassMap[bc].Export = true
				}
			}

			if parser.IsPackedList(f.Output) || parser.IsPackedMap(f.Output) {
				for _, f := range class.Functions {
					if strings.Contains(f.Name, function.Name) &&
						(strings.HasSuffix(f.Name, "_atList") ||
							strings.HasSuffix(f.Name, "_setList") ||
							strings.HasSuffix(f.Name, "_newList") ||
							strings.HasSuffix(f.Name, "_keyList")) {
						f.Export = true
					}
				}
				for _, bc := range append(class.GetAllBases(), class.GetAllDerivations()...) {
					var class, ok = parser.State.ClassMap[bc]
					if !ok {
						continue
					}
					for _, f := range class.Functions {
						if strings.Contains(f.Name, function.Name) &&
							(strings.HasSuffix(f.Name, "_atList") ||
								strings.HasSuffix(f.Name, "_setList") ||
								strings.HasSuffix(f.Name, "_newList") ||
								strings.HasSuffix(f.Name, "_keyList")) {
							f.Export = true
						}
					}
				}
			}

			for _, p := range f.Parameters {
				if parser.IsPackedList(parser.CleanValue(p.Value)) || parser.IsPackedMap(p.Value) {
					for _, f := range class.Functions {
						if strings.Contains(f.Name, function.Name) &&
							(strings.HasSuffix(f.Name, "_atList") ||
								strings.HasSuffix(f.Name, "_setList") ||
								strings.HasSuffix(f.Name, "_newList") ||
								strings.HasSuffix(f.Name, "_keyList")) {
							f.Export = true
						}
					}
					for _, bc := range append(class.GetAllBases(), class.GetAllDerivations()...) {
						var class, ok = parser.State.ClassMap[bc]
						if !ok {
							continue
						}
						for _, f := range class.Functions {
							if strings.Contains(f.Name, function.Name) &&
								(strings.HasSuffix(f.Name, "_atList") ||
									strings.HasSuffix(f.Name, "_setList") ||
									strings.HasSuffix(f.Name, "_newList") ||
									strings.HasSuffix(f.Name, "_keyList")) {
								f.Export = true
							}
						}
					}
				}
			}

			//TODO: packed input
		}
	}
}

func exportClass(c *parser.Class) {
	c.Export = true
	for _, f := range c.Functions {
		if f.Meta == parser.CONSTRUCTOR && !f.Export {
			exportFunction(c, f)
		}
	}
}

func hasPureVirtualFunctions(className string) bool {
	for _, f := range parser.State.ClassMap[className].Functions {
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
