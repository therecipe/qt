package minimal

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

func QmakeMinimal(path, target string) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start QmakeMinimal")

	var files []string
	for _, path := range append([]string{path}, getAllImports(path, 0)...) {
		fileList, err := ioutil.ReadDir(path)
		if err != nil {
			utils.Log.WithError(err).Error("failed to read dir")
			continue
		}

		for _, file := range fileList {
			path := filepath.Join(path, file.Name())
			if !file.IsDir() && filepath.Ext(path) == ".go" {
				files = append(files, utils.Load(path))
				utils.Log.WithField("path", path).Debug("analyse for minimal")
			}
		}
	}

	c := len(files)
	utils.Log.Debugln("found", c, "files to analyse")
	if c == 0 {
		return
	}

	if len(parser.State.ClassMap) == 0 {
		parser.LoadModules()
	} else {
		utils.Log.Debug("modules already cached")
	}

	for _, f := range files {
		for _, c := range parser.State.ClassMap {
			if strings.Contains(f, c.Name) &&
				strings.Contains(f, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(strings.TrimPrefix(c.Module, "Qt")))) {
				QmakeExportClass(c, files)
			}
		}
	}

	/* TODO:
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
	*/

	//TODO: cleanup state
	parser.State.Minimal = true
	for _, m := range parser.GetLibs() {
		templater.GenModule(m, target, templater.MINIMAL)
	}
	parser.State.Minimal = false
	for _, c := range parser.State.ClassMap {
		c.Export = false
		for _, f := range c.Functions {
			f.Export = false
		}
	}
}

func getAllImports(path string, level int) []string {
	utils.Log.WithField("path", path).WithField("level", level).Debug("getAllImports")

	var imports []string

	level++
	if level > 2 {
		return imports
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	var fileList []string
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".go" {
			fileList = append(fileList, filepath.Join(path, f.Name()))
		}
	}

	for _, f := range fileList {
		file, err := goparser.ParseFile(token.NewFileSet(), f, nil, 0)
		if err != nil {
			utils.Log.WithError(err).Debugln("failed to parse", f)
			continue
		}
		for _, i := range file.Imports {
			if strings.Contains(i.Path.Value, "github.com/therecipe/qt") && !strings.Contains(i.Path.Value, "qt/internal") {
				continue
			}

			for _, gopath := range strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator)) {
				path := filepath.Join(gopath, "src", strings.Replace(i.Path.Value, "\"", "", -1))
				if utils.ExistsDir(path) {
					var has bool
					for _, i := range imports {
						if i == path {
							has = true
							break
						}
					}
					if has {
						continue
					}
					imports = append(imports, path)
					for _, path := range getAllImports(path, level) {
						var has bool
						for _, i := range imports {
							if i == path {
								has = true
								break
							}
						}
						if !has {
							imports = append(imports, path)
						}
					}
				}
			}

		}
	}

	return imports
}

func QmakeExportClass(c *parser.Class, files []string) {
	if c.Export {
		return
	}
	c.Export = true

	for _, file := range files {
		for _, f := range c.Functions {

			switch {
			case f.Virtual == parser.IMPURE, f.Virtual == parser.PURE, f.Meta == parser.SIGNAL, f.Meta == parser.SLOT:
				for _, mode := range []string{parser.CONNECT, parser.DISCONNECT, ""} {
					f.SignalMode = mode
					if strings.Contains(file, "."+converter.GoHeaderName(f)+"(") {
						QmakeExportFunction(f, files)
					}
				}

			default:
				if f.Static {
					if strings.Contains(file, "."+converter.GoHeaderName(f)+"(") {
						QmakeExportFunction(f, files)
					}
					f.Static = false
					if strings.Contains(file, "."+converter.GoHeaderName(f)+"(") {
						QmakeExportFunction(f, files)
					}
					f.Static = true
				} else {
					if strings.Contains(file, "."+converter.GoHeaderName(f)+"(") {
						QmakeExportFunction(f, files)
					}
				}
			}

			if strings.HasPrefix(f.Name, "__") || f.Meta == parser.CONSTRUCTOR ||
				f.Meta == parser.DESTRUCTOR || f.Virtual == parser.PURE {
				QmakeExportFunction(f, files)
			}

		}
	}

	for _, b := range c.GetAllBases() {
		if c, ok := parser.State.ClassMap[b]; ok {
			QmakeExportClass(c, files)
		}
	}
}

func QmakeExportFunction(f *parser.Function, files []string) {
	if f.Export {
		return
	}
	f.Export = true

	for _, p := range f.Parameters {
		if c, ok := parser.State.ClassMap[parser.CleanValue(p.Value)]; ok {
			QmakeExportClass(c, files)
		}
		if parser.IsPackedList(p.Value) {
			if c, ok := parser.State.ClassMap[parser.UnpackedList(p.Value)]; ok {
				QmakeExportClass(c, files)
			}
		}
		if parser.IsPackedMap(p.Value) {
			key, value := parser.UnpackedMap(p.Value)
			if c, ok := parser.State.ClassMap[key]; ok {
				QmakeExportClass(c, files)
			}
			if c, ok := parser.State.ClassMap[value]; ok {
				QmakeExportClass(c, files)
			}
		}
	}

	if c, ok := parser.State.ClassMap[parser.CleanValue(f.Output)]; ok {
		QmakeExportClass(c, files)
	}
	if parser.IsPackedList(f.Output) {
		if c, ok := parser.State.ClassMap[parser.UnpackedList(f.Output)]; ok {
			QmakeExportClass(c, files)
		}
	}
	if parser.IsPackedMap(f.Output) {
		key, value := parser.UnpackedMap(f.Output)
		if c, ok := parser.State.ClassMap[key]; ok {
			QmakeExportClass(c, files)
		}
		if c, ok := parser.State.ClassMap[value]; ok {
			QmakeExportClass(c, files)
		}
	}
}
