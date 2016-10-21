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
)

var BuildTarget string

func Minimal(path string) error {

	var (
		imported []string
		cached   []string
	)

	var walkFuncImports = func(path string, info os.FileInfo, err error) error {
		if err == nil && !strings.HasPrefix(info.Name(), "moc") && strings.HasSuffix(info.Name(), ".go") && !info.IsDir() {
			var pFile, errParse = goparser.ParseFile(token.NewFileSet(), path, nil, 0)
			if errParse != nil {
				fmt.Println("minimal.parseFile", errParse)
			} else {
				for _, i := range pFile.Imports {
					if !strings.Contains(i.Path.Value, "github.com/therecipe/qt") {
						var path = filepath.Join(os.Getenv("GOPATH"), "src", strings.Replace(i.Path.Value, "\"", "", -1))
						if _, err := ioutil.ReadDir(path); err == nil {
							if !isImported(imported, path) {
								imported = append(imported, path)
							}
						}
					}
				}
			}
		}
		return nil
	}

	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(info.Name(), ".go") && !info.IsDir() {

			var src, err = ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("minimal.readFile", err)
			}
			var file = string(src)

			if strings.Contains(file, "github.com/therecipe/qt/") && !(strings.Contains(file, "github.com/therecipe/qt/androidextras") && strings.Count(file, "github.com/therecipe/qt/") == 1) {
				cached = append(cached, file)
			}
		}

		return nil
	}

	filepath.Walk(path, walkFuncImports)
	for _, imp := range imported {
		filepath.Walk(imp, walkFuncImports)
	}

	filepath.Walk(path, walkFunc)
	for _, imp := range imported {
		filepath.Walk(imp, walkFunc)
	}

	for _, module := range templater.GetLibs() {
		if _, err := parser.GetModule(strings.ToLower(module)); err != nil {
			return err
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
					if strings.Contains(file, "."+converter.GoHeaderName(f)+"(") {
						exportFunction(c, f)
					}
				}
			}
		}
	}

	if BuildTarget == "sailfish" || BuildTarget == "sailfish-emulator" {
		parser.ClassMap["QQuickWidget"].Export = false
	}

	if BuildTarget == "ios" || BuildTarget == "ios-simulator" {
		parser.ClassMap["QProcess"].Export = false
		parser.ClassMap["QProcessEnvironment"].Export = false
	}

	templater.Minimal = true
	for _, module := range templater.GetLibs() {
		templater.GenModule(module)
	}
	return nil
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

func isImported(imported []string, path string) bool {
	for _, i := range imported {
		if i == path {
			return true
		}
	}
	return false
}
